package config

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	viper.SetConfigFile("../test_data/test_config.yml")
	cfg, err := NewConfig("../test_data/test_config.yml")
	if err != nil {
		t.Errorf("Cannot read Config file")
	}

	if reflect.TypeOf(cfg.EtcdServersUrls).Kind() != reflect.Slice {
		t.Errorf("Error, Server Urls is not Right")
	}
	if reflect.TypeOf(cfg.EtcdServers).Kind() != reflect.Slice {
		t.Errorf("Error, Server List is not Right")
	}
}
