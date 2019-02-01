package contrail

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/agent"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/compilation"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/db/cassandra"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/errutil"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/retry"
	syncp "github.com/Juniper/contrail/pkg/sync"
)

const (
	syncRetryInterval = 3 * time.Second
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
	MaybeStart("replication.cassandra", startCassandraReplicator, wg)
	MaybeStart("replication.amqp", startAmqpReplicator, wg)
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

func startCassandraReplicator(_ *sync.WaitGroup) {
	log.Debug("Cassandra replication service enabled")
	cassandraProcessor := cassandra.NewEventProcessor()
	producer, err := etcd.NewEventProducer(cassandraProcessor, "cassandra-replicator")
	if err != nil {
		pkglog.FatalWithStackTrace(err)
	}
	err = producer.Start(context.Background())
	if err != nil {
		log.Warn(err)
	}
}

func startAmqpReplicator(_ *sync.WaitGroup) {
	log.Debug("AMQP replication service enabled")
	amqpProcessor := cassandra.NewAmqpEventProcessor()
	producer, err := etcd.NewEventProducer(amqpProcessor, "amqp-replicator")
	if err != nil {
		pkglog.FatalWithStackTrace(err)
	}
	err = producer.Start(context.Background())
	if err != nil {
		log.Warn(err)
	}
}

func startCacheService(wg *sync.WaitGroup) {
	log.Debug("Cache service enabled")
	cacheDB = cache.NewDB(uint64(viper.GetInt64("cache.max_history")))
	MaybeStart("cache.cassandra", startCassandraWatcher, wg)
	MaybeStart("cache.etcd", startEtcdWatcher, wg)
	MaybeStart("cache.rdbms", startRDBMSWatcher, wg)
}

func startCassandraWatcher(_ *sync.WaitGroup) {
	log.Debug("Cassandra watcher enabled for cache")
	producer := cassandra.NewEventProducer(cacheDB)
	err := producer.Start(context.Background())
	if err != nil {
		log.Warn(err)
	}
}

func startEtcdWatcher(_ *sync.WaitGroup) {
	log.Debug("etcd watcher enabled for cache")
	producer, err := etcd.NewEventProducer(cacheDB, "cache-service")
	if err != nil {
		pkglog.FatalWithStackTrace(err)
	}
	err = producer.Start(context.Background())
	if err != nil {
		log.Warn(err)
	}
}

func startRDBMSWatcher(_ *sync.WaitGroup) {
	log.Debug("RDBMS watcher enabled for cache")
	producer, err := syncp.NewEventProducer(cacheDB)
	if err != nil {
		pkglog.FatalWithStackTrace(err)
	}
	defer producer.Close()
	err = producer.Start(context.Background())
	if err != nil {
		log.Warn(err)
	}
}

func startServer(_ *sync.WaitGroup) {
	server, err := apisrv.NewServer()
	if err != nil {
		pkglog.FatalWithStackTrace(err)
	}
	server.Cache = cacheDB
	if err = server.Init(); err != nil {
		pkglog.FatalWithStackTrace(err)
	}
	if err = server.Run(); err != nil {
		pkglog.FatalWithStackTrace(err)
	}
}

func startSync(_ *sync.WaitGroup) {
	if err := retry.Do(func() (retry bool, err error) {
		s, err := syncp.NewService()
		if err != nil {
			pkglog.FatalWithStackTrace(err)
		}
		defer s.Close()

		err = s.Run()

		return errutil.ShouldRetry(err), err
	}, retry.WithLog(log.StandardLogger()), retry.WithInterval(syncRetryInterval)); err != nil {
		log.Warn(err)
	}
}

func startCompilationService(_ *sync.WaitGroup) {
	server, err := compilation.NewIntentCompilationService()
	if err != nil {
		pkglog.FatalWithStackTrace(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = server.Run(ctx); err != nil {
		log.Warn(err)
	}
}

func startAgent(_ *sync.WaitGroup) {
	a, err := agent.NewAgentByConfig()
	if err != nil {
		pkglog.FatalWithStackTrace(err)
	}
	for {
		if err := a.Watch(context.Background()); err != nil {
			log.Warn(err)
		}
	}
}
