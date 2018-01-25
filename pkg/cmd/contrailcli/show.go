package contrailcli

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(ShowCmd)
}

// ShowCmd defines show command.
var ShowCmd = &cobra.Command{
	Use:   "show [SchemaID] [ID]",
	Short: "Show resource data",
	Long:  "Invoke command with empty SchemaID in order to show available schemas",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		showResource(args)
	},
}

func showResource(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	output, err := a.ShowCLI(args[0], args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
