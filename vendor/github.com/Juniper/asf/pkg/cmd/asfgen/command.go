package asfgen

import (
	"github.com/spf13/cobra"
)

// Command defines ASF Generate utility command.
var Command = &cobra.Command{
	Use:   "asfgen",
	Short: "ASF Generate Command",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute executes the ASF Generate command.
func Execute() error {
	return Command.Execute()
}
