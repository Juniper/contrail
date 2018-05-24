// Package sync contains functionality that supplies etcd with data from PostgreSQL or MySQL database.
package sync

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/etcd"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/sync/replication"
	"github.com/Juniper/contrail/pkg/sync/sink"
	"github.com/coreos/etcd/clientv3"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	mysqlcanal "github.com/siddontang/go-mysql/canal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
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
	codec      sink.Codec
	sink       sink.Sink
	dbService  *db.Service

	log *logrus.Entry
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
	if err := pkglog.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}
	log := pkglog.NewLogger("sync-service")
	log.WithField("config", fmt.Sprintf("%+v", viper.AllSettings())).Debug("Got configuration")

	// Etcd client
	clientv3.SetLogger(grpclog.NewLoggerV2(ioutil.Discard, os.Stdout, os.Stdout))
	etcdClient, err := etcd.DialByConfig()
	if err != nil {
		return nil, err
	}

	// Etcd sink
	codec := sink.JSONCodec
	s := etcd.NewClient(etcdClient)

	sqlDB, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}

	dbService := db.NewService(sqlDB, viper.GetString("database.dialect"))
	if err != nil {
		return nil, err
	}

	service := &Service{
		etcdClient: etcdClient,
		codec:      codec,
		sink:       s,
		dbService:  dbService,
		log:        log,
	}

	// Replication
	watcher, err := service.createWatcher(log, s)
	if err != nil {
		return nil, err
	}

	service.watcher = watcher

	return service, nil
}

func (s *Service) createWatcher(log *logrus.Entry, si sink.Sink) (watchCloser, error) {
	objectWriter := sink.NewObjectWriter(si)

	driver := viper.GetString("database.type")
	switch driver {
	case db.DriverPostgreSQL:
		return createPostgreSQLWatcher(log, s, s.dbService, objectWriter)
	case db.DriverMySQL:
		return createMySQLWatcher(log, s)
	default:
		return nil, errors.New("undefined database type")
	}
}

func createPostgreSQLWatcher(log *logrus.Entry, sink replication.RowSink, dbService *db.Service, ow db.ObjectWriter) (watchCloser, error) {
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

	return replication.NewPostgresWatcher(conf, dbService, replConn, handler.Handle, ow)
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

func (s *Service) handle(ctx context.Context, r HandleRequest, functions map[string]tableCallback) error {
	callback, ok := functions[r.SchemaID]
	if !ok {
		return fmt.Errorf("invalid table name: %v", r.SchemaID)
	}
	return callback(s, ctx, r)
}

// Create handles INSERT replication message.
func (s *Service) Create(ctx context.Context, schemaID string, pk string, data map[string]interface{}) error {
	if err := s.handle(ctx, HandleRequest{SchemaID: schemaID, PK: pk, Data: data}, createFunctions); err != nil {
		return fmt.Errorf("error handling Create: %v", err)
	}
	return nil
}

// Update handles UPDATE replication message.
func (s *Service) Update(ctx context.Context, schemaID string, pk string, data map[string]interface{}) error {
	if err := s.handle(ctx, HandleRequest{SchemaID: schemaID, PK: pk, Data: data}, updateFunctions); err != nil {
		return fmt.Errorf("error handling Update: %v", err)
	}
	return nil
}

// Delete handles INSERT replication message.
func (s *Service) Delete(ctx context.Context, schemaID string, pk string, data map[string]interface{}) error {
	if err := s.handle(ctx, HandleRequest{SchemaID: schemaID, PK: pk, Data: data}, deleteFunctions); err != nil {
		return fmt.Errorf("error handling Delete: %v", err)
	}
	return nil
}
