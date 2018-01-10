package common

import (
	"github.com/spf13/viper"
)

//InitConfig initializes Viper config.
func InitConfig() error {
	viper.SetConfigName("test_config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../../tools")
	viper.AddConfigPath("../../tools")
	viper.AddConfigPath("../tools")
	viper.AddConfigPath("./tools")
	return viper.ReadInConfig()
}
