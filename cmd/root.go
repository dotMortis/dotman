package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dotman",
	Short: "Dotman is an environment manager",
	Long: `Dotman is a comprehensive environment manager that helps you:
- Manage and sync your dotfiles across different machines
- Track and install system packages and applications
- Backup and restore your system configurations
- Automate environment setup with simple commands

Perfect for developers who want to maintain consistent environments across multiple systems.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
