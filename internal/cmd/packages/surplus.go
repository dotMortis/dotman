package packages

import (
	"dotman/internal/bashcmd"
	"dotman/internal/cmd/packages/workflow"
	"dotman/internal/manager"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	saveSurplus   bool
	ignoreSurplus bool
	listSurplus   bool
	forceSurplus  bool
)

var surplusCmd = &cobra.Command{
	Use:   "surplus",
	Short: "Show surplus packages and manage them",
	Long:  `Get a list of surplus packages, select and manage them if needed`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if forceSurplus && !ignoreSurplus {
			return fmt.Errorf("--force can only be used with --ignore")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		bcmd := bashcmd.NewBashCmd(bashcmd.NewIOWriter(bashcmd.Green))
		pm, err := manager.NewPacmanManager("temp/pacman-packages.toml", bcmd)
		if err != nil {
			log.Fatal(err)
			return
		}

		var action workflow.SurplusAction
		switch {
		case saveSurplus:
			action = workflow.SurplusActionSave
		case ignoreSurplus:
			if forceSurplus {
				action = workflow.SurplusActionForceIgnore
			} else {
				action = workflow.SurplusActionIgnore
			}
		default:
			action = workflow.SurplusActionList
		}

		workflow.Surplus(pm, action)
	},
}

func init() {
	surplusCmd.Flags().BoolVarP(&saveSurplus, "save", "s", false, "Save selected packages")
	surplusCmd.Flags().BoolVarP(&ignoreSurplus, "ignore", "i", false, "Ignore selected packages")
	surplusCmd.Flags().BoolVarP(&listSurplus, "list", "l", false, "List surplus packages (default)")
	surplusCmd.Flags().BoolVarP(&forceSurplus, "force", "f", false, "Force ignore without confirmation")
	surplusCmd.MarkFlagsMutuallyExclusive("save", "ignore", "list")
}
