package apisrv

import "testing"

func TestServer(t *testing.T) {
	err := RunTest("./test_data/test_virtual_network.yml")
	if err != nil {
		t.Fatal(err)
	}
}
