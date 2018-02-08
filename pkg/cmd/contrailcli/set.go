package contrailcli

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(SetCmd)
}

// SetCmd defines set command.
var SetCmd = &cobra.Command{
	Use:   "set [SchemaID] [UUID] [Properties to update in YAML format]",
	Short: "Set properties of specified resource",
	Long:  "Invoke command with empty SchemaID in order to show possible usages",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		setResourceParameter(args)
	},
}

func setResourceParameter(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	output, err := a.SetCLI(args[0], args[1], args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
