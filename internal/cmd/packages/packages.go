package packages

import (
	"github.com/spf13/cobra"
)

var PackagesCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Manage packages",
	Long:  `Manage packages installed on the system`,
}

func init() {
	PackagesCmd.AddCommand(uninstalledCmd)
}
