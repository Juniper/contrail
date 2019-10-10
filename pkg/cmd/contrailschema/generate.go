package contrailschema

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/schema"
)

type templateOption struct {
	SchemasDir        string
	AddonsDir         string
	TemplateConfPath  string
	SchemaOutputPath  string
	OpenAPIOutputPath string
	SkipMissingRefs   bool
	NoRegenerate      bool
}

var option = templateOption{}

func init() {
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	ContrailSchema.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&option.SchemasDir, "schemas", "s", "", "Schema Directory")
	generateCmd.Flags().StringVarP(&option.AddonsDir, "addons", "a", "", "Addons Directory")
	generateCmd.Flags().StringVarP(&option.TemplateConfPath, "templates", "t", "", "Template Configuration")
	generateCmd.Flags().StringVarP(&option.SchemaOutputPath, "schema-output", "", "", "Schema Output path")
	generateCmd.Flags().StringVarP(&option.OpenAPIOutputPath, "openapi-output", "", "", "OpenAPI Output path")
	generateCmd.Flags().BoolVarP(
		&option.NoRegenerate, "no-regenerate", "n", false,
		"Avoid regenerating file if it is newer that its source schema and template files",
	)
	generateCmd.Flags().BoolVarP(
		&option.SkipMissingRefs, "skip-missing-refs", "", false,
		"If this flag is enabled contrailschema skips the references that are missing instead of failing",
	)
}

func generateCode() error {
	logrus.Info("Generating source code from schema")
	api, err := schema.MakeAPI(strings.Split(option.SchemasDir, ","), strings.Split(option.AddonsDir, ","), option.SkipMissingRefs)
	if err != nil {
		return err
	}

	if !option.NoRegenerate {
		api.Timestamp = time.Time{}
	}

	templateConf, err := schema.LoadTemplates(option.TemplateConfPath)
	if err != nil {
		return err
	}
	if err = schema.ApplyTemplates(api, templateConf); err != nil {
		return err
	}

	if option.SchemaOutputPath != "" {
		if err = fileutil.SaveFile(option.SchemaOutputPath, api); err != nil {
			return err
		}
	}

	if option.OpenAPIOutputPath != "" {
		openapi, err := api.ToOpenAPI()
		if err != nil {
			return err
		}

		if err = fileutil.SaveFile(option.OpenAPIOutputPath, openapi); err != nil {
			return err
		}
	}
	return nil
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate code from schema",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if level := viper.GetString("log_level"); level != "" {
			if err := logutil.Configure(level); err != nil {
				logutil.FatalWithStackTrace(err)
			}
		}
		if err := generateCode(); err != nil {
			if logrus.GetLevel() == logrus.DebugLevel {
				logutil.FatalWithStackTrace(err)
			} else {
				logrus.Fatal(err)
			}
		}
	},
}
