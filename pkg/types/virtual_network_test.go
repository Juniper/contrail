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
	VirtualNetworkRefs              []*models.VirtualNetworkVirtualNetworkRef
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

	if len(testVnData.VirtualNetworkRefs) > 0 {
		vn.VirtualNetworkRefs = testVnData.VirtualNetworkRefs
	}

	vn.UUID = "test_vn_uuid"
	vn.VirtualNetworkProperties.ForwardingMode = testVnData.ForwardingMode

	return vn
}

func virtualNetworkSetupDBMocks(s *ContrailTypeLogicService) {
	dbServiceMock := s.DB.(*typesmock.MockDBServiceInterface)

	// Virtual Networks
	dbServiceMock.EXPECT().DB().AnyTimes()
	dbServiceMock.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()),
		&services.GetVirtualNetworkRequest{
			ID: "test_vn_uuid",
		}).Return(
		&services.GetVirtualNetworkResponse{
			VirtualNetwork: models.MakeVirtualNetwork(),
		}, nil).AnyTimes()

	providerNetwork := models.MakeVirtualNetwork()
	providerNetwork.IsProviderNetwork = true
	dbServiceMock.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()),
		&services.GetVirtualNetworkRequest{
			ID: "test_provider_vn_uuid",
		}).Return(
		&services.GetVirtualNetworkResponse{
			VirtualNetwork: providerNetwork,
		}, nil).AnyTimes()

	dbServiceMock.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, common.ErrorNotFound).AnyTimes()

	// BGPVPN
	bgpVPNL3 := models.MakeBGPVPN()
	bgpVPNL3.BGPVPNType = models.L3Mode
	dbServiceMock.EXPECT().GetBGPVPN(gomock.Not(gomock.Nil()),
		&services.GetBGPVPNRequest{
			ID: "bgpvpn_uuid_l3",
		}).Return(
		&services.GetBGPVPNResponse{
			BGPVPN: bgpVPNL3,
		}, nil).AnyTimes()

	bgpVPNAny := models.MakeBGPVPN()
	bgpVPNAny.BGPVPNType = models.L2L3Mode
	dbServiceMock.EXPECT().GetBGPVPN(gomock.Not(gomock.Nil()),
		&services.GetBGPVPNRequest{
			ID: "bgpvpn_uuid_any",
		}).Return(
		&services.GetBGPVPNResponse{
			BGPVPN: bgpVPNAny,
		}, nil).AnyTimes()

	dbServiceMock.EXPECT().GetBGPVPN(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(nil,
		common.ErrorNotFound).AnyTimes()

	// Logical Routers
	logicalRouter := models.MakeLogicalRouter()
	dbServiceMock.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()),
		&services.GetLogicalRouterRequest{
			ID: "logical_router_uuid",
		}).Return(
		&services.GetLogicalRouterResponse{
			LogicalRouter: logicalRouter,
		}, nil).AnyTimes()

	logicalRouterWithBGPVPN := models.MakeLogicalRouter()
	logicalRouterWithBGPVPN.BGPVPNRefs = []*models.LogicalRouterBGPVPNRef{
		{
			UUID: "bgpvpn_uuid_l3",
			To:   []string{"logical_router_with_bgpvpn_uuid"},
		},
	}
	dbServiceMock.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()),
		&services.GetLogicalRouterRequest{
			ID: "logical_router_with_bgpvpn_uuid",
		}).Return(
		&services.GetLogicalRouterResponse{
			LogicalRouter: logicalRouterWithBGPVPN,
		}, nil).AnyTimes()
}

func virtualNetworkSetupNetworkIpam(s *ContrailTypeLogicService, ipamSubnetMethod string) {
	dbServiceMock := s.DB.(*typesmock.MockDBServiceInterface)

	ipamSubnetA := models.MakeIpamSubnetType()
	ipamSubnetA.Subnet = &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}
	ipamSubnetA.AllocationPools = []*models.AllocationPoolType{{Start: "10.0.0.0", End: "10.0.0.10"}}
	ipamSubnetA.SubnetUUID = "5d54b8ca-e5d4-4cac-bdaa-3acc8caac4b1"
	ipamSubnetB := models.MakeIpamSubnetType()
	ipamSubnetB.Subnet = &models.SubnetType{IPPrefix: "10.0.1.0", IPPrefixLen: 24}
	ipamSubnetB.AllocationPools = []*models.AllocationPoolType{{Start: "10.0.1.0", End: "10.0.1.10"}}
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
		{
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
			name:       "check for virtual network id",
			testVnData: &testVn{},
			fails:      false,
		},
		{
			name: "check for provider network ref",
			testVnData: &testVn{
				VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
					{
						UUID: "test_provider_vn_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			fails: false,
		},
		{
			name: "check for non-provider network ref",
			testVnData: &testVn{
				VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
					{
						UUID: "test_vn_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			fails: true,
		},
		{
			name: "check for multiple network refs",
			testVnData: &testVn{
				VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
					{
						UUID: "test_provider_vn_uuid",
						To:   []string{"test_vn_uuid"},
					},
					{
						UUID: "test_vn_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			fails: true,
		},
		{
			name: "check for non-existing network refs",
			testVnData: &testVn{
				VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
					{
						UUID: "test_non-existing_vn_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			fails: true,
		},
		{
			name: "check for user defined subnet",
			testVnData: &testVn{
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
				ForwardingMode: models.L2L3Mode,
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
				ForwardingMode: models.L3Mode,
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
			name: "Check for bgp",
			testVnData: &testVn{
				ForwardingMode: models.L3Mode,
				BGPVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_l3",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			ipamSubnetMethod: "flat-subnet",
			fails:            false,
		},
		{
			name: "Check for bgp with wrong Forwarding mode",
			testVnData: &testVn{
				ForwardingMode: models.L3Mode,
				BGPVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_any",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			ipamSubnetMethod: "flat-subnet",
			fails:            true,
		},
		{
			name: "Check for bgp with non-existing",
			testVnData: &testVn{
				ForwardingMode: models.L3Mode,
				BGPVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_wrong_non_existing",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			ipamSubnetMethod: "flat-subnet",
			fails:            true,
		},
		{
			name: "Check for logical routers refs",
			testVnData: &testVn{
				ForwardingMode: models.L3Mode,
				BGPVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_l3",
						To:   []string{"test_vn_uuid"},
					},
				},

				LogicalRouterRefs: []*models.VirtualNetworkLogicalRouterRef{
					{
						UUID: "logical_router_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			ipamSubnetMethod: "flat-subnet",
			fails:            false,
		},
		{
			name: "Check for logical routers with bgpvpn refs",
			testVnData: &testVn{
				ForwardingMode: models.L3Mode,
				BGPVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_l3",
						To:   []string{"test_vn_uuid"},
					},
				},
				LogicalRouterRefs: []*models.VirtualNetworkLogicalRouterRef{
					{
						UUID: "logical_router_with_bgpvpn_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			ipamSubnetMethod: "flat-subnet",
			fails:            true,
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
