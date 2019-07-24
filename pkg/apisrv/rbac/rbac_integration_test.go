package rbac_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRBAC(t *testing.T) {
	integration.WithTestDBs(func() {
		s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
			RepoRootPath: "../../..",
			EnableRBAC:   true,
		})
		defer func() {
			assert.NoError(t, s.Close())
		}()

		testScenario, err := integration.LoadTest("./test_data/test_rbac.yml", nil)
		require.NoError(t, err, "failed to load test data")
		integration.RunCleanTestScenario(t, testScenario, s)
	})
}
