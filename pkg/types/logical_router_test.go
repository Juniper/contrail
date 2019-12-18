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
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
)

const (
	//mockModes
	vmiListRequest = 1 << iota
	vnListRequest
	bgpvpnListRequest
)

func logicalRouterSetupReadServiceMocks(
	s *ContrailTypeLogicService,
	lr *models.LogicalRouter,
	vxlan bool) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
	project := models.MakeProject()
	project.VxlanRouting = vxlan
	vmi := models.MakeVirtualMachineInterface()

	if lr.GetParentUUID() == "project-uuid-1" {
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

func logicalRouterSetupIntPoolAllocatorMocks(s *ContrailTypeLogicService) {
	intPoolAllocator := s.IntPoolAllocator.(*typesmock.MockIntPoolAllocator) //nolint: errcheck
	intPoolAllocator.EXPECT().AllocateInt(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).Return(int64(1), nil).AnyTimes()
	intPoolAllocator.EXPECT().SetInt(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), int64(2), gomock.Not(gomock.Nil()),
	).Return(nil).AnyTimes()
	intPoolAllocator.EXPECT().DeallocateInt(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).Return(nil).AnyTimes()
}

func TestCreateLogicalRouter(t *testing.T) {
	tests := []struct {
		name              string
		testLogicalRouter models.LogicalRouter
		vxlanEnabled      bool
		errorCode         codes.Code
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
			name: "Create logical-router with vxlan enabled and no vxlan id",
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
		{
			name: "Try to create logical-router with improper vxlan id",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID:             "project-uuid-2",
				VxlanNetworkIdentifier: "id",
				VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
					{
						UUID: "virtual-machine-interface-1",
					},
				},
			},
			vxlanEnabled: true,
			errorCode:    codes.InvalidArgument,
		},
		{
			name: "Create logical-router with vxlan enabled",
			testLogicalRouter: models.LogicalRouter{
				ParentUUID:             "project-uuid-2",
				VxlanNetworkIdentifier: "2",
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
			logicalRouterSetupIntPoolAllocatorMocks(service)

			ctx := context.Background()

			paramRequest := services.CreateLogicalRouterRequest{LogicalRouter: &tt.testLogicalRouter}
			expectedResponse := services.CreateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}

			createLRCall := service.Next().(*servicesmock.MockService).EXPECT().CreateLogicalRouter(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, _ *services.CreateLogicalRouterRequest,
				) (response *services.CreateLogicalRouterResponse, err error) {
					return &services.CreateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}, nil
				},
			)

			createVNCall := service.WriteService.(*servicesmock.MockWriteService).EXPECT( //nolint: errcheck
			).CreateVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
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
	tests := []struct { // nolint: maligned
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

			readService := service.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
			readService.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.GetLogicalRouterResponse{
					LogicalRouter: &tt.dbLogicalRouter,
				},
				nil,
			).AnyTimes()

			updateLRCall := service.Next().(*servicesmock.MockService).EXPECT().UpdateLogicalRouter( //nolint: errcheck
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, _ *services.UpdateLogicalRouterRequest,
				) (response *services.UpdateLogicalRouterResponse, err error) {
					return &services.UpdateLogicalRouterResponse{LogicalRouter: &tt.testLogicalRouter}, nil
				},
			)

			updateVNCall := service.WriteService.(*servicesmock.MockWriteService).EXPECT( //nolint: errcheck
			).UpdateVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
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
		dbLogicalRouter *models.LogicalRouter
		vxlanEnabled    bool
		errorCode       codes.Code
	}{
		{
			name:         "Try to delete logical router when cannot be found",
			vxlanEnabled: false,
			errorCode:    codes.NotFound,
		},
		{
			name: "Delete logical router with vxlan routing disabled",
			dbLogicalRouter: &models.LogicalRouter{
				ParentUUID: "project-uuid-1",
			},
			vxlanEnabled: false,
		},
		{
			name: "Delete logical router with vxlan routing",
			dbLogicalRouter: &models.LogicalRouter{
				VxlanNetworkIdentifier: "1",
				ParentUUID:             "project-uuid-1",
			},
			vxlanEnabled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			logicalRouterSetupReadServiceMocks(service, tt.dbLogicalRouter, tt.vxlanEnabled)
			logicalRouterSetupIntPoolAllocatorMocks(service)

			readService := service.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
			if tt.dbLogicalRouter != nil {
				readService.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					&services.GetLogicalRouterResponse{
						LogicalRouter: tt.dbLogicalRouter,
					}, nil).AnyTimes()
			} else {
				readService.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					nil, errutil.ErrorNotFound).AnyTimes()
			}

			deleteLRCall := service.Next().(*servicesmock.MockService).EXPECT().DeleteLogicalRouter(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).Return(
				&services.DeleteLogicalRouterResponse{}, nil,
			)

			deleteVNRefCall := service.WriteService.(*servicesmock.MockWriteService).EXPECT( //nolint: errcheck
			).DeleteLogicalRouterVirtualNetworkRef(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.DeleteLogicalRouterVirtualNetworkRefResponse{}, nil,
			)

			deleteVNCall := service.WriteService.(*servicesmock.MockWriteService).EXPECT( //nolint: errcheck
			).DeleteVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.DeleteVirtualNetworkResponse{}, nil,
			)

			metadataCall := service.MetadataGetter.(*typesmock.MockMetadataGetter).EXPECT( //nolint: errcheck
			).GetMetadata(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&basemodels.Metadata{
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
				deleteVNRefCall.Times(1)
			} else {
				deleteVNCall.MaxTimes(1)
				metadataCall.MaxTimes(1)
				deleteVNRefCall.MaxTimes(1)
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

func TestCheckRouterSupportsVPNType(t *testing.T) {
	tests := []struct {
		name              string
		testLogicalRouter *models.LogicalRouter
		bgpvpnType        string
		mockMode          uint
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
			mockMode:   bgpvpnListRequest,
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
			mockMode:   bgpvpnListRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			readService := service.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

			if tt.mockMode&bgpvpnListRequest != 0 {
				bgpvpn := models.MakeBGPVPN()
				bgpvpn.BGPVPNType = tt.bgpvpnType
				readService.EXPECT().ListBGPVPN(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					&services.ListBGPVPNResponse{
						BGPVPNs: []*models.BGPVPN{bgpvpn},
					},
					nil,
				)
			}

			ctx := context.Background()
			err := service.checkRouterSupportsVPNType(ctx, tt.testLogicalRouter)
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

func TestCheckRouterHasBGPVPNAssocViaNetwork(t *testing.T) {
	tests := []struct {
		name              string
		testLogicalRouter *models.LogicalRouter
		dbLogicalRouter   *models.LogicalRouter
		vmiVnRefs         []*models.VirtualMachineInterfaceVirtualNetworkRef
		vnBgpvpnRefs      []*models.VirtualNetworkBGPVPNRef
		fieldMask         *types.FieldMask
		errorCode         codes.Code
		mockMode          uint
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
			fieldMask: &types.FieldMask{Paths: []string{models.LogicalRouterFieldBGPVPNRefs}},
		},
		{
			name: "vmi and bgpvpn only in db",
			dbLogicalRouter: &models.LogicalRouter{
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
			fieldMask: &types.FieldMask{},
		},
		{
			name: "delete vmi and bgpvpn refs",
			dbLogicalRouter: &models.LogicalRouter{
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
			fieldMask: &types.FieldMask{
				Paths: []string{
					models.LogicalRouterFieldBGPVPNRefs,
					models.LogicalRouterFieldVirtualMachineInterfaceRefs,
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
			mockMode: vmiListRequest,
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
			mockMode: vmiListRequest | vnListRequest,
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
			mockMode:  vmiListRequest | vnListRequest,
			errorCode: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			readService := service.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

			if tt.mockMode&vmiListRequest != 0 {
				vmi := models.MakeVirtualMachineInterface()
				vmi.VirtualNetworkRefs = tt.vmiVnRefs
				readService.EXPECT().ListVirtualMachineInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					&services.ListVirtualMachineInterfaceResponse{
						VirtualMachineInterfaces: []*models.VirtualMachineInterface{vmi},
					},
					nil,
				)
			}

			if tt.mockMode&vnListRequest != 0 {
				vn := models.MakeVirtualNetwork()
				vn.BGPVPNRefs = tt.vnBgpvpnRefs
				readService.EXPECT().ListVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					&services.ListVirtualNetworkResponse{
						VirtualNetworks: []*models.VirtualNetwork{vn},
					},
					nil,
				)
			}

			ctx := context.Background()
			err := service.checkRouterHasBGPVPNAssocViaNetwork(ctx, tt.testLogicalRouter, tt.dbLogicalRouter, tt.fieldMask)
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
