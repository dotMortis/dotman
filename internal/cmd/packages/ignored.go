package packages

import (
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"

	"github.com/spf13/cobra"
)

func NewIgnoredCmd(packageManager manager.Manager) *cobra.Command {
	var (
		save   bool
		list   bool
		remove bool
	)
	var cmd = &cobra.Command{
		Use:   "ignored",
		Short: "Show ignored packages and manage them",
		Long:  `Get a list of ignored packages, select and manage them if needed`,
		Run: func(cmd *cobra.Command, args []string) {
			var action workflow.IgnoreAction
			switch {
			case save:
				action = workflow.IgnoreActionSave
			case remove:
				action = workflow.IgnoreActionRemove
			default:
				action = workflow.IgnoreActionList
			}
			workflow.Ignored(packageManager, action)
		},
	}
	cmd.Flags().BoolVarP(&save, "save", "s", false, "Save selected packages")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "List ignored packages (default)")
	cmd.Flags().BoolVarP(&remove, "remove", "r", false, "Remove selected packages")
	cmd.MarkFlagsMutuallyExclusive("save", "list", "remove")
	return cmd
}
