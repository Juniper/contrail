package types

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestCreateFirewallPolicy(t *testing.T) {
	tests := []struct {
		name      string
		errorCode codes.Code
	}{
		{
			name: "TODO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			ctx := context.Background()

			createFirewallPolicyResponse, err := service.CreateFirewallPolicy(ctx, &paramRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createFirewallPolicyResponse)
			}
		})
	}
}

func TestUpdateLogicalRouter(t *testing.T) {
	tests := []struct {
		name              string
		testLogicalRouter models.LogicalRouter
		vxlanEnabled      bool
		fieldMaskPaths    []string
		dbLogicalRouter   models.LogicalRouter
		errorCode         codes.Code
	}{
		{
			name: "Try to update logical-router when external gateway with enabled vxlan routing",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID: "project-uuid-1",
				VirtualNetworkRefs: []*models.LogicalRouterVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
						Attr: &models.LogicalRouterVirtualNetworkType{
							LogicalRouterVirtualNetworkType: "ExternalGateway",
						},
					},
				},
			},
			vxlanEnabled:   true,
			fieldMaskPaths: []string{models.LogicalRouterFieldParentUUID, models.LogicalRouterFieldVirtualNetworkRefs},
			errorCode:      codes.InvalidArgument,
		},
		{
			name: "Try to update logical-router when logical router interface and gateway in the same network",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID: "project-uuid-1",
				VirtualNetworkRefs: []*models.LogicalRouterVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-1",
						Attr: &models.LogicalRouterVirtualNetworkType{
							LogicalRouterVirtualNetworkType: "InternalVirtualNetwork",
						},
					},
				},
			},
			fieldMaskPaths: []string{models.LogicalRouterFieldParentUUID, models.LogicalRouterFieldVirtualNetworkRefs},
			dbLogicalRouter: models.LogicalRouter{
				ParentUUID: "project-uuid-2",
				VirtualNetworkRefs: []*models.LogicalRouterVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
						Attr: &models.LogicalRouterVirtualNetworkType{
							LogicalRouterVirtualNetworkType: "InternalVirtualNetwork",
						},
					},
				},
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "virtual-machine-interface-1",
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to update logical-router when port already in use by virtual-machine",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID: "project-uuid-1",
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "virtual-machine-interface-1",
					},
				},
			},
			fieldMaskPaths: []string{models.LogicalRouterFieldParentUUID, models.LogicalRouterFieldVirtualNetworkRefs},
			errorCode:      codes.AlreadyExists,
		},
		{
			name: "try to update logical-router with improper vxlan id",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID:             "project-uuid-1",
				VxlanNetworkIdentifier: "id",
			},
			dbLogicalRouter: models.LogicalRouter{
				ParentUUID:             "project-uuid-1",
				VxlanNetworkIdentifier: "1",
				VirtualNetworkRefs: []*models.LogicalRouterVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
						Attr: &models.LogicalRouterVirtualNetworkType{
							LogicalRouterVirtualNetworkType: "InternalVirtualNetwork",
						},
					},
				},
			},
			vxlanEnabled:   true,
			fieldMaskPaths: []string{models.LogicalRouterFieldVxlanNetworkIdentifier},
			errorCode:      codes.InvalidArgument,
		},
		{
			name: "Update logical-router with vxlan routing enabled",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID:             "project-uuid-1",
				VxlanNetworkIdentifier: "2",
			},
			dbLogicalRouter: models.LogicalRouter{
				ParentUUID:             "project-uuid-1",
				VxlanNetworkIdentifier: "1",
				VirtualNetworkRefs: []*models.LogicalRouterVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid-2",
						Attr: &models.LogicalRouterVirtualNetworkType{
							LogicalRouterVirtualNetworkType: "InternalVirtualNetwork",
						},
					},
				},
			},
			vxlanEnabled:   true,
			fieldMaskPaths: []string{models.LogicalRouterFieldVxlanNetworkIdentifier},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			logicalRouterSetupReadServiceMocks(service, &tt.testLogicalRouter, tt.vxlanEnabled)
			logicalRouterSetupIntPoolAllocatorMocks(service)

			readService := service.ReadService.(*servicesmock.MockReadService)
			readService.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.GetLogicalRouterResponse{
					LogicalRouter: &tt.dbLogicalRouter,
				},
				nil,
			).AnyTimes()

			updateLRCall := service.Next().(*servicesmock.MockService).EXPECT().UpdateLogicalRouter(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, request *services.UpdateLogicalRouterRequest,
				) (response *services.UpdateLogicalRouterResponse, err error) {
					return &services.UpdateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}, nil
				},
			)

			updateVNCall := service.WriteService.(*servicesmock.MockWriteService).EXPECT().UpdateVirtualNetwork(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, request *services.UpdateVirtualNetworkRequest,
				) (response *services.UpdateVirtualNetworkResponse, err error) {
					return &services.UpdateVirtualNetworkResponse{VirtualNetwork: request.VirtualNetwork}, nil
				},
			)

			if tt.errorCode != codes.OK {
				updateLRCall.MaxTimes(1)
			} else {
				updateLRCall.Times(1)
			}

			if tt.errorCode == codes.OK && tt.vxlanEnabled {
				updateVNCall.Times(1)
			} else {
				updateVNCall.MaxTimes(1)
			}

			ctx := context.Background()
			paramRequest := services.UpdateLogicalRouterRequest{
				LogicalRouter: &tt.testLogicalRouter,
				FieldMask: types.FieldMask{
					Paths: tt.fieldMaskPaths,
				},
			}
			expectedResponse := services.UpdateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}
			updateLogicalRouterResponse, err := service.UpdateLogicalRouter(ctx, &paramRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, updateLogicalRouterResponse)
			}
		})
	}
}
