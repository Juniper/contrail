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

func TestCreateRoutingInstanceCreatesRouteTarget(t *testing.T) {
	routingInstance := &models.RoutingInstance{
		UUID:       "a8edf702-7dd6-4cbc-b599-a7f8ace1d22b",
		ParentUUID: "af68b258-6fc4-4959-8181-a2cfb6f93500",
		ParentType: "virtual-network",
		FQName: []string{
			"default-domain",
			"project-blue",
			"vn-blue",
			"vn-blue"},
		RoutingInstanceIsDefault: true,
	}

	expectedRT := &models.RouteTarget{
		FQName:      []string{"target:64512:8000002"},
		DisplayName: "target:64512:8000002",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAPIService := servicesmock.NewMockWriteService(mockCtrl)
	service := NewService(mockAPIService)

	expectCreateRT(mockAPIService, expectedRT)

	compilationif.Init()
	_, err := service.CreateRoutingInstance(context.Background(), &services.CreateRoutingInstanceRequest{
		RoutingInstance: routingInstance,
	})

	assert.NoError(t, err)
}

func expectCreateRT(mockAPIService *servicesmock.MockWriteService, expectedRT *models.RouteTarget) {
	// TODO Revisit code below when route target is allocated properly.
	mockAPIService.EXPECT().CreateRouteTarget(
		testutil.NotNil(), testutil.NotNil()).Return(&services.CreateRouteTargetResponse{
		RouteTarget: expectedRT}, nil).Times(1)

	mockAPIService.EXPECT().CreateRoutingInstanceRouteTargetRef(
		testutil.NotNil(), testutil.NotNil(),
	).Return(nil, nil).Times(1)
}
