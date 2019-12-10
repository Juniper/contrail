package contrailschema

import (
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/schema"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type templateOption struct {
	SchemasDir         string
	AddonsDir          string
	TemplateConfigPath string
	SchemaOutputPath   string
	OpenAPIOutputPath  string
	SkipMissingRefs    bool
	NoRegenerate       bool
	Verbose            bool
}

var option = templateOption{}

func init() {
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	ContrailSchema.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&option.SchemasDir, "schemas", "s", "", "Schema Directory")
	generateCmd.Flags().StringVarP(&option.AddonsDir, "addons", "a", "", "Addons Directory")
	generateCmd.Flags().StringVarP(&option.TemplateConfigPath, "templates", "t", "", "Template Configuration")
	generateCmd.Flags().StringVarP(&option.SchemaOutputPath, "schema-output", "", "", "Schema Output path")
	generateCmd.Flags().StringVarP(&option.OpenAPIOutputPath, "openapi-output", "", "", "OpenAPI Output path")
	generateCmd.Flags().BoolVarP(
		&option.NoRegenerate, "no-regenerate", "n", false,
		"Avoid regenerating file if it is newer that its source schema and template files",
	)
	generateCmd.Flags().BoolVarP(
		&option.SkipMissingRefs, "skip-missing-refs", "", false,
		"Skip references that are missing instead of failing",
	)
	generateCmd.Flags().BoolVarP(&option.Verbose, "verbose", "v", false, "Enable debug logging")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate code from schema",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := logutil.Configure(getLogLevel()); err != nil {
			terminate(err)
		}

		if err := generateCode(); err != nil {
			terminate(err)
		}
	},
}

func getLogLevel() string {
	level := viper.GetString("log_level")
	if option.Verbose {
		level = logrus.DebugLevel.String()
	}
	if level == "" {
		level = logrus.InfoLevel.String()
	}
	return level
}

func terminate(err error) {
	if logrus.GetLevel() == logrus.DebugLevel {
		logutil.FatalWithStackTrace(err)
	} else {
		logrus.Fatal(err)
	}
}

func generateCode() error {
	logrus.Info("Generating source code from schema")
	api, err := schema.MakeAPI(
		strings.Split(option.SchemasDir, ","),
		strings.Split(option.AddonsDir, ","),
		option.SkipMissingRefs,
	)
	if err != nil {
		return err
	}

	if !option.NoRegenerate {
		api.Timestamp = time.Time{}
	}

	tc, err := schema.LoadTemplateConfig(option.TemplateConfigPath)
	if err != nil {
		return err
	}

	if err = schema.GenerateFiles(api, tc); err != nil {
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
