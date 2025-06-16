package packages

import (
	"dotman/internal/bashcmd"
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"
	"log"

	"github.com/spf13/cobra"
)

var (
	ignoreSaved  bool
	listSaved    bool
	removeSaved  bool
	reorderSaved bool
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

		var action workflow.SaveAction
		switch {
		case removeSaved:
			action = workflow.SaveActionRemove
		case ignoreSaved:
			action = workflow.SaveActionIgnore
		case reorderSaved:
			action = workflow.SaveActionReorder
		default:
			action = workflow.SaveActionList
		}

		workflow.Saved(pm, action)
	},
}

func init() {
	savedCmd.Flags().BoolVarP(&ignoreSaved, "ignore", "i", false, "Ignore selected packages")
	savedCmd.Flags().BoolVarP(&listSaved, "list", "l", false, "List saved packages (default)")
	savedCmd.Flags().BoolVarP(&removeSaved, "remove", "r", false, "Remove selected packages")
	savedCmd.Flags().BoolVarP(&reorderSaved, "reorder", "o", false, "Reorder selected packages")
	savedCmd.MarkFlagsMutuallyExclusive("ignore", "list", "remove", "reorder")
}
