package logic_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/neutron/logic"
)

const (
	testDataPath = "./test_data/"

	createNetworkRequestPath = testDataPath + "create_network.json"
)

func TestNetworkCreate(t *testing.T) {
	hc, cleanup := newHTTPClient(t)
	defer cleanup()

	response, err := hc.NeutronPost(context.Background(), loadRequestFromJSONFile(t, createNetworkRequestPath), []int{200})
	assert.NoError(t, err)
	assertEqual(t, logic.NetworkResponse{Name: "ctest-vn-49391908"}, response)
}
