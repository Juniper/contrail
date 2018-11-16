package contrail

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Juniper/contrail/pkg/command"
)

func init() {
	Contrail.AddCommand(commanderCmd)
}

var commanderCmd = &cobra.Command{
	Use:   "command",
	Short: "Start managing contrail cluster",
	Run: func(cmd *cobra.Command, args []string) {
		manageCluster()
	},
}

func manageCluster() {
	manager, err := command.NewCommandManager(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := manager.Manage(); err != nil {
		log.Fatal(err)
	}
}
