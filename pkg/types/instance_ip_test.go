package types

import (
	"context"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"github.com/Juniper/contrail/pkg/types/ipam/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func instanceIDPrepareVirtualNetwork(s *ContrailTypeLogicService) {
	readService := s.ReadService.(*servicesmock.MockService)

	readService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()),
		&services.GetVirtualNetworkRequest{
			ID: "virtual-network-uuid-1",
		}).Return(
		&services.GetVirtualNetworkResponse{
			VirtualNetwork: &models.VirtualNetwork{
				FQName: []string{"default", "ip-fabric", "__link_local__"},
			},
		}, nil).AnyTimes()

	readService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, common.ErrorNotFound).AnyTimes()
}

func instanceIPPrepareReadService(s *ContrailTypeLogicService, instanceIP *models.InstanceIP) {
	readService := s.ReadService.(*servicesmock.MockService)

	if instanceIP != nil {
		readService.EXPECT().GetInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			&services.GetInstanceIPResponse{
				InstanceIP: instanceIP,
			}, nil).AnyTimes()
	} else {
		readService.EXPECT().GetInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			nil, fmt.Errorf("Not found")).AnyTimes()
	}
}

func instanceIDPrepareVirtualRouter(s *ContrailTypeLogicService) {
	readService := s.ReadService.(*servicesmock.MockService)

	readService.EXPECT().GetVirtualRouter(gomock.Not(gomock.Nil()),
		&services.GetVirtualRouterRequest{
			ID: "virtual-router-uuid-1",
		}).Return(
		&services.GetVirtualRouterResponse{
			VirtualRouter: &models.VirtualRouter{
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{
					{
						Attr: &models.VirtualRouterNetworkIpamType{
							AllocationPools: []*models.AllocationPoolType{
								{Start: "10.10.10.12", End: "10.10.10.15"},
								{Start: "10.10.10.17", End: "10.10.10.20"},
							},
						},
					},
				},
			},
		}, nil).AnyTimes()

	readService.EXPECT().GetVirtualRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, common.ErrorNotFound).AnyTimes()
}

func instanceIPSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService)
	nextService.EXPECT().CreateInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.CreateInstanceIPRequest,
		) (response *services.CreateInstanceIPResponse, err error) {
			return &services.CreateInstanceIPResponse{InstanceIP: request.InstanceIP}, nil
		}).AnyTimes()

	nextService.EXPECT().UpdateInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.UpdateInstanceIPRequest,
		) (response *services.UpdateInstanceIPResponse, err error) {
			return &services.UpdateInstanceIPResponse{InstanceIP: request.InstanceIP}, nil
		}).AnyTimes()

	nextService.EXPECT().DeleteInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.DeleteInstanceIPRequest,
		) (response *services.DeleteInstanceIPResponse, err error) {
			return &services.DeleteInstanceIPResponse{ID: request.ID}, nil
		}).AnyTimes()
}

func instanceIPSetupIPAMMocks(s *ContrailTypeLogicService) {
	addressManager := s.AddressManager.(*ipammock.MockAddressManager)
	addressManager.EXPECT().AllocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *ipam.AllocateIPRequest) (address string, subnetUUID string, err error) {
			return "10.10.10.20", "uuid-1", nil
		}).AnyTimes()

	addressManager.EXPECT().DeallocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *ipam.DeallocateIPRequest) error {
			return nil
		}).AnyTimes()

	addressManager.EXPECT().IsIPAllocated(gomock.Not(gomock.Nil()),
		&ipam.IsIPAllocatedRequest{
			IPAddress: "10.10.10.10",
		}).Return(false, nil).AnyTimes()

	addressManager.EXPECT().IsIPAllocated(gomock.Not(gomock.Nil()),
		&ipam.IsIPAllocatedRequest{
			IPAddress: "10.10.10.11",
		}).Return(false, common.ErrorNotFound).AnyTimes()
}

func TestCreateInstanceIP(t *testing.T) {
	tests := []struct {
		name               string
		paramInstanceIP    models.InstanceIP
		expectedInstanceIP models.InstanceIP
		fails              bool
		errorCode          codes.Code
	}{
		{
			name: "Try to create instance-ip when both virtual-router-refs and network-ipam-refs are defined",
			paramInstanceIP: models.InstanceIP{
				VirtualRouterRefs: []*models.InstanceIPVirtualRouterRef{
					{
						UUID: "virtual-router-uuid-1",
					},
				},
				NetworkIpamRefs: []*models.InstanceIPNetworkIpamRef{
					{
						UUID: "network-ipam-uuid-1",
					},
				},
			},
			fails:     true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create instance-ip when both virtual-router-refs and virtual-network-refs are defined",
			paramInstanceIP: models.InstanceIP{
				VirtualRouterRefs: []*models.InstanceIPVirtualRouterRef{
					{
						UUID: "virtual-router-uuid-1",
					},
				},
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
					},
				},
			},
			fails:     true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create instance-ip when cannot find virtual-network with corresponding uuid",
			paramInstanceIP: models.InstanceIP{
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
				},
			},
			fails: true,
		},
		{
			name: "Try to create instance-ip when checking if ip-address is allocated returns error",
			paramInstanceIP: models.InstanceIP{
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
				},
				InstanceIPAddress: "10.10.10.11",
			},
			fails: true,
		},
		{
			name: "Try to create instance-ip when refers to multiple vrouters",
			paramInstanceIP: models.InstanceIP{
				VirtualRouterRefs: []*models.InstanceIPVirtualRouterRef{
					{
						UUID: "virtual-router-uuid-1",
					},
					{
						UUID: "virtual-router-uuid-2",
					},
				},
				InstanceIPAddress: "10.10.10.10",
			},
			fails:     true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create instance-ip when cannot find virtual-router with corresponding uuid",
			paramInstanceIP: models.InstanceIP{
				VirtualRouterRefs: []*models.InstanceIPVirtualRouterRef{
					{
						UUID: "virtual-router-uuid-2",
					},
				},
				InstanceIPAddress: "10.10.10.10",
			},
			fails: true,
		},
		{
			name: "Try to create instance-ip when allocation for requested ip from a network-ipam",
			paramInstanceIP: models.InstanceIP{
				VirtualRouterRefs: []*models.InstanceIPVirtualRouterRef{
					{
						UUID: "virtual-router-uuid-1",
					},
				},
				InstanceIPAddress: "10.10.10.10",
			},
			fails:     true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create instance-ip when allocation-pools in virtual-router are defined",
			paramInstanceIP: models.InstanceIP{
				VirtualRouterRefs: []*models.InstanceIPVirtualRouterRef{
					{
						UUID: "virtual-router-uuid-1",
					},
				},
			},
			expectedInstanceIP: models.InstanceIP{
				VirtualRouterRefs: []*models.InstanceIPVirtualRouterRef{
					{
						UUID: "virtual-router-uuid-1",
					},
				},
				InstanceIPAddress: "10.10.10.20",
				SubnetUUID:        "uuid-1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			instanceIPSetupNextServiceMocks(service)
			instanceIPSetupIPAMMocks(service)
			instanceIDPrepareVirtualNetwork(service)
			instanceIDPrepareVirtualRouter(service)

			ctx := context.Background()

			paramRequest := services.CreateInstanceIPRequest{InstanceIP: &tt.paramInstanceIP}
			expectedResponse := services.CreateInstanceIPResponse{InstanceIP: &tt.expectedInstanceIP}
			createInstanceIPResponse, err := service.CreateInstanceIP(ctx, &paramRequest)

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
			assert.NotNil(t, createInstanceIPResponse)
			assert.EqualValues(t, &expectedResponse, createInstanceIPResponse)
		})
	}
}
func TestUpdateInstanceIP(t *testing.T) {
	tests := []struct {
		name               string
		requestInstanceIP  *models.InstanceIP
		databaseInstanceIP *models.InstanceIP
		expectedInstanceIP *models.InstanceIP
		fails              bool
		errorCode          codes.Code
	}{
		{
			name: "Try to update instance-ip when cannot get instance-ip from database",
			requestInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
			},
			fails: true,
		},
		{
			name: "Try to update instance-ip when instance-ip does not have virtual-network refs",
			requestInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
			},
			databaseInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
			},
			expectedInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
			},
		},
		{
			name: "Try to update instance-ip when cannot get virtual-network from database",
			requestInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
			},
			databaseInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
				},
			},
			fails: true,
		},
		{
			name: "Try to update instance-ip ip-address",
			requestInstanceIP: &models.InstanceIP{
				UUID:              "instance-ip-uuid",
				InstanceIPAddress: "10.10.10.10",
			},
			databaseInstanceIP: &models.InstanceIP{
				UUID:              "instance-ip-uuid",
				InstanceIPAddress: "10.10.10.20",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
					},
				},
			},
			fails:     true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to update instance-ip when fq-name is ip-fablic or link-local",
			requestInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
			},
			databaseInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
					},
				},
			},
			expectedInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
			},
		},
		{
			name: "Try to update instance-ip",
			requestInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
			},
			databaseInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
					},
				},
			},
			expectedInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			instanceIPSetupNextServiceMocks(service)
			instanceIDPrepareVirtualNetwork(service)
			instanceIPPrepareReadService(service, tt.databaseInstanceIP)

			ctx := context.Background()

			paramRequest := services.UpdateInstanceIPRequest{InstanceIP: tt.requestInstanceIP}
			expectedResponse := services.UpdateInstanceIPResponse{InstanceIP: tt.expectedInstanceIP}
			createInstanceIPResponse, err := service.UpdateInstanceIP(ctx, &paramRequest)

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
			assert.NotNil(t, createInstanceIPResponse)
			assert.EqualValues(t, &expectedResponse, createInstanceIPResponse)
		})
	}
}

func TestDeleteInstanceIP(t *testing.T) {
	tests := []struct {
		name       string
		paramID    string
		instanceIP *models.InstanceIP
		fails      bool
		errorCode  codes.Code
	}{
		{
			name:    "Try to delete instance-ip when cannot find instance-ip with param id",
			paramID: "instance-ip-id",
			fails:   true,
		},
		{
			name:       "Try to delete instance-ip when ip-address is not defined",
			paramID:    "instance-ip-id",
			instanceIP: &models.InstanceIP{},
		},
		{
			name:    "Try to delete instance-ip when ipam-refs are defined",
			paramID: "instance-ip-id",
			instanceIP: &models.InstanceIP{
				NetworkIpamRefs: []*models.InstanceIPNetworkIpamRef{
					{
						UUID: "ipam-uuid",
					},
				},
			},
		},
		{
			name:    "Try to delete instance-ip when virtual-network-refs are not defined",
			paramID: "instance-ip-id",
			instanceIP: &models.InstanceIP{
				InstanceIPAddress: "10.10.10.10",
			},
		},
		{
			name:    "Try to delete instance-ip when cannot find virtual-network",
			paramID: "instance-ip-id",
			instanceIP: &models.InstanceIP{
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
				},
				InstanceIPAddress: "10.10.10.10",
			},
			fails: true,
		},
		{
			name:    "Try to delete instance-ip when virtual-network was found",
			paramID: "instance-ip-id",
			instanceIP: &models.InstanceIP{
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
					},
				},
				InstanceIPAddress: "10.10.10.10",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			instanceIPSetupIPAMMocks(service)
			instanceIPSetupNextServiceMocks(service)
			instanceIDPrepareVirtualNetwork(service)
			instanceIPPrepareReadService(service, tt.instanceIP)

			ctx := context.Background()

			paramRequest := services.DeleteInstanceIPRequest{ID: tt.paramID}
			expectedResponse := services.DeleteInstanceIPResponse{ID: tt.paramID}
			createInstanceIPResponse, err := service.DeleteInstanceIP(ctx, &paramRequest)

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
			assert.NotNil(t, createInstanceIPResponse)
			assert.EqualValues(t, &expectedResponse, createInstanceIPResponse)
		})
	}
}
