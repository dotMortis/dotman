package packages

import (
	"dotman/internal/bashcmd"
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var missingCmd = &cobra.Command{
	Use:   "missing",
	Short: "Show missing packages and install them",
	Long:  `Get a list of missing packages, select and mannage them if needed`,
	Run: func(cmd *cobra.Command, args []string) {
		bcmd := bashcmd.NewBashCmd(bashcmd.NewIOWriter(bashcmd.Green))
		pm, err := manager.NewPacmanManager("temp/pacman-packages.toml", bcmd)
		if err != nil {
			log.Fatal(err)
			return
		}
		workflow.Missing(pm)
		fmt.Println("Done...")
		fmt.Println("See you, UwU")
	},
}
