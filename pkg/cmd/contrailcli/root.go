package contrailcli

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/agent"
	"github.com/spf13/cobra"
)

var configFile string

func init() {
	ContrailCLI.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		"Configuration File")
}

// ContrailCLI defines root Contrail CLI command.
var ContrailCLI = &cobra.Command{
	Use:   "contrailcli",
	Short: "Contrail CLI command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func getAuthenticatedAgent(configFilePath string) (*agent.Agent, error) {
	a, err := agent.NewAgentByFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("agent creation failed: %s", err)
	}

	err = a.APIServer.Login()
	if err != nil {
		return nil, fmt.Errorf("agent authentication failed: %s", err)
	}
	return a, nil
}
