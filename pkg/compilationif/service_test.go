package compilationif

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestIntentCompilerHandlesDependentResources(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	compiler := NewCompilationService()
	mockService := servicesmock.NewMockService(mockCtrl)
	services.Chain(compiler, mockService)

	network := &models.VirtualNetwork{
		UUID: "Virtual-Network-1",
	}
	policy := &models.NetworkPolicy{
		UUID: "Network-policy-1",
	}

	ref := &models.VirtualNetworkNetworkPolicyRef{
		UUID: policy.UUID,
		To:   []string{"default-domain", "default-project", network.UUID},
	}
	network.NetworkPolicyRefs = append(network.NetworkPolicyRefs, ref)
	policy.VirtualNetworkBackRefs = append(policy.VirtualNetworkBackRefs, network)

	mockService.EXPECT().CreateVirtualNetwork(gomock.Not(gomock.Nil()),
		&services.CreateVirtualNetworkRequest{network},
	).Return(&services.CreateVirtualNetworkResponse{network}, nil)

	mockService.EXPECT().UpdateNetworkPolicy(gomock.Not(gomock.Nil()),
		&services.UpdateNetworkPolicyRequest{
			NetworkPolicy: policy,
			FieldMask:     types.FieldMask{Paths: []string{}},
		}).Return(&services.UpdateNetworkPolicyResponse{policy}, nil)

	networkJSON, err := json.Marshal(network)
	require.NoError(t, err, "Marshaling the network should succeed")

	err = compiler.handleEtcdMessages(context.Background(), int32(mvccpb.PUT),
		"/a/virtual_network/"+network.UUID, string(networkJSON))
	assert.NoError(t, err)
}

func TestIntentCompilerRunsNextService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	compiler := NewCompilationService()
	mockService := servicesmock.NewMockService(mockCtrl)
	services.Chain(compiler, mockService)

	mockService.EXPECT().CreateVirtualNetwork(gomock.Not(gomock.Nil()), &services.CreateVirtualNetworkRequest{
		VirtualNetwork: &models.VirtualNetwork{},
	}).Return(&services.CreateVirtualNetworkResponse{}, nil)

	err := compiler.handleEtcdMessages(context.Background(), int32(mvccpb.PUT),
		"/a/virtual_network/test_uuid", "{}")
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
