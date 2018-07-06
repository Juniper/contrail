package compilationif

import (
	"context"
	"testing"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

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
