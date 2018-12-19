package neutron_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestFloatingIP(t *testing.T) {
	integration.RunIntegrationTest(t, t.Name(), server)
}

func TestNetwork(t *testing.T) {
	integration.RunIntegrationTest(t, t.Name(), server)
}
