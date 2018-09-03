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
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
)

func logicalRouterSetupReadServiceMocks(
	s *ContrailTypeLogicService,
	lr *models.LogicalRouter,
	vxlan bool) {
	readService := s.ReadService.(*servicesmock.MockReadService)
	project := models.MakeProject()
	project.VxlanRouting = vxlan
	vmi := models.MakeVirtualMachineInterface()

	if lr.ParentUUID == "project-uuid-1" {
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
		},
		nil,
	).AnyTimes()

	readService.EXPECT().GetVirtualMachineInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetVirtualMachineInterfaceResponse{
			VirtualMachineInterface: vmi,
		},
		nil,
	).AnyTimes()
}

func TestCreateLogicalRouter(t *testing.T) {
	tests := []struct {
		name                  string
		testLogicalRouter     models.LogicalRouter
		expectedLogicalRouter models.LogicalRouter
		vxlanEnabled          bool
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
			vxlanEnabled: true,
			errorCode:    codes.InvalidArgument,
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
		{
			name: "Create logical-router with vxlan enabled",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID: "project-uuid-2",
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "virtual-machine-interface-1",
					},
				},
			},
			vxlanEnabled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			logicalRouterSetupReadServiceMocks(service, &tt.testLogicalRouter, tt.vxlanEnabled)

			ctx := context.Background()

			paramRequest := services.CreateLogicalRouterRequest{LogicalRouter: &tt.testLogicalRouter}
			expectedResponse := services.CreateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}

			createLRCall := service.Next().(*servicesmock.MockService).EXPECT().CreateLogicalRouter(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, request *services.CreateLogicalRouterRequest,
				) (response *services.CreateLogicalRouterResponse, err error) {
					return &services.CreateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}, nil
				},
			)

			createVNCall := service.WriteService.(*servicesmock.MockWriteService).EXPECT().CreateVirtualNetwork(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, request *services.CreateVirtualNetworkRequest,
				) (response *services.CreateVirtualNetworkResponse, err error) {
					return &services.CreateVirtualNetworkResponse{VirtualNetwork: request.VirtualNetwork}, nil
				},
			)

			if tt.errorCode != codes.OK {
				createLRCall.MaxTimes(1)
			} else {
				createLRCall.Times(1)
			}

			if tt.errorCode == codes.OK && tt.vxlanEnabled {
				createVNCall.Times(1)
			} else {
				createVNCall.MaxTimes(1)
			}

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
			name: "Update logical-router with vxlan routing enabled",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID:             "project-uuid-1",
				VxlanNetworkIdentifier: "5678",
			},
			dbLogicalRouter: models.LogicalRouter{
				ParentUUID:             "project-uuid-1",
				VxlanNetworkIdentifier: "1234",
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

func TestDeleteLogicalRouter(t *testing.T) {
	tests := []struct {
		name            string
		dbLogicalRouter models.LogicalRouter
		vxlanEnabled    bool
		errorCode       codes.Code
	}{
		{
			name: "Delete logical router with vxlan routing disabled",
			dbLogicalRouter: models.LogicalRouter{
				ParentUUID: "project-uuid-1",
			},
			vxlanEnabled: false,
		},
		{
			name: "Delete logical router with vxlan routing",
			dbLogicalRouter: models.LogicalRouter{
				ParentUUID: "project-uuid-1",
			},
			vxlanEnabled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			logicalRouterSetupReadServiceMocks(service, &tt.dbLogicalRouter, tt.vxlanEnabled)

			readService := service.ReadService.(*servicesmock.MockReadService)
			readService.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.GetLogicalRouterResponse{
					LogicalRouter: &tt.dbLogicalRouter,
				},
				nil,
			).AnyTimes()

			deleteLRCall := service.Next().(*servicesmock.MockService).EXPECT().DeleteLogicalRouter(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).Return(
				&services.DeleteLogicalRouterResponse{}, nil,
			)

			deleteVNCall := service.WriteService.(*servicesmock.MockWriteService).EXPECT().DeleteVirtualNetwork(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).Return(
				&services.DeleteVirtualNetworkResponse{}, nil,
			)

			metadataCall := service.MetadataGetter.(*typesmock.MockMetadataGetter).EXPECT().GetMetaData(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).Return(
				&basemodels.MetaData{
					UUID: "internal-virtual-network-uuid",
				},
				nil,
			)

			if tt.errorCode != codes.OK {
				deleteLRCall.MaxTimes(1)
			} else {
				deleteLRCall.Times(1)
			}

			if tt.errorCode == codes.OK && tt.vxlanEnabled {
				deleteVNCall.Times(1)
				metadataCall.Times(1)
			} else {
				deleteVNCall.MaxTimes(1)
				metadataCall.MaxTimes(1)
			}

			ctx := context.Background()
			paramRequest := services.DeleteLogicalRouterRequest{}
			expectedResponse := services.DeleteLogicalRouterResponse{}
			deleteLogicalRouterResponse, err := service.DeleteLogicalRouter(ctx, &paramRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, deleteLogicalRouterResponse)
			}
		})
	}
}
