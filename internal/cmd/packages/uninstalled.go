package packages

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uninstalledCmd = &cobra.Command{
	Use:   "uninstalled",
	Short: "Manage packages",
	Long:  `Manage packages installed on the system`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uninstalled")
	},
}
