package asfcli

import (
	"strings"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	Command.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Configuration file")
	viper.SetEnvPrefix("asf")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

// Command defines root ASF CLI command.
var Command = &cobra.Command{
	Use:   "asfcli",
	Short: "ASF CLI command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute executes the main command.
func Execute() error {
	return Command.Execute()
}

func initConfig() {
	if configFile == "" {
		configFile = viper.GetString("config")
	}
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
