package contrailcli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
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
			pkglog.FatalWithStackTrace(err)
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
	_, err = client.Read(context.Background(), path(schemaID, uuid), &response)
	if err != nil {
		return "", err
	}
	data, _ := response[dashedCase(schemaID)].(map[string]interface{}) //nolint: errcheck
	event, err := services.NewEvent(&services.EventOption{
		Kind: schemaID,
		Data: data,
	})
	if err != nil {
		return "", err
	}
	eventList := &services.EventList{Events: []*services.Event{event}}
	output, err := yaml.Marshal(eventList)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
