package contrail

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/logutil"
)

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	Contrail.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Configuration File")
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

// Contrail defines root Contrail command.
var Contrail = &cobra.Command{
	Use:   "contrail",
	Short: "Contrail command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func initConfig() {
	if configFile == "" {
		configFile = viper.GetString("config")
	}
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Failed to read config: ", err)
	}

	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		logrus.Fatal(err)
	}
}
