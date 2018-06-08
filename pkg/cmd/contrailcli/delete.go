package contrailcli

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(DeleteCmd)
}

// DeleteCmd defines delete command.
var DeleteCmd = &cobra.Command{
	Use:   "delete [FilePath]",
	Short: "Delete resources specified in given YAML file",
	Long:  "Use resource format just like in 'schema' command output",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := deleteResources(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
	},
}

func deleteResources(dataPath string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", nil
	}
	request, err := readResources(dataPath)
	if err != nil {
		return "", nil
	}
	for i := len(request.Events) - 1; i >= 0; i-- {
		event := request.Events[i]
		resource := event.GetResource()
		uuid := resource.GetUUID()
		if err != nil {
			return "", err
		}
		var output interface{}
		response, err := client.Delete(path(resource.Kind(), uuid), &output)
		if response.StatusCode != http.StatusNotFound && err != nil {
			return "", err
		}
	}
	return "", nil
}
