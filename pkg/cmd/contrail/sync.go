package contrail

import (
	"github.com/Juniper/contrail/pkg/sync"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	Contrail.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Start Sync service",
	Run: func(cmd *cobra.Command, args []string) {
		startSync()
	},
}

func startSync() {
	config := configFile
	if syncConfigFile != "" {
		config = syncConfigFile
	}
	s, err := sync.NewServiceByFile(config)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
