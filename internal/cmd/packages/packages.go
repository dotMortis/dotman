package packages

import (
	"dotman/internal/manager"
	"fmt"

	"github.com/spf13/cobra"
)

func NewPackagesCmd(mainCommand string, fullName string, packageManager manager.Manager) *cobra.Command {
	var PackagesCmd = &cobra.Command{
		Use:   mainCommand,
		Short: fmt.Sprintf("Manage %s packages", fullName),
		Long:  fmt.Sprintf("Manage %s packages installed on the system", fullName),
	}
	PackagesCmd.AddCommand(NewMissingCmd(packageManager))
	PackagesCmd.AddCommand(NewIgnoredCmd(packageManager))
	PackagesCmd.AddCommand(NewSavedCmd(packageManager))
	PackagesCmd.AddCommand(NewSurplusCmd(packageManager))
	return PackagesCmd
}
