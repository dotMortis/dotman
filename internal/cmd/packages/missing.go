package packages

import (
	"dotman/internal/bashcmd"
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"
	"log"

	"github.com/spf13/cobra"
)

var (
	installMissing bool
	ignoreMissing  bool
	listMissing    bool
	removeMissing  bool
)

var missingCmd = &cobra.Command{
	Use:   "missing",
	Short: "Show missing packages and manage them",
	Long:  `Get a list of missing packages, select and manage them if needed`,
	Run: func(cmd *cobra.Command, args []string) {
		bcmd := bashcmd.NewBashCmd(bashcmd.NewIOWriter(bashcmd.Green))
		pm, err := manager.NewPacmanManager("temp/pacman-packages.toml", bcmd)
		if err != nil {
			log.Fatal(err)
			return
		}

		var action workflow.MissingAction
		switch {
		case installMissing:
			action = workflow.MissingActionInstall
		case ignoreMissing:
			action = workflow.MissingActionIgnore
		case removeMissing:
			action = workflow.MissingActionRemove
		default:
			action = workflow.MissingActionList
		}

		workflow.Missing(pm, action)
	},
}

func init() {
	missingCmd.Flags().BoolVarP(&installMissing, "install", "I", false, "Install selected packages")
	missingCmd.Flags().BoolVarP(&ignoreMissing, "ignore", "i", false, "Ignore selected packages")
	missingCmd.Flags().BoolVarP(&listMissing, "list", "l", false, "List missing packages (default)")
	missingCmd.Flags().BoolVarP(&removeMissing, "remove", "r", false, "Remove selected packages")
	missingCmd.MarkFlagsMutuallyExclusive("install", "ignore", "list", "remove")
}
