package contrailutil

import (
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	common.SetLogLevel()
}
