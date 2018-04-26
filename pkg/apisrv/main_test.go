package apisrv

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
)

var testServer *httptest.Server
var server *Server

func TestMain(m *testing.M) {
	err := common.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	common.SetLogLevel()
	dbConfig := viper.GetStringMap("test_database")
	for _, iConfig := range dbConfig {
		config := common.InterfaceToInterfaceMap(iConfig)
		viper.Set("database.type", config["type"])
		viper.Set("database.connection", config["connection"])
		viper.Set("database.dialect", config["dialect"])
		RunTestForDB(m)
	}
}

func RunTestForDB(m *testing.M) {
	var err error
	server, err = NewServer()
	if err != nil {
		log.Fatal(err)
	}
	testServer = httptest.NewUnstartedServer(server.Echo)
	testServer.TLS = new(tls.Config)
	testServer.TLS.NextProtos = append(testServer.TLS.NextProtos, "h2")
	testServer.StartTLS()
	defer testServer.Close()

	viper.Set("keystone.authurl", testServer.URL+"/v3")
	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer logFatalIfErr(server.Close)
	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	if code != 0 {
		os.Exit(code)
	}
}

func logFatalIfErr(f func() error) {
	if err := f(); err != nil {
		log.Fatal(err)
	}
}

func RunTest(t *testing.T, file string) {
	testScenario, err := LoadTest(file)
	assert.NoError(t, err, "failed to load test data")
	RunTestScenario(t, testScenario)
}

func RunTestScenario(t *testing.T, testScenario *TestScenario) {
	clients := map[string]*Client{}

	for key, client := range testScenario.Clients {
		//Rewrite endpoint for test server
		client.Endpoint = testServer.URL
		client.AuthURL = testServer.URL + "/v3"
		client.InSecure = true
		client.Init()

		clients[key] = client

		err := clients[key].Login()
		assert.NoError(t, err, "client failed to login")
	}
	for _, cleanTask := range testScenario.Cleanup {
		fmt.Println(cleanTask)
		clientID := cleanTask["client"]
		if clientID == "" {
			clientID = "default"
		}
		client := clients[clientID]
		// delete existing resources.
		log.Debug(cleanTask["path"])
		_, err := client.Delete(cleanTask["path"], nil) // nolint
		if err != nil {
			log.Debug(err)
		}
	}
	for _, task := range testScenario.Workflow {
		log.Debug("[Step] ", task.Name)
		task.Request.Data = common.YAMLtoJSONCompat(task.Request.Data)
		clientID := "default"
		if task.Client != "" {
			clientID = task.Client
		}
		client := clients[clientID]
		_, err := client.DoRequest(task.Request)
		assert.NoError(t, err, fmt.Sprintf("task %v failed", task))
		task.Expect = common.YAMLtoJSONCompat(task.Expect)
		ok := common.AssertEqual(t, task.Expect, task.Request.Output, fmt.Sprintf("task %v failed", task))
		if !ok {
			break
		}
	}
}

func GetTestFromTemplate(t *testing.T, templateFile string, ctx map[string]interface{}) string {
	// create test data yaml from the template
	template, err := pongo2.FromFile(templateFile)
	assert.NoError(t, err, "failed to read test data template")

	content, err := template.ExecuteBytes(ctx)
	assert.NoError(t, err, "failed to apply test data template")

	fileName := filepath.Base(templateFile)
	var extension = filepath.Ext(fileName)
	var prefix = fileName[0 : len(fileName)-len(extension)]
	tmpfile, err := ioutil.TempFile("", prefix)
	assert.NoError(t, err, "failed to create test data tempfile")

	_, err = tmpfile.Write(content)
	assert.NoError(t, err, "failed to write test data to tempfile")

	err = tmpfile.Close()
	assert.NoError(t, err, "failed to close tempfile")
	testFile := tmpfile.Name() + ".yml"
	err = os.Rename(tmpfile.Name(), testFile)
	assert.NoError(t, err, "failed to rename test data file to yml file")
	return testFile
}

func LoadTest(file string) (*TestScenario, error) {
	var testScenario TestScenario
	err := LoadTestScenario(&testScenario, file)
	return &testScenario, err

}

func LoadTestScenario(testScenario *TestScenario, file string) error {
	err := common.LoadFile(file, &testScenario)
	return err
}
