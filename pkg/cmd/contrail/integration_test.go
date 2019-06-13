package contrail_test

import (
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
)

func TestFQNameCleanup(t *testing.T) {
	cleanup := runDirtyTest(t, t.Name())
	cleanup = runDirtyTest(t, t.Name())
	cleanup()
}

func runDirtyTest(t *testing.T, name string) func() {
	testScenario, err := integration.LoadTest(fmt.Sprintf("./tests/%s.yml", format.CamelToSnake(name)), nil)
	assert.NoError(t, err, "failed to load test data")
	return integration.RunDirtyTestScenario(t, testScenario, server)
}

func TestProject(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestSecurityGroup(t *testing.T) {
	t.Skip("Fix buggy test scenario / implementation")
	integration.RunTest(t, t.Name(), server)
}

func TestLogicalRouterPing(t *testing.T) {
	t.Skip("Intent compiler methods must be done in transaction otherwise this test is flaky.")
	integration.RunTest(t, t.Name(), server)
}

func TestWaiter(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestReferredSecurityGroups(t *testing.T) {
	t.Skip("Fix buggy test scenario / implementation")
	integration.RunTest(t, t.Name(), server)
}
