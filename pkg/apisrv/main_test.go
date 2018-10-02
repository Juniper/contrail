package apisrv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

var Server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &Server)
}

func RunTest(t *testing.T, file string) {
	testScenario, err := integration.LoadTest(file, nil)
	assert.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, testScenario, Server)
}
