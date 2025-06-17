package packages

import (
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"
	"fmt"

	"github.com/spf13/cobra"
)

func NewSurplusCmd(packageManager manager.Manager) *cobra.Command {
	var (
		save   bool
		ignore bool
		list   bool
		force  bool
	)
	var cmd = &cobra.Command{
		Use:   "surplus",
		Short: "Show surplus packages and manage them",
		Long:  `Get a list of surplus packages, select and manage them if needed`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if force && !ignore {
				return fmt.Errorf("--force can only be used with --ignore")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			var action workflow.SurplusAction
			switch {
			case save:
				action = workflow.SurplusActionSave
			case ignore:
				if force {
					action = workflow.SurplusActionForceIgnore
				} else {
					action = workflow.SurplusActionIgnore
				}
			default:
				action = workflow.SurplusActionList
			}

			workflow.Surplus(packageManager, action)
		},
	}
	cmd.Flags().BoolVarP(&save, "save", "s", false, "Save selected packages")
	cmd.Flags().BoolVarP(&ignore, "ignore", "i", false, "Ignore selected packages")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "List surplus packages (default)")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force ignore without confirmation")
	cmd.MarkFlagsMutuallyExclusive("save", "ignore", "list")
	return cmd
}
