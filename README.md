# DeployBot cli

This is a crude cli for [DeployBot](http://deploybot.com/)'s API written in GoLang. It exposes all methods in a somewhat
useful way and permit triggering resfresh and deploy from the command line.

## Installation

Until I provide binaries, you can do a `go get github.com/epicagency/deploybot-cli` and that should set everything up
providing you have a working Go environment.

## Usage

Before anything you have to generate a token in DeployBot's account settings and create a configuration file named
`.deploybot.toml` containing 2 keys: the token and your domain.

Example:

```toml
Token = "0000000000000000000000000000000000000000"
Domain = "my-domain"
```

Once done you can just run `deploybot-cli` for a list of commands and parameters.

```
usage: deploybot-cli [<flags>] <command> [<args> ...]

DeployBot command line client

Flags:
  --help         Show context-sensitive help (also try --help-long and
                 --help-man).
  -v, --verbose  Verbose
  -r, --repository_id=REPOSITORY_ID  
                 Repository id (applies to select commands)
  -e, --environment_id=ENVIRONMENT_ID  
                 Environment id (applies to select commands)
  -u, --user_id=USER_ID  
                 User id (applies to select commands)
  -s, --server_id=SERVER_ID  
                 Server id (applies to select commands)

Commands:
  help [<command>...]
    Show help.

  list repositories
    List repositories.

  list environments
    List environments (optionnaly filter by repository).

  list servers
    List servers (optionnaly filter by repository or environment).

  list users
    List users.

  show repository
    Show repository details (and environments if verbose is set).

  show environment
    Show environment details (and servers if verbose is set).

  show server
    Show server details.

  show user
    Show user details.
```

## Contributions

Most welcome! :) Just fork and PR away!

## Authors

* Hugues Lismonde [@hlidotbe](https://twitter.com/hlidotbe), [hlidotbe](https://github.com/hlidotbe),
  [epic.net](http://epic.net)

## License

MIT: See LICENSE
