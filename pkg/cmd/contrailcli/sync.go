package contrailcli

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/services"
)

func init() {
	ContrailCLI.AddCommand(SyncCmd)
}

// SyncCmd defines sync command.
var SyncCmd = &cobra.Command{
	Use:   "sync [FilePath]",
	Short: "Synchronise resources with data defined in given YAML file",
	Long: `
Sync creates new resource for every not already existing resource
Use resource format just like in 'schema' command output`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := syncResources(args[0])
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println(response)
	},
}

func syncResources(dataPath string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", nil
	}
	request, err := readResources(dataPath)
	if err != nil {
		return "", err
	}
	response := []*services.Event{}
	_, err = client.Create(context.Background(), "/sync", request, &response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	output, err := yaml.Marshal(&services.EventList{
		Events: response})
	if err != nil {
		return "", err
	}
	return string(output), nil
}
