package types

import (
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
)

//Structure testVn is used to pass vn parameters during VirtualNetwork object creation
type testVn struct {
	MultiPolicyServiceChainsEnabled bool
	ImportRouteTargetList           string
	ExportRouteTargetList           string
	ForwardingMode                  string
	VirtualNetworkNetworkID         int64
	NetworkIpamRefs                 []*models.VirtualNetworkNetworkIpamRef
	BGPVPNRefs                      []*models.VirtualNetworkBGPVPNRef
	LogicalRouterRefs               []*models.VirtualNetworkLogicalRouterRef
}

func createTestVn(testVnData *testVn) *models.VirtualNetwork {
	vn := models.MakeVirtualNetwork()
	models.MakeIpamSubnetType()
	vn.MultiPolicyServiceChainsEnabled = testVnData.MultiPolicyServiceChainsEnabled
	vn.ImportRouteTargetList = &models.RouteTargetList{RouteTarget: []string{testVnData.ImportRouteTargetList}}
	vn.ExportRouteTargetList = &models.RouteTargetList{RouteTarget: []string{testVnData.ExportRouteTargetList}}
	vn.VirtualNetworkNetworkID = testVnData.VirtualNetworkNetworkID
	if len(testVnData.NetworkIpamRefs) > 0 {
		vn.NetworkIpamRefs = testVnData.NetworkIpamRefs
	}
	if len(testVnData.BGPVPNRefs) > 0 {
		vn.BGPVPNRefs = testVnData.BGPVPNRefs
	}

	if len(testVnData.LogicalRouterRefs) > 0 {
		vn.LogicalRouterRefs = testVnData.LogicalRouterRefs
	}

	vn.UUID = "test_vn_uuid"
	vn.VirtualNetworkProperties.ForwardingMode = testVnData.ForwardingMode

	return vn
}

func virtualNetworkSetupDBMocks(s *ContrailTypeLogicService) {
	dbServiceMock := s.DB.(*typesmock.MockDBServiceInterface)

	dbServiceMock.EXPECT().DB().AnyTimes()
	dbServiceMock.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()),
		&services.GetVirtualNetworkRequest{
			ID: "test_vn_uuid",
		}).Return(
		&services.GetVirtualNetworkResponse{
			VirtualNetwork: models.MakeVirtualNetwork(),
		}, nil).AnyTimes()

	dbServiceMock.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, common.ErrorBadRequest("Not found")).AnyTimes()
}

func virtualNetworkSetupNetworkIpam(s *ContrailTypeLogicService, ipamSubnetMethod string) {
	dbServiceMock := s.DB.(*typesmock.MockDBServiceInterface)

	ipamSubnetA := models.MakeIpamSubnetType()
	ipamSubnetA.Subnet = &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}
	ipamSubnetA.AllocationPools = []*models.AllocationPoolType{&models.AllocationPoolType{Start: "10.0.0.0", End: "10.0.0.10"}}
	ipamSubnetA.SubnetUUID = "5d54b8ca-e5d4-4cac-bdaa-3acc8caac4b1"
	ipamSubnetB := models.MakeIpamSubnetType()
	ipamSubnetB.Subnet = &models.SubnetType{IPPrefix: "10.0.1.0", IPPrefixLen: 24}
	ipamSubnetB.AllocationPools = []*models.AllocationPoolType{&models.AllocationPoolType{Start: "10.0.1.0", End: "10.0.1.10"}}
	ipamSubnetB.SubnetUUID = "5d54b8ca-e5d4-4cac-bdaa-beefbeefbee2"

	networkIpam := models.MakeNetworkIpam()
	networkIpam.IpamSubnetMethod = ipamSubnetMethod
	networkIpam.IpamSubnets.Subnets = []*models.IpamSubnetType{
		ipamSubnetA,
		ipamSubnetB,
	}

	dbServiceMock.EXPECT().GetNetworkIpam(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetNetworkIpamResponse{
			NetworkIpam: networkIpam,
		}, nil).AnyTimes()
}

func virtualNetworkSetupIntPoolAllocatorMocks(s *ContrailTypeLogicService) {
	intPoolAllocator := s.IntPoolAllocator.(*ipammock.MockIntPoolAllocator)
	intPoolAllocator.EXPECT().AllocateInt(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		int64(13), nil).AnyTimes()
	intPoolAllocator.EXPECT().DeallocateInt(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), int64(0)).Return(
		nil).AnyTimes()
}

func virtualNetworkSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService)
	nextService.EXPECT().DeleteVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).AnyTimes()
	nextService.EXPECT().CreateVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(ctx context.Context, request *services.CreateVirtualNetworkRequest) (*services.CreateVirtualNetworkResponse, error) {
			return &services.CreateVirtualNetworkResponse{
				VirtualNetwork: request.VirtualNetwork,
			}, nil
		}).AnyTimes()
}

func TestCreateVirtualNetwork(t *testing.T) {
	ipamSubnetUserDefined := models.MakeIpamSubnetType()
	ipamSubnetUserDefined.Subnet = &models.SubnetType{
		IPPrefix:    "10.0.0.0",
		IPPrefixLen: 24,
	}
	ipamSubnetUserDefined.AllocationPools = []*models.AllocationPoolType{
		&models.AllocationPoolType{
			Start: "10.0.0.5",
			End:   "10.0.0.20",
		},
	}
	ipamSubnetUserDefined.SubnetUUID = "5d54b8ca-e5d4-4cac-bdaa-beefbeefbee3"

	var tests = []struct {
		name             string
		testVnData       *testVn
		ipamSubnetMethod string
		fails            bool
	}{
		{
			name: "check for rt",
			testVnData: &testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:101",
			},
			fails: true,
		},
		{
			name: "check for virtual network id",
			testVnData: &testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:102",
				VirtualNetworkNetworkID:         9999,
			},
			fails: true,
		},
		{
			name: "check for virtual network id",
			testVnData: &testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:102",
			},
			fails: false,
		},
		{
			name: "check for user defined subnet",
			testVnData: &testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:102",
				NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "ipam_uuid",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								ipamSubnetUserDefined,
							},
						},
					},
				},
			},
			fails: false,
		},
		{
			name: "Check for flat subnet with user defined subnet",
			testVnData: &testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:102",
				NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "ipam_uuid",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								ipamSubnetUserDefined,
							},
						},
					},
				},
			},
			ipamSubnetMethod: "flat-subnet",
			fails:            true,
		},
		{
			name: "Check for flat subnet with l2_l3 network",
			testVnData: &testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:102",
				ForwardingMode:                  "l2_l3",
				NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "ipam_uuid",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								models.MakeIpamSubnetType(),
							},
						},
					},
				},
			},
			ipamSubnetMethod: "flat-subnet",
			fails:            true,
		},
		{
			name: "Check for flat subnet",
			testVnData: &testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:102",
				ForwardingMode:                  "l3",
				NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "ipam_uuid",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								models.MakeIpamSubnetType(),
							},
						},
					},
				},
			},
			ipamSubnetMethod: "flat-subnet",
			fails:            false,
		},
		{
			name: "Check for flat subnet",
			testVnData: &testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:102",
				ForwardingMode:                  "l3",
				NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "ipam_uuid",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								models.MakeIpamSubnetType(),
							},
						},
					},
				},
			},
			ipamSubnetMethod: "flat-subnet",
			fails:            false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(t, mockCtrl)

			virtualNetworkSetupDBMocks(service)
			virtualNetworkSetupIntPoolAllocatorMocks(service)
			virtualNetworkSetupNextServiceMocks(service)
			virtualNetworkSetupNetworkIpam(service, tt.ipamSubnetMethod)

			vn := createTestVn(tt.testVnData)
			vnRef := createTestVn(tt.testVnData)
			vnRef.VirtualNetworkNetworkID = 13

			// Put an empty transaction into context so we could call DoInTransaction() without access to the real db
			ctx := context.WithValue(nil, db.Transaction, &sql.Tx{})

			res, err := service.CreateVirtualNetwork(ctx,
				&services.CreateVirtualNetworkRequest{
					VirtualNetwork: vn,
				})
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, vnRef, res.GetVirtualNetwork())
			}
			mockCtrl.Finish()
		})
	}
	//TODO Remaining tests
}

func TestDeleteVirtualNetwork(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	service := makeMockedContrailTypeLogicService(t, mockCtrl)
	virtualNetworkSetupDBMocks(service)
	virtualNetworkSetupIntPoolAllocatorMocks(service)
	virtualNetworkSetupNextServiceMocks(service)

	// Put an empty transaction into context so we could call DoInTransaction() without access to the real db
	ctx := context.WithValue(nil, db.Transaction, &sql.Tx{})

	//Check missing VirtualNetwork in DB (negative)
	_, err := service.DeleteVirtualNetwork(ctx,
		&services.DeleteVirtualNetworkRequest{
			ID: "nonexistent_uuid",
		})
	assert.Error(t, err)

	//Check DeleteVirtualNetwork (positive)
	_, err = service.DeleteVirtualNetwork(ctx,
		&services.DeleteVirtualNetworkRequest{
			ID: "test_vn_uuid",
		})
	assert.NoErrorf(t, err, "DeleteVirtualNetwork Failed %v", err)
	mockCtrl.Finish()
}
