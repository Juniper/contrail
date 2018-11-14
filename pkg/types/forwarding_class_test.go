package types

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type MockedListForwardingClassResponse struct {
	forwardingClassResponse *services.ListForwardingClassResponse
	retrunedError           error
}

type MockedGetForwardingClassResponse struct {
	forwardingClassResponse *services.GetForwardingClassResponse
	retrunedError           error
}

func TestCreateForwardingClass(t *testing.T) {
	tests := []struct {
		name                              string
		testForwardingClass               *models.ForwardingClass
		mockedListForwardingClassResponse MockedListForwardingClassResponse // values of mocked ReadService.ListForwardingClass()
		errorCode                         codes.Code
	}{
		{
			name: "Create successfully ForwardingClass",
			testForwardingClass: &models.ForwardingClass{
				UUID:              "forwarding-class-1",
				ForwardingClassID: 1,
			},
			mockedListForwardingClassResponse: MockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{
					ForwardingClasss:     []*models.ForwardingClass{},
					ForwardingClassCount: 0,
				},
				retrunedError: nil,
			},
			errorCode: codes.OK,
		},
		{
			name: "Fail creating ForwardingClass on alredy used ForwardingClassID param",
			testForwardingClass: &models.ForwardingClass{
				UUID:              "forwarding-class-1",
				ForwardingClassID: 1,
			},
			mockedListForwardingClassResponse: MockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{
					ForwardingClasss: []*models.ForwardingClass{
						{DisplayName: "Mocked_obj_in_db", ForwardingClassID: 1, UUID: "forwarding-class-2"},
					},
					ForwardingClassCount: 1,
				},
				retrunedError: nil,
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Fail creating ForwardingClass on db error",
			testForwardingClass: &models.ForwardingClass{
				UUID:              "forwarding-class-1",
				ForwardingClassID: 1,
			},
			mockedListForwardingClassResponse: MockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{},
				retrunedError:           errutil.ErrorInternal, // simulate internal db error
			},
			errorCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			ctx := context.Background()
			paramRequest := services.CreateForwardingClassRequest{ForwardingClass: tt.testForwardingClass}
			expectedResponse := services.CreateForwardingClassResponse{ForwardingClass: tt.testForwardingClass}

			// Mock  sv.Next().UpdateForwardingClass(ctx, request) response
			service.Next().(*servicesmock.MockService).EXPECT().CreateForwardingClass(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, _ *services.CreateForwardingClassRequest,
				) (response *services.CreateForwardingClassResponse, err error) {
					return &services.CreateForwardingClassResponse{ForwardingClass: tt.testForwardingClass}, nil
				},
			).AnyTimes()

			// Mock sv.ReadService.ListForwardingClass() response
			service.ReadService.(*servicesmock.MockReadService).EXPECT().ListForwardingClass(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, _ *services.ListForwardingClassRequest) (*services.ListForwardingClassResponse, error) {
					return tt.mockedListForwardingClassResponse.forwardingClassResponse, tt.mockedListForwardingClassResponse.retrunedError
				},
			).AnyTimes()

			createForwardingClassResponse, err := service.CreateForwardingClass(ctx, &paramRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createForwardingClassResponse)
			}
		})
	}

}

func TestUpdateForwardingClass(t *testing.T) {
	tests := []struct {
		name                              string
		request                           services.UpdateForwardingClassRequest // new data to be updated of FrowardingClass defined as "mockedGetForwardingClassResponse"
		mockedListForwardingClassResponse MockedListForwardingClassResponse     // values of mocked ReadService.ListForwardingClass()
		mockedGetForwardingClassResponse  MockedGetForwardingClassResponse      // this is mock of ForwardingClass that exist in DB
		errorCode                         codes.Code
	}{
		{
			name: "Update successfully ForwardingClass",
			request: services.UpdateForwardingClassRequest{
				ForwardingClass: &models.ForwardingClass{
					UUID:              "forwarding-class-1",
					ForwardingClassID: 2,
				},
			},
			mockedListForwardingClassResponse: MockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{
					ForwardingClasss:     []*models.ForwardingClass{},
					ForwardingClassCount: 0,
				},
				retrunedError: nil,
			},
			mockedGetForwardingClassResponse: MockedGetForwardingClassResponse{
				forwardingClassResponse: &services.GetForwardingClassResponse{
					ForwardingClass: &models.ForwardingClass{
						UUID:              "forwarding-class-1",
						ForwardingClassID: 1,
					},
				},
				retrunedError: nil,
			},
			errorCode: codes.OK,
		},
		{
			name: "Fail updating ForwardingClass on alredy used ForwardingClassID param",
			request: services.UpdateForwardingClassRequest{
				ForwardingClass: &models.ForwardingClass{
					UUID:              "forwarding-class-1",
					ForwardingClassID: 1,
				},
			},
			mockedListForwardingClassResponse: MockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{
					ForwardingClasss: []*models.ForwardingClass{
						{DisplayName: "Mocked_obj_in_db", ForwardingClassID: 1, UUID: "forwarding-class-2"},
					},
					ForwardingClassCount: 1,
				},
				retrunedError: nil,
			},
			mockedGetForwardingClassResponse: MockedGetForwardingClassResponse{
				forwardingClassResponse: &services.GetForwardingClassResponse{
					ForwardingClass: &models.ForwardingClass{
						UUID:              "forwarding-class-1",
						ForwardingClassID: 2,
					},
				},
				retrunedError: nil,
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Fail updating ForwardingClass on disability to find this object in db",
			request: services.UpdateForwardingClassRequest{
				ForwardingClass: &models.ForwardingClass{
					UUID:              "forwarding-class-1",
					ForwardingClassID: 1,
				},
			},
			mockedListForwardingClassResponse: MockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{
					ForwardingClasss:     []*models.ForwardingClass{},
					ForwardingClassCount: 0,
				},
				retrunedError: nil,
			},
			mockedGetForwardingClassResponse: MockedGetForwardingClassResponse{
				forwardingClassResponse: nil,
				retrunedError:           errutil.ErrorNotFound,
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Fail updating ForwardingClass on db error",
			request: services.UpdateForwardingClassRequest{
				ForwardingClass: &models.ForwardingClass{
					UUID:              "forwarding-class-1",
					ForwardingClassID: 1,
				},
			},
			mockedListForwardingClassResponse: MockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{},
				retrunedError:           errutil.ErrorInternal, // simulate internal db error
			},
			mockedGetForwardingClassResponse: MockedGetForwardingClassResponse{
				forwardingClassResponse: &services.GetForwardingClassResponse{
					ForwardingClass: &models.ForwardingClass{
						UUID:              "forwarding-class-1",
						ForwardingClassID: 2,
					},
				},
				retrunedError: nil,
			},
			errorCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			ctx := context.Background()
			expectedResponse := services.UpdateForwardingClassResponse{ForwardingClass: tt.request.ForwardingClass}

			// Mock sv.Next().UpdateForwardingClass(ctx, request) response
			service.Next().(*servicesmock.MockService).EXPECT().UpdateForwardingClass(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, _ *services.UpdateForwardingClassRequest,
				) (*services.UpdateForwardingClassResponse, error) {
					return &services.UpdateForwardingClassResponse{ForwardingClass: tt.request.ForwardingClass}, nil
				},
			).AnyTimes()

			// Mock sv.ReadService.ListForwardingClass() response
			service.ReadService.(*servicesmock.MockReadService).EXPECT().ListForwardingClass(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, _ *services.ListForwardingClassRequest) (*services.ListForwardingClassResponse, error) {
					return tt.mockedListForwardingClassResponse.forwardingClassResponse, tt.mockedListForwardingClassResponse.retrunedError
				},
			).AnyTimes()

			// Mock sv.ReadService.GetForwardingClass()
			service.ReadService.(*servicesmock.MockReadService).EXPECT().GetForwardingClass(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, _ *services.GetForwardingClassRequest,
				) (*services.GetForwardingClassResponse, error) {
					return tt.mockedGetForwardingClassResponse.forwardingClassResponse, tt.mockedGetForwardingClassResponse.retrunedError
				},
			).AnyTimes()

			updateForwardingClassResponse, err := service.UpdateForwardingClass(ctx, &tt.request)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, updateForwardingClassResponse)
			}
		})
	}

}
