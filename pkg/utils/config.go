package utils

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("test_config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../../tools")
	viper.AddConfigPath("../../tools")
	viper.AddConfigPath("../tools")
	viper.AddConfigPath("./tools")
	return viper.ReadInConfig()
}
