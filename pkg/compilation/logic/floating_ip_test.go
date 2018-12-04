package logic

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilation/dependencies"
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
)

func TestCreateFloatingIP(t *testing.T) {
	tests := []struct {
		name           string
		testFloatingIP *models.FloatingIP
		fipIntent      *FloatingIPIntent
		fails          bool
	}{
		{
			name: "add ip version(ipv4) to intent properly",
			testFloatingIP: &models.FloatingIP{
				UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
				FloatingIPAddress: "10.10.10.10",
			},
			fipIntent: &FloatingIPIntent{
				FloatingIP: &models.FloatingIP{
					UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
					FloatingIPAddress: "10.10.10.10",
				},
				ipVersion: 4,
			},
		},
		{
			name: "add ip version(ipv6) to intent properly",
			testFloatingIP: &models.FloatingIP{
				UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
				FloatingIPAddress: "::1",
			},
			fipIntent: &FloatingIPIntent{
				FloatingIP: &models.FloatingIP{
					UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
					FloatingIPAddress: "::1",
				},
				ipVersion: 6,
			},
		},
		{
			name: "try to add floating ip with invalid ip address",
			testFloatingIP: &models.FloatingIP{
				UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
				FloatingIPAddress: "10.10.10.10.10",
			},
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAPIClient := servicesmock.NewMockWriteService(mockCtrl)
			mockReadService := servicesmock.NewMockReadService(mockCtrl)
			mockIntPoolAllocator := typesmock.NewMockIntPoolAllocator(mockCtrl)
			cache := intent.NewCache()
			service := NewService(
				mockAPIClient,
				mockReadService,
				mockIntPoolAllocator,
				cache,
				dependencies.NewDependencyProcessor(parseReactions(t)),
			)

			_, err := service.CreateFloatingIP(context.Background(), &services.CreateFloatingIPRequest{
				FloatingIP: tt.testFloatingIP,
			})

			fipIntent := LoadFloatingIPIntent(cache, intent.ByUUID(tt.testFloatingIP.GetUUID()))
			if tt.fails {
				assert.Error(t, err)
				assert.Nil(t, fipIntent)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, fipIntent)
				assert.Equal(t, tt.fipIntent, fipIntent)
			}
		})
	}
}

func TestUpdateFloatingIP(t *testing.T) {
	tests := []struct {
		name           string
		testFloatingIP *models.FloatingIP
		cacheIntent    *FloatingIPIntent
		updatedIntent  *FloatingIPIntent
		fails          bool
	}{
		{
			name: "update ip address(ipv4)",
			testFloatingIP: &models.FloatingIP{
				UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
				FloatingIPAddress: "10.10.10.10",
			},
			cacheIntent: &FloatingIPIntent{
				FloatingIP: &models.FloatingIP{
					UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
					FloatingIPAddress: "::1",
				},
				ipVersion: 6,
			},
			updatedIntent: &FloatingIPIntent{
				FloatingIP: &models.FloatingIP{
					UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
					FloatingIPAddress: "10.10.10.10",
				},
				ipVersion: 4,
			},
		},
		{
			name: "update ip address(ipv6)",
			testFloatingIP: &models.FloatingIP{
				UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
				FloatingIPAddress: "::1",
			},
			cacheIntent: &FloatingIPIntent{
				FloatingIP: &models.FloatingIP{
					UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
					FloatingIPAddress: "10.10.10.10",
				},
				ipVersion: 4,
			},
			updatedIntent: &FloatingIPIntent{
				FloatingIP: &models.FloatingIP{
					UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
					FloatingIPAddress: "::1",
				},
				ipVersion: 6,
			},
		},
		{
			name: "try to update fip not existing in cache",
			testFloatingIP: &models.FloatingIP{
				UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
				FloatingIPAddress: "::1",
			},
			fails: true,
		},
		{
			name: "try to update with improper ip address",
			testFloatingIP: &models.FloatingIP{
				UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
				FloatingIPAddress: "10.10.10.10.10",
			},
			cacheIntent: &FloatingIPIntent{
				FloatingIP: &models.FloatingIP{
					UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
					FloatingIPAddress: "::1",
				},
				ipVersion: 6,
			},
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAPIClient := servicesmock.NewMockWriteService(mockCtrl)
			mockReadService := servicesmock.NewMockReadService(mockCtrl)
			mockIntPoolAllocator := typesmock.NewMockIntPoolAllocator(mockCtrl)
			cache := intent.NewCache()
			service := NewService(
				mockAPIClient,
				mockReadService,
				mockIntPoolAllocator,
				cache,
				dependencies.NewDependencyProcessor(parseReactions(t)),
			)

			if tt.cacheIntent != nil {
				cache.Store(tt.cacheIntent)
			}

			_, err := service.UpdateFloatingIP(context.Background(), &services.UpdateFloatingIPRequest{
				FloatingIP: tt.testFloatingIP,
			})

			fipIntent := LoadFloatingIPIntent(cache, intent.ByUUID(tt.testFloatingIP.GetUUID()))
			if tt.fails {
				assert.Error(t, err)
				assert.Equal(t, tt.cacheIntent, fipIntent)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, fipIntent)
				assert.Equal(t, tt.updatedIntent, fipIntent)
			}
		})
	}
}

func TestDeleteFloatingIP(t *testing.T) {
	tests := []struct {
		name        string
		uuid        string
		cacheIntent *FloatingIPIntent
		fails       bool
	}{
		{
			name: "delete floating ip from cache",
			uuid: "d91369ad-6bdb-42e3-8659-f33e3b76612a",
			cacheIntent: &FloatingIPIntent{
				FloatingIP: &models.FloatingIP{
					UUID:              "d91369ad-6bdb-42e3-8659-f33e3b76612a",
					FloatingIPAddress: "10.10.10.10",
				},
				ipVersion: 4,
			},
		},
		{
			name:  "delete floating ip not existing in cache",
			uuid:  "d91369ad-6bdb-42e3-8659-f33e3b76612a",
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAPIClient := servicesmock.NewMockWriteService(mockCtrl)
			mockReadService := servicesmock.NewMockReadService(mockCtrl)
			mockIntPoolAllocator := typesmock.NewMockIntPoolAllocator(mockCtrl)
			cache := intent.NewCache()
			service := NewService(
				mockAPIClient,
				mockReadService,
				mockIntPoolAllocator,
				cache,
				dependencies.NewDependencyProcessor(parseReactions(t)),
			)

			if tt.cacheIntent != nil {
				cache.Store(tt.cacheIntent)
			}

			_, err := service.DeleteFloatingIP(context.Background(), &services.DeleteFloatingIPRequest{
				ID: tt.uuid,
			})

			fipIntent := LoadFloatingIPIntent(cache, intent.ByUUID(tt.uuid))
			assert.Nil(t, fipIntent)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
