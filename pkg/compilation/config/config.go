package config

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
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
	return ReadCompilerConfig("")
}

// ReadCompilerConfig gets configuration from Viper and logs its values.
func ReadCompilerConfig(compiler string) Config {
	viper.SetDefault(
		strings.Join([]string{compiler, "compilation.service_name"}, "."),
		"intent-compilation-service")

	c := Config{
		DefaultCfg: DefaultConfig{
			PluginDirectory: viper.GetString(
				strings.Join([]string{compiler, "compilation.plugin_directory"}, ".")),
			NumberOfWorkers: viper.GetInt(
				strings.Join([]string{compiler, "compilation.number_of_workers"}, ".")),
			MaxJobQueueLen: viper.GetInt(
				strings.Join([]string{compiler, "compilation.max_job_queue_len"}, ".")),
		},
		EtcdNotifierCfg: EtcdNotifierConfig{
			EtcdServers: viper.GetStringSlice("etcd.endpoints"),
			WatchPath:   viper.GetString("etcd.path"),
			MsgQueueLockTime: viper.GetInt(
				strings.Join([]string{compiler, "compilation.msg_queue_lock_time"}, ".")), // TODO(Michal): Change to GetDuration
			MsgIndexString: viper.GetString(
				strings.Join([]string{compiler, "compilation.msg_index_string"}, ".")),
			ReadLockString: viper.GetString(
				strings.Join([]string{compiler, "compilation.read_lock_string"}, ".")),
			MasterElection: viper.GetBool(
				strings.Join([]string{compiler, "compilation.master_election"}, ".")),
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
