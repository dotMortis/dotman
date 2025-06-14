package main

import (
	"dotman/cmd"
	"dotman/config"
	"fmt"
	"os"
)

func main() {
	config, err := config.BaseConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("error initializing config: %v", err))
		os.Exit(1)
	}
	fmt.Printf("VALUES: %+v\n", config.Values.Giturl.Value())
	os.Exit(1)
	cmd.Execute()
}
