package contrailcli

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(RmCmd)
}

// RmCmd defines rm command.
var RmCmd = &cobra.Command{
	Use:   "rm [SchemaID] [ID]",
	Short: "Remove a resource",
	Long:  "Invoke command with empty SchemaID in order to show available schemas",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		deleteResource(args)
	},
}

func deleteResource(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	output, err := a.RemoveCLI(args[0], args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
