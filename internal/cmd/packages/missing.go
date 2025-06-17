package packages

import (
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"
	"fmt"

	"github.com/spf13/cobra"
)

func NewMissingCmd(packageManager manager.Manager) *cobra.Command {
	var (
		isntall      bool
		installForce bool
		ignore       bool
		list         bool
		remove       bool
	)
	var cmd = &cobra.Command{
		Use:   "missing",
		Short: "Show missing packages and manage them",
		Long:  `Get a list of missing packages, select and manage them if needed`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if installForce && !isntall {
				return fmt.Errorf("--yes can only be used with --install")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

			var action workflow.MissingAction
			switch {
			case isntall:
				if installForce {
					action = workflow.MissingActionForceInstall
				} else {
					action = workflow.MissingActionInstall
				}
			case ignore:
				action = workflow.MissingActionIgnore
			case remove:
				action = workflow.MissingActionRemove
			default:
				action = workflow.MissingActionList
			}

			workflow.Missing(packageManager, action)
		},
	}
	cmd.Flags().BoolVarP(&isntall, "install", "I", false, "Install selected packages")
	cmd.Flags().BoolVarP(&isntall, "yes", "y", false, "Auto confirm installations")
	cmd.Flags().BoolVarP(&ignore, "ignore", "i", false, "Ignore selected packages")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "List missing packages (default)")
	cmd.Flags().BoolVarP(&remove, "remove", "r", false, "Remove selected packages")
	cmd.MarkFlagsMutuallyExclusive("install", "ignore", "list", "remove")
	return cmd
}
