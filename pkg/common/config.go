package common

import (
	"github.com/spf13/viper"
)

//InitConfig initializes Viper config.
func InitConfig() error {
	viper.SetConfigName("server")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../../sample")
	viper.AddConfigPath("../../sample")
	viper.AddConfigPath("../sample")
	viper.AddConfigPath("./sample")
	return viper.ReadInConfig()
}
