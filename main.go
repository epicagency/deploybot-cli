package main

import (
	"fmt"
	"os"
)

func main() {
	config := &Config{}
	if err := config.Load(""); err != nil {
		fmt.Println("Unable to read config")
		fmt.Println(err)
		os.Exit(-1)
	}

	//bot := &DeployBot{Config: config}

	fmt.Printf("Token: %s, domain: %s\n", config.Token, config.Domain)
}
