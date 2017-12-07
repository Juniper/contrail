package contrailutil

import (
	"github.com/spf13/cobra"
)

//Cmd for utility command
var Cmd = &cobra.Command{
	Use:   "contrail_util",
	Short: "Contrail Utility Command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
