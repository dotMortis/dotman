package packages

import (
	"dotman/internal/bashcmd"
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var savedCmd = &cobra.Command{
	Use:   "saved",
	Short: "Show saved packages and manage them",
	Long:  `Get a list of saved packages, select and manage them if needed`,
	Run: func(cmd *cobra.Command, args []string) {
		bcmd := bashcmd.NewBashCmd(bashcmd.NewIOWriter(bashcmd.Green))
		pm, err := manager.NewPacmanManager("temp/pacman-packages.toml", bcmd)
		if err != nil {
			log.Fatal(err)
			return
		}

		workflow.Saved(pm)
		fmt.Println("Done...")
		fmt.Println("See you, UwU")
	},
}
