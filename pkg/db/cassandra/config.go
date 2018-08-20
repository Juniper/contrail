package cassandra

import (
	"os"
	"time"

	"github.com/gocql/gocql"

	"github.com/spf13/viper"
)

const (
	defaultCassandraVersion  = "3.4.4"
	defaultCassandraKeyspace = "config_db_uuid"

	exchangeName = "vnc_config.object-update"
)

// Config fields for cassandra
type Config struct {
	Host           string
	Port           int
	Timeout        time.Duration
	ConnectTimeout time.Duration
}

// AmqpConfig groups config fields for AMQP
type AmqpConfig struct {
	host      string
	queueName string
}

//GetConfig returns cassandra Config filled with data from config file.
func GetConfig() Config {
	return Config{
		Host:           viper.GetString("cassandra.host"),
		Port:           viper.GetInt("cassandra.port"),
		Timeout:        viper.GetDuration("cassandra.timeout"),
		ConnectTimeout: viper.GetDuration("cassandra.connect_timeout"),
	}
}

func getQueueName() string {
	name, _ := os.Hostname() // nolint: noerror
	return "contrail_process_" + name
}

func getCluster(cfg Config) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(cfg.Host)
	if cfg.Port != 0 {
		cluster.Port = cfg.Port
	}
	cluster.Timeout = cfg.Timeout
	cluster.ConnectTimeout = cfg.ConnectTimeout
	cluster.Keyspace = defaultCassandraKeyspace
	cluster.CQLVersion = defaultCassandraVersion
	return cluster
}
