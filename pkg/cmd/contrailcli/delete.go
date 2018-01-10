package contrailcli

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(DeleteCmd)
}

// DeleteCmd defines delete command.
var DeleteCmd = &cobra.Command{
	Use:   "delete [FilePath]",
	Short: "Delete resources defined in given file",
	Long:  "Use resource format just like in 'schema' command output or 'list' command output",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteResources(args)
	},
}

func deleteResources(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = a.DeleteCLI(args[0])
	if err != nil {
		log.Fatal(err)
	}
}
