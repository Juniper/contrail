package contrailcli

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/flosch/pongo2"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

const serverSchemaRoot = "/public/"
const serverSchemaFile = "schema.json"

const schemaTemplate = `
{% for schema in schemas %}
# {{ schema.Title }} {{ schema.Description }}
- kind: {{ schema.ID }}
  data: {% for key, value in schema.JSONSchema.Properties %}
    {{ key }}: {{ value.Default }} # {{ value.Title }} ({{ value.Type }}) {% endfor %}
{% endfor %}`

func init() {
	ContrailCLI.AddCommand(SchemaCmd)
}

// SchemaCmd defines schema command.
var SchemaCmd = &cobra.Command{
	Use:   "schema [SchemaID]",
	Short: "Show schema for specified resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		schemaID := ""
		if len(args[0]) > 0 {
			schemaID = args[0]
		}
		output, err := showSchema(schemaID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	},
}

func showSchema(schemaID string) (string, error) {
	return showHelp(schemaID, schemaTemplate)
}

func showHelp(schemaID string, template string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", nil
	}
	serverSchema := filepath.Join(serverSchemaRoot, serverSchemaFile)
	api, err := fetchServerAPI(client, serverSchema)
	if err != nil {
		return "", err
	}
	schemas := api.Schemas
	if schemaID != "" {
		s := api.SchemaByID(schemaID)
		if s == nil {
			return "", fmt.Errorf("schema %s not found", schemaID)
		}
		schemas = []*schema.Schema{s}
	}
	tpl, err := pongo2.FromString(template)
	if err != nil {
		return "", err
	}
	o, err := tpl.Execute(pongo2.Context{"schemas": schemas})
	if err != nil {
		return "", err
	}
	return o, nil
}

func fetchServerAPI(client *apisrv.Client, serverSchema string) (*schema.API, error) {
	var api schema.API
	for {
		_, err := client.Read(serverSchema, &api)
		if err == nil {
			break
		}
		log.Warn("failed to connect server %d. reconnecting...", err)
		time.Sleep(time.Second)
	}
	return &api, nil
}
