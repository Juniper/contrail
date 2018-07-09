package types

import (
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func floatingIPPoolSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService)
	nextService.EXPECT().CreateFloatingIPPool(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.CreateFloatingIPPoolRequest) (response *services.CreateFloatingIPPoolResponse, err error) {
			return &services.CreateFloatingIPPoolResponse{FloatingIPPool: request.FloatingIPPool}, nil
		}).AnyTimes()
}

func createTestVirtualNetwork() *models.VirtualNetwork {
	return &models.VirtualNetwork{
		UUID: "uuid",
		NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
			&models.VirtualNetworkNetworkIpamRef{
				Attr: &models.VnSubnetsType{
					IpamSubnets: []*models.IpamSubnetType{
						&models.IpamSubnetType{
							SubnetUUID: "subnet-uuid-1",
						},
						&models.IpamSubnetType{
							SubnetUUID: "subnet-uuid-2",
						},
					},
				},
			},
			&models.VirtualNetworkNetworkIpamRef{
				Attr: &models.VnSubnetsType{
					IpamSubnets: []*models.IpamSubnetType{
						&models.IpamSubnetType{
							SubnetUUID: "subnet-uuid-3",
						},
						&models.IpamSubnetType{
							SubnetUUID: "subnet-uuid-4",
						},
					},
				},
			},
			&models.VirtualNetworkNetworkIpamRef{},
		},
	}
}

func floatingIPPoolPrepareNetwork(s *ContrailTypeLogicService, virtualNetwork *models.VirtualNetwork) {
	dataService := s.DataService.(*servicesmock.MockService)
	if len(virtualNetwork.GetUUID()) > 0 {
		dataService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			&services.GetVirtualNetworkResponse{
				VirtualNetwork: virtualNetwork,
			}, nil).AnyTimes()
	} else {
		dataService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			nil, fmt.Errorf("Incomplete info to create a floating-ip-pool")).AnyTimes()
	}
}

func TestCreateFloatingIPPool(t *testing.T) {
	tests := []struct {
		name             string
		paramRequest     services.CreateFloatingIPPoolRequest
		expectedResponse services.CreateFloatingIPPoolResponse
		fails            bool
		errorCode        codes.Code
	}{
		{
			name: "Try to create floating ip pool not corresponding to a virtual-network",
			paramRequest: services.CreateFloatingIPPoolRequest{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "instance-ip",
			}},
			expectedResponse: services.CreateFloatingIPPoolResponse{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "instance-ip",
			}},
		},
		{
			name: "Try to create floating ip pool with no subnets attached",
			paramRequest: services.CreateFloatingIPPoolRequest{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "virtual-network",
			}},
			expectedResponse: services.CreateFloatingIPPoolResponse{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "virtual-network",
			}},
		},
		{
			name: "Try to create floating ip pool when list of subnets is empty",
			paramRequest: services.CreateFloatingIPPoolRequest{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{},
				},
			}},
			expectedResponse: services.CreateFloatingIPPoolResponse{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{},
				},
			}},
		},
		{
			name: "Try to create floating ip pool when cannot find virtual-network with corresponding uuid",
			paramRequest: services.CreateFloatingIPPoolRequest{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"subnet-uuid"},
				},
			}},
			fails: true,
		},
		{
			name: "Try to create floating ip pool when  corresponding subnets in virtual-network cannot be found",
			paramRequest: services.CreateFloatingIPPoolRequest{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"subnet-uuid-5"},
				},
				ParentUUID: "uuid",
			}},
			fails:     true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create floating ip pool with subnet-uuids that can be found on virtual-network",
			paramRequest: services.CreateFloatingIPPoolRequest{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"subnet-uuid-1", "subnet-uuid-4"},
				},
				ParentUUID: "uuid",
			}},
			expectedResponse: services.CreateFloatingIPPoolResponse{FloatingIPPool: &models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"subnet-uuid-1", "subnet-uuid-4"},
				},
				ParentUUID: "uuid",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(mockCtrl)
			virtualNetwork := createTestVirtualNetwork()
			floatingIPPoolSetupNextServiceMocks(service)
			floatingIPPoolPrepareNetwork(service, virtualNetwork)

			ctx := context.Background()

			createFloatingIPPoolResponse, err := service.CreateFloatingIPPool(ctx, &tt.paramRequest)

			if tt.fails {
				assert.Error(t, err)
				if tt.errorCode != codes.OK {
					status, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, tt.errorCode, status.Code())
				}
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, createFloatingIPPoolResponse)
			assert.EqualValues(t, &tt.expectedResponse, createFloatingIPPoolResponse)
			mockCtrl.Finish()
		})
	}
}
