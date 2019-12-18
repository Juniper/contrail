package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

func floatingIPPoolSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService) //nolint: errcheck
	nextService.EXPECT().CreateFloatingIPPool(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *services.CreateFloatingIPPoolRequest,
		) (response *services.CreateFloatingIPPoolResponse, err error) {
			return &services.CreateFloatingIPPoolResponse{FloatingIPPool: request.FloatingIPPool}, nil
		}).AnyTimes()
}

func createTestVirtualNetwork() *models.VirtualNetwork {
	virtualNetwork := models.MakeVirtualNetwork()
	virtualNetwork.UUID = "uuid-1" // nolint: goconst
	virtualNetwork.NetworkIpamRefs = []*models.VirtualNetworkNetworkIpamRef{
		{
			Attr: &models.VnSubnetsType{
				IpamSubnets: []*models.IpamSubnetType{
					{
						SubnetUUID: "subnet-uuid-1",
					},
					{
						SubnetUUID: "subnet-uuid-2",
					},
				},
			},
		},
		{
			Attr: &models.VnSubnetsType{
				IpamSubnets: []*models.IpamSubnetType{
					{
						SubnetUUID: "subnet-uuid-3",
					},
					{
						SubnetUUID: "subnet-uuid-4",
					},
				},
			},
		},
		{},
	}
	return virtualNetwork
}

func floatingIPPoolPrepareNetwork(s *ContrailTypeLogicService, virtualNetwork *models.VirtualNetwork) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	readService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()),
		&services.GetVirtualNetworkRequest{
			ID: virtualNetwork.GetUUID(),
		}).Return(
		&services.GetVirtualNetworkResponse{
			VirtualNetwork: virtualNetwork,
		}, nil).AnyTimes()

	readService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, errutil.ErrorNotFound).AnyTimes()
}

func TestCreateFloatingIPPool(t *testing.T) {
	tests := []struct {
		name                   string
		paramFloatingIPPool    models.FloatingIPPool
		expectedFloatingIPPool models.FloatingIPPool
		fails                  bool
		errorCode              codes.Code
	}{
		{
			name: "Try to create floating ip pool not corresponding to a virtual-network",
			paramFloatingIPPool: models.FloatingIPPool{
				ParentType: "instance-ip",
			},
			expectedFloatingIPPool: models.FloatingIPPool{
				ParentType: "instance-ip",
			},
		},
		{
			name: "Try to create floating ip pool with no subnets attached",
			paramFloatingIPPool: models.FloatingIPPool{
				ParentType: "virtual-network",
			},
			expectedFloatingIPPool: models.FloatingIPPool{
				ParentType: "virtual-network",
			},
		},
		{
			name: "Try to create floating ip pool when list of subnets is empty",
			paramFloatingIPPool: models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{},
				},
			},
			expectedFloatingIPPool: models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{},
				},
			},
		},
		{
			name: "Try to create floating ip pool when cannot find virtual-network with corresponding uuid",
			paramFloatingIPPool: models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"subnet-uuid"},
				},
				ParentUUID: "uuid-2",
			},
			fails:     true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create floating ip pool when corresponding subnets in virtual-network cannot be found",
			paramFloatingIPPool: models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"subnet-uuid-5"},
				},
				ParentUUID: "uuid-1",
			},
			fails:     true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create floating ip pool with subnet-uuids that can be found on virtual-network",
			paramFloatingIPPool: models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"subnet-uuid-1", "subnet-uuid-4"},
				},
				ParentUUID: "uuid-1",
			},
			expectedFloatingIPPool: models.FloatingIPPool{
				ParentType: "virtual-network",
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"subnet-uuid-1", "subnet-uuid-4"},
				},
				ParentUUID: "uuid-1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			virtualNetwork := createTestVirtualNetwork()
			floatingIPPoolSetupNextServiceMocks(service)
			floatingIPPoolPrepareNetwork(service, virtualNetwork)

			ctx := context.Background()

			paramRequest := services.CreateFloatingIPPoolRequest{FloatingIPPool: &tt.paramFloatingIPPool}
			expectedResponse := services.CreateFloatingIPPoolResponse{FloatingIPPool: &tt.expectedFloatingIPPool}
			createFloatingIPPoolResponse, err := service.CreateFloatingIPPool(ctx, &paramRequest)

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
			assert.EqualValues(t, &expectedResponse, createFloatingIPPoolResponse)
		})
	}
}
