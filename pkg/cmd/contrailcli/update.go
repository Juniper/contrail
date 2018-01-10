package contrailcli

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailCLI.AddCommand(UpdateCmd)
}

// UpdateCmd defines update command.
var UpdateCmd = &cobra.Command{
	Use:   "update [FilePath]",
	Short: "Update resources defined in given file",
	Long:  "Use resource format just like in 'schema' command output or 'list' command output",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		updateResources(args)
	},
}

func updateResources(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	output, err := a.UpdateCLI(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
