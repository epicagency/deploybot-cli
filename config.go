package main

import (
	"github.com/BurntSushi/toml"
	"os/user"
	"path/filepath"
)

type Config struct {
	Token  string
	Domain string
}

func (c *Config) Load(path string) error {
	if path == "" {
		usr, _ := user.Current()
		path = filepath.Join(usr.HomeDir, ".deploybot.toml")
	}

	if _, err := toml.DecodeFile(path, c); err != nil {
		return err
	}

	return nil
}
