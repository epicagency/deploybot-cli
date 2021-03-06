package main

import (
	"fmt"
	"github.com/extemporalgenome/slug"
	"os"
	"time"
)

func listRepositories() {
	var (
		repositories *Repositories
		err          error
	)
	if repositories, err = bot.GetRepositories(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, repository := range repositories.Entries {
		fmt.Printf("Repository %d: %s (%s)\n", repository.Id, repository.Title, repository.Name)
		if *verbose {
			fmt.Printf("\tType: %s\n\tRefresh: %s\n\tColor label: %s\n\tCreated: %s\n\tUpdated: %s\n", repository.Type, repository.RefreshWebhookUrl, repository.ColorLabel, repository.CreatedAt, repository.UpdatedAt)
		}
	}
}

func showRepository() {
	var (
		repository *Repository
		err        error
		id         int
	)
	if id, err = config.Alias(*repositoryIdFlag); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if repository, err = bot.GetRepository(id); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("Repository %d: %s (%s)\n", repository.Id, repository.Title, repository.Name)
	if *verbose {
		fmt.Printf("\tType: %s\n\tRefresh: %s\n\tColor label: %s\n\tCreated: %s\n\tUpdated: %s\n", repository.Type, repository.RefreshWebhookUrl, repository.ColorLabel, repository.CreatedAt, repository.UpdatedAt)
		listEnvironments()
	}
}

func listEnvironments() {
	var (
		environments *Environments
		err          error
		id           int
	)
	if id, err = config.Alias(*repositoryIdFlag); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if id != 0 {
		environments, err = bot.GetEnvironmentsByRepository(id)
	} else {
		environments, err = bot.GetEnvironments()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, environment := range environments.Entries {
		fmt.Printf("Environment %d: %s\n", environment.Id, environment.Name)
		if *verbose {
			fmt.Printf("\tCurrent version: %s\n\tIs automatic?: %t\n\tRepository Id: %d\n\tBranch name: %s\n\tDeploy webhook url: %s\n\tCreated: %s\n\tUpdated: %s\n", environment.CurrentVersion, environment.IsAutomatic, environment.RepositoryId, environment.BranchName, environment.DeployWebhookUrl, environment.CreatedAt, environment.UpdatedAt)
		}
	}
}

func showEnvironment() {
	var (
		environment *Environment
		err         error
	)
	id, _ := config.Alias(*environmentIdFlag)
	environment, err = bot.GetEnvironment(id)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("Environment %d: %s\n", environment.Id, environment.Name)
	if *verbose {
		fmt.Printf("\tCurrent version: %s\n\tIs automatic?: %t\n\tRepository Id: %d\n\tBranch name: %s\n\tDeploy webhook url: %s\n\tCreated: %s\n\tUpdated: %s\n", environment.CurrentVersion, environment.IsAutomatic, environment.RepositoryId, environment.BranchName, environment.DeployWebhookUrl, environment.CreatedAt, environment.UpdatedAt)
		listServers()
	}
}

func listUsers() {
	var (
		users *Users
		err   error
	)
	users, err = bot.GetUsers()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, user := range users.Entries {
		fmt.Printf("User %d: %s %s (%s)\n", user.Id, user.FirstName, user.LastName, user.Email)
		if *verbose {
			fmt.Printf("\tTimezone: %s\n\tIs admin?: %t\n\tCreated at: %s\n\tUpdated at: %s\n", user.Timezone, user.IsAdmin, user.CreatedAt, user.UpdatedAt)
		}
	}
}

func showUser() {
	var (
		user *User
		err  error
	)
	user, err = bot.GetUser(*userIdFlag)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("User %d: %s %s (%s)\n", user.Id, user.FirstName, user.LastName, user.Email)
	if *verbose {
		fmt.Printf("\tTimezone: %s\n\tIs admin?: %t\n\tCreated at: %s\n\tUpdated at: %s\n", user.Timezone, user.IsAdmin, user.CreatedAt, user.UpdatedAt)
	}
}

func listServers() {
	var (
		servers *Servers
		err     error
		id      int
	)

	if *environmentIdFlag != "" {
		if id, err = config.Alias(*environmentIdFlag); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		servers, err = bot.GetServersByEnvironment(id)
	} else if *repositoryIdFlag != "" {
		if id, err = config.Alias(*repositoryIdFlag); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		servers, err = bot.GetServersByRepository(id)
	} else {
		servers, err = bot.GetServers()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, server := range servers.Entries {
		fmt.Printf("Server %d: %s\n", server.Id, server.Name)
		if *verbose {
			fmt.Printf("\tProtocol: %s\n\tRepository id: %d\n\tEnvironment Id: %d\n\tCreated: %s\n\tUpdated: %s\n", server.Protocol, server.RepositoryId, server.EnvironmentId, server.CreatedAt, server.UpdatedAt)
		}
	}
}

func showServer() {
	var (
		server *Server
		err    error
	)

	server, err = bot.GetServer(*serverIdFlag)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("Server %d: %s\n", server.Id, server.Name)
	if *verbose {
		fmt.Printf("\tProtocol: %s\n\tRepository id: %d\n\tEnvironment Id: %d\n\tCreated: %s\n\tUpdated: %s\n", server.Protocol, server.RepositoryId, server.EnvironmentId, server.CreatedAt, server.UpdatedAt)
	}
}

func refreshRepository() {
	var (
		repository *Repository
		err        error
		id         int
	)

	if id, err = config.Alias(*repositoryIdFlag); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if repository, err = bot.GetRepository(id); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	bot.Refresh(repository)
}

func deployEnvironment() {
	var deploySetting DeploymentSetting
	id, _ := config.Alias(*environmentIdFlag)
	deploySetting.EnvironmentId = id
	if *userIdFlag != 0 {
		deploySetting.UserId = *userIdFlag
	} else if config.User != 0 {
		deploySetting.UserId = config.User
	}
	deploySetting.DeployedVersion = *deployDeployedVersionFlag
	deploySetting.DeployFromScratch = *deployFromScratchFlag
	deploySetting.TriggerNotifications = !*deployDontTriggerNotificationsFlag
	deploySetting.Comment = *deployCommentArg

	deployment, err := bot.TriggerDeployment(deploySetting)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("Deployment %d: %s (%s), retries: %d\n", deployment.Id, deployment.DeployedVersion, deployment.State, deployment.Retries)
	if deployment.Comment != "" {
		fmt.Printf("\t%s\n", deployment.Comment)
	}
	lastState := deployment.State
	if *deployWaitForCompletionFlag {
		for {
			deployment, err = bot.GetDeployment(deployment.Id)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			fmt.Print(".")
			if lastState != deployment.State {
				fmt.Printf("\nCurrent state: %s\n", deployment.State)
				lastState = deployment.State
			}
			time.Sleep(1000 * time.Millisecond)
			if !deployment.IsRunning() {
				break
			}
		}
		fmt.Printf("\nLast state: %s on %s\n", deployment.State, deployment.DeployedAt)
	}
}

func exportAliases() {
	var (
		repositories *Repositories
		environments *Environments
		environment  *Environment
		err          error
	)
	if repositories, err = bot.GetRepositories(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if environments, err = bot.GetEnvironments(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	mappedEnvironments := make(map[int][]*Environment)
	for _, environment = range environments.Entries {
		mappedEnvironments[environment.RepositoryId] = append(mappedEnvironments[environment.RepositoryId], environment)
	}

	fmt.Printf("[aliases]\n")
	for _, repository := range repositories.Entries {
		fmt.Printf("# %s\n", repository.Title)
		fmt.Printf("%s = %d\n", slug.Slug(repository.Title), repository.Id)
		for _, environment := range mappedEnvironments[repository.Id] {
			fmt.Printf("\t%s-%s = %d\n", slug.Slug(repository.Title), slug.Slug(environment.Name), environment.Id)
		}
	}

}
