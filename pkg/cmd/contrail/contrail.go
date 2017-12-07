package contrail

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var configFile string

func init() {
	cobra.OnInitialize()
	cobra.OnInitialize(initConfig)
	Cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Configuraion File")
}

//Cmd for main process commands
var Cmd = &cobra.Command{
	Use:   "contrail",
	Short: "Contrail command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't read config:", err)
	}
}
