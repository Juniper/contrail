// Package watcher contains functionality that supplies etcd with data from MySQL database.
// It uses mysqldump and MySQL binlog replication protocol to achieve that.
package watcher

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/watcher/etcd"
	"github.com/Juniper/contrail/pkg/watcher/replication"
	"github.com/coreos/etcd/clientv3"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib" // allows using of pgx sql driver
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

const (
	dbDriverMySQL      = "mysql"
	dbDriverPostgreSQL = "pgx"

	dbDSNFormatMySQL      = "%s:%s@tcp(%s)/%s"
	dbDSNFormatPostgreSQL = "user=%s password=%s host=%s dbname=%s"
)

type config interface {
	GetString(string) string
	GetStringSlice(string) []string
	GetInt(string) int
	GetDuration(string) time.Duration

	AllSettings() map[string]interface{}

	SetDefault(string, interface{})
}

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

// NewServiceByFile creates Watcher service with configuration from file.
// Close needs to be explicitly called on service teardown.
func NewServiceByFile(configFilePath string) (*Service, error) {
	c := viper.New()
	c.SetConfigFile(configFilePath)
	if err := c.ReadInConfig(); err != nil {
		return nil, err
	}
	setDefaults(c)

	return NewService(c)
}

func setDefaults(c config) {
	c.SetDefault("log_level", "debug")
	c.SetDefault("etcd.dial_timeout", "60s")
	c.SetDefault("database.retry_period", "1s")
	c.SetDefault("database.connection_retries", 60)
}

// NewService creates Watcher service with given configuration.
// Close needs to be explicitly called on service teardown.
func NewService(conf *viper.Viper) (*Service, error) {
	// Logging
	if err := pkglog.Configure(conf.GetString("log_level")); err != nil {
		return nil, err
	}
	log := pkglog.NewLogger("watcher-service")
	log.WithField("config", fmt.Sprintf("%+v", conf.AllSettings())).Debug("Got configuration")

	// Etcd client
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.GetStringSlice("etcd.endpoints"),
		Username:    conf.GetString("etcd.username"),
		Password:    conf.GetString("etcd.password"),
		DialTimeout: conf.GetDuration("etcd.dial_timeout"),
	})
	if err != nil {
		return nil, err
	}

	// Etcd sink
	var sink replication.Sink
	switch conf.GetString("storage") {
	case StorageJSON:
		sink = etcd.NewJSONSink(clientv3.NewKV(etcdClient))
	case StorageNested:
		sink = etcd.NewNestingSink(clientv3.NewKV(etcdClient))
	default:
		return nil, errors.New("undefined storage strategy")
	}

	// Replication
	watcher, err := createWatcher(conf, log, sink)
	if err != nil {
		return nil, err
	}

	return &Service{
		etcdClient: etcdClient,
		watcher:    watcher,
		log:        log,
	}, nil
}

func createWatcher(dbConfig config, log *logrus.Entry, sink replication.Sink) (watchCloser, error) {
	switch dbConfig.GetString("database.driver") {
	case dbDriverPostgreSQL:
		return createPostgreSQLWatcher(dbConfig, log, sink)
	case dbDriverMySQL:
		return createMySQLWatcher(dbConfig, log, sink)
	default:
		return nil, errors.New("undefined database type")
	}
}

func createPostgreSQLWatcher(c config, log *logrus.Entry, sink replication.Sink) (watchCloser, error) {
	if err := awaitDB(c, log, dbDSNFormatPostgreSQL); err != nil {
		return nil, err
	}
	config := pgx.ConnConfig{
		Host:     c.GetString("database.host"),
		Database: c.GetString("database.name"),
		User:     c.GetString("database.user"),
		Password: c.GetString("database.password"),
	}
	conn, err := pgx.ReplicationConnect(config)
	if err != nil {
		return nil, err
	}

	handler := replication.NewPgoutputEventHandler(sink)
	return replication.NewSubscriptionWatcher(
		conn,
		replication.PostgreSQLReplicationSlotName,
		replication.PostgreSQLPublicationName,
		handler,
	), nil
}

func createMySQLWatcher(c config, log *logrus.Entry, sink replication.Sink) (watchCloser, error) {
	if err := awaitDB(c, log, dbDSNFormatMySQL); err != nil {
		return nil, err
	}
	canal, err := mysqlcanal.NewCanal(canalConfig(c, log))
	if err != nil {
		return nil, err
	}
	canal.SetEventHandler(replication.NewCanalEventHandler(sink))

	return replication.NewBinlogWatcher(canal), nil
}

func awaitDB(c config, log *logrus.Entry, dbDSNFormat string) error {
	db, err := sql.Open(c.GetString("database.driver"), dataSourceName(c, dbDSNFormat))
	if err != nil {
		return fmt.Errorf("database connection: %s", err)
	}
	for i := 0; i < c.GetInt("database.connection_retries"); i++ {
		if err = db.Ping(); err == nil {
			return nil
		}
		time.Sleep(c.GetDuration("database.retry_period"))
		log.WithField("error", err).Debug("Cannot establish database connection, retrying")
	}
	return fmt.Errorf("reached database connection retry limit: %s", err)
}

func dataSourceName(c config, format string) string {
	return fmt.Sprintf(
		format,
		c.GetString("database.user"),
		c.GetString("database.password"),
		c.GetString("database.host"),
		c.GetString("database.name"),
	)
}

func canalConfig(conf config, log *logrus.Entry) *mysqlcanal.Config {
	c := mysqlcanal.NewDefaultConfig()
	c.Addr = conf.GetString("host")
	c.User = conf.GetString("user")
	c.Password = conf.GetString("password")
	c.Dump.Databases = []string{conf.GetString("name")}
	c.Dump.DiscardErr = false
	c.ServerID = randomServerID()
	log.WithField("config", fmt.Sprintf("%+v", c)).Debug("Got Canal config")
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
