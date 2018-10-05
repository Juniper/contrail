package contrail_test

import (
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func RunTest(t *testing.T, name string) {
	integration.AddKeystoneProjectAndUser(server.APIServer, name)
	testScenario, err := integration.LoadTest(fmt.Sprintf("./tests/%s.yml", common.CamelToSnake(name)), nil)
	assert.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, testScenario, server)
}
