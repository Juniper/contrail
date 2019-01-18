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
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
)

func TestCreateTag(t *testing.T) {
	tests := []struct {
		name        string
		paramTag    models.Tag
		expectedTag models.Tag
		status      codes.Code
	}{
		{
			name: "Create tag without tag_type_name should failed",
			paramTag: models.Tag{
				FQName:   []string{"namespace=ctest-namespace-95268437"},
				TagValue: "ctest-namespace-95268437",
			},
			status: codes.InvalidArgument,
		},
		{
			name: "Create tag without tag_value should failed",
			paramTag: models.Tag{
				TagTypeName: "namespace",
				FQName:      []string{"namespace=ctest-namespace-95268437"},
			},
			status: codes.InvalidArgument,
		},
		{
			name: "Create tag with tag_id should failed",
			paramTag: models.Tag{
				TagID:       "0x00000001",
				TagTypeName: "namespace",
				FQName:      []string{"namespace=ctest-namespace-95268437"},
				TagValue:    "ctest-namespace-95268437",
			},
			expectedTag: models.Tag{},
			status:      codes.InvalidArgument,
		},
		{
			name: "Create Tag with correct request should succeed",
			paramTag: models.Tag{
				TagTypeName: "namespace",
				FQName:      []string{"namespace=ctest-namespace-95268437"},
				TagValue:    "ctest-namespace-95268437",
			},
			expectedTag: models.Tag{
				TagID:       "0x00ff0001",
				TagTypeName: "namespace",
				FQName:      []string{"namespace=ctest-namespace-95268437"},
				TagValue:    "ctest-namespace-95268437",
			},
			status: codes.OK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			tagSetupNextServiceMocks(service)
			tagSetupIntPoolAllocMock(service)
			tagSetUpListTagTypeMock(service)
			tagSetUpWriteServiceMock(service)

			ctx := context.Background()

			paramRequest := services.CreateTagRequest{Tag: &test.paramTag}
			expectedResponse := services.CreateTagResponse{Tag: &test.expectedTag}
			createTagResponse, err := service.CreateTag(ctx, &paramRequest)

			if test.status != codes.OK {
				stat, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, test.status, stat.Code())
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, &expectedResponse, createTagResponse)
			}
		})
	}
}

func tagSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService) //nolint: errcheck
	nextService.EXPECT().CreateTag(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.CreateTagRequest) (
			response *services.CreateTagResponse, err error,
		) {
			return &services.CreateTagResponse{Tag: request.Tag}, nil
		},
	).AnyTimes()
}

func tagSetUpListTagTypeMock(s *ContrailTypeLogicService) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
	readService.EXPECT().ListTagType(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _ *services.ListTagTypeRequest) (*services.ListTagTypeResponse, error) {
			res := &services.ListTagTypeResponse{
				TagTypeCount: 1,
				TagTypes: []*models.TagType{
					{
						Name:      "namespace",
						FQName:    []string{"namespace"},
						TagTypeID: "0x00ff",
					},
				},
			}
			return res, nil
		},
	).MaxTimes(1)
}

func tagSetUpWriteServiceMock(s *ContrailTypeLogicService) {
	writeService := s.WriteService.(*servicesmock.MockWriteService) //nolint: errcheck
	writeService.EXPECT().CreateTagType(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _ *services.CreateTagTypeRequest) (*services.CreateTagTypeResponse, error) {
			res := &services.CreateTagTypeResponse{
				TagType: &models.TagType{
					Name:      "namespace",
					FQName:    []string{"namespace"},
					TagTypeID: "0x00ff",
				},
			}
			return res, nil
		},
	).MinTimes(0).MaxTimes(1)

	writeService.EXPECT().UpdateTagType(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _ *services.UpdateTagTypeRequest) (*services.UpdateTagTypeResponse, error) {
			res := &services.UpdateTagTypeResponse{}
			return res, nil
		},
	).MaxTimes(1)
}

func tagSetupIntPoolAllocMock(s *ContrailTypeLogicService) {
	intPoolAllocator := s.IntPoolAllocator.(*typesmock.MockIntPoolAllocator) //nolint: errcheck
	intPoolAllocator.EXPECT().AllocateInt(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, tagTypeIDPoolKey string) (int64, error) {
			return 1, nil
		},
	).MaxTimes(1)
}
