package contrail

import (
	"github.com/spf13/cobra"

	"github.com/Juniper/asf/pkg/logutil"
	// TODO(buoto): Decouple from below packages
	//"github.com/Juniper/asf/pkg/deploy"
)

func init() {
	Contrail.AddCommand(deployerCmd)
}

var deployerCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Start managing contrail cluster",
	Run: func(cmd *cobra.Command, args []string) {
		manageCluster()
	},
}

func manageCluster() {
	manager, err := deploy.NewDeployManager(configFile)
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}

	if err := manager.Manage(); err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
