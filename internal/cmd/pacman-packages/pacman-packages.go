package pacmanPackages

import (
	"github.com/spf13/cobra"
)

var PackagesCmd = &cobra.Command{
	Use:   "pac",
	Short: "Manage pacman packages",
	Long:  `Manage pacman packages installed on the system`,
}

func init() {
	PackagesCmd.AddCommand(missingCmd)
	PackagesCmd.AddCommand(ignoredCmd)
	PackagesCmd.AddCommand(savedCmd)
	PackagesCmd.AddCommand(surplusCmd)
}
