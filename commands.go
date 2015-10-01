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
		fmt.Printf("%d: %s (%s)\n", repository.Id, repository.Title, repository.Name)
		if *verbose {
			fmt.Printf("\tType: %s\n\tRefresh: %s\n\tColor label: %s\n\tCreated: %s\n\tUpdated: %s\n", repository.Type, repository.RefreshWebhookUrl, repository.ColorLabel, repository.CreatedAt, repository.UpdatedAt)
		}
	}
}

func listEnvironments() {
	var (
		environments *Environments
		err          error
	)
	if *listEnvironmentRepositoryId != 0 {
		environments, err = bot.GetEnvironmentsByRepository(*listEnvironmentRepositoryId)
	} else {
		environments, err = bot.GetEnvironments()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, environment := range environments.Entries {
		fmt.Printf("%d: %s\n", environment.Id, environment.Name)
		if *verbose {
			fmt.Printf("\tCurrent version: %s\n\tIs automatic?: %t\n\tRepository Id: %d\n\tBranch name: %s\n\tDeploy webhook url: %s\n\tCreated: %s\n\tUpdated: %s\n", environment.CurrentVersion, environment.IsAutomatic, environment.RepositoryId, environment.BranchName, environment.DeployWebhookUrl, environment.CreatedAt, environment.UpdatedAt)
		}
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
		fmt.Printf("%d: %s %s (%s)\n", user.Id, user.FirstName, user.LastName, user.Email)
		if *verbose {
			fmt.Printf("\tTimezone: %s\n\tIs admin?: %t\n\tCreated at: %s\n\tUpdated at: %s\n", user.Timezone, user.IsAdmin, user.CreatedAt, user.UpdatedAt)
		}
	}
}
