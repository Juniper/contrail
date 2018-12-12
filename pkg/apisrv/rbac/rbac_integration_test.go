package rbac_test

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestRBAC(t *testing.T) {
	viper.SetConfigType("yml")
	viper.SetConfigName("test_config")
	viper.AddConfigPath("../../../sample")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:     basedb.DriverPostgreSQL,
		RepoRootPath: "../../..",
		EnableRBAC:   true,
	})
	defer func() {
		assert.NoError(t, s.Close())
	}()

	testScenario, err := integration.LoadTest("./test_data/test_rbac.yml", nil)
	assert.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, testScenario, s)
}
