package main

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Token   string
	Domain  string
	User    int
	Aliases map[string]int
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

func (c *Config) Alias(v string) (int, error) {
	if c.Aliases[v] != 0 {
		return c.Aliases[v], nil
	} else {
		return strconv.Atoi(v)
	}
}
