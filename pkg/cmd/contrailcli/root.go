package contrailcli

import (
	"context"
	"strings"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	ContrailCLI.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Configuration file")
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

// Package-common functions below

func newHTTPClient() (*client.HTTP, error) {
	c := client.NewHTTP(
		viper.GetString("client.endpoint"),
		viper.GetString("keystone.authurl"),
		viper.GetString("client.id"),
		viper.GetString("client.password"),
		viper.GetBool("insecure"),
		keystone.NewScope(
			viper.GetString("client.domain_id"),
			viper.GetString("client.domain_name"),
			viper.GetString("client.project_id"),
			viper.GetString("client.project_name"),
		),
	)

	var err error
	if viper.GetString("keystone.authurl") != "" {
		_, err = c.Login(context.Background())
	}
	return c, err
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

func dashedCase(schemaID string) string {
	return strings.Replace(schemaID, "_", "-", -1)
}
