package main

import (
	"dotman/bashcmd"
	"dotman/bashcmd/writer"
	"dotman/cmd"
	"dotman/config"
	"fmt"
	"os"
)

func main() {

	os.Setenv("TERM", "xterm-256color")

	bcmd := bashcmd.NewBashCmd(writer.NewIOWriter())
	if err := bcmd.Execute("sudo", "apt", "remove", "htop"); err != nil {
		fmt.Println("Error: ", err)
	}
	os.Exit(0)

	config, err := config.BaseConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("error initializing config: %v", err))
		os.Exit(1)
	}
	fmt.Printf("VALUES: %+v\n", config)
	os.Exit(0)
	cmd.Execute()
}
