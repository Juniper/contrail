package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

func logicalInterfaceNextServMocks(t *testing.T, service *ContrailTypeLogicService) {
	nextServiceMock, ok := service.Next().(*servicesmock.MockService)
	assert.True(t, ok)
	nextServiceMock.EXPECT().CreateLogicalInterface(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.CreateLogicalInterfaceRequest,
		) (response *services.CreateLogicalInterfaceResponse, err error) {
			return &services.CreateLogicalInterfaceResponse{LogicalInterface: request.LogicalInterface}, nil
		}).AnyTimes()
	nextServiceMock.EXPECT().UpdateLogicalInterface(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.UpdateLogicalInterfaceRequest,
		) (response *services.UpdateLogicalInterfaceResponse, err error) {
			return &services.UpdateLogicalInterfaceResponse{LogicalInterface: request.LogicalInterface}, nil
		}).AnyTimes()
}

func logicalInterfaceReadServiceMocks(
	t *testing.T,
	s *ContrailTypeLogicService,
	parentRouter *models.PhysicalRouter,
	listPhysicalInterface []*models.PhysicalInterface,
) {
	readService, ok := s.ReadService.(*servicesmock.MockReadService)
	assert.True(t, ok)
	// Use empty physical router structure instead nil parent physical rounter
	if parentRouter == nil {
		parentRouter = new(models.PhysicalRouter)
	}
	readService.EXPECT().GetPhysicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetPhysicalRouterResponse{
			PhysicalRouter: parentRouter,
		}, nil).AnyTimes()
	// Try to find requested physical interface within parent router data
	readService.EXPECT().GetPhysicalInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(
			_ context.Context, request *services.GetPhysicalInterfaceRequest,
		) (response *services.GetPhysicalInterfaceResponse, err error) {
			for _, pi := range parentRouter.PhysicalInterfaces {
				if pi.UUID == request.ID {
					return &services.GetPhysicalInterfaceResponse{PhysicalInterface: pi}, nil
				}
			}
			return nil, grpc.Errorf(codes.NotFound, "physical interface with uuid %s not found", request.ID)
		}).AnyTimes()
	// Use empty physical interface slice instead nil others physical interface list
	readService.EXPECT().ListPhysicalInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.ListPhysicalInterfaceResponse{
			PhysicalInterfaces: listPhysicalInterface,
		}, nil).AnyTimes()
	// Try to find requested logical interface within current physical router and/or list of physical interface data
	readService.EXPECT().GetLogicalInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(
			_ context.Context, request *services.GetLogicalInterfaceRequest,
		) (response *services.GetLogicalInterfaceResponse, err error) {
			for _, pi := range parentRouter.PhysicalInterfaces {
				for _, li := range pi.LogicalInterfaces {
					if li.UUID == request.ID {
						return &services.GetLogicalInterfaceResponse{LogicalInterface: li}, nil
					}
				}
			}
			for _, pi := range listPhysicalInterface {
				for _, li := range pi.LogicalInterfaces {
					if li.UUID == request.ID {
						return &services.GetLogicalInterfaceResponse{LogicalInterface: li}, nil
					}
				}
			}
			return nil, grpc.Errorf(codes.NotFound, "logical interface with uuid %s not found", request.ID)
		}).AnyTimes()
}

func TestCreateLogicalInterface(t *testing.T) {
	tests := []struct {
		name                  string
		createRequest         *services.CreateLogicalInterfaceRequest
		parentRouter          *models.PhysicalRouter
		listPhysicalInterface []*models.PhysicalInterface
		errorCode             codes.Code
	}{
		{
			name: "Try create logical interface without display name and vlan tag",
			createRequest: &services.CreateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID: "uuid",
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try create logical interface with display name and vlan tag",
			createRequest: &services.CreateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:                    "uuid",
				DisplayName:             "display_name",
				LogicalInterfaceVlanTag: 1024,
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try create logical interface with wrong vlan tag",
			createRequest: &services.CreateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:                    "uuid",
				LogicalInterfaceVlanTag: 9999,
			}},
			errorCode: codes.PermissionDenied,
		},
		{
			name: "Try create logical interface with wrong display name",
			createRequest: &services.CreateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:        "uuid",
				DisplayName: "double",
			}},
			parentRouter: &models.PhysicalRouter{
				UUID: "uuid",
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:        "uuid",
					DisplayName: "double",
				}},
			},
			errorCode: codes.AlreadyExists,
		},
		{
			name: "Try create logical interface with QFX type and valid vlan tag",
			createRequest: &services.CreateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:                    "uuid",
				DisplayName:             "display_name",
				LogicalInterfaceType:    "l2",
				LogicalInterfaceVlanTag: 1024,
				ParentType:              "physical-interface",
				ParentUUID:              "physical_interface_uuid",
			}},
			parentRouter: &models.PhysicalRouter{
				UUID:                      "router_uuid",
				PhysicalRouterProductName: "QFX_all",
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID: "physical_interface_uuid",
				}},
			},
			errorCode: codes.OK,
		},
		{
			name: "Try create logical interface with QFX type and wrong vlan tag",
			createRequest: &services.CreateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:                    "uuid",
				DisplayName:             "display_name",
				LogicalInterfaceType:    "l2",
				LogicalInterfaceVlanTag: 2,
				ParentType:              "physical-interface",
				ParentUUID:              "physical_interface_uuid",
			}},
			parentRouter: &models.PhysicalRouter{
				UUID:                      "router_uuid",
				PhysicalRouterProductName: "QFX_all",
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID: "physical_interface_uuid",
				}},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try create logical interface with same VMI refs",
			createRequest: &services.CreateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:        "uuid",
				ParentType:  "physical-interface",
				ParentUUID:  "physical_interface_uuid",
				DisplayName: "display_name",
				VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
					UUID: "vmi0",
				}},
			}},
			parentRouter: &models.PhysicalRouter{
				UUID: "physical_interface_uuid",
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:                      "physical_interface_uuid",
					EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0",
				}},
			},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID:                      "physical_interface_uuid2",
				EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID: "logical_interface_uuid",
					VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
						UUID: "vmi0",
					}},
				}},
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try create logical interface with not equal VMI refs",
			createRequest: &services.CreateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:        "uuid",
				ParentType:  "physical-interface",
				ParentUUID:  "physical_interface_uuid",
				DisplayName: "display_name",
				VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
					UUID: "vmi0",
				}},
			}},
			parentRouter: &models.PhysicalRouter{
				UUID: "physical_interface_uuid",
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:                      "physical_interface_uuid",
					EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0",
				}},
			},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID:                      "physical_interface_uuid2",
				EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID: "logical_interface_uuid",
					VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
						UUID: "vmi0",
					}, {
						UUID: "vmi9",
					}},
				}},
			}},
			errorCode: codes.PermissionDenied,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			logicalInterfaceNextServMocks(t, service)
			ctx := context.Background()
			logicalInterfaceReadServiceMocks(t, service, tt.parentRouter, tt.listPhysicalInterface)

			createLogicalInterfaceResponse, err := service.CreateLogicalInterface(ctx, tt.createRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err, "create succeeded but shouldn't")
				assert.Nil(t, createLogicalInterfaceResponse)
				assert.Equal(t, tt.errorCode, errutil.CauseCode(err))
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, createLogicalInterfaceResponse)
			}
		})
	}
}

func TestUpdateLogicalInterface(t *testing.T) {
	tests := []struct {
		name                  string
		updateRequest         *services.UpdateLogicalInterfaceRequest
		parentRouter          *models.PhysicalRouter
		listPhysicalInterface []*models.PhysicalInterface
		errorCode             codes.Code
	}{
		{
			name: "Try update logical interface without display name and vlan tag",
			updateRequest: &services.UpdateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID: "uuid",
			}},
			parentRouter: &models.PhysicalRouter{},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID: "uuid",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID: "uuid",
				}},
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try update logical interface with wrong vlan tag",
			updateRequest: &services.UpdateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:                    "uuid",
				LogicalInterfaceVlanTag: 9999,
			}},
			parentRouter: &models.PhysicalRouter{},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID: "uuid",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID:                    "uuid",
					LogicalInterfaceVlanTag: 1024,
				}},
			}},
			errorCode: codes.PermissionDenied,
		},
		{
			name: "Try update logical interface without vlan tag",
			updateRequest: &services.UpdateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:                    "uuid",
				LogicalInterfaceVlanTag: 1024,
			}},
			parentRouter: &models.PhysicalRouter{},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID: "uuid",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID:                    "uuid",
					LogicalInterfaceVlanTag: 0, // need "not defined" flag
				}},
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try update logical interface with non equal vlan tag",
			updateRequest: &services.UpdateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:                    "uuid",
				LogicalInterfaceVlanTag: 2048,
			}},
			parentRouter: &models.PhysicalRouter{},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID: "uuid",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID:                    "uuid",
					LogicalInterfaceVlanTag: 1024,
				}},
			}},
			errorCode: codes.PermissionDenied,
		},
		{
			name: "Try update logical interface with non equal display name",
			updateRequest: &services.UpdateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:        "uuid",
				DisplayName: "second",
			}},
			parentRouter: &models.PhysicalRouter{},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID: "uuid",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID:        "uuid",
					DisplayName: "first",
				}},
			}},
			errorCode: codes.PermissionDenied,
		},
		{
			name: "Try update logical interface with same VMI refs",
			updateRequest: &services.UpdateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:        "uuid",
				ParentType:  "physical-interface",
				ParentUUID:  "physical_interface_uuid",
				DisplayName: "display_name",
				VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
					UUID: "vmi0",
				}},
			}},
			parentRouter: &models.PhysicalRouter{
				UUID: "physical_interface_uuid",
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:                      "physical_interface_uuid",
					EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0",
				}},
			},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID:                      "physical_interface_uuid2",
				EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID: "logical_interface_uuid",
					VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
						UUID: "vmi0",
					}},
				}, {
					UUID:        "uuid",
					ParentType:  "physical-interface",
					ParentUUID:  "physical_interface_uuid",
					DisplayName: "display_name",
				}},
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try update logical interface with not equal VMI refs",
			updateRequest: &services.UpdateLogicalInterfaceRequest{LogicalInterface: &models.LogicalInterface{
				UUID:        "uuid",
				ParentType:  "physical-interface",
				ParentUUID:  "physical_interface_uuid",
				DisplayName: "display_name",
				VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
					UUID: "vmi0",
				}},
			}},
			parentRouter: &models.PhysicalRouter{
				UUID: "physical_interface_uuid",
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:                      "physical_interface_uuid",
					EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0",
				}},
			},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID:                      "physical_interface_uuid2",
				EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID: "logical_interface_uuid",
					VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
						UUID: "vmi0",
					}, {
						UUID: "vmi9",
					}},
				}, {
					UUID:        "uuid",
					ParentType:  "physical-interface",
					ParentUUID:  "physical_interface_uuid",
					DisplayName: "display_name",
				}},
			}},
			errorCode: codes.PermissionDenied,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			logicalInterfaceNextServMocks(t, service)
			ctx := context.Background()
			logicalInterfaceReadServiceMocks(t, service, tt.parentRouter, tt.listPhysicalInterface)

			updateLogicalInterfaceResponse, err := service.UpdateLogicalInterface(ctx, tt.updateRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err, "update succeeded but shouldn't")
				assert.Nil(t, updateLogicalInterfaceResponse)
				assert.Equal(t, tt.errorCode, errutil.CauseCode(err))
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, updateLogicalInterfaceResponse)
			}
		})
	}
}
