package contrailcli

import (
	"fmt"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/client/baseclient"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set [SchemaID] [UUID] [Properties to update in YAML format]",
	Short: "Set properties of specified resource",
	Long:  "Invoke command with empty SchemaID in order to show possible usages",
	Run: func(cmd *cobra.Command, args []string) {
		schemaID := ""
		uuid := ""
		yaml := ""
		if len(args) >= 3 {
			schemaID = args[0]
			uuid = args[1]
			yaml = args[2]
		}

		c, err := baseclient.NewCLIByViper()
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}

		r, err := c.SetResourceParameter(schemaID, uuid, yaml)
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}

		fmt.Println(r)
	},
}
