package neutron_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestFloatingIP(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestNetwork(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestNetworkReferredByPort(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestPort(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestProject(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestProjectDefaultSecurityGroup(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestRouter(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestSecurityGroup(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestSecurityGroupRule(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestSubnet(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestHealthCheck(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}
