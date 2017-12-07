package apisrv

import (
	"fmt"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var testServer *httptest.Server
var server *Server
var testURL string
var client *Client

func TestMain(m *testing.M) {
	common.InitConfig()
	common.SetLogLevel()
	var err error
	server, err = NewServer()
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	testServer = httptest.NewServer(server.Echo)
	if err != nil {
		log.Fatal(err)
	}
	testURL = testServer.URL
	client = NewClient(testURL)

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
	defer func() {
		_, err := server.DB.Exec("SET FOREIGN_KEY_CHECKS=0;")
		if err != nil {
			log.Error(err)
		}
		for _, table := range testData.Cleanup {
			_, err := server.DB.Exec("truncate table " + table)
			if err != nil {
				log.Error(err)
			}
		}
		_, err = server.DB.Exec("SET FOREIGN_KEY_CHECKS=1;")
		if err != nil {
			log.Error(err)
		}
		if r := recover(); r != nil {
			log.Fatal("panic", r)
		}
	}()
	for _, task := range testData.Workflow {
		task.Request.Data = common.YAMLtoJSONCompat(task.Request.Data)
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
	Request *Request    `yaml:"request"`
	Expect  interface{} `yaml:"expect"`
}

type TestScenario struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Cleanup     []string `yaml:"cleanup"`
	Workflow    []*Task  `yaml:"workflow"`
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
	default:
		if t != actual {
			return fmt.Errorf("expected %s but actually we got %s for path %s", t, actual, path)
		}
	}
	return nil
}
