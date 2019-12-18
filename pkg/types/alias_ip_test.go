package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam"
	ipammock "github.com/Juniper/contrail/pkg/types/ipam/mock"
)

func aliasIPSetupReadServiceMocks(t *testing.T, s *ContrailTypeLogicService) {
	readService, ok := s.ReadService.(*servicesmock.MockReadService)
	assert.True(t, ok)
	readService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetVirtualNetworkResponse{
			VirtualNetwork: &models.VirtualNetwork{},
		}, nil).AnyTimes()
}

func aliasIPSetupIPAMMocks(t *testing.T, s *ContrailTypeLogicService) {
	const okAddress = "10.0.0.1"
	addressManager, ok := s.AddressManager.(*ipammock.MockAddressManager)
	assert.True(t, ok)
	addressManager.EXPECT().AllocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *ipam.AllocateIPRequest) (address string, subnetUUID string, err error) {
			return okAddress, "uuid-0", nil
		}).AnyTimes()

	addressManager.EXPECT().IsIPAllocated(gomock.Not(gomock.Nil()),
		&ipam.IsIPAllocatedRequest{
			VirtualNetwork: &models.VirtualNetwork{},
			IPAddress:      okAddress,
		}).Return(false, nil).AnyTimes()
	addressManager.EXPECT().IsIPAllocated(gomock.Not(gomock.Nil()),
		&ipam.IsIPAllocatedRequest{
			VirtualNetwork: &models.VirtualNetwork{},
			IPAddress:      "10.0.0.2",
		}).Return(true, nil).AnyTimes()
}

func aliasIPSetupNextServiceMocks(t *testing.T, s *ContrailTypeLogicService) {
	nextService, ok := s.Next().(*servicesmock.MockService)
	assert.True(t, ok)
	// CreateAliasIP - response
	nextService.EXPECT().CreateAliasIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(
			_ context.Context, request *services.CreateAliasIPRequest,
		) (response *services.CreateAliasIPResponse, err error) {
			return &services.CreateAliasIPResponse{
				AliasIP: request.AliasIP,
			}, nil
		}).AnyTimes()

	// DeleteAliasIP - response
	nextService.EXPECT().DeleteAliasIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.DeleteAliasIPRequest,
		) (response *services.DeleteAliasIPResponse, err error) {
			return &services.DeleteAliasIPResponse{
				ID: request.ID,
			}, nil
		}).AnyTimes()
}

func aliasIPPrepareParent(t *testing.T, s *ContrailTypeLogicService, aliasIPPool *models.AliasIPPool) {
	readService, ok := s.ReadService.(*servicesmock.MockReadService)
	assert.True(t, ok)
	if aliasIPPool != nil {
		readService.EXPECT().GetAliasIPPool(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			&services.GetAliasIPPoolResponse{
				AliasIPPool: aliasIPPool,
			}, nil).AnyTimes()
	} else {
		readService.EXPECT().GetAliasIPPool(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			nil, errutil.ErrorNotFound).AnyTimes()
	}
}

func TestCreateAliasIP(t *testing.T) {
	tests := []struct {
		name             string
		aliasIPParent    *models.AliasIPPool
		createRequest    services.CreateAliasIPRequest
		expectedResponse services.CreateAliasIPResponse
		errorCode        codes.Code
	}{
		{
			name: "Create alias ip without alias ip pool",
			createRequest: services.CreateAliasIPRequest{AliasIP: &models.AliasIP{
				ParentType:     "alias-ip-pool",
				AliasIPAddress: "10.0.0.0",
			}},
			aliasIPParent: nil,
			errorCode:     codes.NotFound,
		},
		{
			name: "Create alias ip with a free ip address",
			createRequest: services.CreateAliasIPRequest{AliasIP: &models.AliasIP{
				ParentType:     "alias-ip-pool",
				AliasIPAddress: "10.0.0.1",
			}},
			aliasIPParent: &models.AliasIPPool{},
			expectedResponse: services.CreateAliasIPResponse{AliasIP: &models.AliasIP{
				ParentType:     "alias-ip-pool",
				AliasIPAddress: "10.0.0.1",
			}},
		},
		{
			name: "Try to create alias ip with IP address which is already allocated",
			createRequest: services.CreateAliasIPRequest{AliasIP: &models.AliasIP{
				ParentType:     "alias-ip-pool",
				AliasIPAddress: "10.0.0.2",
			}},
			aliasIPParent: &models.AliasIPPool{},
			errorCode:     codes.AlreadyExists,
		},
		{
			name: "Create alias ip without IP address",
			createRequest: services.CreateAliasIPRequest{AliasIP: &models.AliasIP{
				ParentType: "alias-ip-pool",
			}},
			aliasIPParent: &models.AliasIPPool{},
			expectedResponse: services.CreateAliasIPResponse{AliasIP: &models.AliasIP{
				ParentType:     "alias-ip-pool",
				AliasIPAddress: "10.0.0.1",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			aliasIPSetupReadServiceMocks(t, service)
			aliasIPSetupIPAMMocks(t, service)
			aliasIPSetupNextServiceMocks(t, service)

			aliasIPPrepareParent(t, service, tt.aliasIPParent)
			ctx := context.Background()
			createAliasIPResponse, err := service.CreateAliasIP(ctx, &tt.createRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				assert.Equal(t, tt.errorCode, errutil.CauseCode(err))
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, createAliasIPResponse)
				assert.EqualValues(t, &tt.expectedResponse, createAliasIPResponse)
			}
		})
	}
}

func TestDeleteAliasIP(t *testing.T) {
	tests := []struct {
		name             string
		aliasIP          *models.AliasIP
		aliasIPParent    *models.AliasIPPool
		deleteRequest    *services.DeleteAliasIPRequest
		expectedResponse *services.DeleteAliasIPResponse
		fails            bool
		deallocatesIP    bool
		errorCode        codes.Code
	}{
		{
			name:          "Try to delete non-existent alias ip",
			deleteRequest: &services.DeleteAliasIPRequest{ID: "uuid-1"},
			deallocatesIP: false,
			errorCode:     codes.NotFound,
		},
		{
			name:             "Try to delete allocated alias ip",
			aliasIP:          &models.AliasIP{AliasIPAddress: "10.0.0.1"},
			aliasIPParent:    &models.AliasIPPool{},
			deleteRequest:    &services.DeleteAliasIPRequest{ID: "uuid-2"},
			expectedResponse: &services.DeleteAliasIPResponse{ID: "uuid-2"},
			deallocatesIP:    true,
		},
		{
			name:          "Try to delete unallocated alias ip",
			aliasIP:       &models.AliasIP{AliasIPAddress: "10.0.0.1"},
			aliasIPParent: &models.AliasIPPool{},
			deleteRequest: &services.DeleteAliasIPRequest{ID: "uuid-3"},
			deallocatesIP: false,
			errorCode:     codes.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			aliasIPSetupReadServiceMocks(t, service)
			aliasIPSetupIPAMMocks(t, service)
			aliasIPSetupNextServiceMocks(t, service)
			aliasIPPrepareParent(t, service, tt.aliasIPParent)

			readService, ok := service.ReadService.(*servicesmock.MockReadService)
			assert.True(t, ok)
			if tt.aliasIP != nil {
				readService.EXPECT().GetAliasIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					&services.GetAliasIPResponse{
						AliasIP: tt.aliasIP,
					}, nil).AnyTimes()
			} else {
				readService.EXPECT().GetAliasIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					nil, errutil.ErrorNotFound).AnyTimes()
			}

			addressManager, ok := service.AddressManager.(*ipammock.MockAddressManager)
			assert.True(t, ok)
			if tt.deallocatesIP {
				addressManager.EXPECT().DeallocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(nil)
			} else {
				addressManager.EXPECT().DeallocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					errutil.ErrorNotFound).AnyTimes()
			}

			ctx := context.Background()
			deleteAliasIPResponse, err := service.DeleteAliasIP(ctx, tt.deleteRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				assert.Equal(t, tt.errorCode, errutil.CauseCode(err))
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, deleteAliasIPResponse)
				assert.EqualValues(t, tt.expectedResponse, deleteAliasIPResponse)
			}
		})
	}
}
