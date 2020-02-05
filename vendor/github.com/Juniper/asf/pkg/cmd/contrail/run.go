package contrail

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/services"

	"github.com/Juniper/asf/pkg/apisrv"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/retry"
	syncp "github.com/Juniper/asf/pkg/sync"
	// TODO(buoto): Decouple from below packages
	//"github.com/Juniper/asf/pkg/agent"
	//"github.com/Juniper/asf/pkg/collector/analytics"
	//"github.com/Juniper/asf/pkg/compilation"
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
	MaybeStart("cache.etcd", startEtcdWatcher, wg)
	MaybeStart("cache.rdbms", startRDBMSWatcher, wg)
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

func startServer() {
	server, err := apisrv.NewServer()
	if err != nil {
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
