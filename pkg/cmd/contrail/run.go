package contrail

import (
	"context"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/agent"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/compilation"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/db/cassandra"
	"github.com/Juniper/contrail/pkg/db/etcd"
	syncp "github.com/Juniper/contrail/pkg/sync"
)

var cacheDB *cache.DB

func init() {
	Contrail.AddCommand(processCmd)
}

func getQueueName() string {
	name, _ := os.Hostname() // nolint: noerror
	return "contrail_process_" + name
}

var processCmd = &cobra.Command{
	Use:   "run",
	Short: "Start Contrail Processes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}
		StartProcesses(wg)
		wg.Wait()
	},
}

//StartProcesses starts processes based on config.
func StartProcesses(wg *sync.WaitGroup) {
	startCacheService(wg)
	startServer(wg)
	startAgent(wg)
	startSync(wg)
	startCompilationService(wg)
}

func startCacheService(wg *sync.WaitGroup) {
	if !viper.GetBool("cache.enabled") {
		return
	}
	log.Debug("cache service enabled")
	cacheDB = cache.New(uint64(viper.GetInt64("cache.max_history")))
	startCassandraWatcher(wg)
	startEtcdWatcher(wg)
	startRDBMSWatcher(wg)
}

func startCassandraWatcher(wg *sync.WaitGroup) {
	ctx := context.Background()
	if !viper.GetBool("cache.cassandra.enabled") {
		return
	}
	log.Debug("cassandra watcher enabled for cache")
	processor := cassandra.NewEventProducer(
		cacheDB,
		getQueueName(),
		viper.GetString("cache.cassandra.host"),
		viper.GetInt("cache.cassandra.port"),
		viper.GetDuration("cache.cassandra.timeout"),
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

func startEtcdWatcher(wg *sync.WaitGroup) {
	ctx := context.Background()
	if !viper.GetBool("cache.etcd.enabled") {
		return
	}
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

func startRDBMSWatcher(wg *sync.WaitGroup) {
	ctx := context.Background()
	if !viper.GetBool("cache.rdbms.enabled") {
		return
	}
	log.Debug("rdbms watcher enabled for cache")
	processor, err := syncp.NewEventProducer(cacheDB)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go func() {
		defer processor.Close()
		defer wg.Done()
		err := processor.Start(ctx)
		if err != nil {
			log.Warn(err)
		}
	}()
}

func startServer(wg *sync.WaitGroup) {
	if !viper.GetBool("server.enabled") {
		return
	}
	server, err := apisrv.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Cache = cacheDB
	if err = server.Init(); err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err = server.Run(); err != nil {
			log.Fatal(err)
		}
	}()
}

func startSync(wg *sync.WaitGroup) {
	if !viper.GetBool("sync.enabled") {
		return
	}
	wg.Add(1)

	s, err := syncp.NewService()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer s.Close()
		defer wg.Done()
		err = s.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func startCompilationService(wg *sync.WaitGroup) {
	if !viper.GetBool("compilation.enabled") {
		return
	}

	server, err := compilation.NewIntentCompilationService()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = server.Run(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

func startAgent(wg *sync.WaitGroup) {
	if !viper.GetBool("agent.enabled") {
		return
	}
	wg.Add(1)
	a, err := agent.NewAgentByConfig()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer wg.Done()
		for {
			if err := a.Watch(); err != nil {
				log.Fatal(err)
			}
		}
	}()
}
