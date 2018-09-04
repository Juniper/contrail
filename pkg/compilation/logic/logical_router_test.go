package logic

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/testutil"
)

func TestCreateLogicalRouterCreatesRouteTarget(t *testing.T) {
	tests := []struct {
		name              string
		createsRT         bool
		testLogicalRouter *models.LogicalRouter
		expectedRT        *models.RouteTarget
	}{
		{
			name:      "Try to process creation of logical router without route target refs",
			createsRT: false,
			testLogicalRouter: &models.LogicalRouter{
				UUID:       "ffe0e3e8-b035-11e8-8981-529269fb1459",
				ParentUUID: "ffe0e050-b035-11e8-8981-529269fb1459",
				ParentType: "project",
				FQName: []string{
					"default-domain",
					"project-blue",
					"logical_router_blue"},
			},
			expectedRT: &models.RouteTarget{
				FQName:      []string{"target:64512:8000002"},
				DisplayName: "target:64512:8000002",
			},
		},
		{
			name:      "Try to process creation of logical router with existing route target ref",
			createsRT: true,
			testLogicalRouter: &models.LogicalRouter{
				UUID:       "ffe0e3e8-b035-11e8-8981-529269fb1659",
				ParentUUID: "ffe0e050-b035-11e8-8981-529269fb1559",
				ParentType: "project",
				FQName: []string{
					"default-domain",
					"project-blue",
					"logical_router_red"},
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
			compilationif.Init()
			mockAPIService := servicesmock.NewMockWriteService(mockCtrl)
			service := NewService(mockAPIService)

			if tt.createsRT {
				expectCreateRTinLR(mockAPIService, tt.expectedRT, 0)
			} else {
				expectCreateRTinLR(mockAPIService, tt.expectedRT, 1)
			}

			_, err := service.CreateLogicalRouter(context.Background(), &services.CreateLogicalRouterRequest{
				LogicalRouter: tt.testLogicalRouter,
			})
			assert.NoError(t, err)

		})
	}
}

func expectCreateRTinLR(mockAPIService *servicesmock.MockWriteService, expectedRT *models.RouteTarget, times int) {
	mockAPIService.EXPECT().CreateRouteTarget(
		testutil.NotNil(),
		testutil.NotNil(),
	).Return(&services.CreateRouteTargetResponse{RouteTarget: expectedRT},
		nil,
	).Times(times)

	mockAPIService.EXPECT().CreateLogicalRouterRouteTargetRef(
		testutil.NotNil(), testutil.NotNil(),
	).Return(nil, nil).Times(times)
}
