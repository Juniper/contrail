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

func logicalRouterSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService)
	nextService.EXPECT().CreateLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.CreateLogicalRouterRequest,
		) (response *services.CreateLogicalRouterResponse, err error) {
			return &services.CreateLogicalRouterResponse{LogicalRouter: request.LogicalRouter}, nil
		}).MaxTimes(1)

	nextService.EXPECT().UpdateLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.UpdateLogicalRouterRequest,
		) (response *services.UpdateLogicalRouterResponse, err error) {
			return &services.UpdateLogicalRouterResponse{LogicalRouter: request.LogicalRouter}, nil
		}).MaxTimes(1)

	nextService.EXPECT().DeleteLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.DeleteLogicalRouterRequest,
		) (response *services.DeleteLogicalRouterResponse, err error) {
			return &services.DeleteLogicalRouterResponse{ID: request.ID}, nil
		}).MaxTimes(1)
}

func logicalRouterSetupReadServiceMocks(s *ContrailTypeLogicService, lr *models.LogicalRouter) {
	readService := s.ReadService.(*servicesmock.MockReadService)
	project := models.MakeProject()
	vmi := models.MakeVirtualMachineInterface()

	if lr.ParentUUID == "project-uuid-1" {
		project.VxlanRouting = true
		vmi.ParentType = "virtual-machine"
		vmi.VirtualNetworkRefs = []*models.VirtualMachineInterfaceVirtualNetworkRef{
			{
				UUID: "virtual-network-uuid-1",
			},
		}
	}

	readService.EXPECT().GetProject(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetProjectResponse{
			Project: project,
		}, nil).AnyTimes()

	readService.EXPECT().GetVirtualMachineInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetVirtualMachineInterfaceResponse{
			VirtualMachineInterface: vmi,
		}, nil).AnyTimes()
}

func TestCreateLogicalRouter(t *testing.T) {
	tests := []struct {
		name                  string
		testLogicalRouter     models.LogicalRouter
		expectedLogicalRouter models.LogicalRouter
		errorCode             codes.Code
	}{
		{
			name:              "Try to create logical-router when cannot find parent project",
			testLogicalRouter: models.LogicalRouter{},
			errorCode:         codes.InvalidArgument,
		},
		{
			name: "Try to create logical-router when external gateway with enabled vxlan routing",
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
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create logical-router when logical router interface and gateway in the same network",
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
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "virtual-machine-interface-1",
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create logical-router when port already in use by virtual-machine",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID: "project-uuid-1",
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "virtual-machine-interface-1",
					},
				},
			},
			errorCode: codes.AlreadyExists,
		},
		{
			name: "Create logical-router properly",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID: "project-uuid-2",
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "virtual-machine-interface-1",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			logicalRouterSetupNextServiceMocks(service)
			logicalRouterSetupReadServiceMocks(service, &tt.testLogicalRouter)

			ctx := context.Background()

			paramRequest := services.CreateLogicalRouterRequest{LogicalRouter: &tt.testLogicalRouter}
			expectedResponse := services.CreateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}
			createLogicalRouterResponse, err := service.CreateLogicalRouter(ctx, &paramRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createLogicalRouterResponse)
			}
		})
	}
}

func TestUpdateLogicalRouter(t *testing.T) {
	tests := []struct {
		name              string
		testLogicalRouter models.LogicalRouter
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
			fieldMaskPaths: []string{"parent_uuid", "virtual_network_refs"},
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
			fieldMaskPaths: []string{"parent_uuid", "virtual_network_refs"},
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
			fieldMaskPaths: []string{"parent_uuid", "virtual_machine_interface_refs"},
			errorCode:      codes.AlreadyExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			logicalRouterSetupNextServiceMocks(service)
			logicalRouterSetupReadServiceMocks(service, &tt.testLogicalRouter)

			readService := service.ReadService.(*servicesmock.MockReadService)
			readService.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.GetLogicalRouterResponse{
					LogicalRouter: &tt.dbLogicalRouter,
				}, nil).AnyTimes()

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
