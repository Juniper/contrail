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
			logicalRouterSetupReadServiceMocks(service, &tt.testLogicalRouter)

			ctx := context.Background()

			paramRequest := services.CreateLogicalRouterRequest{LogicalRouter: &tt.testLogicalRouter}
			expectedResponse := services.CreateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}

			createCall := service.Next().(*servicesmock.MockService).EXPECT().CreateLogicalRouter(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, request *services.CreateLogicalRouterRequest,
				) (response *services.CreateLogicalRouterResponse, err error) {
					return &services.CreateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}, nil
				},
			)

			if tt.errorCode != codes.OK {
				createCall.MaxTimes(1)
			} else {
				createCall.Times(1)
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			logicalRouterSetupReadServiceMocks(service, &tt.testLogicalRouter)

			readService := service.ReadService.(*servicesmock.MockReadService)
			readService.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.GetLogicalRouterResponse{
					LogicalRouter: &tt.dbLogicalRouter,
				},
				nil,
			).AnyTimes()

			updateCall := service.Next().(*servicesmock.MockService).EXPECT().UpdateLogicalRouter(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, request *services.UpdateLogicalRouterRequest,
				) (response *services.UpdateLogicalRouterResponse, err error) {
					return &services.UpdateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}, nil
				},
			)

			if tt.errorCode != codes.OK {
				updateCall.MaxTimes(1)
			} else {
				updateCall.Times(1)
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

func TestCheckRouterSupportsVpnType(t *testing.T) {
	tests := []struct {
		name              string
		testLogicalRouter *models.LogicalRouter
		bgpvpnType        string
		errorCode         codes.Code
	}{
		{
			name: "no bgpvpn",
		},
		{
			name: "l2 type bgpvpn",
			testLogicalRouter: &models.LogicalRouter{
				BGPVPNRefs: []*models.LogicalRouterBGPVPNRef{
					{
						UUID: "bgpvpn-1",
					},
				},
			},
			bgpvpnType: models.L2VPNType,
			errorCode:  codes.InvalidArgument,
		},
		{
			name: "l3 type bgpvpn",
			testLogicalRouter: &models.LogicalRouter{
				BGPVPNRefs: []*models.LogicalRouterBGPVPNRef{
					{
						UUID: "bgpvpn-1",
					},
				},
			},
			bgpvpnType: models.L3VPNType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			readService := service.ReadService.(*servicesmock.MockReadService)

			bgpvpn := models.MakeBGPVPN()
			bgpvpn.BGPVPNType = tt.bgpvpnType
			readService.EXPECT().ListBGPVPN(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.ListBGPVPNResponse{
					BGPVPNs: []*models.BGPVPN{bgpvpn},
				}, nil).AnyTimes()

			ctx := context.Background()
			err := service.checkRouterSupportsVpnType(ctx, tt.testLogicalRouter)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCheckRouterHasBgpvpnAssocViaNetwork(t *testing.T) {
	tests := []struct {
		name              string
		testLogicalRouter *models.LogicalRouter
		dbLogicalRouter   *models.LogicalRouter
		vmiVnRefs         []*models.VirtualMachineInterfaceVirtualNetworkRef
		vnBgpvpnRefs      []*models.VirtualNetworkBGPVPNRef
		errorCode         codes.Code
	}{
		{
			name: "no bgpvpn refs",
		},
		{
			name: "no vmi refs",
			testLogicalRouter: &models.LogicalRouter{
				BGPVPNRefs: []*models.LogicalRouterBGPVPNRef{
					{
						UUID: "bgpvpn-1",
					},
				},
			},
			dbLogicalRouter: &models.LogicalRouter{
				BGPVPNRefs: []*models.LogicalRouterBGPVPNRef{
					{
						UUID: "bgpvpn-2",
					},
				},
			},
		},
		{
			name: "no vn refs in vmi",
			testLogicalRouter: &models.LogicalRouter{
				BGPVPNRefs: []*models.LogicalRouterBGPVPNRef{
					{
						UUID: "bgpvpn-1",
					},
				},
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "vmi-1",
					},
				},
			},
		},
		{
			name: "vn with no bgpvpn in vmi",
			testLogicalRouter: &models.LogicalRouter{
				BGPVPNRefs: []*models.LogicalRouterBGPVPNRef{
					{
						UUID: "bgpvpn-1",
					},
				},
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "vmi-1",
					},
				},
			},
			vmiVnRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
				{
					UUID: "vn-1",
				},
			},
		},
		{
			name: "vn with bgpvpn in vmi refs",
			testLogicalRouter: &models.LogicalRouter{
				BGPVPNRefs: []*models.LogicalRouterBGPVPNRef{
					{
						UUID: "bgpvpn-1",
					},
				},
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "vmi-1",
					},
				},
			},
			vmiVnRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
				{
					UUID: "vn-1",
				},
			},
			vnBgpvpnRefs: []*models.VirtualNetworkBGPVPNRef{
				{
					UUID: "bgpvpn-2",
				},
			},
			errorCode: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			readService := service.ReadService.(*servicesmock.MockReadService)

			vmi := models.MakeVirtualMachineInterface()
			vmi.VirtualNetworkRefs = tt.vmiVnRefs
			readService.EXPECT().ListVirtualMachineInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.ListVirtualMachineInterfaceResponse{
					VirtualMachineInterfaces: []*models.VirtualMachineInterface{vmi},
				},
				nil,
			).AnyTimes()

			vn := models.MakeVirtualNetwork()
			vn.BGPVPNRefs = tt.vnBgpvpnRefs
			readService.EXPECT().ListVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.ListVirtualNetworkResponse{
					VirtualNetworks: []*models.VirtualNetwork{vn},
				},
				nil,
			).AnyTimes()

			ctx := context.Background()
			err := service.checkRouterHasBgpvpnAssocViaNetwork(ctx, tt.testLogicalRouter, tt.dbLogicalRouter)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
