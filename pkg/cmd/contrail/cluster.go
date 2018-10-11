package contrail

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Juniper/contrail/pkg/cluster"
)

func init() {
	Contrail.AddCommand(clusterCmd)
}

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Start managing contrail cluster",
	Run: func(cmd *cobra.Command, args []string) {
		manageCluster()
	},
}

func manageCluster() {
	manager, err := cluster.NewClusterManager(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := manager.Manage(); err != nil {
		log.Fatal(err)
	}
}
