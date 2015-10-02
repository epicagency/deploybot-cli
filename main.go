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
	userIdFlag              = app.Flag("user_id", "User id (applies to select commands)").Short('u').Int()
	serverIdFlag            = app.Flag("server_id", "Server id (applies to select commands)").Short('s').Int()
	listCommand             = app.Command("list", "List repositories, environments, servers, users.")
	listRepositoriesCommand = listCommand.Command("repositories", "List repositories.")
	listEnvironmentsCommand = listCommand.Command("environments", "List environments (optionnaly filter by repository).")
	listServersCommand      = listCommand.Command("servers", "List servers (optionnaly filter by repository or environment).")
	listUsersCommand        = listCommand.Command("users", "List users.")
	showCommand             = app.Command("show", "Show repository, environment, server, user.")
	showRepositoryCommand   = showCommand.Command("repository", "Show repository details (and environments if verbose is set).")
	showEnvironmentCommand  = showCommand.Command("environment", "Show environment details (and servers if verbose is set).")
	showServerCommand       = showCommand.Command("server", "Show server details.")
	showUserCommand         = showCommand.Command("user", "Show user details.")
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
	case showRepositoryCommand.FullCommand():
		showRepository()
	case showEnvironmentCommand.FullCommand():
		showEnvironment()
	case showServerCommand.FullCommand():
		showServer()
	case showUserCommand.FullCommand():
		showUser()
	}
}
