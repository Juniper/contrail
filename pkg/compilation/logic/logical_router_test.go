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
			cache := intent.NewCache()
			mockAPIService := servicesmock.NewMockWriteService(mockCtrl)
			service := NewService(mockAPIService, cache)

			if tt.returnedRT == nil {
				expectCreateRTinLR(mockAPIService, tt.returnedRT, 0)
			} else {
				expectCreateRTinLR(mockAPIService, tt.returnedRT, 1)
			}

			_, err := service.CreateLogicalRouter(context.Background(), &services.CreateLogicalRouterRequest{
				LogicalRouter: tt.testLogicalRouter,
			})
			assert.NoError(t, err)

		})
	}
}

func expectCreateRTinLR(
	mockAPIService *servicesmock.MockWriteService,
	returnedRT *models.RouteTarget,
	times int) {
	mockAPIService.EXPECT().CreateRouteTarget(
		testutil.NotNil(),
		testutil.NotNil(),
	).Return(&services.CreateRouteTargetResponse{RouteTarget: returnedRT},
		nil,
	).Times(times)

	mockAPIService.EXPECT().CreateLogicalRouterRouteTargetRef(
		testutil.NotNil(), testutil.NotNil(),
	).Return(nil, nil).Times(times)
}
