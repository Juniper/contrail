// Package watcher contains functionality that supplies etcd with data from MySQL database.
// It uses mysqldump and MySQL binlog replication protocol to achieve that.
package watcher

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Juniper/contrail/pkg/db"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/watcher/etcd"
	"github.com/Juniper/contrail/pkg/watcher/replication"
	"github.com/coreos/etcd/clientv3"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib" // allows using of pgx sql driver
	"github.com/kyleconroy/pgoutput"
	"github.com/pkg/errors"
	mysqlcanal "github.com/siddontang/go-mysql/canal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

// Storage strategies
const (
	StorageJSON   = "json"
	StorageNested = "nested"
)

type watchCloser interface {
	Watch(context.Context) error
	Close()
}

// Service represents Watcher service.
type Service struct {
	etcdClient *clientv3.Client
	watcher    watchCloser
	log        *logrus.Entry
}

// NewService creates Watcher service with given configuration.
// Close needs to be explicitly called on service teardown.
func NewService() (*Service, error) {
	// Logging
	if err := pkglog.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}
	log := pkglog.NewLogger("watcher-service")
	log.WithField("config", fmt.Sprintf("%+v", viper.AllSettings())).Debug("Got configuration")

	// Etcd client
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   viper.GetStringSlice("etcd.endpoints"),
		Username:    viper.GetString("etcd.username"),
		Password:    viper.GetString("etcd.password"),
		DialTimeout: viper.GetDuration("etcd.dial_timeout"),
	})
	if err != nil {
		return nil, err
	}

	// Etcd sink
	var sink replication.Sink
	switch viper.GetString("watcher.storage") {
	case StorageJSON:
		sink = etcd.NewJSONSink(clientv3.NewKV(etcdClient))
	case StorageNested:
		sink = etcd.NewNestingSink(clientv3.NewKV(etcdClient))
	default:
		return nil, errors.New("undefined storage strategy")
	}

	// Replication
	watcher, err := createWatcher(log, sink)
	if err != nil {
		return nil, err
	}

	return &Service{
		etcdClient: etcdClient,
		watcher:    watcher,
		log:        log,
	}, nil
}

func createWatcher(log *logrus.Entry, sink replication.Sink) (watchCloser, error) {
	driver := viper.GetString("database.type")
	if _, err := db.ConnectDB(); err != nil {
		return nil, err
	}

	switch driver {
	case db.DriverPostgreSQL:
		return createPostgreSQLWatcher(log, sink)
	case db.DriverMySQL:
		return createMySQLWatcher(log, sink)
	default:
		return nil, errors.New("undefined database type")
	}
}

func createPostgreSQLWatcher(log *logrus.Entry, sink replication.Sink) (watchCloser, error) {
	conf := pgx.ConnConfig{
		Host:     viper.GetString("database.host"),
		Database: viper.GetString("database.name"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
	}
	log.WithField("config", fmt.Sprintf("%+v", conf)).Debug("Got pgx config")
	conn, err := pgx.ReplicationConnect(conf)
	if err != nil {
		return nil, err
	}

	handler := replication.NewPgoutputEventHandler(sink)
	return replication.NewSubscriptionWatcher(
		conn,
		pgoutput.NewSubscription(replication.PostgreSQLReplicationSlotName, replication.PostgreSQLPublicationName),
		handler,
	), nil
}

func createMySQLWatcher(log *logrus.Entry, sink replication.Sink) (watchCloser, error) {
	conf := canalConfig()
	log.WithField("config", fmt.Sprintf("%+v", conf)).Debug("Got Canal config")

	canal, err := mysqlcanal.NewCanal(conf)
	if err != nil {
		return nil, err
	}
	canal.SetEventHandler(replication.NewCanalEventHandler(sink))

	return replication.NewBinlogWatcher(canal), nil
}

func canalConfig() *mysqlcanal.Config {
	c := mysqlcanal.NewDefaultConfig()
	c.Addr = viper.GetString("host")
	c.User = viper.GetString("user")
	c.Password = viper.GetString("password")
	c.Dump.Databases = []string{viper.GetString("name")}
	c.Dump.DiscardErr = false
	c.ServerID = randomServerID()
	return c
}

func randomServerID() uint32 {
	rand.Seed(time.Now().UnixNano())
	return uint32(rand.Intn(1000)) + 1001
}

// Run runs Watcher service.
func (s *Service) Run() error {
	s.log.Info("Running Watcher service")
	return s.watcher.Watch(context.Background())
}

// Close closes Watcher service.
func (s *Service) Close() {
	s.log.Info("Closing Watcher service")
	s.watcher.Close()
	if err := s.etcdClient.Close(); err != nil {
		s.log.WithField("error", err).Error("Error closing etcd connection")
	}
}
