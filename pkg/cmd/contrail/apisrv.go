package contrail

import (
	"sync"

	"github.com/Juniper/contrail/pkg/apisrv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	Contrail.AddCommand(apiServerCmd)
}

var apiServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start API Server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := apisrv.NewServer()
		if err != nil {
			log.Fatal(err)
		}
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err = server.Init(); err != nil {
				log.Fatal(err)
			}

			if err = server.Run(); err != nil {
				log.Fatal(err)
			}
		}()
		if agentConfigFile != "" {
			wg.Add(1)
			go func() {
				startAgent()
				wg.Done()
			}()
		}
		if watcherConfigFile != "" {
			wg.Add(1)
			go func() {
				startWatcher()
				wg.Done()
			}()
		}
		wg.Wait()
	},
}
