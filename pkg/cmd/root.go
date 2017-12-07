package cmd

import (
	"github.com/spf13/cobra"
)

var ContrailUtilCmd = &cobra.Command{
	Use:   "contrail_util",
	Short: "Contrail Utility Command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
