package contrailcli

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/Juniper/contrail/pkg/logutil"
)

func init() {
	ContrailCLI.AddCommand(rmCmd)
}

const removeHelpTemplate = `Remove command possible usages:
{% for schema in schemas %}contrail rm {{ schema.ID }} $UUID
{% endfor %}`

// rmCmd defines rm command.
var rmCmd = &cobra.Command{
	Use:   "rm [SchemaID] [UUID]",
	Short: "Remove a resource with specified UUID",
	Long:  "Invoke command with empty SchemaID in order to show possible usages",
	Run: func(cmd *cobra.Command, args []string) {
		schemaID, uuid := "", ""
		if len(args) >= 2 {
			schemaID, uuid = args[0], args[1]
		}
		output, err := deleteResource(schemaID, uuid)
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}
		fmt.Println(output)
	},
}

func deleteResource(schemaID, uuid string) (string, error) {
	if schemaID == "" || uuid == "" {
		return showHelp(schemaID, removeHelpTemplate)
	}
	client, err := newHTTPClient()
	if err != nil {
		return "", nil
	}
	response, err := client.Delete(context.Background(), path(schemaID, uuid), nil)
	if response.StatusCode != http.StatusNotFound && err != nil {
		return "", err
	}
	return "", nil
}
