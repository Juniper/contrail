package logic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	TestDataPath = "./test_data/"

	CreateNetworkRequestPath = TestDataPath + "create_network.json"
)

func TestNetworkCreate(t *testing.T) {
	hc := integration.NewTestingHTTPClient(t, server.URL())

	r := loadRequestFromJSONFile(t, CreateNetworkRequestPath)
	response, err := hc.NeutronPost(r, []int{200})
	assert.NoError(t, err)
	assertEqual(t, logic.NetworkResponse{Name: "ctest-vn-49391908"}, response)
}
