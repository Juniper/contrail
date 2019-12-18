package types

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam"
	ipammock "github.com/Juniper/contrail/pkg/types/ipam/mock"
)

func instanceIPPrepareVirtualNetwork(s *ContrailTypeLogicService) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	virtualNetworks := []*models.VirtualNetwork{models.MakeVirtualNetwork(), models.MakeVirtualNetwork()}
	virtualNetworks[0].UUID = "virtual-network-uuid-1"
	virtualNetworks[0].FQName = []string{"default", "ip-fabric", "__link_local__"}
	virtualNetworks[1].UUID = "virtual-network-uuid-2"
	virtualNetworks[1].FQName = []string{"default"}

	mockedReadServiceAddVirtualNetwork(s, virtualNetworks[0])
	mockedReadServiceAddVirtualNetwork(s, virtualNetworks[1])

	readService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, errutil.ErrorNotFound).AnyTimes()
}

func instanceIPPrepareReadService(s *ContrailTypeLogicService, instanceIP *models.InstanceIP) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	if instanceIP != nil {
		readService.EXPECT().GetInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			&services.GetInstanceIPResponse{
				InstanceIP: instanceIP,
			}, nil).AnyTimes()
	} else {
		readService.EXPECT().GetInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			nil, errutil.ErrorNotFound).AnyTimes()
	}
}

func instanceIPPrepareVirtualRouter(s *ContrailTypeLogicService) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

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
		nil, errutil.ErrorNotFound).AnyTimes()
}

func instanceIPSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService) //nolint: errcheck
	nextService.EXPECT().CreateInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.CreateInstanceIPRequest,
		) (response *services.CreateInstanceIPResponse, err error) {
			return &services.CreateInstanceIPResponse{InstanceIP: request.InstanceIP}, nil
		}).MaxTimes(1)

	nextService.EXPECT().UpdateInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.UpdateInstanceIPRequest,
		) (response *services.UpdateInstanceIPResponse, err error) {
			return &services.UpdateInstanceIPResponse{InstanceIP: request.InstanceIP}, nil
		}).MaxTimes(1)

	nextService.EXPECT().DeleteInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.DeleteInstanceIPRequest,
		) (response *services.DeleteInstanceIPResponse, err error) {
			return &services.DeleteInstanceIPResponse{ID: request.ID}, nil
		}).MaxTimes(1)
}

func instanceIPSetupIPAMMocks(s *ContrailTypeLogicService) {
	addressManager := s.AddressManager.(*ipammock.MockAddressManager) //nolint: errcheck
	virtualNetwork := models.MakeVirtualNetwork()
	virtualNetwork.UUID = "virtual-network-uuid-2"
	virtualNetwork.FQName = []string{"default"}

	addressManager.EXPECT().AllocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		"10.10.10.20", "uuid-1", nil,
	).AnyTimes()

	addressManager.EXPECT().DeallocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *ipam.DeallocateIPRequest) error {
			return nil
		}).AnyTimes()

	addressManager.EXPECT().IsIPAllocated(gomock.Not(gomock.Nil()),
		&ipam.IsIPAllocatedRequest{
			VirtualNetwork: virtualNetwork,
			IPAddress:      "10.10.10.10",
		}).Return(false, nil).AnyTimes()

	addressManager.EXPECT().IsIPAllocated(gomock.Not(gomock.Nil()),
		&ipam.IsIPAllocatedRequest{
			VirtualNetwork: virtualNetwork,
			IPAddress:      "10.10.10.11",
		}).Return(false, errutil.ErrorNotFound).AnyTimes()
}

func TestCreateInstanceIP(t *testing.T) {
	tests := []struct {
		name               string
		paramInstanceIP    models.InstanceIP
		expectedInstanceIP models.InstanceIP
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
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create instance-ip when cannot find virtual-network with corresponding uuid",
			paramInstanceIP: models.InstanceIP{
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-3",
					},
				},
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Create instance-ip when should ignore ip allocation",
			paramInstanceIP: models.InstanceIP{
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
					},
				},
			},
			expectedInstanceIP: models.InstanceIP{
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
					},
				},
			},
		},
		{
			name: "Try to create instance-ip when checking if ip-address is allocated returns error",
			paramInstanceIP: models.InstanceIP{
				InstanceIPAddress: "10.10.10.11",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
				},
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Create instance-ip with virtual-network-refs",
			paramInstanceIP: models.InstanceIP{
				InstanceIPAddress: "10.10.10.10",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
				},
			},
			expectedInstanceIP: models.InstanceIP{
				InstanceIPAddress: "10.10.10.20",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
				},
				SubnetUUID: "uuid-1",
			},
		},
		{
			name: "Try to create instance-ip when refers to multiple vrouters",
			paramInstanceIP: models.InstanceIP{
				InstanceIPAddress: "10.10.10.10",
				VirtualRouterRefs: []*models.InstanceIPVirtualRouterRef{
					{
						UUID: "virtual-router-uuid-1",
					},
					{
						UUID: "virtual-router-uuid-2",
					},
				},
			},
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
			errorCode: codes.NotFound,
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
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Create instance-ip when allocation-pools in virtual-router are defined",
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
		{
			name: "Create instance-ip with network-ipam-refs",
			paramInstanceIP: models.InstanceIP{
				NetworkIpamRefs: []*models.InstanceIPNetworkIpamRef{
					{
						UUID: "network-ipam-uuid-1",
					},
				},
			},
			expectedInstanceIP: models.InstanceIP{
				NetworkIpamRefs: []*models.InstanceIPNetworkIpamRef{
					{
						UUID: "network-ipam-uuid-1",
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
			instanceIPPrepareVirtualNetwork(service)
			instanceIPPrepareVirtualRouter(service)

			ctx := context.Background()

			paramRequest := services.CreateInstanceIPRequest{InstanceIP: &tt.paramInstanceIP}
			expectedResponse := services.CreateInstanceIPResponse{InstanceIP: &tt.expectedInstanceIP}
			createInstanceIPResponse, err := service.CreateInstanceIP(ctx, &paramRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createInstanceIPResponse)
			}
		})
	}
}

func TestUpdateInstanceIP(t *testing.T) {
	tests := []struct {
		name               string
		request            *services.UpdateInstanceIPRequest
		databaseInstanceIP *models.InstanceIP
		expectedInstanceIP *models.InstanceIP
		errorCode          codes.Code
	}{
		{
			name: "Try to update instance-ip when cannot get instance-ip from database",
			request: &services.UpdateInstanceIPRequest{
				InstanceIP: &models.InstanceIP{
					UUID: "instance-ip-uuid",
				},
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Try to update instance-ip when instance-ip does not have virtual-network refs",
			request: &services.UpdateInstanceIPRequest{
				InstanceIP: &models.InstanceIP{
					UUID: "instance-ip-uuid",
				},
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
			request: &services.UpdateInstanceIPRequest{
				InstanceIP: &models.InstanceIP{
					UUID: "instance-ip-uuid",
				},
			},
			databaseInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-3",
					},
				},
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Try to update instance-ip ip-address",
			request: &services.UpdateInstanceIPRequest{
				InstanceIP: &models.InstanceIP{
					UUID:              "instance-ip-uuid",
					InstanceIPAddress: "10.10.10.10",
				},
				FieldMask: types.FieldMask{
					Paths: []string{models.InstanceIPFieldInstanceIPAddress},
				},
			},
			databaseInstanceIP: &models.InstanceIP{
				UUID:              "instance-ip-uuid",
				InstanceIPAddress: "10.10.10.20",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to update instance-ip when fq-name is ip-fablic or link-local",
			request: &services.UpdateInstanceIPRequest{
				InstanceIP: &models.InstanceIP{
					UUID: "instance-ip-uuid",
				},
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
			request: &services.UpdateInstanceIPRequest{
				InstanceIP: &models.InstanceIP{
					UUID: "instance-ip-uuid",
				},
			},
			databaseInstanceIP: &models.InstanceIP{
				UUID: "instance-ip-uuid",
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
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
			instanceIPPrepareVirtualNetwork(service)
			instanceIPPrepareReadService(service, tt.databaseInstanceIP)

			ctx := context.Background()

			expectedResponse := services.UpdateInstanceIPResponse{InstanceIP: tt.expectedInstanceIP}
			createInstanceIPResponse, err := service.UpdateInstanceIP(ctx, tt.request)

			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createInstanceIPResponse)
			}
		})
	}
}

func TestDeleteInstanceIP(t *testing.T) {
	tests := []struct {
		name       string
		paramID    string
		instanceIP *models.InstanceIP
		errorCode  codes.Code
	}{
		{
			name:      "Try to delete instance-ip when cannot find instance-ip with param id",
			paramID:   "instance-ip-id",
			errorCode: codes.NotFound,
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
						UUID: "virtual-network-uuid-3",
					},
				},
				InstanceIPAddress: "10.10.10.10",
			},
			errorCode: codes.NotFound,
		},
		{
			name:    "Try to delete instance-ip when virtual-network was found",
			paramID: "instance-ip-id",
			instanceIP: &models.InstanceIP{
				VirtualNetworkRefs: []*models.InstanceIPVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
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
			instanceIPPrepareVirtualNetwork(service)
			instanceIPPrepareReadService(service, tt.instanceIP)

			ctx := context.Background()

			paramRequest := services.DeleteInstanceIPRequest{ID: tt.paramID}
			expectedResponse := services.DeleteInstanceIPResponse{ID: tt.paramID}
			createInstanceIPResponse, err := service.DeleteInstanceIP(ctx, &paramRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createInstanceIPResponse)
			}
		})
	}
}
