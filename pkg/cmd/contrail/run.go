package contrail

import (
	"context"
	"crypto/tls"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Juniper/asf/pkg/db/basedb"
	asfkeystone "github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron"
	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/Juniper/contrail/pkg/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/asf/pkg/apisrv/baseapisrv"
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
	"github.com/Juniper/contrail/pkg/endpoint"
	"github.com/Juniper/contrail/pkg/replication"

	syncp "github.com/Juniper/contrail/pkg/sync"
)

const (
	syncRetryInterval = 3 * time.Second
)

// TODO: remove
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

// TODO: move to startServer
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

// TODO: move to startServer
func startCassandraWatcher() {
	logrus.Debug("Cassandra watcher enabled for cache")
	producer := cassandra.NewEventProducer(cacheDB)
	err := producer.Start(context.Background())
	if err != nil {
		logrus.Warn(err)
	}
}

// TODO: move to startServer
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

// TODO: move to startServer
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

func startServer() {
	analyticsCollector, err := analytics.NewCollectorFromGlobalConfig()
	logutil.FatalWithStackTraceIfError(err)
	analytics.AddLoggerHook(analyticsCollector)

	sqlDB, err := basedb.ConnectDB(analytics.WithCommitLatencyReporting(analyticsCollector))
	logutil.FatalWithStackTraceIfError(err)
	dbService := db.NewService(sqlDB)
	defer func() {
		if cErr := dbService.Close(); cErr != nil {
			logrus.WithError(err).Error("Failed to close DB service")
		}
	}()

	serviceChain, err := newServiceChain(dbService, analyticsCollector)
	logutil.FatalWithStackTraceIfError(err)

	staticProxyPlugin, err := proxy.NewStaticByViper()
	logutil.FatalWithStackTraceIfError(err)

	es := endpoint.NewStore()
	dynamicProxy := proxy.NewDynamicFromViper(es, dbService) // TODO(dfurman): it could use head of service chain
	// TODO(dfurman): move to proxy constructor and use context for cancellation
	dynamicProxy.StartEndpointsSync()
	defer dynamicProxy.StopEndpointsSync()

	plugins := []baseapisrv.APIPlugin{
		serviceChain,
		staticProxyPlugin,
		dynamicProxy,
		services.UploadCloudKeysPlugin{},
		analytics.BodyDumpPlugin{Collector: analyticsCollector},
	}

	if viper.GetBool("keystone.local") {
		k, initErr := keystone.Init(es)
		logutil.FatalWithStackTraceIfError(initErr)
		plugins = append(plugins, k)
	}

	if cacheDB != nil {
		plugins = append(plugins, cacheDB)
	}

	if viper.GetBool("server.enable_vnc_neutron") {
		plugins = append(plugins, &neutron.Server{
			ReadService:       serviceChain,
			WriteService:      serviceChain,
			UserAgentKV:       serviceChain,
			IDToFQNameService: serviceChain,
			FQNameToIDService: serviceChain,
			InTransactionDoer: dbService,
			Log:               logutil.NewLogger("neutron-server"),
		})
	}

	server, err := baseapisrv.NewServer(plugins, noAuthPaths())
	logutil.FatalWithStackTraceIfError(err)

	r, err := startVNCReplicator(es)
	logutil.FatalWithStackTraceIfError(err)
	defer r.Stop()

	err = server.Run()
	logutil.FatalWithStackTraceIfError(err)
}

func newServiceChain(dbService *db.Service, c collector.Collector) (*services.ContrailService, error) {
	var extraServices []services.Service
	var neutronService *neutron.Service
	if viper.GetBool("server.enable_vnc_neutron") {
		neutronService = &neutron.Service{
			Keystone: &asfkeystone.Client{
				URL: viper.GetString("keystone.authurl"),
				HTTPDoer: analytics.LatencyReportingDoer{
					Doer: &http.Client{
						Transport: &http.Transport{
							TLSClientConfig: &tls.Config{InsecureSkipVerify: viper.GetBool("keystone.insecure")},
						},
					},
					Collector:   c,
					Operation:   "VALIDATE",
					Application: "KEYSTONE",
				},
			},
			ReadService:       dbService,
			MetadataGetter:    dbService,
			InTransactionDoer: dbService,
		}
		extraServices = append(extraServices, neutronService)
	}

	typeLogicService := &types.ContrailTypeLogicService{
		ReadService:       dbService,
		InTransactionDoer: dbService,
		AddressManager:    dbService,
		IntPoolAllocator:  dbService,
		MetadataGetter:    dbService,
	}
	extraServices = append(extraServices, typeLogicService)

	serviceChain, err := apisrv.SetupServiceChain(dbService, extraServices...)
	if err != nil {
		return nil, err
	}

	if neutronService != nil {
		neutronService.WriteService = &services.InternalContextWriteServiceWrapper{
			WriteService: serviceChain,
		}
	}
	typeLogicService.WriteService = &services.InternalContextWriteServiceWrapper{
		WriteService: serviceChain,
	}
	return serviceChain, nil
}

func noAuthPaths() []string {
	return []string{
		"/v3/auth/tokens", // TODO(mblotniak): Is this ever used?
		strings.Join([]string{
			models.ContrailClusterPluralPath,
			"?fields=",
			models.ContrailClusterFieldUUID,
			",",
			models.ContrailClusterFieldName,
		}, ""),
	}
}

func startVNCReplicator(es *endpoint.Store) (vncReplicator *replication.Replicator, err error) {
	vncReplicator, err = replication.New(es)
	if err != nil {
		return nil, err
	}
	err = vncReplicator.Start()
	if err != nil {
		return nil, err
	}
	return vncReplicator, nil
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
