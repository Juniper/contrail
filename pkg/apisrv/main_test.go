package apisrv

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
)

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
	server, testServer := LaunchTestAPIServer()
	defer testServer.Close()
	defer LogFatalIfErr(server.Close)
	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	if code != 0 {
		os.Exit(code)
	}
}

func RunTest(t *testing.T, file string) {
	testScenario, err := LoadTest(file, nil)
	assert.NoError(t, err, "failed to load test data")
	RunTestScenario(t, testScenario)
}
