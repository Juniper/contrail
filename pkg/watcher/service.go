// Package watcher contains functionality that supplies etcd with data from MySQL database.
// It uses mysqldump and MySQL binlog replication protocol to achieve that.
package watcher

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/watcher/etcd"
	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	mysqlcanal "github.com/siddontang/go-mysql/canal"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/grpclog"
	"gopkg.in/yaml.v2"
)

// Storage strategies
const (
	StorageJSON   = "json"
	StorageNested = "nested"
)

const (
	dbConnectionRetries = 60
	dbRetryPeriod       = 1 * time.Second
	dbDriverName        = "mysql"
	defaultDialTimeoutS = 60
)

// Config represents configuration of Watcher service.
type Config struct {
	// Database represents configuration of database.
	Database DBConfig `yaml:"database"`
	// Etcd represents configuration of etcd service.
	Etcd EtcdConfig `yaml:"etcd"`
	// Storage specifies strategy of storing data in etcd (nested, json).
	Storage string `yaml:"storage"`
	// LogLevel is minimal log level, f.e. "info" (default "debug")
	LogLevel string `yaml:"log_level"`
}

// DBConfig represents configuration of source database.
type DBConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// EtcdConfig represents configuration of etcd service.
type EtcdConfig struct {
	Endpoints []string `yaml:"endpoints"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
	// DialTimeoutMs is timeout for establishing connection in seconds (default 60).
	DialTimeoutS uint `yaml:"dial_timeout_s"`
}

// Service represents Watcher service.
type Service struct {
	etcdClient *clientv3.Client
	watcher    *binlogWatcher
	log        *logrus.Entry
}

// NewServiceByFile creates Watcher service with configuration from file.
// Close needs to be explicitly called on service teardown.
func NewServiceByFile(configFilePath string) (*Service, error) {
	c, err := loadConfig(configFilePath)
	if err != nil {
		return nil, err
	}

	return NewService(c)
}

func loadConfig(configFilePath string) (*Config, error) {
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("read config file from path %s: %s", configFilePath, err)
	}

	var c Config
	if err = yaml.UnmarshalStrict(data, &c); err != nil {
		return nil, fmt.Errorf("unmarshal yaml data from config file: %s", err)
	}

	return &c, nil
}

// NewService creates Watcher service with given configuration.
// Close needs to be explicitly called on service teardown.
func NewService(c *Config) (*Service, error) {
	c = setDefaults(c)

	// Logging
	if err := pkglog.Configure(c.LogLevel); err != nil {
		return nil, err
	}
	log := pkglog.NewLogger("watcher-service")
	log.WithField("config", fmt.Sprintf("%+v", c)).Debug("Got configuration")

	// Etcd client
	clientv3.SetLogger(grpclog.NewLoggerV2(ioutil.Discard, os.Stdout, os.Stdout))
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Endpoints,
		Username:    c.Etcd.Username,
		Password:    c.Etcd.Password,
		DialTimeout: time.Duration(c.Etcd.DialTimeoutS) * time.Second,
	})
	if err != nil {
		return nil, err
	}

	// Etcd sink
	var etcdSink sink
	switch c.Storage {
	case StorageJSON:
		etcdSink = etcd.NewJSONSink(clientv3.NewKV(etcdClient))
	case StorageNested:
		etcdSink = etcd.NewNestingSink(clientv3.NewKV(etcdClient))
	default:
		return nil, errors.New("undefined storage strategy")
	}

	// Canal
	if err = awaitDB(&c.Database, log); err != nil {
		return nil, err
	}
	canal, err := mysqlcanal.NewCanal(canalConfig(&c.Database, log))
	if err != nil {
		return nil, err
	}
	canal.SetEventHandler(newHandler(etcdSink))

	return &Service{
		etcdClient: etcdClient,
		watcher:    newBinlogWatcher(canal),
		log:        log,
	}, nil
}

func setDefaults(c *Config) *Config {
	if c.LogLevel == "" {
		c.LogLevel = "debug"
	}
	if c.Etcd.DialTimeoutS == 0 {
		c.Etcd.DialTimeoutS = defaultDialTimeoutS
	}

	return c
}

func awaitDB(c *DBConfig, log *logrus.Entry) error {
	db, err := sql.Open(dbDriverName, dataSourceName(c))
	if err != nil {
		return fmt.Errorf("database connection: %s", err)
	}
	for i := 0; i < dbConnectionRetries; i++ {
		if err = db.Ping(); err == nil {
			return nil
		}
		time.Sleep(dbRetryPeriod)
		log.WithField("error", err).Debug("Cannot establish database connection, retrying")
	}
	return fmt.Errorf("reached database connection retry limit: %s", err)
}

func dataSourceName(c *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Password, c.Host, c.Name)
}

func canalConfig(db *DBConfig, log *logrus.Entry) *mysqlcanal.Config {
	c := mysqlcanal.NewDefaultConfig()
	c.Addr = db.Host
	c.User = db.User
	c.Password = db.Password
	c.Dump.Databases = []string{db.Name}
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
	return s.watcher.watch()
}

// Close closes Watcher service.
func (s *Service) Close() {
	s.log.Info("Closing Watcher service")
	s.watcher.close()
	if err := s.etcdClient.Close(); err != nil {
		s.log.WithField("error", err).Error("Error closing etcd connection")
	}
}
