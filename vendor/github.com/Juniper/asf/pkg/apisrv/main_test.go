package apisrv_test

import (
	"testing"

	"github.com/Juniper/asf/pkg/testutil/integration"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}
