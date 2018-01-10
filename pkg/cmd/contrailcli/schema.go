package contrailcli

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(SchemaCmd)
}

// SchemaCmd defines schema command.
var SchemaCmd = &cobra.Command{
	Use:   "schema [SchemaID]",
	Short: "Show schema data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		schema(args)
	},
}

func schema(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	output, err := a.SchemaCLI(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
