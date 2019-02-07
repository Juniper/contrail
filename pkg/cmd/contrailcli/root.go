package contrailcli

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
)

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	ContrailCLI.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		"Configuration File")
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

// ContrailCLI defines root Contrail CLI command.
var ContrailCLI = &cobra.Command{
	Use:   "contrailcli",
	Short: "Contrail CLI command",
	Run:   func(cmd *cobra.Command, args []string) {},
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

func getClient() (*client.HTTP, error) {
	authURL := viper.GetString("keystone.auth_url")
	scope := client.GetKeystoneScope(
		viper.GetString("client.domain_id"),
		viper.GetString("client.domain_name"),
		viper.GetString("client.project_id"),
		viper.GetString("client.project_name"),
	)
	client := client.NewHTTP(
		viper.GetString("client.endpoint"),
		authURL,
		viper.GetString("client.id"),
		viper.GetString("client.password"),
		viper.GetBool("insecure"),
		scope,
	)
	var err error
	if authURL != "" {
		_, err = client.Login(context.Background())
	}
	return client, err
}

// readResources decodes single or array of input data from YAML.
func readResources(file string) (*services.EventList, error) {
	request := &services.EventList{}
	err := fileutil.LoadFile(file, request)
	return request, err
}

func path(schemaID, uuid string) string {
	return "/" + dashedCase(schemaID) + "/" + uuid
}

func pluralPath(schemaID string) string {
	return "/" + dashedCase(schemaID) + "s"
}
