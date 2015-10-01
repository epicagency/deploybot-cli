package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app                     = kingpin.New("deploybot-cli", "DeployBot command line client")
	verbose                 = app.Flag("verbose", "Verbose").Short('v').Bool()
	repositoryIdFlag        = app.Flag("repository_id", "Repository id (applies to select commands)").Short('r').Int()
	environmentIdFlag       = app.Flag("environment_id", "Environment id (applies to select commands)").Short('e').Int()
	listCommand             = app.Command("list", "List repositories, environments, servers, users.")
	listRepositoriesCommand = listCommand.Command("repositories", "List repositories.")
	listEnvironmentsCommand = listCommand.Command("environments", "List environments (optionnaly filter by repository).")
	listServersCommand      = listCommand.Command("servers", "List servers (optionnaly filter by repository or environment).")
	listUsersCommand        = listCommand.Command("users", "List users.")
	bot                     = &DeployBot{}
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
		listServers()
	case listUsersCommand.FullCommand():
		listUsers()
	}
}
