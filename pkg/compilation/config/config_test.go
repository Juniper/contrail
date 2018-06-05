package config

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	viper.AddConfigPath("../test_data/")
	viper.SetConfigFile("test_config")
	cfg, err := NewConfig()
	if err != nil {
		t.Errorf("Cannot read Config file")
	}

	if reflect.TypeOf(cfg.EtcdNotifierCfg.EtcdServers).Kind() != reflect.Slice {
		t.Errorf("Error, Server Urls is not Right")
	}
	if reflect.TypeOf(cfg.EtcdNotifierCfg.WatchPath).Kind() != reflect.String {
		t.Errorf("Error, Server List is not Right")
	}
}
