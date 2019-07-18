package apisrv_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func RunTest(t *testing.T, file string) {
	ts, err := integration.LoadTest(file, nil)
	require.NoError(t, err)
	integration.RunCleanTestScenario(t, ts, server)
}

func RunTestTemplate(t *testing.T, file string, context map[string]interface{}) {
	ts, err := integration.LoadTest(file, context)
	require.NoError(t, err)
	integration.RunCleanTestScenario(t, ts, server)
}
