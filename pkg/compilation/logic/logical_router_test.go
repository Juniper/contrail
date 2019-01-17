package logic

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilation/dependencies"
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/testutil"
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
)

func TestCreateLogicalRouterCreatesRouteTarget(t *testing.T) {
	tests := []struct {
		name                string
		testLogicalRouter   *models.LogicalRouter
		returnedRT          *models.RouteTarget
		logicalRouterIntent *LogicalRouterIntent
	}{
		{
			name: "without route target refs",
			testLogicalRouter: &models.LogicalRouter{
				UUID:       "ffe0e3e8-b035-11e8-8981-529269fb1459",
				ParentUUID: "aae0e050-b035-11e8-8981-529269fb1459",
				ParentType: "project",
				FQName: []string{
					"default-domain",
					"project-blue",
					"logical_router_blue",
				},
			},
			returnedRT: &models.RouteTarget{
				FQName:      []string{"target:64512:8000002"},
				UUID:        "default-route-target",
				DisplayName: "target:64512:8000002",
			},
			logicalRouterIntent: &LogicalRouterIntent{
				defaultRouteTargetUUID: "default-route-target",
			},
		},
		{
			name: "with existing route target ref",
			testLogicalRouter: &models.LogicalRouter{
				UUID:       "ffe0e3e8-b035-11e8-8981-529269fb1659",
				ParentUUID: "ffe0e050-b035-11e8-8981-529269fb1559",
				ParentType: "project",
				FQName: []string{
					"default-domain",
					"project-blue",
					"logical_router_red",
				},
				RouteTargetRefs: []*models.LogicalRouterRouteTargetRef{
					{
						UUID: "default-route-target",
						To:   []string{"target:64512:8000003"},
					},
				},
			},
			logicalRouterIntent: &LogicalRouterIntent{
				defaultRouteTargetUUID: "default-route-target",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAPIClient := servicesmock.NewMockWriteService(mockCtrl)
			mockIntPoolAllocator := typesmock.NewMockIntPoolAllocator(mockCtrl)
			mockReadService := servicesmock.NewMockReadService(mockCtrl)
			cache := intent.NewCache()
			service := NewService(
				mockAPIClient,
				mockReadService,
				mockIntPoolAllocator,
				cache,
				dependencies.NewDependencyProcessor(parseReactions(t)),
			)

			mockReadService.EXPECT().GetProject(
				testutil.NotNil(),
				testutil.NotNil(),
			).Return(&services.GetProjectResponse{Project: &models.Project{VxlanRouting: false}},
				nil,
			).AnyTimes()

			if tt.returnedRT == nil {
				expectCreateRTinLR(mockAPIClient, tt.returnedRT, 0)
			} else {
				expectAllocateInt(mockIntPoolAllocator, routeTargetIntPoolID)
				expectCreateRTinLR(mockAPIClient, tt.returnedRT, 1)
			}

			_, err := service.CreateLogicalRouter(context.Background(), &services.CreateLogicalRouterRequest{
				LogicalRouter: tt.testLogicalRouter,
			})

			assert.NoError(t, err)

			lrIntent := LoadLogicalRouterIntent(cache, intent.ByUUID(tt.testLogicalRouter.GetUUID()))

			if tt.logicalRouterIntent != nil {
				if assert.NotNil(t, lrIntent) {
					assert.Equal(t, tt.logicalRouterIntent.defaultRouteTargetUUID, lrIntent.defaultRouteTargetUUID)
				}
			}
		})
	}
}

func expectCreateRTinLR(
	mockAPIClient *servicesmock.MockWriteService,
	returnedRT *models.RouteTarget,
	times int,
) {
	mockAPIClient.EXPECT().CreateRouteTarget(
		testutil.NotNil(),
		testutil.NotNil(),
	).Return(
		&services.CreateRouteTargetResponse{RouteTarget: returnedRT},
		nil,
	).Times(times)

	mockAPIClient.EXPECT().CreateLogicalRouterRouteTargetRef(
		testutil.NotNil(), testutil.NotNil(),
	).Return(nil, nil).Times(times)
}

func TestCreateRefToDefaultRouteTargetInRoutingInstance(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAPIClient := servicesmock.NewMockWriteService(mockCtrl)
	mockReadService := servicesmock.NewMockReadService(mockCtrl)
	mockIntPoolAllocator := typesmock.NewMockIntPoolAllocator(mockCtrl)
	cache := intent.NewCache()
	service := NewService(
		mockAPIClient,
		mockReadService,
		mockIntPoolAllocator,
		cache,
		dependencies.NewDependencyProcessor(parseReactions(t)),
	)

	rt := &models.RouteTarget{
		FQName:      []string{"target:64512:8000002"},
		DisplayName: "target:64512:8000002",
		UUID:        "default-route-target",
	}

	ri := &models.RoutingInstance{
		UUID: "test-ri",
		FQName: []string{
			"default-domain",
			"project-blue",
			"test-vn",
			"test-vn",
		},
	}

	vn := &models.VirtualNetwork{
		UUID: "test-vn",
		FQName: []string{
			"default-domain",
			"project-blue",
			"test-vn",
		},
	}

	vmi := &models.VirtualMachineInterface{
		UUID: "test-vmi",
		VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
			{
				UUID: vn.UUID,
			},
		},
	}

	lr := &models.LogicalRouter{
		UUID:       "ffe0e3e8-b035-11e8-8981-529269fb1659",
		ParentUUID: "ffe0e050-b035-11e8-8981-529269fb1559",
		ParentType: "project",
		FQName: []string{
			"default-domain",
			"project-blue",
			"logical_router_red",
		},
		VirtualMachineInterfaceRefs: []*models.LogicalRouterVirtualMachineInterfaceRef{
			{
				UUID: vmi.UUID,
			},
		},
	}

	_, err := service.CreateRoutingInstance(context.Background(), &services.CreateRoutingInstanceRequest{
		RoutingInstance: ri,
	})
	assert.NoError(t, err)

	mockIntPoolAllocator.EXPECT().AllocateInt(testutil.NotNil(), routeTargetIntPoolID).Return(int64(800002), "", nil)

	_, err = service.CreateVirtualNetwork(context.Background(), &services.CreateVirtualNetworkRequest{
		VirtualNetwork: vn,
	})
	assert.NoError(t, err)

	_, err = service.CreateVirtualMachineInterface(context.Background(), &services.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
	})
	assert.NoError(t, err)

	mockReadService.EXPECT().GetProject(
		testutil.NotNil(),
		testutil.NotNil(),
	).Return(&services.GetProjectResponse{Project: &models.Project{VxlanRouting: false}},
		nil,
	).Times(1)

	mockAPIClient.EXPECT().CreateRouteTarget(
		testutil.NotNil(),
		testutil.NotNil(),
	).Return(&services.CreateRouteTargetResponse{RouteTarget: rt},
		nil,
	).Times(1)

	mockAPIClient.EXPECT().CreateLogicalRouterRouteTargetRef(
		testutil.NotNil(), testutil.NotNil(),
	).Return(nil, nil).Times(1)

	mockAPIClient.EXPECT().CreateRoutingInstanceRouteTargetRef(
		testutil.NotNil(),
		&services.CreateRoutingInstanceRouteTargetRefRequest{
			ID: ri.UUID,
			RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
				UUID: rt.GetUUID(),
			},
		},
	).Times(1)

	_, err = service.CreateLogicalRouter(context.Background(), &services.CreateLogicalRouterRequest{
		LogicalRouter: lr,
	})
	assert.NoError(t, err)
}

func expectAllocateInt(mock *typesmock.MockIntPoolAllocator, poolKey string) {
	mock.EXPECT().AllocateInt(testutil.NotNil(), poolKey).Return(int64(800002), "", nil)
}
