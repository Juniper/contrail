package common

import (
	"strings"

	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

//InitConfig initializes Viper config.
func InitConfig() error {
	viper.SetConfigName("contrail")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../../sample")
	viper.AddConfigPath("../../sample")
	viper.AddConfigPath("../sample")
	viper.AddConfigPath("./sample")
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

//LoadConfig load data from data and bind to struct.
func LoadConfig(path string, dest interface{}) error {
	config := viper.Get(path)
	configYAML, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = yaml.UnmarshalStrict(configYAML, dest)
	if err != nil {
		return err
	}
	return nil
}
