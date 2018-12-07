package types

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/models"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/mock/gomock"

	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

type testPhysicalInterfaceParams struct {
	UUID                      string
	DisplayName               string
	EthernetSegmentIdentifier string
}

func physicalInterfaceNextServMocks(service *ContrailTypeLogicService) {
	nextServiceMock := service.Next().(*servicesmock.MockService) //nolint: errcheck
	nextServiceMock.EXPECT().CreatePhysicalInterface(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.CreatePhysicalInterfaceRequest,
		) (response *services.CreatePhysicalInterfaceResponse, err error) {
			return &services.CreatePhysicalInterfaceResponse{PhysicalInterface: request.PhysicalInterface}, nil
		}).AnyTimes()
	nextServiceMock.EXPECT().UpdatePhysicalInterface(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.UpdatePhysicalInterfaceRequest,
		) (response *services.UpdatePhysicalInterfaceResponse, err error) {
			return &services.UpdatePhysicalInterfaceResponse{PhysicalInterface: request.PhysicalInterface}, nil
		}).AnyTimes()
}

func physicalInterfaceReadServiceMocks(
	s *ContrailTypeLogicService,
	parentRouter *models.PhysicalRouter,
) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
	if parentRouter == nil {
		parentRouter = new(models.PhysicalRouter)
	}
	readService.EXPECT().GetPhysicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetPhysicalRouterResponse{
			PhysicalRouter: parentRouter,
		}, nil).AnyTimes()
}

func TestCreatePhysicalInterface(t *testing.T) {
	tests := []struct {
		name                   string
		createRequest          *services.CreatePhysicalInterfaceRequest
		parentRouter           *models.PhysicalRouter
		errorCode              codes.Code
	}{
		{
			name: "Try create physical interface with empty display name and ESI",
			createRequest: &services.CreatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID: "uuid",
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try create physical interface with valid display name and ESI",
			createRequest: &services.CreatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:                      "uuid",
				DisplayName:               "unique",
				EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0:01",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{&models.PhysicalInterface{
					UUID:        "uuid",
					DisplayName: "other",
				}},
			},
			errorCode: codes.OK,
		},
		{
			name: "Try create physical interface with same display name",
			createRequest: &services.CreatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:        "uuid",
				DisplayName: "double",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{&models.PhysicalInterface{
					UUID:        "uuid",
					DisplayName: "double",
				}},
			},
			errorCode: codes.AlreadyExists,
		},
		{
			name: "Try create physical interface with wrong ESI",
			createRequest: &services.CreatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:                      "uuid",
				EthernetSegmentIdentifier: "0123:4567:891A:BCDE:F001",
			}},
			errorCode: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			physicalInterfaceNextServMocks(service)
			ctx := context.Background()

			physicalInterfaceReadServiceMocks(service, tt.parentRouter)

			createPhysicalInterfaceResponse, err := service.CreatePhysicalInterface(ctx, tt.createRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err, "create succeeded but shouldn't")
				assert.Nil(t, createPhysicalInterfaceResponse)

				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, createPhysicalInterfaceResponse)
			}
		})
	}
}

func TestUpdatePhysicalInterface(t *testing.T) {

}
