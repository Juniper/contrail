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

func TestPort(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}
