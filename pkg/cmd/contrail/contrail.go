package contrail

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var configFile string
var agentConfigFile string

func init() {
	cobra.OnInitialize(initConfig)
	Contrail.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Configuration File")
	Contrail.PersistentFlags().StringVarP(&agentConfigFile, "agent", "a", "", "Agent Config File")
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
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't read config:", err)
	}
}
