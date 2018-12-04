package types

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam"
	ipammock "github.com/Juniper/contrail/pkg/types/ipam/mock"
)

type addrMgrSubnetExhausted int

func (e addrMgrSubnetExhausted) SubnetExhausted() {
}

func (e addrMgrSubnetExhausted) Error() string {
	return ""
}

func floatingIPSetupReadServiceMocks(s *ContrailTypeLogicService) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
	readService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetVirtualNetworkResponse{
			VirtualNetwork: &models.VirtualNetwork{},
		}, nil).AnyTimes()
}

func floatingIPSetupIPAMMocks(s *ContrailTypeLogicService) {
	addressManager := s.AddressManager.(*ipammock.MockAddressManager) //nolint: errcheck
	addressManager.EXPECT().AllocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *ipam.AllocateIPRequest) (address string, subnetUUID string, err error) {

			if request.SubnetUUID == "uuid-1" {
				return "10.0.0.1", "uuid-1", nil
			}

			exhaustedError := addrMgrSubnetExhausted(0)
			if request.SubnetUUID == "uuid-2" {
				return "", "", &exhaustedError
			}
			if request.SubnetUUID == "uuid-3" {
				return "", "", &exhaustedError
			}
			if request.SubnetUUID == "uuid-4" {
				return "", "", fmt.Errorf("generic error")
			}

			return "10.0.0.1", "uuid-1", nil
		}).AnyTimes()

	addressManager.EXPECT().IsIPAllocated(gomock.Not(gomock.Nil()),
		&ipam.IsIPAllocatedRequest{
			VirtualNetwork: &models.VirtualNetwork{},
			IPAddress:      "10.0.0.2",
		}).Return(true, nil).AnyTimes()
}

func floatingIPSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService) //nolint: errcheck
	// CreateFloatingIP - response
	nextService.EXPECT().CreateFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(
			_ context.Context, request *services.CreateFloatingIPRequest,
		) (response *services.CreateFloatingIPResponse, err error) {
			return &services.CreateFloatingIPResponse{
				FloatingIP: request.FloatingIP,
			}, nil
		}).AnyTimes()

	// DeleteFloatingIP - response
	nextService.EXPECT().DeleteFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.DeleteFloatingIPRequest,
		) (response *services.DeleteFloatingIPResponse, err error) {
			return &services.DeleteFloatingIPResponse{
				ID: request.ID,
			}, nil
		}).AnyTimes()
}

func floatingIPPrepareParent(s *ContrailTypeLogicService, floatingIPPool *models.FloatingIPPool) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
	if floatingIPPool != nil {
		readService.EXPECT().GetFloatingIPPool(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			&services.GetFloatingIPPoolResponse{
				FloatingIPPool: floatingIPPool,
			}, nil).AnyTimes()
	} else {
		readService.EXPECT().GetFloatingIPPool(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			nil, fmt.Errorf("error from DB")).AnyTimes()
	}
}

func TestCreateFloatingIP(t *testing.T) {
	tests := []struct {
		name             string
		floatingIPParent *models.FloatingIPPool
		createRequest    services.CreateFloatingIPRequest
		expectedResponse services.CreateFloatingIPResponse
		fails            bool
		errorCode        codes.Code
	}{
		{
			name:             "Create floating ip when parent type is instance-ip",
			createRequest:    services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{ParentType: "instance-ip"}},
			floatingIPParent: &models.FloatingIPPool{},
			expectedResponse: services.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{ParentType: "instance-ip"}},
		},
		{
			name:             "Create floating ip with a free ip address",
			createRequest:    services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{ParentType: "floating-ip-pool"}},
			floatingIPParent: &models.FloatingIPPool{},
			expectedResponse: services.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.1",
			}},
		},
		{
			name: "Try to create floating ip with IP address which is already allocated",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.2",
			}},
			floatingIPParent: &models.FloatingIPPool{},
			fails:            true,
			errorCode:        codes.AlreadyExists,
		},
		{
			name: "Create floating ip without IP address",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			floatingIPParent: &models.FloatingIPPool{},
			expectedResponse: services.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.1",
			}},
		},
		{
			name: "Create floating ip with subnets from floating ip pool",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			floatingIPParent: &models.FloatingIPPool{
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"uuid-2", "uuid-1"},
				},
			},
			expectedResponse: services.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.1",
			}},
		},
		{
			name: "Try to create floating ip with exhausted subnets from floating ip pool",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			floatingIPParent: &models.FloatingIPPool{
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"uuid-3", "uuid-2"},
				},
			},
			fails:     true,
			errorCode: codes.ResourceExhausted,
		},
		{
			name: "Try to create floating ip with subnets from floating ip pool with generic error",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			floatingIPParent: &models.FloatingIPPool{
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"uuid-3", "uuid-4"},
				},
			},
			fails: true,
		},
		{
			name: "Try to create floating ip when parent can't be get from DB ",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			floatingIPSetupReadServiceMocks(service)
			floatingIPSetupIPAMMocks(service)
			floatingIPSetupNextServiceMocks(service)

			floatingIPPrepareParent(service, tt.floatingIPParent)
			ctx := context.Background()
			createFloatingIPResponse, err := service.CreateFloatingIP(ctx, &tt.createRequest)

			if tt.fails {
				assert.Error(t, err)
				if tt.errorCode != codes.OK {
					status, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, tt.errorCode, status.Code())
				}
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, createFloatingIPResponse)
			assert.EqualValues(t, &tt.expectedResponse, createFloatingIPResponse)
		})
	}
}

func TestDeleteFloatingIP(t *testing.T) {

	tests := []struct {
		name             string
		floatingIPParent *models.FloatingIPPool
		floatingIP       *models.FloatingIP
		deleteRequest    *services.DeleteFloatingIPRequest
		expectedResponse *services.DeleteFloatingIPResponse
		fails            bool
		deallocatesIP    bool
		errorCode        codes.Code
	}{
		{
			name:             "Delete floating ip when parent type is instance-ip",
			deleteRequest:    &services.DeleteFloatingIPRequest{ID: "uuid-1"},
			floatingIPParent: &models.FloatingIPPool{},
			floatingIP:       &models.FloatingIP{UUID: "uuid-1", ParentType: "instance-ip"},
			expectedResponse: &services.DeleteFloatingIPResponse{ID: "uuid-1"},
			deallocatesIP:    false,
		},
		{
			name:             "Delete floating ip when parent type is floating-ip-pool",
			deleteRequest:    &services.DeleteFloatingIPRequest{ID: "uuid-1"},
			floatingIPParent: &models.FloatingIPPool{},
			floatingIP:       &models.FloatingIP{UUID: "uuid-1", ParentType: "floating-ip-pool"},
			expectedResponse: &services.DeleteFloatingIPResponse{ID: "uuid-1"},
			deallocatesIP:    true,
		},
		{
			name:          "Try to delete floating ip if it doesn't exist in DB",
			deleteRequest: &services.DeleteFloatingIPRequest{ID: "uuid-1"},
			deallocatesIP: false,
			fails:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			floatingIPSetupReadServiceMocks(service)
			floatingIPSetupIPAMMocks(service)
			floatingIPSetupNextServiceMocks(service)
			floatingIPPrepareParent(service, tt.floatingIPParent)

			readService := service.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
			if tt.floatingIP != nil {
				readService.EXPECT().GetFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					&services.GetFloatingIPResponse{
						FloatingIP: tt.floatingIP,
					}, nil).AnyTimes()
			} else {
				readService.EXPECT().GetFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					nil, fmt.Errorf("not found")).AnyTimes()
			}

			if tt.deallocatesIP {
				addressManager := service.AddressManager.(*ipammock.MockAddressManager) //nolint: errcheck
				addressManager.EXPECT().DeallocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(nil)
			}

			ctx := context.Background()
			deleteFloatingIPResponse, err := service.DeleteFloatingIP(ctx, tt.deleteRequest)

			if tt.fails {
				assert.Error(t, err)
				if tt.errorCode != codes.OK {
					status, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, tt.errorCode, status.Code())
				}
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, deleteFloatingIPResponse)
			assert.EqualValues(t, tt.expectedResponse, deleteFloatingIPResponse)
		})
	}
}
