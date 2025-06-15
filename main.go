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

	bcmd := bashcmd.NewBashCmd(writer.NewIOWriter())
	if err := bcmd.Execute("sudo", "ls", "-al"); err != nil {
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
