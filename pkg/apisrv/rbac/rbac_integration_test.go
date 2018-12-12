package rbac_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestRBAC(t *testing.T) {
	integration.WithTestDBs(func(dbType string) {
		s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
			DBDriver:     dbType,
			RepoRootPath: "../../..",
			EnableRBAC:   true,
		})
		defer func() {
			assert.NoError(t, s.Close())
		}()

		testScenario, err := integration.LoadTest("./test_data/test_rbac.yml", nil)
		assert.NoError(t, err, "failed to load test data")
		integration.RunCleanTestScenario(t, testScenario, s)
	})
}
