package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app                         = kingpin.New("deploybot-cli", "DeployBot command line client")
	verbose                     = app.Flag("verbose", "Verbose").Short('v').Bool()
	listCommand                 = app.Command("list", "List repositories, environments, servers, users.")
	listRepositoriesCommand     = listCommand.Command("repositories", "List repositories.")
	listEnvironmentsCommand     = listCommand.Command("environments", "List environments.")
	listEnvironmentRepositoryId = listEnvironmentsCommand.Flag("repository_id", "Repository id to list environments from").Short('r').Int()
	listServersCommand          = listCommand.Command("servers", "List servers.")
	listUsersCommand            = listCommand.Command("users", "List users.")
	bot                         = &DeployBot{}
)

func main() {
	config := &Config{}
	if err := config.Load(""); err != nil {
		fmt.Println("Unable to read config, please create a ~/.deploybot.toml with a token and a domain")
		fmt.Println(err)
		os.Exit(-1)
	}
	bot.Config = config

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case listRepositoriesCommand.FullCommand():
		listRepositories()
	case listEnvironmentsCommand.FullCommand():
		listEnvironments()
	case listServersCommand.FullCommand():
	case listUsersCommand.FullCommand():
		listUsers()
	}
}
