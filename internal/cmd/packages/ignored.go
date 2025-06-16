package packages

import (
	"dotman/internal/bashcmd"
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"
	"log"

	"github.com/spf13/cobra"
)

var (
	saveIgnored   bool
	listIgnored   bool
	removeIgnored bool
)

var ignoredCmd = &cobra.Command{
	Use:   "ignored",
	Short: "Show ignored packages and manage them",
	Long:  `Get a list of ignored packages, select and manage them if needed`,
	Run: func(cmd *cobra.Command, args []string) {
		bcmd := bashcmd.NewBashCmd(bashcmd.NewIOWriter(bashcmd.Green))
		pm, err := manager.NewPacmanManager("temp/pacman-packages.toml", bcmd)
		if err != nil {
			log.Fatal(err)
			return
		}

		var action workflow.IgnoreAction
		switch {
		case saveIgnored:
			action = workflow.IgnoreActionSave
		case removeIgnored:
			action = workflow.IgnoreActionRemove
		default:
			action = workflow.IgnoreActionList
		}
		workflow.Ignored(pm, action)
	},
}

func init() {
	ignoredCmd.Flags().BoolVarP(&saveIgnored, "save", "s", false, "Save selected packages")
	ignoredCmd.Flags().BoolVarP(&listIgnored, "list", "l", false, "List ignored packages (default)")
	ignoredCmd.Flags().BoolVarP(&removeIgnored, "remove", "r", false, "Remove selected packages")
	ignoredCmd.MarkFlagsMutuallyExclusive("save", "list", "remove")
}
