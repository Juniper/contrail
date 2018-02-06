package contrailcli

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(CreateCmd)
}

// CreateCmd defines create command.
var CreateCmd = &cobra.Command{
	Use:   "create [FilePath]",
	Short: "Create resources defined in given YAML file",
	Long:  "Use resource format just like in 'schema' command output",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createResources(args)
	},
}

func createResources(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	output, err := a.CreateCLI(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
