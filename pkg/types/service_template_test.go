package types

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
)

func TestCreateServiceTemplate(t *testing.T) {
	tests := []struct {
		name                string
		testServiceTemplate *models.ServiceTemplate
		errorCode           codes.Code
	}{
		{
			name: "create service template with parent UUID",
			testServiceTemplate: &models.ServiceTemplate{
				ParentUUID: "beefbeef-beef-beef-beef-beefbeef0002",
				ParentType: "domain",
			},
			errorCode: codes.OK,
		},
		{
			name: "create service template with fq name",
			testServiceTemplate: &models.ServiceTemplate{
				FQName:     []string{"default-domain", "service-template-name"},
				ParentType: "domain",
			},
			errorCode: codes.OK,
		},
		{
			name: "create service template without fq name or parent UUID",
			testServiceTemplate: &models.ServiceTemplate{
				ParentType: "domain",
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "create service template with incorrect domain in fq name",
			testServiceTemplate: &models.ServiceTemplate{
				FQName:     []string{"bad-domain", "service-template-name"},
				ParentType: "domain",
			},
			errorCode: codes.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			ctx := context.Background()
			paramRequest := services.CreateServiceTemplateRequest{ServiceTemplate: tt.testServiceTemplate}
			expectedResponse := services.CreateServiceTemplateResponse{ServiceTemplate: tt.testServiceTemplate}

			serviceTemplateSetupCreateMock(service, tt.testServiceTemplate)
			serviceTemplateSetupMetadataMock(service)

			createServiceTemplateResponse, err := service.CreateServiceTemplate(ctx, &paramRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createServiceTemplateResponse)
			}
		})
	}
}

func serviceTemplateSetupCreateMock(service *ContrailTypeLogicService, r *models.ServiceTemplate) {
	service.Next().(*servicesmock.MockService).EXPECT().CreateServiceTemplate(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, _ *services.CreateServiceTemplateRequest,
		) (response *services.CreateServiceTemplateResponse, err error) {
			return &services.CreateServiceTemplateResponse{ServiceTemplate: r}, nil
		},
	).AnyTimes()
}

func serviceTemplateSetupMetadataMock(s *ContrailTypeLogicService) {
	s.MetadataGetter.(*typesmock.MockMetadataGetter).EXPECT().GetMetadata(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, requested basemodels.Metadata) (
			response *basemodels.Metadata, err error,
		) {
			return serviceTemplateSetupMetadata(requested)
		},
	).AnyTimes()
}

func serviceTemplateSetupMetadata(requested basemodels.Metadata) (
	*basemodels.Metadata, error) {
	if len(requested.FQName) == 0 {
		return nil, errutil.ErrorBadRequest("FQ name is empty")
	}

	if requested.Type == models.KindDomain {
		if requested.FQName[0] != "default-domain" {
			return nil, errutil.ErrorNotFound
		}

		return &basemodels.Metadata{
			UUID:   "beefbeef-beef-beef-beef-beefbeef0002",
			FQName: requested.FQName,
		}, nil
	}

	return nil, errutil.ErrorBadRequest("No parent found")
}
