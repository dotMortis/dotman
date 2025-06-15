package main

import (
	"dotman/internal/bashcmd"
	"dotman/internal/bashcmd/writer"
	"dotman/internal/cmd"
	"dotman/internal/config"
	"dotman/internal/manager/pacman"
	"fmt"
	"os"
)

func main() {

	bcmd := bashcmd.NewBashCmd(writer.NewIOWriter())

	pm, err := pacman.NewPacmanManager("temp/pacman-packages.toml", bcmd)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	installed, _ := pm.Packages.Installed()

	fmt.Println("Installed: ", installed)
	fmt.Println("Saved: ", pm.Packages.Saved())
	fmt.Println("Surplus: ", pm.Packages.Surplus)
	fmt.Println("Uninstalled: ", pm.Packages.Uninstalled)
	os.Exit(0)

	config, err := config.Config()
	if err != nil {
		fmt.Println(fmt.Errorf("error initializing config: %v", err))
		os.Exit(1)
	}
	fmt.Printf("VALUES: %+v\n", config)
	os.Exit(0)
	cmd.Execute()
}
