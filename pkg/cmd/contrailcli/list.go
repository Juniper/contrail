package contrailcli

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var filter string
var offset string
var limit string

func init() {
	ContrailCLI.AddCommand(ListCmd)

	ListCmd.Flags().StringVarP(&filter, "filter", "f", "",
		"Comma separated filter parameters (e.g. name=john,status=active)")
	ListCmd.Flags().StringVarP(&offset, "offset", "o", "0", "Start offset of output")
	ListCmd.Flags().StringVarP(&limit, "limit", "l", "50", "Number of elements in output")
}

// ListCmd defines list command.
var ListCmd = &cobra.Command{
	Use:   "list [SchemaID]",
	Short: "List resources data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listResources(args)
	},
}

func listResources(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	output, err := a.ListCLI(args[0], filter, offset, limit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
