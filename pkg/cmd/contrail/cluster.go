package contrail

import (
	"github.com/Juniper/contrail/pkg/cluster"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clusterId string

func init() {
	Contrail.AddCommand(clusterCmd)
	clusterCmd.Flags().StringVarP(&clusterId, "clusterid", "c", "", "uuid of the cluster to be built.")
}

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Start building contrail cluster",
	Run: func(cmd *cobra.Command, args []string) {
		buildCluster()
	},
}

func buildCluster() {
	builder, err := cluster.NewClusterBuilder(configFile, clusterId)
	if err != nil {
		log.Fatal(err)
	}

	if err := builder.Build(); err != nil {
		log.Fatal(err)
	}
}
