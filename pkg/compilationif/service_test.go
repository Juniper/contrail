package compilationif

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestIntentCompilerRunsNextService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	compiler := NewCompilationService()
	mockService := servicesmock.NewMockService(mockCtrl)
	services.Chain(compiler, mockService)

	network := &models.VirtualNetwork{
		UUID: "test_uuid",
	}
	mockService.EXPECT().CreateVirtualNetwork(gomock.Not(gomock.Nil()), &services.CreateVirtualNetworkRequest{
		VirtualNetwork: network,
	}).Return(&services.CreateVirtualNetworkResponse{VirtualNetwork: network}, nil)

	networkJSON, err := json.Marshal(network)
	require.NoError(t, err)
	err = compiler.handleEtcdMessages(context.Background(), int32(mvccpb.PUT),
		"/a/virtual_network/"+network.UUID, string(networkJSON))
	assert.NoError(t, err)
}

func TestIntentCompilerFailsForBadJSON(t *testing.T) {
	compiler := NewCompilationService()
	err := compiler.handleEtcdMessages(context.Background(), int32(mvccpb.PUT), "/a/virtual_network/test_uuid", "")
	assert.Error(t, err)
}

func TestIntentCompilerFailsForUnknownResource(t *testing.T) {
	compiler := NewCompilationService()
	err := compiler.handleEtcdMessages(context.Background(), int32(mvccpb.PUT), "/a/anything/test_uuid", "")
	assert.Error(t, err)
}
