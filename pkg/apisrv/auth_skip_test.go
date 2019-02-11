package apisrv_test

import (
	"testing"
)

const (
	authSkipTestFile = "./test_data/test_auth_skip.yml"
)

func TestContrailClusterAuthSkip(t *testing.T) {
	RunTest(t, authSkipTestFile)
}
