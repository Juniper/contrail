package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

func domainNextServMocks(t *testing.T, service *ContrailTypeLogicService) {
	nextServiceMock, ok := service.Next().(*servicesmock.MockService)
	assert.True(t, ok)
	nextServiceMock.EXPECT().CreateDomain(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.CreateDomainRequest,
		) (response *services.CreateDomainResponse, err error) {
			return &services.CreateDomainResponse{Domain: request.Domain}, nil
		}).AnyTimes()
}

func TestCreateDomain(t *testing.T) {
	tests := []struct {
		name          string
		createRequest *services.CreateDomainRequest
		errorCode     codes.Code
	}{
		{
			name: "Try create domain without perms2.share",
			createRequest: &services.CreateDomainRequest{Domain: &models.Domain{
				UUID:   "uuid",
				Perms2: &models.PermType2{},
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try create domain with empty perms2.share",
			createRequest: &services.CreateDomainRequest{Domain: &models.Domain{
				UUID: "uuid",
				Perms2: &models.PermType2{
					Share: []*models.ShareType{},
				},
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try create domain with any perms2.share",
			createRequest: &services.CreateDomainRequest{Domain: &models.Domain{
				UUID: "uuid",
				Perms2: &models.PermType2{
					Share: []*models.ShareType{{
						TenantAccess: basemodels.PermsRW,
						Tenant:       "domain:any_uuid",
					}},
				},
			}},
			errorCode: codes.OK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			domainNextServMocks(t, service)
			ctx := context.Background()
			//logicalInterfaceReadServiceMocks(t, service, tt.parentRouter, tt.listPhysicalInterface)

			createDomainResponse, err := service.CreateDomain(ctx, tt.createRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err, "create succeeded but shouldn't")
				assert.Nil(t, createDomainResponse)
				assert.Equal(t, tt.errorCode, errutil.CauseCode(err))
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, createDomainResponse)
			}
		})
	}
}
