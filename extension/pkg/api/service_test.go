package api

import (
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
)

func TestMain(m *testing.M) {
	apisrv.RegisterExtension(Init)
	apisrv.SetupAndRunTest(m)
}

func TestInit(t *testing.T) {
	//Test for see we can compile.

}
