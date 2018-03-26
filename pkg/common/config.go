package common

import (
	"strings"

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
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}
