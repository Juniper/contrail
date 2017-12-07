package cmd

import (
	"path/filepath"

	"github.com/Juniper/contrail/pkg/utils"
	"github.com/ngaut/log"
	"github.com/spf13/cobra"
)

var schemasDir string
var templateConfPath string

func init() {
	cobra.OnInitialize()
	ContrailUtilCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&schemasDir, "schemas", "s", "", "Schema Directory")
	generateCmd.Flags().StringVarP(&templateConfPath, "templates", "t", "", "Template Configuraion")
}

func generateCode() {
	log.Info("Generating source code from schema")
	api, err := utils.MakeAPI(schemasDir)
	if err != nil {
		log.Fatal(err)
	}
	templateConf, err := utils.LoadTemplates(templateConfPath)
	if err != nil {
		log.Fatal(err)
	}
	err = utils.ApplyTemplates(api, filepath.Dir(templateConfPath), templateConf)
	if err != nil {
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
