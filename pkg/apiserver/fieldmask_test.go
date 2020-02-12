package apiserver_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestFieldMask(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{})
	defer s.CloseT(t)
	integration.RunTest(t, "./test_data/test_fieldmask.yml", s)
}
