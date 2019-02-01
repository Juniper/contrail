package contrailcli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/log"
)

func init() {
	ContrailCLI.AddCommand(SetCmd)
}

const setHelpTemplate = `Set command possible usages:
{% for schema in schemas %}contrail set {{ schema.ID }} $UUID $YAML
{% endfor %}`

// SetCmd defines set command.
var SetCmd = &cobra.Command{
	Use:   "set [SchemaID] [UUID] [Properties to update in YAML format]",
	Short: "Set properties of specified resource",
	Long:  "Invoke command with empty SchemaID in order to show possible usages",
	Run: func(cmd *cobra.Command, args []string) {
		schemaID := ""
		uuid := ""
		yaml := ""
		if len(args) >= 3 {
			schemaID = args[0]
			uuid = args[1]
			yaml = args[2]
		}
		output, err := setResourceParameter(schemaID, uuid, yaml)
		if err != nil {
			log.FatalWithStackTrace(err)
		}
		fmt.Println(output)
	},
}

func setResourceParameter(schemaID, uuid, yamlString string) (string, error) {
	if schemaID == "" || uuid == "" {
		return showHelp(schemaID, setHelpTemplate)
	}
	client, err := getClient()
	if err != nil {
		return "", nil
	}

	var data map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlString), &data)
	if err != nil {
		return "", err
	}

	data["uuid"] = uuid
	_, err = client.Update(context.Background(), path(schemaID, uuid), map[string]interface{}{
		dashedCase(schemaID): fileutil.YAMLtoJSONCompat(data),
	}, nil)
	if err != nil {
		return "", err
	}
	return showResource(schemaID, uuid)
}
