package contrail

import (
	"github.com/Juniper/contrail/pkg/agent"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	Contrail.AddCommand(agentCmd)
}

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Start Agent",
	Run: func(cmd *cobra.Command, args []string) {
		startAgent()
	},
}

func startAgent() {
	a, err := agent.NewAgentByFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Watch(); err != nil {
		log.Fatal(err)
	}
}
