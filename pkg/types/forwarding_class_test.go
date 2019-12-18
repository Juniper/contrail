package types

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

type mockedListForwardingClassResponse struct {
	forwardingClassResponse *services.ListForwardingClassResponse
	returnedError           error
}

type mockedGetForwardingClassResponse struct {
	forwardingClassResponse *services.GetForwardingClassResponse
	returnedError           error
}

func TestCreateForwardingClass(t *testing.T) {
	tests := []struct {
		name                string
		testForwardingClass *models.ForwardingClass
		*mockedListForwardingClassResponse
		errorCode codes.Code
	}{
		{
			name: "Create successfully ForwardingClass",
			testForwardingClass: &models.ForwardingClass{
				UUID:              "forwarding-class-1",
				ForwardingClassID: 1,
			},
			mockedListForwardingClassResponse: &mockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{
					ForwardingClasss:     []*models.ForwardingClass{},
					ForwardingClassCount: 0,
				},
				returnedError: nil,
			},
			errorCode: codes.OK,
		},
		{
			name: "Fail creating ForwardingClass on alredy used ForwardingClassID param",
			testForwardingClass: &models.ForwardingClass{
				UUID:              "forwarding-class-1",
				ForwardingClassID: 1,
			},
			mockedListForwardingClassResponse: &mockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{
					ForwardingClasss: []*models.ForwardingClass{
						{DisplayName: "Mocked_obj_in_db", ForwardingClassID: 1, UUID: "forwarding-class-2"},
					},
					ForwardingClassCount: 1,
				},
				returnedError: nil,
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Fail creating ForwardingClass on db error",
			testForwardingClass: &models.ForwardingClass{
				UUID:              "forwarding-class-1",
				ForwardingClassID: 1,
			},
			mockedListForwardingClassResponse: &mockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{},
				returnedError:           errutil.ErrorInternal, // simulate internal db error
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

			fails := tt.errorCode != codes.OK

			if !fails {
				initCreateForwardingClassMock(service, tt.testForwardingClass)
			}
			initListForwardingClassMock(service, tt.mockedListForwardingClassResponse)

			createForwardingClassResponse, err := service.CreateForwardingClass(ctx, &paramRequest)
			if fails {
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
		name    string
		request services.UpdateForwardingClassRequest // params to be updated
		*mockedListForwardingClassResponse
		*mockedGetForwardingClassResponse // retruns existing obj in DB
		errorCode                         codes.Code
	}{
		{
			name: "Update successfully ForwardingClass",
			request: services.UpdateForwardingClassRequest{
				ForwardingClass: &models.ForwardingClass{
					UUID:              "forwarding-class-1",
					ForwardingClassID: 2,
				},
				FieldMask: types.FieldMask{Paths: []string{models.ForwardingClassFieldForwardingClassID}},
			},
			mockedListForwardingClassResponse: &mockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{
					ForwardingClasss:     []*models.ForwardingClass{},
					ForwardingClassCount: 0,
				},
				returnedError: nil,
			},
			mockedGetForwardingClassResponse: &mockedGetForwardingClassResponse{
				forwardingClassResponse: &services.GetForwardingClassResponse{
					ForwardingClass: &models.ForwardingClass{
						UUID:              "forwarding-class-1",
						ForwardingClassID: 1,
					},
				},
				returnedError: nil,
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
				FieldMask: types.FieldMask{Paths: []string{models.ForwardingClassFieldForwardingClassID}},
			},
			mockedListForwardingClassResponse: &mockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{
					ForwardingClasss: []*models.ForwardingClass{
						{DisplayName: "Mocked_obj_in_db", ForwardingClassID: 1, UUID: "forwarding-class-2"},
					},
					ForwardingClassCount: 1,
				},
				returnedError: nil,
			},
			mockedGetForwardingClassResponse: &mockedGetForwardingClassResponse{
				forwardingClassResponse: &services.GetForwardingClassResponse{
					ForwardingClass: &models.ForwardingClass{
						UUID:              "forwarding-class-1",
						ForwardingClassID: 2,
					},
				},
				returnedError: nil,
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
				FieldMask: types.FieldMask{Paths: []string{models.ForwardingClassFieldForwardingClassID}},
			},
			mockedGetForwardingClassResponse: &mockedGetForwardingClassResponse{
				forwardingClassResponse: nil,
				returnedError:           errutil.ErrorNotFound,
			},
			mockedListForwardingClassResponse: nil, // list shouldn't be called
			errorCode:                         codes.NotFound,
		},
		{
			name: "Fail updating ForwardingClass on db error",
			request: services.UpdateForwardingClassRequest{
				ForwardingClass: &models.ForwardingClass{
					UUID:              "forwarding-class-1",
					ForwardingClassID: 1,
				},
				FieldMask: types.FieldMask{Paths: []string{models.ForwardingClassFieldForwardingClassID}},
			},
			mockedListForwardingClassResponse: &mockedListForwardingClassResponse{
				forwardingClassResponse: &services.ListForwardingClassResponse{},
				returnedError:           errutil.ErrorInternal, // simulate internal db error
			},
			mockedGetForwardingClassResponse: &mockedGetForwardingClassResponse{
				forwardingClassResponse: &services.GetForwardingClassResponse{
					ForwardingClass: &models.ForwardingClass{
						UUID:              "forwarding-class-1",
						ForwardingClassID: 2,
					},
				},
				returnedError: nil,
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

			fails := tt.errorCode != codes.OK

			if !fails {
				initUpdateForwardingClassMock(service, tt.request.ForwardingClass)
			}
			initListForwardingClassMock(service, tt.mockedListForwardingClassResponse)
			initGetForwardingClassMock(service, tt.mockedGetForwardingClassResponse)

			updateForwardingClassResponse, err := service.UpdateForwardingClass(ctx, &tt.request)
			if fails {
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

func initCreateForwardingClassMock(service *ContrailTypeLogicService, returnedForwardingClass *models.ForwardingClass) {
	service.Next().(*servicesmock.MockService).EXPECT().CreateForwardingClass(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _ *services.CreateForwardingClassRequest,
		) (response *services.CreateForwardingClassResponse, err error) {
			return &services.CreateForwardingClassResponse{ForwardingClass: returnedForwardingClass}, nil
		},
	).Times(1)
}

func initUpdateForwardingClassMock(service *ContrailTypeLogicService, returnedForwardingClass *models.ForwardingClass) {
	service.Next().(*servicesmock.MockService).EXPECT().UpdateForwardingClass(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _ *services.UpdateForwardingClassRequest,
		) (*services.UpdateForwardingClassResponse, error) {
			return &services.UpdateForwardingClassResponse{ForwardingClass: returnedForwardingClass}, nil
		},
	).Times(1)
}

func initListForwardingClassMock(service *ContrailTypeLogicService, r *mockedListForwardingClassResponse) {
	if r == nil {
		return
	}
	service.ReadService.(*servicesmock.MockReadService).EXPECT().ListForwardingClass(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _ *services.ListForwardingClassRequest,
		) (*services.ListForwardingClassResponse, error) {
			return r.forwardingClassResponse,
				r.returnedError
		},
	).Times(1)
}

func initGetForwardingClassMock(service *ContrailTypeLogicService, r *mockedGetForwardingClassResponse) {
	if r == nil {
		return
	}
	service.ReadService.(*servicesmock.MockReadService).EXPECT().GetForwardingClass(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _ *services.GetForwardingClassRequest,
		) (*services.GetForwardingClassResponse, error) {
			return r.forwardingClassResponse,
				r.returnedError
		},
	).Times(1)
}
