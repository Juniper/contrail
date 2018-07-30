package config

import (
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

// DefaultConfig section.
type DefaultConfig struct {
	ServiceName     string
	PluginDirectory string
	NumberOfWorkers int
	MaxJobQueueLen  int
}

// EtcdNotifierConfig section.
type EtcdNotifierConfig struct {
	EtcdServers      []string
	WatchPath        string
	MsgQueueLockTime int
	MsgIndexString   string
	ReadLockString   string
	MasterElection   bool
}

// PluginConfig section.
type PluginConfig struct {
	Handlers map[string]interface{}
}

// config object.
type Config struct {
	DefaultCfg      DefaultConfig
	EtcdNotifierCfg EtcdNotifierConfig
	PluginCfg       PluginConfig
	PluginNames     []string
}

// ReadConfig gets configuration from Viper and logs its values.
func ReadConfig() Config {
	viper.SetDefault("compilation.service_name", "intent-compilation-service")

	c := Config{
		DefaultCfg: DefaultConfig{
			ServiceName:     viper.GetString("compilation.service_name"),
			PluginDirectory: viper.GetString("compilation.plugin_directory"),
			NumberOfWorkers: viper.GetInt("compilation.number_of_workers"),
			MaxJobQueueLen:  viper.GetInt("compilation.max_job_queue_len"),
		},
		EtcdNotifierCfg: EtcdNotifierConfig{
			EtcdServers:      viper.GetStringSlice("etcd.endpoints"),
			WatchPath:        viper.GetString("etcd.path"),
			MsgQueueLockTime: viper.GetInt("compilation.msg_queue_lock_time"), // TODO(Michal): Change to GetDuration
			MsgIndexString:   viper.GetString("compilation.msg_index_string"),
			ReadLockString:   viper.GetString("compilation.read_lock_string"),
			MasterElection:   viper.GetBool("compilation.master_election"),
		},
		PluginCfg: PluginConfig{
			Handlers: viper.GetStringMap("plugin.handlers"),
		},
	}

	log.Println("Intent compilation service name:", c.DefaultCfg.ServiceName)
	log.Println("Plugin Directory:", c.DefaultCfg.PluginDirectory)
	log.Println("Number of Workers:", c.DefaultCfg.NumberOfWorkers)
	log.Println("Maximum Job Queue Len:", c.DefaultCfg.MaxJobQueueLen)

	log.Println("ETCD Notifier Servers List:", c.EtcdNotifierCfg.EtcdServers)
	log.Println("ETCD Notifier WatchPath :", "/"+c.EtcdNotifierCfg.WatchPath)
	log.Println("ETCD Notifier MsgQueueLockTime:", c.EtcdNotifierCfg.MsgQueueLockTime)
	log.Println("ETCD Notifier MsgIndexString:", c.EtcdNotifierCfg.MsgIndexString)
	log.Println("ETCD Notifier ReadLockString:", c.EtcdNotifierCfg.ReadLockString)
	log.Println("ETCD Notifier MasterElection:", c.EtcdNotifierCfg.MasterElection)

	log.Println("Plugin Handlers:", c.PluginCfg.Handlers)
	return c
}
