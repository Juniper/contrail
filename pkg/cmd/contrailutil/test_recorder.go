package contrailutil

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/common"
)

var inputPath string
var outputPath string
var variablePath string
var endpoint string
var authURL string

func init() {
	ContrailUtil.AddCommand(recordTestCmd)
	recordTestCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input test scenario path")
	recordTestCmd.Flags().StringVarP(&variablePath, "vars", "v", "", "test variables")
	recordTestCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output test scenario path")
	recordTestCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "Endpoint")
	recordTestCmd.Flags().StringVarP(&authURL, "auth_url", "a", "", "AuthURL")
}

func assertError(err error, message string) {
	if err != nil {
		log.Fatalf("%s (%s)", message, err)
	}
}

func recordTest() {
	ctx := context.Background()
	log.Info("Recording API beheivior")
	var vars map[string]interface{}
	if variablePath != "" {
		err := common.LoadFile(variablePath, &vars)
		if err != nil {
			log.Fatal(err)
		}
	}

	testScenario, err := integration.LoadTest(inputPath, vars)
	assertError(err, "failed to load test scenario")
	clients := map[string]*client.HTTP{}

	for key, client := range testScenario.Clients {
		//Rewrite endpoint for test server
		client.Endpoint = endpoint
		client.AuthURL = authURL
		client.InSecure = true
		client.Init()

		clients[key] = client

		err = clients[key].Login(ctx)
		assertError(err, "client can't login")
	}

	for _, task := range testScenario.Workflow {
		log.Debug("[Step] ", task.Name)
		task.Request.Data = common.YAMLtoJSONCompat(task.Request.Data)
		clientID := "default"
		if task.Client != "" {
			clientID = task.Client
		}
		client := clients[clientID]
		_, err = client.DoRequest(ctx, task.Request)
		assertError(err, fmt.Sprintf("Task %s failed", task.Name))
		task.Expect = task.Request.Output
		task.Request.Output = nil
	}

	err = common.SaveFile(outputPath, testScenario)
	if err != nil {
		log.Fatal(err)
	}
}

var recordTestCmd = &cobra.Command{
	Use:   "record_test",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		recordTest()
	},
}
