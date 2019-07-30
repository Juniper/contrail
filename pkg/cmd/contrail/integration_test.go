package contrail_test

import (
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/require"
)

func TestFQNameCleanup(t *testing.T) {
	cleanup := runDirtyTest(t, t.Name())
	defer cleanup()

	runDirtyTest(t, t.Name())
}

func runDirtyTest(t *testing.T, name string) func() {
	ts, err := integration.LoadTest(fmt.Sprintf("./tests/%s.yml", format.CamelToSnake(name)), nil)
	require.NoError(t, err, "failed to load test data")
	return integration.RunDirtyTestScenario(t, ts, server)
}

func TestProject(t *testing.T) {
	integration.RunTestFromTestsDirectory(t, t.Name(), server)
}

func TestSecurityGroup(t *testing.T) {
	t.Skip("The test scenario or the implementation is buggy")
	integration.RunTestFromTestsDirectory(t, t.Name(), server)
}

func TestLogicalRouterPing(t *testing.T) {
	t.Skip("Intent compiler methods must be done in transaction otherwise this test is flaky.")
	integration.RunTestFromTestsDirectory(t, t.Name(), server)
}

func TestWaiter(t *testing.T) {
	integration.RunTestFromTestsDirectory(t, t.Name(), server)
}

func TestReferredSecurityGroups(t *testing.T) {
	t.Skip("The test scenario or the implementation is buggy")
	integration.RunTestFromTestsDirectory(t, t.Name(), server)
}
