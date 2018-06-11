package contrail

import (
	"context"
	"sync"

	"github.com/Juniper/contrail/pkg/agent"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/compilation"
	syncp "github.com/Juniper/contrail/pkg/sync"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Contrail.AddCommand(processCmd)
}

var processCmd = &cobra.Command{
	Use:   "run",
	Short: "Start Contrail Processes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}

		if viper.GetBool("server.enabled") {
			server, err := apisrv.NewServer()
			if err != nil {
				log.Fatal(err)
			}
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
		}
		if viper.GetBool("agent.enabled") {
			wg.Add(1)
			go func() {
				startAgent()
				wg.Done()
			}()
		}
		if viper.GetBool("sync.enabled") {
			wg.Add(1)
			go func() {
				startSync()
				wg.Done()
			}()
		}
		if viper.GetBool("compilation.enabled") {
			wg.Add(1)
			go func() {
				startCompilationService()
				wg.Done()
			}()
		}
		wg.Wait()
	},
}

func startSync() {
	s, err := syncp.NewService()
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func startCompilationService() {
	server, err := compilation.NewIntentCompilationService()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = server.Init(ctx); err != nil {
		log.Fatal(err)
	}

	if err = server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func startAgent() {
	a, err := agent.NewAgentByConfig()
	if err != nil {
		log.Fatal(err)
	}
	for {
		if err := a.Watch(); err != nil {
			log.Error(err)
		}
	}
}
