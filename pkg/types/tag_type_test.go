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

func TestCreateTagType(t *testing.T) {
	tests := []struct {
		name                    string
		requestedTagType        *models.TagType
		expectedTagTypeResponse *services.CreateTagTypeResponse
		status                  codes.Code
	}{
		{
			name: "should succeed given correct properties",
			requestedTagType: &models.TagType{
				FQName: []string{"default", "default-project", "test-tag-type"}},
			expectedTagTypeResponse: &services.CreateTagTypeResponse{
				TagType: &models.TagType{
					FQName: []string{"default", "default-project", "test-tag-type"},
				},
			},
			status: codes.OK,
		},
		{
			name: "should return InvalidArgument error given TagTypeID property",
			requestedTagType: &models.TagType{
				FQName:    []string{"default", "default-project", "test-tag-type"},
				TagTypeID: "BEEF",
			},
			status: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctls := makeMockedContrailTypeLogicService(mockCtrl)
			expectCreateTagTypeToBeCalledOnNextService(ctls.Next().(*servicesmock.MockService)) // nolint: errcheck
			//expectAllocateIntToBeCalledOnIntPoolAllocator(
			//	ctls.IntPoolAllocator.(*typesmock.MockIntPoolAllocator)) // nolint: errcheck

			r, err := ctls.CreateTagType(
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

//func expectAllocateIntToBeCalledOnIntPoolAllocator(s *ContrailTypeLogicService) {
//	s.EXPECT().AllocateInt( // nolint: errcheck
//		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
//	).DoAndReturn(
//		func(_ context.Context, tagTypeIDPoolKey string) (int64, error) {
//			return 1, nil
//		},
//	).MaxTimes(1)
//}
