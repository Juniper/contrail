package contrailcli

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

func init() {
	ContrailCLI.AddCommand(ShowCmd)
}

const showHelpTemplate = `Show command possible usages:
{% for schema in schemas %}contrail show {{ schema.ID }} $UUID
{% endfor %}`

// ShowCmd defines show command.
var ShowCmd = &cobra.Command{
	Use:   "show [SchemaID] [UUID]",
	Short: "Show data of specified resource",
	Long:  "Invoke command with empty SchemaID in order to show possible usages",
	Run: func(cmd *cobra.Command, args []string) {
		schemaID := ""
		uuid := ""
		if len(args) >= 2 {
			schemaID = args[0]
			uuid = args[1]
		}
		output, err := showResource(schemaID, uuid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	},
}

func showResource(schemaID, uuid string) (string, error) {
	if schemaID == "" || uuid == "" {
		return showHelp(schemaID, showHelpTemplate)
	}
	client, err := getClient()
	if err != nil {
		return "", nil
	}
	var response map[string]interface{}
	_, err = client.Read(path(schemaID, uuid), &response)
	if err != nil {
		return "", err
	}
	data, _ := response[dashedCase(schemaID)].(map[string]interface{})
	eventList := &services.EventList{
		Events: []*services.Event{
			services.NewEvent(&services.EventOption{
				Kind: schemaID,
				Data: data,
			}),
		},
	}
	output, err := yaml.Marshal(eventList)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
