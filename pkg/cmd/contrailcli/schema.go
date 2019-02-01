package contrailcli

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/schema"
)

const serverSchemaFile = "schema.json"

const schemaTemplate = `
{% for schema in schemas %}
# {{ schema.Title }} {{ schema.Description }}
- kind: {{ schema.ID }}
  data: {% for key, value in schema.JSONSchema.Properties %}
    {{ key }}: {{ value.Default }} # {{ value.Title }} ({{ value.Type }}) {% endfor %}
{% endfor %}`

const retryMax = 5

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
			logutil.FatalWithStackTrace(err)
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
	serverSchemaRoot := viper.GetString("client.schema_root")
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

func fetchServerAPI(client *client.HTTP, serverSchema string) (*schema.API, error) {
	var api schema.API
	ctx := context.Background()
	for i := 0; i < retryMax; i++ {
		_, err := client.Read(ctx, serverSchema, &api)
		if err == nil {
			break
		}
		logrus.WithError(err).Warn("Failed to connect API Server - reconnecting")
		time.Sleep(time.Second)
	}
	return &api, nil
}
