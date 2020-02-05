package config

import (
	"github.com/Juniper/asf/pkg/etcd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// DefaultConfig section.
type DefaultConfig struct {
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

// APIClientConfig is the configuration for intent compiler's REST API client.
type APIClientConfig struct {
	URL                     string
	AuthURL                 string
	ID, Password            string
	DomainID, ProjectID     string
	DomainName, ProjectName string
	Insecure                bool
}

// PluginConfig section.
type PluginConfig struct {
	Handlers map[string]interface{}
}

// Config object.
type Config struct {
	DefaultCfg      DefaultConfig
	EtcdNotifierCfg EtcdNotifierConfig
	APIClientConfig APIClientConfig
	PluginCfg       PluginConfig
	PluginNames     []string
}

// ReadConfig gets configuration from Viper and logs its values.
func ReadConfig() Config {
	viper.SetDefault("compilation.service_name", "intent-compilation-service")

	c := Config{
		DefaultCfg: DefaultConfig{
			PluginDirectory: viper.GetString("compilation.plugin_directory"),
			NumberOfWorkers: viper.GetInt("compilation.number_of_workers"),
			MaxJobQueueLen:  viper.GetInt("compilation.max_job_queue_len"),
		},
		EtcdNotifierCfg: EtcdNotifierConfig{
			EtcdServers:      viper.GetStringSlice(etcd.ETCDEndpointsVK),
			WatchPath:        viper.GetString(etcd.ETCDPathVK),
			MsgQueueLockTime: viper.GetInt("compilation.msg_queue_lock_time"), // TODO(Michal): Change to GetDuration
			MsgIndexString:   viper.GetString("compilation.msg_index_string"),
			ReadLockString:   viper.GetString("compilation.read_lock_string"),
			MasterElection:   viper.GetBool("compilation.master_election"),
		},
		APIClientConfig: APIClientConfig{
			URL:         viper.GetString("client.endpoint"),
			ID:          viper.GetString("client.id"),
			Password:    viper.GetString("client.password"),
			ProjectID:   viper.GetString("client.project_id"),
			ProjectName: viper.GetString("client.project_name"),
			DomainID:    viper.GetString("client.domain_id"),
			DomainName:  viper.GetString("client.domain_name"),

			AuthURL:  viper.GetString("keystone.authurl"),
			Insecure: viper.GetBool("insecure"),
		},
		PluginCfg: PluginConfig{
			Handlers: viper.GetStringMap("plugin.handlers"),
		},
	}

	logrus.Println("Plugin Directory:", c.DefaultCfg.PluginDirectory)
	logrus.Println("Number of Workers:", c.DefaultCfg.NumberOfWorkers)
	logrus.Println("Maximum Job Queue Len:", c.DefaultCfg.MaxJobQueueLen)

	logrus.Println("etcd Notifier Servers List:", c.EtcdNotifierCfg.EtcdServers)
	logrus.Println("etcd Notifier WatchPath :", "/"+c.EtcdNotifierCfg.WatchPath)
	logrus.Println("etcd Notifier MsgQueueLockTime:", c.EtcdNotifierCfg.MsgQueueLockTime)
	logrus.Println("etcd Notifier MsgIndexString:", c.EtcdNotifierCfg.MsgIndexString)
	logrus.Println("etcd Notifier ReadLockString:", c.EtcdNotifierCfg.ReadLockString)
	logrus.Println("etcd Notifier MasterElection:", c.EtcdNotifierCfg.MasterElection)

	logrus.Println("Plugin Handlers:", c.PluginCfg.Handlers)
	return c
}
