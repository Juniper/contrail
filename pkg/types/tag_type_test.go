package types

import (
	"context"
	"fmt"
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
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
)

func TestCreateTagType(t *testing.T) {
	tests := []struct {
		name                    string
		requestedTagType        *models.TagType
		expectedTagTypeResponse *services.CreateTagTypeResponse
		status                  codes.Code
		allocSucceeded          bool
	}{
		{
			name: "should succeed given correct properties",
			requestedTagType: &models.TagType{
				FQName: []string{"test-tag-type"}},
			expectedTagTypeResponse: &services.CreateTagTypeResponse{
				TagType: &models.TagType{
					FQName:    []string{"test-tag-type"},
					TagTypeID: "0x0100",
				},
			},
			allocSucceeded: true,
			status:         codes.OK,
		},
		{
			name: "should fail given TagTypeID property",
			requestedTagType: &models.TagType{
				FQName:    []string{"test-tag-type"},
				TagTypeID: "BEEF",
			},
			allocSucceeded: true,
			status:         codes.InvalidArgument,
		},
		{
			name: "should fail when int allocation fails",
			requestedTagType: &models.TagType{
				FQName:    []string{"test-tag-type"},
				TagTypeID: "BEEF",
			},
			allocSucceeded: false,
			status:         codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			service := makeMockedContrailTypeLogicService(mockCtrl)
			expectCreateTagTypeToBeCalledOnNextService(service.Next().(*servicesmock.MockService)) // nolint: errcheck
			expectAllocateIntToBeCalledOnIntPoolAllocator(
				service.IntPoolAllocator.(*typesmock.MockIntPoolAllocator), tt.allocSucceeded) // nolint: errcheck

			r, err := service.CreateTagType(
				context.Background(),
				&services.CreateTagTypeRequest{TagType: tt.requestedTagType},
			)

			s, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.status, s.Code())
			assert.Equal(t, tt.expectedTagTypeResponse, r)
		})
	}
}

func TestUpdateTagType(t *testing.T) {
	tests := []struct {
		name    string
		request *services.UpdateTagTypeRequest
		status  codes.Code
	}{
		{
			name: "should fail given TagTypeID in request",
			request: &services.UpdateTagTypeRequest{
				TagType: &models.TagType{
					FQName:    []string{"test-tag-type"},
					TagTypeID: "0x0001",
				},
				FieldMask: types.FieldMask{Paths: []string{models.TagTypeFieldTagTypeID}},
			},
			status: codes.InvalidArgument,
		},
		{
			name: "should fail given DisplayName in request",
			request: &services.UpdateTagTypeRequest{
				TagType: &models.TagType{
					FQName:      []string{"test-tag-type"},
					DisplayName: "0x0001",
				},
				FieldMask: types.FieldMask{Paths: []string{models.TagTypeFieldTagTypeID}},
			},
			status: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			service := makeMockedContrailTypeLogicService(mockCtrl)
			expectUpdateTagTypeToBeCalledOnNextService(service.Next().(*servicesmock.MockService)) // nolint: errcheck

			_, err := service.UpdateTagType(
				context.Background(),
				tt.request,
			)

			s, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.status, s.Code())
		})
	}
}

func TestDeleteTagType(t *testing.T) {
	tests := []struct {
		name             string
		dbTagType        *models.TagType
		request          *services.DeleteTagTypeRequest
		status           codes.Code
		deallocSucceeded bool
	}{
		{
			name: "should fail when cannot deallocate int",
			dbTagType: &models.TagType{
				UUID:   "test-tag-type",
				FQName: []string{"test-tag-type"},
			},
			request: &services.DeleteTagTypeRequest{
				ID: "test-tag-type",
			},
			deallocSucceeded: false,
			status:           codes.InvalidArgument,
		},
		{
			name: "should fail when cannot parse hex to int",
			dbTagType: &models.TagType{
				UUID:      "test-tag-type",
				FQName:    []string{"test-tag-type"},
				TagTypeID: "wrong-hex",
			},
			request: &services.DeleteTagTypeRequest{
				ID: "test-tag-type",
			},
			deallocSucceeded: true,
			status:           codes.InvalidArgument,
		},
		{
			name:      "should fail when cannot get Tag Type from DB",
			dbTagType: nil,
			request: &services.DeleteTagTypeRequest{
				ID: "test-tag-type",
			},
			deallocSucceeded: true,
			status:           codes.NotFound,
		},
		{
			name: "should succeed given correct properties",
			dbTagType: &models.TagType{
				FQName:    []string{"test-tag-type"},
				TagTypeID: "0x100",
			},
			request: &services.DeleteTagTypeRequest{
				ID: "test-tag-type",
			},
			deallocSucceeded: true,
			status:           codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			service := makeMockedContrailTypeLogicService(mockCtrl)
			expectDeleteTagTypeToBeCalledOnNextService(service.Next().(*servicesmock.MockService)) // nolint: errcheck
			tagTypePrepareReadService(service, tt.dbTagType)
			expectDeallocateIntToBeCalledOnIntPoolAllocator(
				service.IntPoolAllocator.(*typesmock.MockIntPoolAllocator), tt.deallocSucceeded) // nolint: errcheck

			_, err := service.DeleteTagType(
				context.Background(),
				tt.request,
			)

			s, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.status, s.Code())
		})
	}
}

func expectCreateTagTypeToBeCalledOnNextService(s *servicesmock.MockService) {
	s.EXPECT().CreateTagType(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.CreateTagTypeRequest) (
			response *services.CreateTagTypeResponse, err error,
		) {
			return &services.CreateTagTypeResponse{TagType: request.TagType}, nil
		},
	).MaxTimes(1)
}

func expectUpdateTagTypeToBeCalledOnNextService(s *servicesmock.MockService) {
	s.EXPECT().UpdateTagType(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.UpdateTagTypeRequest) (
			response *services.UpdateTagTypeResponse, err error,
		) {
			return &services.UpdateTagTypeResponse{TagType: request.TagType}, nil
		},
	).MaxTimes(1)
}

func expectDeleteTagTypeToBeCalledOnNextService(s *servicesmock.MockService) {
	s.EXPECT().DeleteTagType(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.DeleteTagTypeRequest) (
			response *services.DeleteTagTypeResponse, err error,
		) {
			return &services.DeleteTagTypeResponse{ID: request.ID}, nil
		},
	).MaxTimes(1)
}

func expectAllocateIntToBeCalledOnIntPoolAllocator(s *typesmock.MockIntPoolAllocator, allocSucceed bool) {
	s.EXPECT().AllocateInt(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _, _ string) (int64, error) {
			if allocSucceed {
				return 256, nil
			}
			return 0, fmt.Errorf("cannot allocate int")
		},
	).MaxTimes(1)
}

func expectDeallocateIntToBeCalledOnIntPoolAllocator(s *typesmock.MockIntPoolAllocator, deallocSucceed bool) {
	s.EXPECT().DeallocateInt(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _ string, _ int64) error {
			if deallocSucceed {
				return nil
			}
			return fmt.Errorf("cannot deallocate int")
		},
	).MaxTimes(1)
}

func tagTypePrepareReadService(s *ContrailTypeLogicService, tt *models.TagType) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	if tt != nil {
		readService.EXPECT().GetTagType(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			&services.GetTagTypeResponse{
				TagType: tt,
			}, nil).AnyTimes()
	} else {
		readService.EXPECT().GetTagType(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			nil, errutil.ErrorNotFound).AnyTimes()
	}
}
