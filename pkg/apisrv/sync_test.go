package apisrv

import "testing"

func TestSyncAPI(t *testing.T) {
	err := RunTest("./test_data/test_sync.yml")
	if err != nil {
		t.Fatal(err)
	}
}
