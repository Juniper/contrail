package api

import (
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/integration"
)

func TestMain(m *testing.M) {
	apisrv.RegisterExtension(Init)
	integration.SetupAndRunTest(m)
}

func TestInit(t *testing.T) {
	//Test for see we can compile.
	//TODO(nati) add actual test here.
}
