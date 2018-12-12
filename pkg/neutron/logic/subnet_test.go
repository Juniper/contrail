package logic_test

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	subnetReadAllReq = testDataPath + "subnet_readall_req.json"
	subnetReadAllRes = testDataPath + "subnet_readall_res.json"
)

func TestSubnet_ReadAll(t *testing.T) {
	var expected []*logic.SubnetResponse
	require.NoError(t, fileutil.LoadFile(subnetReadAllRes, expected))

	runTest(t, func(t *testing.T, client *integration.HTTPAPIClient) {
		response, err := client.NeutronPost(
			context.Background(),
			loadRequestFromJSONFile(t, subnetReadAllReq),
			[]int{200})
		assert.NoError(t, err)
		assertEqual(t, expected, response)
	})
}
