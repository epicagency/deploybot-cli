package main

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

type Config struct {
	Token  string
	Domain string
	User   int
}

func (c *Config) Load(path string) error {
	if path == "" {
		path = filepath.Join(os.Getenv("HOME"), ".deploybot.toml")
	}

	if _, err := toml.DecodeFile(path, c); err != nil {
		return err
	}

	return nil
}
