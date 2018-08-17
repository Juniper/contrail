package contrail

import (
	"context"
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
	MaybeStart("replication", startCassandraReplicator, wg)
	MaybeStart("replication", startAmqpReplicator, wg)
	MaybeStart("cache", startCacheService, wg)
	MaybeStart("server", startServer, wg)
	MaybeStart("agent", startAgent, wg)
	MaybeStart("sync", startSync, wg)
	MaybeStart("compilation", startCompilationService, wg)
}

//MaybeStart runs process if it is enabled.
func MaybeStart(serviceName string, f func(wg *sync.WaitGroup), wg *sync.WaitGroup) {
	if !viper.GetBool(serviceName + ".enabled") {
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		f(wg)
	}()
}

func startCassandraReplicator(wg *sync.WaitGroup) {
	ctx := context.Background()
	log.Debug("cassandra replication service enabled")
	cassandraProcessor := cassandra.NewEventProcessor()
	producer, err := etcd.NewEventProducer(cassandraProcessor)
	if err != nil {
		log.Fatal(err)
	}
	err = producer.Start(ctx)
	if err != nil {
		log.Warn(err)
	}
}

func startAmqpReplicator(wg *sync.WaitGroup) {
	ctx := context.Background()
	log.Debug("amqp replication service enabled")
	amqpProcessor := cassandra.NewAmqpEventProcessor()
	producer, err := etcd.NewEventProducer(amqpProcessor)
	if err != nil {
		log.Fatal(err)
	}
	err = producer.Start(ctx)
	if err != nil {
		log.Warn(err)
	}
}

func startCacheService(wg *sync.WaitGroup) {
	log.Debug("cache service enabled")
	cacheDB = cache.NewDB(uint64(viper.GetInt64("cache.max_history")))
	MaybeStart("cache.cassandra", startCassandraWatcher, wg)
	MaybeStart("cache.etcd", startEtcdWatcher, wg)
	MaybeStart("cache.rdbms", startRDBMSWatcher, wg)
}

func startCassandraWatcher(wg *sync.WaitGroup) {
	ctx := context.Background()
	log.Debug("cassandra watcher enabled for cache")
	producer := cassandra.NewEventProducer(cacheDB)
	err := producer.Start(ctx)
	if err != nil {
		log.Warn(err)
	}
}

func startEtcdWatcher(wg *sync.WaitGroup) {
	ctx := context.Background()
	log.Debug("etcd watcher enabled for cache")
	producer, err := etcd.NewEventProducer(cacheDB)
	if err != nil {
		log.Fatal(err)
	}
	err = producer.Start(ctx)
	if err != nil {
		log.Warn(err)
	}
}

func startRDBMSWatcher(wg *sync.WaitGroup) {
	ctx := context.Background()
	log.Debug("rdbms watcher enabled for cache")
	producer, err := syncp.NewEventProducer(cacheDB)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()
	err = producer.Start(ctx)
	if err != nil {
		log.Warn(err)
	}
}

func startServer(wg *sync.WaitGroup) {
	server, err := apisrv.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Cache = cacheDB
	if err = server.Init(); err != nil {
		log.Fatal(err)
	}
	if err = server.Run(); err != nil {
		log.Warn(err)
	}
}

func startSync(wg *sync.WaitGroup) {
	s, err := syncp.NewService()
	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()
	err = s.Run()
	if err != nil {
		log.Warn(err)
	}
}

func startCompilationService(wg *sync.WaitGroup) {
	server, err := compilation.NewIntentCompilationService()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = server.Run(ctx); err != nil {
		log.Warn(err)
	}
}

func startAgent(wg *sync.WaitGroup) {
	a, err := agent.NewAgentByConfig()
	if err != nil {
		log.Fatal(err)
	}
	for {
		if err := a.Watch(context.Background()); err != nil {
			log.Warn(err)
		}
	}
}
