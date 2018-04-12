package contrailcli

import (
	"strings"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/services"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
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
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't read config:", err)
	}

}

func getClient() (*apisrv.Client, error) {
	authURL := viper.GetString("keystone.auth_url")
	client := apisrv.NewClient(
		viper.GetString("client.endpoint"),
		authURL,
		viper.GetString("client.id"),
		viper.GetString("client.password"),
		viper.GetBool("insecure"),
		&keystone.Scope{
			Project: &keystone.Project{
				ID: viper.GetString("client.project_id"),
			},
		},
	)
	var err error
	if authURL != "" {
		err = client.Login()
	}
	return client, err
}

// readResources decodes single or array of input data from YAML.
func readResources(file string) (*services.RESTSyncRequest, error) {
	request := &services.RESTSyncRequest{}
	err := common.LoadFile(file, request)
	for _, resource := range request.Resources {
		resource.Data = common.YAMLtoJSONCompat(resource.Data)
	}
	return request, err
}

func path(schemaID, uuid string) string {
	return "/" + dashedCase(schemaID) + "/" + uuid
}

func pluralPath(schemaID string) string {
	return "/" + dashedCase(schemaID) + "s"
}
