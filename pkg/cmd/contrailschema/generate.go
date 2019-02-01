package contrailschema

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Juniper/contrail/pkg/logutil"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/schema"
)

var option = &schema.TemplateOption{}

func init() {
	ContrailSchema.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&option.SchemasDir, "schemas", "s", "", "Schema Directory")
	generateCmd.Flags().StringVarP(&option.TemplateConfPath, "templates", "t", "", "Template Configuration")
	generateCmd.Flags().StringVarP(&option.OutputDir, "output-dir", "", "./", "output dir")
	generateCmd.Flags().StringVarP(&option.SchemaOutputPath, "schema-output", "", "", "Schema Output path")
	generateCmd.Flags().StringVarP(&option.OpenapiOutputPath, "openapi-output", "", "", "OpenAPI Output path")
}

func generateCode() {
	logrus.Info("Generating source code from schema")
	api, err := schema.MakeAPI(strings.Split(option.SchemasDir, ","), "overrides")
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}

	templateConf, err := schema.LoadTemplates(option.TemplateConfPath)
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	if err = schema.ApplyTemplates(api, templateConf, option); err != nil {
		logutil.FatalWithStackTrace(err)
	}

	if err = fileutil.SaveFile(option.SchemaOutputPath, api); err != nil {
		logutil.FatalWithStackTrace(err)
	}

	openapi, err := api.ToOpenAPI()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}

	if err = fileutil.SaveFile(option.OpenapiOutputPath, openapi); err != nil {
		logutil.FatalWithStackTrace(err)
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
