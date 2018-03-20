package apisrv

import (
	"testing"
)

func TestProxyEndpoint(t *testing.T) {
	RunTest(t, "./test_data/test_endpoint.yml")
}
