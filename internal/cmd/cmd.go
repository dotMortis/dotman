package cmd

import (
	"dotman/internal/bashcmd"
	"dotman/internal/cmd/packages"
	"dotman/internal/manager"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const logo = `     _       _                         
    | |     | |                        
  __| | ___ | |_ _ __ ___   __ _ _ __  
 / _` + "`" + ` |/ _ \| __| '_ ` + "`" + ` _ \ / _` + "`" + ` | '_ \ 
| (_| | (_) | |_| | | | | | (_| | | | |
 \__,_|\___/ \__|_| |_| |_|\__,_|_| |_|`

var rootCmd = &cobra.Command{
	Use:   "dotman",
	Short: logo,
	Long: logo + `
a comprehensive environment manager

- Manage and sync your dotfiles across different machines
- Track and install system packages and applications
- Backup and restore your system configurations
- Automate environment setup with simple commands`,
}

func init() {
	bcmd := bashcmd.NewBashCmd(bashcmd.NewIOWriter(bashcmd.Green))
	pacmanPM, err := manager.NewPacmanManager("temp/pacman-packages.toml", bcmd)
	if err != nil {
		log.Fatal(err)
		return
	}
	rootCmd.AddCommand(packages.NewPackagesCmd("pac", "Pacman", pacmanPM))

	yayPM, err := manager.NewYayManager("temp/yay-packages.toml", bcmd)
	if err != nil {
		log.Fatal(err)
		return
	}
	rootCmd.AddCommand(packages.NewPackagesCmd("yay", "Yay", yayPM))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
