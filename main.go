package main

import (
	"dotman/cmd"
	"dotman/config"
	"fmt"
)

func main() {
	_, err := config.BaseConfig()
	if err != nil {
		panic(fmt.Errorf("error initializing config: %v", err))
	}
	cmd.Execute()
}
