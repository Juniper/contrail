package contrail

import (
	"context"
	"crypto/tls"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/retry"
	"github.com/Juniper/contrail/pkg/apiserver"
	"github.com/Juniper/contrail/pkg/cache"
	"github.com/Juniper/contrail/pkg/cassandra"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/collector/analytics"
	"github.com/Juniper/contrail/pkg/compilation"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/etcd"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron"
	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/Juniper/contrail/pkg/rbac"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	asfapiserver "github.com/Juniper/asf/pkg/apiserver"
	asfkeystone "github.com/Juniper/asf/pkg/keystone"
	syncp "github.com/Juniper/contrail/pkg/sync"
)

const (
	syncRetryInterval = 3 * time.Second
)

// TODO: remove
var cacheDB *cache.DB

// Run runs Contrail services.
func Run() error {
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	contrailCmd := &cobra.Command{
		Use:   "contrail",
		Short: "Contrail command",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	var configFile string
	contrailCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Configuration File")
	cobra.OnInitialize(func() {
		initConfig(configFile)
	})

	// TODO(dfurman): use RunE() instead of Run() to return errors to contrailCmd.Execute() caller
	contrailCmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "Start Contrail Processes",
		Run: func(cmd *cobra.Command, args []string) {
			wg := &sync.WaitGroup{}
			StartProcesses(wg)
			wg.Wait()
		},
	})

	return contrailCmd.Execute()
}

func initConfig(configFile string) {
	if configFile == "" {
		configFile = viper.GetString("config") // TODO(dfurman): does it make sense?
	}
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		logutil.FatalWithStackTrace(err)
	}

	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		logutil.FatalWithStackTrace(err)
	}
}

//StartProcesses starts processes based on config.
func StartProcesses(wg *sync.WaitGroup) {
	maybeStartCacheService(wg)
	MaybeStart("replication.cassandra", startCassandraReplicator, wg)
	MaybeStart("replication.amqp", startAmqpReplicator, wg)
	MaybeStart("server", startServer, wg)
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
	producer, err := syncp.NewEventProducer("cache-watcher", cacheDB, services.NoTransaction)
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
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
		cErr := dbService.Close()
		logutil.FatalWithStackTraceIfError(cErr)
	}()

	serviceChain, err := NewServiceChain(dbService, analyticsCollector)
	logutil.FatalWithStackTraceIfError(err)

	staticProxyPlugin, err := proxy.NewStaticByViper()
	logutil.FatalWithStackTraceIfError(err)

	plugins := []asfapiserver.APIPlugin{
		serviceChain,
		staticProxyPlugin,
		analytics.BodyDumpPlugin{Collector: analyticsCollector},
	}
	plugins = append(plugins, services.ContrailPlugins(
		serviceChain,
		dbService,
		dbService,
		serviceChain,
	)...)

	if viper.GetBool("keystone.local") {
		k, initErr := keystone.Init()
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

	server, err := asfapiserver.NewServer(plugins, NoAuthPaths())
	logutil.FatalWithStackTraceIfError(err)

	err = server.Run()
	logutil.FatalWithStackTraceIfError(err)
}

// NewServiceChain creates a new service chain which can be plugged into API Server.
// It returns the first service of the built service chain.
func NewServiceChain(dbService *db.Service, c collector.Collector) (*services.ContrailService, error) {
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

	serviceChain, err := apiserver.SetupServiceChain(dbService, rbac.NewContrailAccessGetter(dbService), extraServices...)
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

// NoAuthPaths returns HTTP paths that require no authentication.
func NoAuthPaths() []string {
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

const (
	syncID = "sync-service"
)

func startSync() {
	if err := retry.Do(func() (retry bool, err error) {
		producer, err := syncp.NewEtcdFeeder(syncID)
		if err != nil {
			return false, err
		}
		err = producer.Start(context.Background())

		return errutil.ShouldRetry(err), err
	}, retry.WithLog(logrus.StandardLogger()), retry.WithInterval(syncRetryInterval)); err != nil {
		logutil.FatalWithStackTrace(err)
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
