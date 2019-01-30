package types

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

func nextServMock(t *testing.T, service *ContrailTypeLogicService, getResp *services.GetBGPAsAServiceResponse) {
	nextServiceMock, ok := service.Next().(*servicesmock.MockService)
	assert.True(t, ok)
	readServiceMock, ok := service.ReadService.(*servicesmock.MockReadService)
	assert.True(t, ok)
	nextServiceMock.EXPECT().CreateBGPAsAService(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, _ *services.CreateBGPAsAServiceRequest,
		) (response *services.CreateBGPAsAServiceResponse, err error) {
			return &services.CreateBGPAsAServiceResponse{}, nil
		}).AnyTimes()
	nextServiceMock.EXPECT().UpdateBGPAsAService(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, _ *services.UpdateBGPAsAServiceRequest,
		) (response *services.UpdateBGPAsAServiceResponse, err error) {
			return &services.UpdateBGPAsAServiceResponse{}, nil
		}).AnyTimes()
	readServiceMock.EXPECT().GetBGPAsAService(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, _ *services.GetBGPAsAServiceRequest,
		) (response *services.GetBGPAsAServiceResponse, err error) {
			return getResp, nil
		}).AnyTimes()
}

func TestCreateBGPAsAService(t *testing.T) {
	tests := []struct {
		name    string
		request *services.CreateBGPAsAServiceRequest
		passing bool
	}{
		{
			name:    "Create empty bgpass",
			request: &services.CreateBGPAsAServiceRequest{},
			passing: true,
		},
		{
			name: "Create shared bgpaas",
			request: &services.CreateBGPAsAServiceRequest{
				BGPAsAService: &models.BGPAsAService{
					BgpaasShared:    true,
					BgpaasIPAddress: "1.1.1.1",
				},
				FieldMask: types.FieldMask{
					Paths: []string{models.BGPAsAServiceFieldBgpaasShared},
				},
			},
			passing: true,
		},
		{
			name: "Fail to create shared bgpaas",
			request: &services.CreateBGPAsAServiceRequest{
				BGPAsAService: &models.BGPAsAService{
					BgpaasShared:    true,
					BgpaasIPAddress: "",
				},
			},
			passing: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			nextServMock(t, service, nil)
			ctx := context.Background()

			response, err := service.CreateBGPAsAService(ctx, tt.request)
			if tt.passing {
				assert.NoError(t, err)
				assert.NotNil(t, response)
			} else {
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}

func TestUpdateBGPAsAService(t *testing.T) {
	tests := []struct {
		name        string
		request     *services.UpdateBGPAsAServiceRequest
		getResponse *services.GetBGPAsAServiceResponse
		passing     bool
	}{
		{
			name:    "Empty Update bgpass",
			request: &services.UpdateBGPAsAServiceRequest{},
			passing: true,
		},
		{
			name: "Update shared bgpaas",
			request: &services.UpdateBGPAsAServiceRequest{
				BGPAsAService: &models.BGPAsAService{
					BgpaasShared:    true,
					BgpaasIPAddress: "1.1.1.1",
				},
				FieldMask: types.FieldMask{
					Paths: []string{models.BGPAsAServiceFieldBgpaasShared},
				},
			},
			getResponse: &services.GetBGPAsAServiceResponse{
				BGPAsAService: &models.BGPAsAService{
					BgpaasShared: true,
				},
			},
			passing: true,
		},
		{
			name: "Fail to update shared bgpaas (overwrite shared flag)",
			request: &services.UpdateBGPAsAServiceRequest{
				BGPAsAService: &models.BGPAsAService{
					BgpaasShared:    true,
					BgpaasIPAddress: "",
				},
				FieldMask: types.FieldMask{
					Paths: []string{models.BGPAsAServiceFieldBgpaasShared},
				},
			},
			getResponse: &services.GetBGPAsAServiceResponse{
				BGPAsAService: &models.BGPAsAService{
					BgpaasShared: false,
				},
			},
			passing: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			nextServMock(t, service, tt.getResponse)
			ctx := context.Background()

			response, err := service.UpdateBGPAsAService(ctx, tt.request)
			if tt.passing {
				assert.NoError(t, err)
				assert.NotNil(t, response)
			} else {
				assert.Error(t, err)
				assert.Nil(t, response)
			}
		})
	}
}
