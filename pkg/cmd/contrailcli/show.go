package contrailcli

import (
	"fmt"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show [SchemaID] [UUID]",
	Short: "Show data of specified resource",
	Long:  "Invoke command with empty SchemaID in order to show possible usages",
	Run: func(cmd *cobra.Command, args []string) {
		schemaID := ""
		uuid := ""
		if len(args) >= 2 {
			schemaID = args[0]
			uuid = args[1]
		}

		cli, err := client.NewCLIByViper()
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}

		r, err := cli.ShowResource(schemaID, uuid)
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}

		fmt.Println(r)
	},
}
