package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app                                = kingpin.New("deploybot-cli", "DeployBot command line client")
	verbose                            = app.Flag("verbose", "Verbose").Short('v').Bool()
	repositoryIdFlag                   = app.Flag("repository_id", "Repository id (applies to select commands)").Short('r').String()
	environmentIdFlag                  = app.Flag("environment_id", "Environment id (applies to select commands)").Short('e').String()
	userIdFlag                         = app.Flag("user_id", "User id (applies to select commands)").Short('u').Int()
	serverIdFlag                       = app.Flag("server_id", "Server id (applies to select commands)").Short('s').Int()
	listCommand                        = app.Command("list", "List repositories, environments, servers, users.")
	listRepositoriesCommand            = listCommand.Command("repositories", "List repositories.")
	listEnvironmentsCommand            = listCommand.Command("environments", "List environments (optionnaly filter by repository).")
	listServersCommand                 = listCommand.Command("servers", "List servers (optionnaly filter by repository or environment).")
	listUsersCommand                   = listCommand.Command("users", "List users.")
	showCommand                        = app.Command("show", "Show repository, environment, server, user.")
	showRepositoryCommand              = showCommand.Command("repository", "Show repository details (and environments if verbose is set).")
	showEnvironmentCommand             = showCommand.Command("environment", "Show environment details (and servers if verbose is set).")
	showServerCommand                  = showCommand.Command("server", "Show server details.")
	showUserCommand                    = showCommand.Command("user", "Show user details.")
	refreshCommand                     = app.Command("refresh", "Refresh a repository")
	deployCommand                      = app.Command("deploy", "Deploy an environment, user is choosen from (in order) user_id flag, User config key, account owner")
	deployFromScratchFlag              = deployCommand.Flag("from_scratch", "Deploy everything again").Short('a').Bool()
	deployDeployedVersionFlag          = deployCommand.Flag("version", "Version to deploy (default to latest)").Short('c').String()
	deployDontTriggerNotificationsFlag = deployCommand.Flag("no_trigger", "Don't trigger notifications").Short('t').Bool()
	deployWaitForCompletionFlag        = deployCommand.Flag("wait", "Wait until deploy is over").Short('w').Bool()
	deployCommentArg                   = deployCommand.Arg("comment", "Add a comment to deploy").String()
	aliasesCommand                     = app.Command("aliases", "Build a list of aliases from repositories and environments")
	dumpConfigCommand                  = app.Command("dump-config", "Dump current configuration")
	bot                                = &DeployBot{}
	config                             = &Config{}
)

func main() {
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
	case refreshCommand.FullCommand():
		refreshRepository()
	case deployCommand.FullCommand():
		deployEnvironment()
	case aliasesCommand.FullCommand():
		exportAliases()
	case dumpConfigCommand.FullCommand():
		if buf, err := config.Dump(); err == nil {
			fmt.Println(buf.String())
		} else {
			fmt.Println(err)
		}
	}
}
