package main

import (
	"fmt"
	"os"
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
	)
	if repository, err = bot.GetRepository(*repositoryIdFlag); err != nil {
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
	)
	if *repositoryIdFlag != 0 {
		environments, err = bot.GetEnvironmentsByRepository(*repositoryIdFlag)
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
	environment, err = bot.GetEnvironment(*environmentIdFlag)

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
	)
	if *environmentIdFlag != 0 {
		servers, err = bot.GetServersByEnvironment(*environmentIdFlag)
	} else if *repositoryIdFlag != 0 {
		servers, err = bot.GetServersByRepository(*repositoryIdFlag)
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
	)

	repository, err = bot.GetRepository(*repositoryIdFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	bot.Refresh(repository)
}
