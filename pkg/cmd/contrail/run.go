package contrail

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/retry"
	"github.com/Juniper/contrail/pkg/agent"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/collector/analytics"
	"github.com/Juniper/contrail/pkg/compilation"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/db/cassandra"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/replication"
	"github.com/pkg/errors"

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
	maybeStartCacheService(wg)
	MaybeStart("replication.cassandra", startCassandraReplicator, wg)
	MaybeStart("replication.amqp", startAmqpReplicator, wg)
	MaybeStart("server", startServer, wg)
	MaybeStart("agent", startAgent, wg)
	MaybeStart("sync", startSync, wg)
	MaybeStart("compilation", startCompilationService, wg)
	MaybeStart("collector", startCollectorWatcher, wg)
}

//MaybeStart runs process if it is enabled.
func MaybeStart(serviceName string, f func(), wg *sync.WaitGroup) {
	if !viper.GetBool(serviceName + ".enabled") {
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

func maybeStartCacheService(wg *sync.WaitGroup) {
	if !viper.GetBool("cache.enabled") {
		return
	}
	logrus.Debug("Cache service enabled")
	cacheDB = cache.NewDB(uint64(viper.GetInt64("cache.max_history")))
	MaybeStart("cache.cassandra", startCassandraWatcher, wg)
	MaybeStart("cache.etcd", startEtcdWatcher, wg)
	MaybeStart("cache.rdbms", startRDBMSWatcher, wg)
}

func startCassandraWatcher() {
	logrus.Debug("Cassandra watcher enabled for cache")
	producer := cassandra.NewEventProducer(cacheDB)
	err := producer.Start(context.Background())
	if err != nil {
		logrus.Warn(err)
	}
}

func startEtcdWatcher() {
	logrus.Debug("etcd watcher enabled for cache")
	producer, err := etcd.NewEventProducer(cacheDB, "cache-service")
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	err = producer.Start(context.Background())
	if err != nil {
		logrus.Warn(err)
	}
}

func startRDBMSWatcher() {
	logrus.Debug("RDBMS watcher enabled for cache")
	processor := &services.EventListProcessor{
		EventProcessor:    cacheDB,
		InTransactionDoer: services.NoTransaction,
	}
	producer, err := syncp.NewEventProducer("cache-watcher", processor)
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	defer producer.Close()
	err = producer.Start(context.Background())
	if err != nil {
		logrus.Warn(err)
	}
}

func startCassandraReplicator() {
	logrus.Debug("Cassandra replication service enabled")
	cassandraProcessor := cassandra.NewEventProcessor()
	producer, err := etcd.NewEventProducer(cassandraProcessor, "cassandra-replicator")
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	err = producer.Start(context.Background())
	if err != nil {
		logrus.Warn(err)
	}
}

func startAmqpReplicator() {
	logrus.Debug("AMQP replication service enabled")
	amqpProcessor := cassandra.NewAmqpEventProcessor()
	producer, err := etcd.NewEventProducer(amqpProcessor, "amqp-replicator")
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	err = producer.Start(context.Background())
	if err != nil {
		logrus.Warn(err)
	}
}

type EndpointStore interface {
	Range(f func(scope, value interface{}) bool)
	InStore(scope string) bool
	ReadEndpointsUrls(scope string, endpointKey string) []string
	Remove(scope string)
	UpdateEndpoint(scope string, endpoint *models.Endpoint) error
	InitScope(scope string)
	RemoveDeleted(scope, value interface{}, endpoints map[string]*models.Endpoint, log *logrus.Entry) bool
	GetEndpoints(scope string, endpointKey string) bool
	ReadAuthURLs(scope string, endpointKey string, authApiVersion string) []string
	ReadUsername(scope string, endpointKey string) string
	ReadPassword(scope string, endpointKey string) string
	GetAuthEndpoint(scope string, endpointKey string) (username string, password string, err error)
	GetEndpointUrl(clusterID, prefix string) (string, bool)
}

func initKeystone(s *apisrv.Server, endpointStore EndpointStore) error {
	if viper.GetBool("keystone.local") {
		k := apisrv.NewKeystone()
		err := k.Init(s.Echo, endpointStore)
		if err != nil {
			return errors.Wrap(err, "Failed to init local keystone server")
		}
		s.Keystone = k
	}
	return nil
}

func startVNCReplicator(s *apisrv.Server, endpointStore EndpointStore) (err error) {
	if viper.GetBool("server.enable_vnc_replication") {

		s.VNCReplicator, err = replication.New(endpointStore, s.Keystone)
		if err != nil {
			return err
		}
		err = s.VNCReplicator.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func startServer() {
	server, err := apisrv.NewServer()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	var endpointStore EndpointStore
	if err = initKeystone(server, endpointStore); err != nil {
		logutil.FatalWithStackTrace(err)
	}
	if err = startVNCReplicator(server, endpointStore); err != nil {
		logutil.FatalWithStackTrace(err)
	}
	server.Cache = cacheDB
	if err = server.Init(); err != nil {
		logutil.FatalWithStackTrace(err)
	}
	if err = server.Run(); err != nil {
		logutil.FatalWithStackTrace(err)
	}
}

func startSync() {
	if err := retry.Do(func() (retry bool, err error) {
		s, err := syncp.NewService()
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}
		defer s.Close()

		err = s.Run()

		return errutil.ShouldRetry(err), err
	}, retry.WithLog(logrus.StandardLogger()), retry.WithInterval(syncRetryInterval)); err != nil {
		logrus.Warn(err)
	}
}

func startCompilationService() {
	server, err := compilation.NewIntentCompilationService()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = server.Run(ctx); err != nil {
		logrus.Warn(err)
	}
}

func startAgent() {
	a, err := agent.NewAgentByConfig()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	for {
		if err := a.Watch(context.Background()); err != nil {
			logrus.Warn(err)
		}
	}
}

func startCollectorWatcher() {
	cfg := &analytics.Config{}
	if err := viper.UnmarshalKey("collector", cfg); err != nil {
		logrus.WithError(err).Warn("failed to unmarshal collector config")
		return
	}
	c, err := analytics.NewCollector(cfg)
	if err != nil {
		logrus.WithError(err).Warn("failed to create collector")
		return
	}
	if err = analytics.NewMessageBusProcessor(c); err != nil {
		logrus.WithError(err).Warn("failed to create collector")
	}
}
