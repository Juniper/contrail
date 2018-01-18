package apisrv

import (
	"fmt"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var testServer *httptest.Server
var server *Server
var testURL string

func TestMain(m *testing.M) {
	common.InitConfig()
	common.SetLogLevel()
	var err error
	server, err = NewServer()
	if err != nil {
		log.Fatal(err)
	}

	testServer = httptest.NewServer(server.Echo)
	if err != nil {
		log.Fatal(err)
	}
	viper.Set("keystone.authurl", testServer.URL+"/v3")
	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	os.Exit(code)
}

func RunTest(file string) error {
	testData, err := LoadTest(file)
	if err != nil {
		return errors.Wrap(err, "failed to load test data")
	}
	for _, table := range testData.Tables {
		common.UseTable(server.DB, table)
		defer common.ClearTable(server.DB, table)
	}
	clients := map[string]*Client{}
	for key, client := range testData.Clients {
		//Rewrite endpoint for test server
		client.Endpoint = testServer.URL
		client.AuthURL = testServer.URL + "/v3"
		clients[key] = client

		client.Init()
		err := clients[key].Login()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("client %s failed to login", key))
		}
	}
	for _, task := range testData.Workflow {
		log.Debug("[Step] ", task.Name)
		task.Request.Data = common.YAMLtoJSONCompat(task.Request.Data)
		clientID := "default"
		if task.Client != "" {
			clientID = task.Client
		}
		client := clients[clientID]
		_, err := client.DoRequest(task.Request)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("task %v failed", task))
		}
		task.Expect = common.YAMLtoJSONCompat(task.Expect)
		err = checkDiff("", task.Expect, task.Request.Output)
		if err != nil {
			log.WithFields(
				log.Fields{
					"scenario": testData.Name,
					"step":     task.Name,
					"expected": task.Expect,
					"actual":   task.Request.Output,
					"err":      err,
				}).Debug("output mismatch")
			return errors.Wrap(err, fmt.Sprintf("task %v failed", task))
		}
	}
	return nil
}

func LoadTest(file string) (*TestScenario, error) {
	var testScenario TestScenario
	err := common.LoadFile(file, &testScenario)
	return &testScenario, err
}

type Task struct {
	Name    string      `yaml:"name"`
	Client  string      `yaml:"client"`
	Request *Request    `yaml:"request"`
	Expect  interface{} `yaml:"expect"`
}

type TestScenario struct {
	Name        string             `yaml:"name"`
	Description string             `yaml:"description"`
	Tables      []string           `yaml:"tables"`
	Clients     map[string]*Client `yaml:"clients"`
	Workflow    []*Task            `yaml:"workflow"`
}

func checkDiff(path string, expected, actual interface{}) error {
	if expected == nil {
		return nil
	}
	switch t := expected.(type) {
	case map[string]interface{}:
		actualMap, ok := actual.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
		for key, value := range t {
			err := checkDiff(path+"."+key, value, actualMap[key])
			if err != nil {
				return err
			}
		}
	case []interface{}:
		actualList, ok := actual.([]interface{})
		if !ok {
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
		if len(t) != len(actualList) {
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
		for index, value := range t {
			err := checkDiff(path+"."+strconv.Itoa(index), value, actualList[index])
			if err != nil {
				return err
			}
		}
	case int:
		if float64(t) != actual {
			fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
	default:
		if t != actual {
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
	}
	return nil
}
