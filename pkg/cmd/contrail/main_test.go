package contrail_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func runTest(t *testing.T, name string) {
	testScenario, err := integration.LoadTest(fmt.Sprintf("./tests/%s.yml", common.CamelToSnake(name)), nil)
	assert.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, testScenario, server)
}
