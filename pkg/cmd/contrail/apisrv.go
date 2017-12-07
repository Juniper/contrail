package contrail

import (
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var conf string

func init() {
	cobra.OnInitialize()
	Cmd.AddCommand(apiServerCmd)
}

var apiServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start API Server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := apisrv.NewServer()
		if err != nil {
			log.Fatal(err)
		}
		err = server.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}
