package logic

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/types/mock"
)

func TestCreateLogicalRouterCreatesRouteTarget(t *testing.T) {
	tests := []struct {
		name              string
		testLogicalRouter *models.LogicalRouter
		returnedRT        *models.RouteTarget
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
				DisplayName: "target:64512:8000002",
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
						To: []string{"target:64512:8000003"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAPIClient := servicesmock.NewMockWriteService(mockCtrl)
			mockIntPoolAllocator := typesmock.NewMockIntPoolAllocator(mockCtrl)
			service := NewService(mockAPIClient, mockIntPoolAllocator, intent.NewCache())

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

func expectAllocateInt(mock *typesmock.MockIntPoolAllocator, poolKey string) {
	mock.EXPECT().AllocateInt(testutil.NotNil(), poolKey).Return(int64(800002), nil)
}
