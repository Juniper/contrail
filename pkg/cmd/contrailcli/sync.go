package contrailcli

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(SyncCmd)
}

// SyncCmd defines sync command.
var SyncCmd = &cobra.Command{
	Use:   "sync [FilePath]",
	Short: "Update resources or create new ones if they do not already exist",
	Long:  "Use resource format just like in 'schema' command output or 'list' command output",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		synchronizeResources(args)
	},
}

func synchronizeResources(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	output, err := a.SyncCLI(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
