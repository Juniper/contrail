package logic_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	createNetworkRequestPath = testDataPath + "create_network.json"
)

func TestNetworkCreate(t *testing.T) {
	runTest(t, func(t *testing.T, client *integration.HTTPAPIClient) {
		response, err := client.NeutronPost(
			context.Background(),
			loadRequestFromJSONFile(t, createNetworkRequestPath),
			[]int{200})
		assert.NoError(t, err)
		assertEqual(t, logic.NetworkResponse{Name: "ctest-vn-49391908"}, response)
	})
}
