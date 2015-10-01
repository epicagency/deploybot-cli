package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app                     = kingpin.New("deploybot-cli", "DeployBot command line client")
	verbose                 = app.Flag("verbose", "Verbose").Short('v').Bool()
	listCommand             = app.Command("list", "List repositories, environments, servers, users.")
	listRepositoriesCommand = listCommand.Command("repositories", "List repositories.")
	listEnvironmentsCommand = listCommand.Command("environments", "List environments.")
	//listEnvironmentRepositoryId = listEnvironmentsCommand.Flag("repository_id", "Repository id to list environments from").Short('r').Required().Int()
	listServersCommand = listCommand.Command("servers", "List servers.")
	listUsersCommand   = listCommand.Command("usrs", "List users.")
)

func main() {
	config := &Config{}
	if err := config.Load(""); err != nil {
		fmt.Println("Unable to read config, please create a ~/.deploybot.toml with a token and a domain")
		fmt.Println(err)
		os.Exit(-1)
	}

	bot := &DeployBot{Config: config}

	var err error

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case listRepositoriesCommand.FullCommand():
		var repositories *Repositories
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
	case listEnvironmentsCommand.FullCommand():
		var environments *Environments
		if environments, err = bot.GetEnvironments(); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		for _, environment := range environments.Entries {
			fmt.Printf("%d: %s\n", environment.Id, environment.Name)
			if *verbose {
				fmt.Printf("\tCurrent version: %s\n\tIs automatic?: %t\n\tRepository Id: %d\n\tBranch name: %s\n\tDeploy webhook url: %s\n\tCreated: %s\n\tUpdated: %s\n", environment.CurrentVersion, environment.IsAutomatic, environment.RepositoryId, environment.BranchName, environment.DeployWebhookUrl, environment.CreatedAt, environment.UpdatedAt)
			}
		}
	case listServersCommand.FullCommand():
	case listUsersCommand.FullCommand():
	}
}
