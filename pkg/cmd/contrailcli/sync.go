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
	Short: "Synchronise resources with data defined in given YAML file",
	Long: `
Sync creates new resource for every not already existing resource
Use resource format just like in 'schema' command output`,
	Args: cobra.ExactArgs(1),
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
