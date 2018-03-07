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
	Short: "Start Agent service",
	Run: func(cmd *cobra.Command, args []string) {
		startAgent()
	},
}

func startAgent() {
	config := configFile
	if agentConfigFile != "" {
		config = agentConfigFile
	}
	a, err := agent.NewAgentByFile(config)
	if err != nil {
		log.Fatal(err)
	}
	for {
		if err := a.Watch(); err != nil {
			log.Error(err)
		}
	}
}
