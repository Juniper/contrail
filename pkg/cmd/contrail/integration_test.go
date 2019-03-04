package contrail_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestFQNameCleanup(t *testing.T) {
	runDirtyTest(t, t.Name())
	integration.RunTest(t, t.Name(), server)
}

func TestProject(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestSecurityGroup(t *testing.T) {
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
	integration.RunTest(t, t.Name(), server)
}
