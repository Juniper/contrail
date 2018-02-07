package contrail

import (
	"github.com/Juniper/contrail/pkg/watcher"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	Contrail.AddCommand(watcherCmd)
}

var watcherCmd = &cobra.Command{
	Use:   "watcher",
	Short: "Start Watcher service",
	Run: func(cmd *cobra.Command, args []string) {
		startWatcher()
	},
}

func startWatcher() {
	s, err := watcher.NewServiceByFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
