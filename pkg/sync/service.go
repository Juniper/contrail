// Package sync contains functionality that supplies etcd with data from PostgreSQL or MySQL database.
package sync

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	mysqlcanal "github.com/siddontang/go-mysql/canal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/sync/replication"
	"github.com/Juniper/contrail/pkg/sync/sink"
)

const (
	mysqlDefaultPort = "3306"
)

type watchCloser interface {
	Watch(context.Context) error
	Close()
}

// Service represents Sync service.
type Service struct {
	etcdClient *clientv3.Client
	watcher    watchCloser
	log        *logrus.Entry
}

// NewServiceByFile creates Sync service with configuration from file.
// Close needs to be explicitly called on service teardown.
func NewServiceByFile(configFilePath string) (*Service, error) {
	viper.SetConfigFile(configFilePath)
	if err := viper.MergeInConfig(); err != nil {
		return nil, err
	}

	return NewService()
}

func setDefaults() {
	viper.SetDefault("log_level", "debug")
	viper.SetDefault("etcd.dial_timeout", "60s")
	viper.SetDefault("database.retry_period", "1s")
	viper.SetDefault("database.connection_retries", 10)
	viper.SetDefault("database.replication_status_timeout", "10s")
}

// NewService creates Sync service with given configuration.
// Close needs to be explicitly called on service teardown.
func NewService() (*Service, error) {
	setDefaults()

	// Logging
	if err := log.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}
	log := log.NewLogger("sync-service")
	log.WithField("config", fmt.Sprintf("%+v", viper.AllSettings())).Debug("Got configuration")

	// Etcd client
	clientv3.SetLogger(grpclog.NewLoggerV2(ioutil.Discard, os.Stdout, os.Stdout))
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   viper.GetStringSlice("etcd.endpoints"),
		Username:    viper.GetString("etcd.username"),
		Password:    viper.GetString("etcd.password"),
		DialTimeout: viper.GetDuration("etcd.dial_timeout"),
	})
	if err != nil {
		return nil, err
	}

	c := determineCodecType()
	if c == nil {
		return nil, errors.New(`unknown codec set as "sync.storage"`)
	}

	// Etcd sink
	etcdNotifierService, err := etcd.NewNotifierService(viper.GetString("etcd.path"), c)

	if err != nil {
		return nil, err
	}
	processor := &services.ServiceEventProcessor{Service: etcdNotifierService}

	// Replication
	watcher, err := createWatcher(log, processor)
	if err != nil {
		return nil, err
	}

	return &Service{
		etcdClient: etcdClient,
		watcher:    watcher,
		log:        log,
	}, nil
}

func determineCodecType() models.Codec {
	switch viper.GetString("sync.storage") {
	case models.JSONCodec.Key():
		return models.JSONCodec
	case models.ProtoCodec.Key():
		return models.ProtoCodec
	default:
		return nil
	}
}

func createWatcher(log *logrus.Entry, processor services.EventProcessor) (watchCloser, error) {
	driver := viper.GetString("database.type")
	sqlDB, err := basedb.ConnectDB()
	if err != nil {
		return nil, err
	}

	dbService := db.NewService(sqlDB, viper.GetString("database.dialect"))
	if err != nil {
		return nil, err
	}

	s := sink.NewEventProcessorSink(processor)
	rowSink := replication.NewObjectMappingAdapter(s, dbService)

	switch driver {
	case basedb.DriverPostgreSQL:
		return createPostgreSQLWatcher(log, rowSink, dbService, processor)
	case basedb.DriverMySQL:
		return createMySQLWatcher(log, rowSink)
	default:
		return nil, errors.New("undefined database type")
	}
}

func createPostgreSQLWatcher(
	log *logrus.Entry, sink replication.RowSink, dbService *db.Service, processor services.EventProcessor,
) (watchCloser, error) {
	handler := replication.NewPgoutputEventHandler(sink)

	connConfig := pgx.ConnConfig{
		Host:     viper.GetString("database.host"),
		Database: viper.GetString("database.name"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
	}

	replConn, err := pgx.ReplicationConnect(connConfig)
	if err != nil {
		return nil, err
	}

	conf := replication.PostgresSubscriptionConfig{
		Slot:          replication.PostgreSQLReplicationSlotName,
		Publication:   replication.PostgreSQLPublicationName,
		StatusTimeout: viper.GetDuration("database.replication_status_timeout"),
	}
	log.WithField("config", fmt.Sprintf("%+v", conf)).Debug("Got pgx config")

	return replication.NewPostgresWatcher(conf, dbService, replConn, handler.Handle, processor)
}

func createMySQLWatcher(log *logrus.Entry, sink replication.RowSink) (watchCloser, error) {
	conf := canalConfig()
	log.WithField("config", fmt.Sprintf("%+v", conf)).Debug("Got Canal config")

	canal, err := mysqlcanal.NewCanal(conf)
	if err != nil {
		return nil, fmt.Errorf("error creating canal: %v", err)
	}
	canal.SetEventHandler(replication.NewCanalEventHandler(sink))

	return replication.NewMySQLWatcher(canal), nil
}

func canalConfig() *mysqlcanal.Config {
	c := mysqlcanal.NewDefaultConfig()
	c.Addr = viper.GetString("database.host") + ":" + mysqlDefaultPort
	c.User = viper.GetString("database.user")
	c.Password = viper.GetString("database.password")
	c.Dump.Databases = []string{viper.GetString("database.name")}
	c.Dump.DiscardErr = false
	c.ServerID = randomServerID()
	return c
}

func randomServerID() uint32 {
	rand.Seed(time.Now().UnixNano())
	return uint32(rand.Intn(1000)) + 1001
}

// Run runs Sync service.
func (s *Service) Run() error {
	s.log.Info("Running Sync service")
	return s.watcher.Watch(context.Background())
}

// Close closes Sync service.
func (s *Service) Close() {
	s.log.Info("Closing Sync service")
	s.watcher.Close()
	if err := s.etcdClient.Close(); err != nil {
		s.log.WithField("error", err).Error("Error closing etcd connection")
	}
}
