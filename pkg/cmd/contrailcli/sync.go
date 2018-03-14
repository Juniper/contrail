package contrailcli

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
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
			log.Fatal(err)
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
	response := &services.RESTSyncRequest{}
	_, err = client.Create("/sync", request, &response.Resources)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	output, err := yaml.Marshal(&response)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
