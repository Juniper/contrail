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
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/Juniper/asf/pkg/retry"
	"github.com/Juniper/contrail/pkg/agent"
	"github.com/Juniper/contrail/pkg/apiserver"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/collector/analytics"
	"github.com/Juniper/contrail/pkg/compilation"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/db/cassandra"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/deploy"
	"github.com/Juniper/contrail/pkg/endpoint"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/neutron"
	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/Juniper/contrail/pkg/rbac"
	"github.com/Juniper/contrail/pkg/replication"
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
	contrailCmd.AddCommand(&cobra.Command{
		Use:   "cloud",
		Short: "sub command cloud is used to manage public cloud infra",
		Long: `Cloud is a sub command used to manage
            public cloud infra. Currently
            supported infra are Azure`,
		Run: func(cmd *cobra.Command, args []string) {
			manageCloud(configFile)
		},
	})
	contrailCmd.AddCommand(&cobra.Command{
		Use:   "deploy",
		Short: "Start managing contrail cluster",
		Run: func(cmd *cobra.Command, args []string) {
			manageCluster(configFile)
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

func manageCloud(configFile string) {
	manager, err := cloud.NewCloudManager(configFile)
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	if err = manager.Manage(); err != nil {
		logutil.FatalWithStackTrace(err)
	}
}

// TODO(apasyniuk) Export this into asf/osutils
type osCommandExecutor struct{}

func (e *osCommandExecutor) ExecCmdAndWait(
	r *report.Reporter, cmd string, args []string, dir string, envVars ...string,
) error {
	return osutil.ExecCmdAndWait(r, cmd, args, dir, envVars...)
}

func manageCluster(configFile string) {
	manager, err := deploy.NewDeployManager(configFile)
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}

	if err = manager.Manage(); err != nil {
		logutil.FatalWithStackTrace(err)
	}
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
		cErr := dbService.Close()
		logutil.FatalWithStackTraceIfError(cErr)
	}()

	serviceChain, err := NewServiceChain(dbService, analyticsCollector)
	logutil.FatalWithStackTraceIfError(err)

	staticProxyPlugin, err := proxy.NewStaticByViper()
	logutil.FatalWithStackTraceIfError(err)

	es := endpoint.NewStore()
	dynamicProxy := proxy.NewDynamicFromViper(es, dbService) // TODO(dfurman): it could use head of service chain
	// TODO(dfurman): move to proxy constructor and use context for cancellation
	dynamicProxy.StartEndpointsSync()
	defer dynamicProxy.StopEndpointsSync()

	plugins := []asfapiserver.APIPlugin{
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

	server, err := asfapiserver.NewServer(plugins, NoAuthPaths())
	logutil.FatalWithStackTraceIfError(err)

	r, err := startVNCReplicator(es)
	logutil.FatalWithStackTraceIfError(err)
	defer r.Stop()

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
		if err = a.Watch(context.Background()); err != nil {
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
