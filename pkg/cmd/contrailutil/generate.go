package contrailutil

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/Juniper/contrail/pkg/schema"
	log "github.com/sirupsen/logrus"
)

var schemasDir string
var templateConfPath string
var schemaOutputPath string
var openapiOutputPath string

func init() {
	ContrailUtil.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&schemasDir, "schemas", "s", "", "Schema Directory")
	generateCmd.Flags().StringVarP(&templateConfPath, "templates", "t", "", "Template Configuration")
	generateCmd.Flags().StringVarP(&schemaOutputPath, "schema-output", "", "", "Schema Output path")
	generateCmd.Flags().StringVarP(&openapiOutputPath, "openapi-output", "", "", "OpenAPI Output path")
}

func generateCode() {
	log.Info("Generating source code from schema")
	api, err := schema.MakeAPI(schemasDir)
	if err != nil {
		log.Fatal(err)
	}

	templateConf, err := schema.LoadTemplates(templateConfPath)
	if err != nil {
		log.Fatal(err)
	}
	if err = schema.ApplyTemplates(api, filepath.Dir(templateConfPath), templateConf); err != nil {
		log.Fatal(err)
	}

	if err = schema.SaveFile(schemaOutputPath, api); err != nil {
		log.Fatal(err)
	}

	openapi, err := api.ToOpenAPI()
	if err != nil {
		log.Fatal(err)
	}

	if err = schema.SaveFile(openapiOutputPath, openapi); err != nil {
		log.Fatal(err)
	}
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate code from schema",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		generateCode()
	},
}
