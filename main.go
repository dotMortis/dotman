package main

import (
	"dotman/bashcmd"
	"dotman/bashcmd/writer"
	"dotman/cmd"
	"dotman/config"
	"dotman/metafile/pacman"
	"fmt"
	"os"
)

func main() {

	ppm, err := pacman.NewPacmanPackages("temp/pacman-packages.toml")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("CONTENT: ", fmt.Sprintf("%+v", ppm.Content()))
	ppm.Content().Packages.Add("first", "second")
	fmt.Println("CONTENT: ", fmt.Sprintf("%+v", ppm.Content()))
	if err := ppm.Save(); err != nil {
		fmt.Println("Error: ", fmt.Errorf("error saving pacman packages: %w", err))
		os.Exit(1)
	}
	ppm.Content().Packages.Remove("first")
	fmt.Println("CONTENT: ", fmt.Sprintf("%+v", ppm.Content()))
	if err := ppm.Save(); err != nil {
		fmt.Println("Error: ", fmt.Errorf("error saving pacman packages: %w", err))
		os.Exit(1)
	}
	os.Exit(0)

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
