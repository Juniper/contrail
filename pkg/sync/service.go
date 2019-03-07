// Package sync contains functionality that supplies etcd with data from PostgreSQL or MySQL database.
package sync

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	mysqlcanal "github.com/siddontang/go-mysql/canal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/constants"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/sync/replication"
	"github.com/Juniper/contrail/pkg/sync/sink"
)

const (
	mysqlDefaultPort = "3306"
	syncID           = "sync-service"
)

type watchCloser interface {
	Watch(context.Context) error
	DumpDone() <-chan struct{}
	Close()
}

// Service represents Sync service.
type Service struct {
	watcher watchCloser
	log     *logrus.Entry
}

// NewService creates Sync service with given configuration.
// Close needs to be explicitly called on service teardown.
func NewService() (*Service, error) {

	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}

	c := determineCodecType()
	if c == nil {
		return nil, errors.New(`unknown codec set as "sync.storage"`)
	}

	etcdNotifierService, err := etcd.NewNotifierService(viper.GetString(constants.ETCDPathVK), c)
	if err != nil {
		return nil, err
	}

	watcher, err := createWatcher(syncID, &services.ServiceEventProcessor{Service: etcdNotifierService})
	if err != nil {
		return nil, err
	}

	return &Service{
		watcher: watcher,
		log:     logutil.NewLogger(syncID),
	}, nil
}

func setViperDefaults() {
	viper.SetDefault("log_level", "debug")
	viper.SetDefault(constants.ETCDDialTimeoutVK, "60s")
	viper.SetDefault("database.retry_period", "1s")
	viper.SetDefault("database.connection_retries", 10)
	viper.SetDefault("database.replication_status_timeout", "10s")
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

func createWatcher(id string, processor services.EventProcessor) (watchCloser, error) {
	setViperDefaults()
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

	driver := viper.GetString("database.type")
	switch driver {
	case basedb.DriverPostgreSQL:
		return createPostgreSQLWatcher(id, rowSink, dbService, processor)
	case basedb.DriverMySQL:
		return createMySQLWatcher(rowSink)
	default:
		return nil, errors.Errorf("invalid database driver: %v", driver)
	}
}

func createPostgreSQLWatcher(
	id string, sink replication.RowSink, dbService *db.Service, processor services.EventProcessor,
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
		Slot:          replication.SlotName(id),
		Publication:   replication.PostgreSQLPublicationName,
		StatusTimeout: viper.GetDuration("database.replication_status_timeout"),
	}

	return replication.NewPostgresWatcher(
		conf,
		dbService,
		replConn,
		handler.Handle,
		processor,
		viper.GetBool("sync.dump"),
	)
}

func createMySQLWatcher(sink replication.RowSink) (watchCloser, error) {
	conf := canalConfig()

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

// DumpDone returns a channel that is closed when dump is done.
func (s *Service) DumpDone() <-chan struct{} {
	return s.watcher.DumpDone()
}

// Close closes Sync service.
func (s *Service) Close() {
	s.log.Info("Closing Sync service")
	s.watcher.Close()
}
