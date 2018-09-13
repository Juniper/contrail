package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

type testRTParams struct {
	uuid   string
	rtName string
}

func createTestRouteTarget(testParams *testRTParams) *models.RouteTarget {
	routeTarget := models.MakeRouteTarget()
	routeTarget.UUID = testParams.uuid
	routeTarget.Name = testParams.rtName
	return routeTarget
}

func routeTargetNextServMocks(service *ContrailTypeLogicService) {
	nextServiceMock := service.Next().(*servicesmock.MockService) //nolint: errcheck
	nextServiceMock.EXPECT().CreateRouteTarget(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.CreateRouteTargetRequest,
		) (response *services.CreateRouteTargetResponse, err error) {
			return &services.CreateRouteTargetResponse{RouteTarget: request.RouteTarget}, nil
		}).AnyTimes()
}

func TestCreateRouteTarget(t *testing.T) {
	tests := []struct {
		name       string
		testParams *testRTParams
		errorCode  codes.Code
	}{
		{
			name: "Try route target without name",
			testParams: &testRTParams{
				uuid: "uuid",
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try route target with 'target:<ip>:<number>' format",
			testParams: &testRTParams{
				uuid:   "uuid",
				rtName: "target:10.0.0.0:50",
			},
		},
		{
			name: "Try route target with 'target:<asn>:<number>'format ",
			testParams: &testRTParams{
				uuid:   "uuid",
				rtName: "target:100:200",
			},
		},
		{
			name: "Try route target without wrong prefix",
			testParams: &testRTParams{
				uuid:   "uuid",
				rtName: "rt:100:200",
			},
			errorCode: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(mockCtrl)
			routeTargetNextServMocks(service)

			ctx := context.Background()

			routeTarget := createTestRouteTarget(tt.testParams)
			createRouteTargetRequest := &services.CreateRouteTargetRequest{RouteTarget: routeTarget}
			createRouteTargetResponse, err := service.CreateRouteTarget(ctx, createRouteTargetRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err, "create succeeded but shouldn't")
				assert.Nil(t, createRouteTargetResponse)

				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, createRouteTargetResponse)
			}

			mockCtrl.Finish()
		})
	}

}
