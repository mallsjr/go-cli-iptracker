package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "iptracker",
		Short: "A CLI tool for tracking IP addresses",
		Long:  "A CLI tool for tracking IP addresses",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}