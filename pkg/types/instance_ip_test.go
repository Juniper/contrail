package types

import (
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func instanceIPSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService)
	nextService.EXPECT().CreateInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.CreateInstanceIPRequest,
		) (response *services.CreateInstanceIPResponse, err error) {
			return &services.CreateInstanceIPResponse{InstanceIP: request.InstanceIP}, nil
		}).AnyTimes()

	nextService.EXPECT().DeleteInstanceIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.DeleteInstanceIPRequest,
		) (response *services.DeleteInstanceIPResponse, err error) {
			return &services.DeleteInstanceIPResponse{ID: request.ID}, nil
		}).AnyTimes()
}

func TestCreateInstanceIP(t *testing.T) {
	tests := []struct {
		name               string
		paramInstanceIP    models.InstanceIP
		expectedInstanceIP models.InstanceIP
		fails              bool
		errorCode          codes.Code
	}{
		{
			name:               "Try create",
			paramInstanceIP:    models.InstanceIP{},
			expectedInstanceIP: models.InstanceIP{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			instanceIPSetupNextServiceMocks(service)

			ctx := context.Background()

			paramRequest := services.CreateInstanceIPRequest{InstanceIP: &tt.paramInstanceIP}
			expectedResponse := services.CreateInstanceIPResponse{InstanceIP: &tt.expectedInstanceIP}
			createInstanceIPResponse, err := service.CreateInstanceIP(ctx, &paramRequest)

			if tt.fails {
				assert.Error(t, err)
				if tt.errorCode != codes.OK {
					status, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, tt.errorCode, status.Code())
				}
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, createInstanceIPResponse)
			assert.EqualValues(t, &expectedResponse, createInstanceIPResponse)
		})
	}
}

func TestDeleteInstanceIP(t *testing.T) {
	tests := []struct {
		name       string
		paramID    string
		expectedID string
		fails      bool
		errorCode  codes.Code
	}{
		{
			name:       "Try delete",
			paramID:    "id",
			expectedID: "id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			instanceIPSetupNextServiceMocks(service)

			ctx := context.Background()

			paramRequest := services.DeleteInstanceIPRequest{ID: tt.paramID}
			expectedResponse := services.DeleteInstanceIPResponse{ID: tt.expectedID}
			createInstanceIPResponse, err := service.DeleteInstanceIP(ctx, &paramRequest)

			if tt.fails {
				assert.Error(t, err)
				if tt.errorCode != codes.OK {
					status, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, tt.errorCode, status.Code())
				}
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, createInstanceIPResponse)
			assert.EqualValues(t, &expectedResponse, createInstanceIPResponse)
		})
	}
}
