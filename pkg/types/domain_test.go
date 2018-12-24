package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

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

func getTenantAccess(createResponse *services.CreateDomainResponse, uuid string) int64 {
	for _, share := range createResponse.GetDomain().GetPerms2().Share {
		if share.GetTenant() == "domain:"+uuid {
			return share.GetTenantAccess()
		}
	}
	return 0
}

func TestCreateDomain(t *testing.T) {
	tests := []struct {
		name          string
		createRequest *services.CreateDomainRequest
	}{
		{
			name: "Try create domain without perms2.share",
			createRequest: &services.CreateDomainRequest{Domain: &models.Domain{
				UUID:   "uuid",
				Perms2: &models.PermType2{},
			}},
		},
		{
			name: "Try create domain with empty perms2.share",
			createRequest: &services.CreateDomainRequest{Domain: &models.Domain{
				UUID: "uuid",
				Perms2: &models.PermType2{
					Share: []*models.ShareType{},
				},
			}},
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			domainNextServMocks(t, service)
			ctx := context.Background()

			createDomainResponse, err := service.CreateDomain(ctx, tt.createRequest)
			assert.NoError(t, err)
			assert.NotNil(t, createDomainResponse)
			assert.Equal(t, getTenantAccess(createDomainResponse, tt.createRequest.Domain.UUID), int64(basemodels.PermsRW))
		})
	}
}
