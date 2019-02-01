package contrailutil

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/log"
)

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	ContrailUtil.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Configuration File")
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

// ContrailUtil defines root Contrail utility command.
var ContrailUtil = &cobra.Command{
	Use:   "contrailutil",
	Short: "Contrail Utility Command",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func initConfig() {
	if configFile == "" {
		return
	}
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.FatalWithStackTrace(err)
	}

	if err := log.Configure(viper.GetString("log_level")); err != nil {
		log.FatalWithStackTrace(err)
	}
}
