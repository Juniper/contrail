package contrailutil

import (
	"github.com/spf13/cobra"
)

// ContrailUtil defines root Contrail utility command.
var ContrailUtil = &cobra.Command{
	Use:   "contrailutil",
	Short: "Contrail Utility Command",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
