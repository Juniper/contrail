package api

import (
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

var Server *integration.APIServer

func TestMain(m *testing.M) {
	apisrv.RegisterExtension(Init)
	integration.TestMain(m, &Server)
}

func TestInit(t *testing.T) {
	//Test for see we can compile.
	//TODO(nati) add actual test here.
}
