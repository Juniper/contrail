package contrail

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/agent"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/compilation"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/db/cassandra"
	"github.com/Juniper/contrail/pkg/db/etcd"
	syncp "github.com/Juniper/contrail/pkg/sync"
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
		var cacheDB *cache.DB
		ctx := context.Background()
		if viper.GetBool("cache.enabled") {
			log.Debug("cache service enabled")
			cacheDB = cache.New()
			if viper.GetBool("cache.cassandra.enabled") {
				log.Debug("cassandra watcher enabled for cache")
				processor := cassandra.NewEventProducer(
					cacheDB,
					"go_api",
					viper.GetString("cache.cassandra.host"),
					viper.GetString("cache.cassandra.amqp"),
				)
				wg.Add(1)
				go func() {
					defer wg.Done()
					err := processor.Start(ctx)
					if err != nil {
						log.Warn(err)
					}
				}()
			}
			if viper.GetBool("cache.etcd.enabled") {
				log.Debug("etcd watcher enabled for cache")
				processor, err := etcd.NewEventProducer(cacheDB)
				if err != nil {
					log.Fatal(err)
				}
				wg.Add(1)
				go func() {
					defer wg.Done()
					err := processor.Start(ctx)
					if err != nil {
						log.Warn(err)
					}
				}()
			}
		}

		if viper.GetBool("server.enabled") {
			server, err := apisrv.NewServer()
			if err != nil {
				log.Fatal(err)
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				server.Cache = cacheDB
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

	if err = server.Run(); err != nil {
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
