package packages

import (
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"

	"github.com/spf13/cobra"
)

func NewSavedCmd(packageManager manager.Manager) *cobra.Command {
	var (
		ignore  bool
		list    bool
		remove  bool
		reorder bool
	)
	var cmd = &cobra.Command{
		Use:   "saved",
		Short: "Show saved packages and manage them",
		Long:  `Get a list of saved packages, select and manage them if needed`,
		Run: func(cmd *cobra.Command, args []string) {
			var action workflow.SaveAction
			switch {
			case remove:
				action = workflow.SaveActionRemove
			case ignore:
				action = workflow.SaveActionIgnore
			case reorder:
				action = workflow.SaveActionReorder
			default:
				action = workflow.SaveActionList
			}

			workflow.Saved(packageManager, action)
		},
	}
	cmd.Flags().BoolVarP(&ignore, "ignore", "i", false, "Ignore selected packages")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "List saved packages (default)")
	cmd.Flags().BoolVarP(&remove, "remove", "r", false, "Remove selected packages")
	cmd.Flags().BoolVarP(&reorder, "reorder", "o", false, "Reorder selected packages")
	cmd.MarkFlagsMutuallyExclusive("ignore", "list", "remove", "reorder")
	return cmd
}
