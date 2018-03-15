package contrail

import (
	"github.com/Juniper/contrail/pkg/compilation"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	Contrail.AddCommand(compilationCmd)
}

var compilationCmd = &cobra.Command{
	Use:   "compilation",
	Short: "Start Intent Compilation service",
	Run: func(cmd *cobra.Command, args []string) {
		startCompilationService()
	},
}

func startCompilationService() {
	server, err := compilation.NewIntentCompilationService()
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Init(configFile); err != nil {
		log.Fatal(err)
	}

	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}
