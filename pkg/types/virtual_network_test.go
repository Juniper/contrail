package types

import (
	"context"
	"net/http"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam"
	ipammock "github.com/Juniper/contrail/pkg/types/ipam/mock"
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
)

//Structure testVn is used to pass vn parameters during VirtualNetwork object creation
type testVn struct {
	isProviderNetwork               bool
	multiPolicyServiceChainsEnabled bool
	importRouteTargetList           string
	exportRouteTargetList           string
	forwardingMode                  string
	vxlanID                         int64
	virtualNetworkNetworkID         int64
	networkIpamRefs                 []*models.VirtualNetworkNetworkIpamRef
	bgpVPNRefs                      []*models.VirtualNetworkBGPVPNRef
	logicalRouterRefs               []*models.VirtualNetworkLogicalRouterRef
	virtualNetworkRefs              []*models.VirtualNetworkVirtualNetworkRef
	instanceIPBackRefs              []*models.InstanceIP
}

func TestCreateVirtualNetwork(t *testing.T) {
	ipamSubnetUserDefined := virtualNetworkMakeUserDefinedSubnet("5d54b8ca-e5d4-4cac-bdaa-beefbeefbee3")
	var tests = []struct {
		name                  string
		testVnData            *testVn
		ipamSubnetMethod      string
		createsSubnet         bool
		fails                 bool
		expectedHTTPErrorCode int
	}{
		{
			name:  "check for rt",
			fails: true,
			testVnData: &testVn{
				multiPolicyServiceChainsEnabled: true,
				importRouteTargetList:           "100:101",
				exportRouteTargetList:           "100:101",
			},
			expectedHTTPErrorCode: http.StatusBadRequest,
		},
		{
			name: "check for rt",
			testVnData: &testVn{
				multiPolicyServiceChainsEnabled: true,
				importRouteTargetList:           "100:101",
				exportRouteTargetList:           "100:102",
			},
		},
		{
			name: "check for MultiPolicyServiceChainsEnabled",
			testVnData: &testVn{
				multiPolicyServiceChainsEnabled: false,
			},
		},
		{
			name:  "check for is provider network",
			fails: true,
			testVnData: &testVn{
				isProviderNetwork: true,
			},
		},
		{
			name:  "check for virtual network id",
			fails: true,
			testVnData: &testVn{
				virtualNetworkNetworkID: 9999,
			},
			expectedHTTPErrorCode: http.StatusForbidden,
		},
		{
			name:       "check for virtual network id",
			testVnData: &testVn{},
		},
		{
			name: "check for provider network ref",
			testVnData: &testVn{
				virtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
					{
						UUID: "test_provider_vn_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
		},
		{
			name:  "check for non-provider network ref",
			fails: true,
			testVnData: &testVn{
				virtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
					{
						UUID: "test_vn_red_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			expectedHTTPErrorCode: http.StatusBadRequest,
		},
		{
			name:  "check for multiple network refs",
			fails: true,
			testVnData: &testVn{
				virtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
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
			expectedHTTPErrorCode: http.StatusBadRequest,
		},
		{
			name:  "check for multiple provider network refs",
			fails: true,
			testVnData: &testVn{
				virtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
					{
						UUID: "test_provider_vn_uuid",
						To:   []string{"test_vn_uuid"},
					},
					{
						UUID: "test_provider_vn_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			expectedHTTPErrorCode: http.StatusBadRequest,
		},
		{
			name:  "check for non-existing network refs",
			fails: true,
			testVnData: &testVn{
				virtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
					{
						UUID: "test_non-existing_vn_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			expectedHTTPErrorCode: http.StatusNotFound,
		},
		{
			name: "check for user defined subnet",
			testVnData: &testVn{
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "network_ipam_a",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								ipamSubnetUserDefined,
							},
						},
					},
				},
			},
			ipamSubnetMethod: models.UserDefinedSubnet,
			createsSubnet:    true,
		},
		{
			name: "check for flat subnet",
			testVnData: &testVn{
				forwardingMode: models.L3Mode,
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "network_ipam_a",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								models.MakeIpamSubnetType(),
							},
						},
					},
				},
			},
			ipamSubnetMethod: models.FlatSubnet,
			createsSubnet:    true,
		},
		{
			name:  "check for flat subnet with user defined subnet",
			fails: true,
			testVnData: &testVn{
				forwardingMode: models.L3Mode,
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "network_ipam_a",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								ipamSubnetUserDefined,
							},
						},
					},
				},
			},
			ipamSubnetMethod:      models.FlatSubnet,
			expectedHTTPErrorCode: http.StatusBadRequest,
		},
		{
			name:  "check for flat subnet with l2_l3 network",
			fails: true,
			testVnData: &testVn{
				forwardingMode: models.L2L3Mode,
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "network_ipam_a",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								models.MakeIpamSubnetType(),
							},
						},
					},
				},
			},
			ipamSubnetMethod:      models.FlatSubnet,
			expectedHTTPErrorCode: http.StatusBadRequest,
		},
		{
			name:  "check for overlapping ip addresses",
			fails: true,
			testVnData: &testVn{
				forwardingMode: models.L3Mode,
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "network_ipam_a",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								ipamSubnetUserDefined,
							},
						},
					},
					{
						UUID: "network_ipam_b",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								ipamSubnetUserDefined,
							},
						},
					},
				},
			},
			ipamSubnetMethod:      models.UserDefinedSubnet,
			expectedHTTPErrorCode: http.StatusBadRequest,
		},
		{
			name: "check for bgp",
			testVnData: &testVn{
				forwardingMode: models.L3Mode,
				bgpVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_l3",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
		},
		{
			name:  "check for bgp with wrong Forwarding mode",
			fails: true,
			testVnData: &testVn{
				forwardingMode: models.L3Mode,
				bgpVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_any",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			expectedHTTPErrorCode: http.StatusBadRequest,
		},
		{
			name:  "check for bgp with non-existing",
			fails: true,
			testVnData: &testVn{
				forwardingMode: models.L3Mode,
				bgpVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_wrong_non_existing",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			expectedHTTPErrorCode: http.StatusNotFound,
		},
		{
			name: "check for logical routers refs",
			testVnData: &testVn{
				forwardingMode: models.L3Mode,
				bgpVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_l3",
						To:   []string{"test_vn_uuid"},
					},
				},

				logicalRouterRefs: []*models.VirtualNetworkLogicalRouterRef{
					{
						UUID: "logical_router_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
		},
		{
			name: "check for logical routers with bgpvpn refs",
			testVnData: &testVn{
				forwardingMode: models.L3Mode,
				bgpVPNRefs: []*models.VirtualNetworkBGPVPNRef{
					{
						UUID: "bgpvpn_uuid_l3",
						To:   []string{"test_vn_uuid"},
					},
				},
				logicalRouterRefs: []*models.VirtualNetworkLogicalRouterRef{
					{
						UUID: "logical_router_with_bgpvpn_uuid",
						To:   []string{"test_vn_uuid"},
					},
				},
			},
			fails:                 true,
			expectedHTTPErrorCode: http.StatusBadRequest,
		},
		{
			name: "allocate vxlan id",
			testVnData: &testVn{
				vxlanID: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			virtualNetwork := models.MakeVirtualNetwork()
			virtualNetwork.UUID = "test_vn_red_uuid"
			virtualNetwork.FQName = []string{"test_vn_red"}
			mockedReadServiceAddVirtualNetwork(service, virtualNetwork)

			virtualNetworkSetupReadServiceMocks(service)
			virtualNetworkSetupIntPoolAllocatorMocks(service)
			virtualNetworkSetupNetworkIpam(service, tt.ipamSubnetMethod)

			vn := createTestVn(tt.testVnData)
			resultingVn := createTestVn(tt.testVnData)
			resultingVn.VirtualNetworkNetworkID = 13

			ctx := context.Background()
			// In case of successful flow:
			// CreateVirtualNetwork should be called once on next service
			// CreateRoutingInstance should be called once on the Write service
			if !tt.fails {
				service.Next().(*servicesmock.MockService).
					EXPECT().CreateVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
					func(
						_ context.Context, request *services.CreateVirtualNetworkRequest,
					) (*services.CreateVirtualNetworkResponse, error) {
						return &services.CreateVirtualNetworkResponse{
							VirtualNetwork: request.VirtualNetwork,
						}, nil
					}).Times(1)
				service.WriteService.(*servicesmock.MockWriteService). //nolint: errcheck
											EXPECT().CreateRoutingInstance(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
					func(
						_ context.Context, request *services.CreateRoutingInstanceRequest,
					) (*services.CreateRoutingInstanceResponse, error) {
						return &services.CreateRoutingInstanceResponse{
							RoutingInstance: request.RoutingInstance,
						}, nil
					}).Times(1)
			}

			if tt.createsSubnet {
				virtualNetworkMustCreateSubnet(service)
			}

			if tt.testVnData.vxlanID != 0 {
				intPoolAllocator := service.IntPoolAllocator.(*typesmock.MockIntPoolAllocator) //nolint: errcheck
				intPoolAllocator.EXPECT().SetInt(
					gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
					tt.testVnData.vxlanID, gomock.Not(gomock.Nil()),
				).Return(nil).Times(1)
			}

			res, err := service.CreateVirtualNetwork(ctx,
				&services.CreateVirtualNetworkRequest{
					VirtualNetwork: vn,
				})
			if tt.fails {
				assert.Error(t, err)
				if tt.expectedHTTPErrorCode != 0 {
					httpError, ok := errutil.ToHTTPError(err).(*echo.HTTPError)
					assert.True(t, ok, "Expected http error")
					assert.Equal(t, tt.expectedHTTPErrorCode, httpError.Code, "Expected different http status")
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, resultingVn, res.GetVirtualNetwork())
			}
		})
	}
}

func TestUpdateVirtualNetwork(t *testing.T) {
	var tests = []struct {
		name                  string
		testVnData            *testVn
		updateRequest         *services.UpdateVirtualNetworkRequest
		ipamSubnetMethod      string
		fails                 bool
		addressManagerSetup   func(s *ContrailTypeLogicService)
		intPoolAllocatorSetup func(s *ContrailTypeLogicService)
		expectedHTTPErrorCode int
	}{
		{
			name:                  "check for update with a different VirtualNetworkID",
			fails:                 false, //TODO: This test should fail
			expectedHTTPErrorCode: http.StatusForbidden,
			testVnData: &testVn{
				virtualNetworkNetworkID: 13,
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID:                    "test_vn_uuid",
					VirtualNetworkNetworkID: 16,
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualNetworkFieldVirtualNetworkNetworkID,
					},
				},
			},
		},
		{
			name:                  "check for update with the same VirtualNetworkID",
			fails:                 false,
			expectedHTTPErrorCode: http.StatusForbidden,
			testVnData: &testVn{
				virtualNetworkNetworkID: 13,
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID:                    "test_vn_uuid",
					VirtualNetworkNetworkID: 13,
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualNetworkFieldVirtualNetworkNetworkID,
					},
				},
			},
		},
		{
			name:                  "check is_provider_network update",
			fails:                 true,
			expectedHTTPErrorCode: http.StatusBadRequest,
			testVnData:            &testVn{},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID:              "test_vn_uuid",
					IsProviderNetwork: true,
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualNetworkFieldIsProviderNetwork,
					},
				},
			},
		},
		{
			name:       "check is_provider_network update",
			testVnData: &testVn{},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "test_vn_uuid",
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualNetworkFieldIsProviderNetwork,
					},
				},
			},
		},
		{
			name:                  "check if provider network can be linked to a provider network",
			fails:                 true,
			expectedHTTPErrorCode: http.StatusBadRequest,
			testVnData: &testVn{
				isProviderNetwork: true,
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "test_vn_uuid",
					VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
						{
							UUID: "test_provider_vn_uuid",
							To:   []string{"test_vn_uuid"},
						},
						{
							UUID: "test_non_provider_vn_uuid",
							To:   []string{"test_vn_uuid"},
						},
					},
				},
				FieldMask: types.FieldMask{},
			},
		},
		{
			name:                  "check if provider network can be linked to non-provider networks",
			expectedHTTPErrorCode: http.StatusBadRequest,
			testVnData: &testVn{
				isProviderNetwork: true,
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "test_vn_uuid",
					VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
						{
							UUID: "test_non_provider_vn_uuid",
							To:   []string{"test_vn_uuid"},
						},
						{
							UUID: "test_non_provider_vn_uuid",
							To:   []string{"test_vn_uuid"},
						},
					},
				},
				FieldMask: types.FieldMask{},
			},
		},
		{
			name:                  "add a new subnet to virtual network",
			expectedHTTPErrorCode: http.StatusBadRequest,
			testVnData: &testVn{
				isProviderNetwork: true,
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "test_vn_uuid",
					NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
						{
							UUID: "network_ipam_a",
							To:   []string{"test_vn_uuid"},
							Attr: &models.VnSubnetsType{
								IpamSubnets: []*models.IpamSubnetType{
									virtualNetworkMakeUserDefinedSubnet("5d54b8ca-e5d4-4cac-bdaa-beefbeefbee3"),
								},
							},
						},
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualNetworkFieldNetworkIpamRefs,
					},
				},
			},
			addressManagerSetup: func(s *ContrailTypeLogicService) {
				virtualNetworkMustCreateSubnet(s)
			},
		},
		{
			name:                  "delete subnet on virtual network update",
			expectedHTTPErrorCode: http.StatusBadRequest,
			testVnData: &testVn{
				isProviderNetwork: true,
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "network_ipam_a",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								virtualNetworkMakeUserDefinedSubnet("5d54b8ca-e5d4-4cac-bdaa-beefbeefbee3"),
							},
						},
					},
				},
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "test_vn_uuid",
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualNetworkFieldNetworkIpamRefs,
					},
				},
			},
			addressManagerSetup: func(s *ContrailTypeLogicService) {
				virtualNetworkMustDeleteSubnet(s)
			},
		},
		{
			name:                  "try to delete subnet on virtual network update without fieldmask",
			expectedHTTPErrorCode: http.StatusBadRequest,
			testVnData: &testVn{
				isProviderNetwork: true,
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "network_ipam_a",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								virtualNetworkMakeUserDefinedSubnet("5d54b8ca-e5d4-4cac-bdaa-beefbeefbee3"),
							},
						},
					},
				},
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "test_vn_uuid",
				},
				FieldMask: types.FieldMask{
					Paths: []string{},
				},
			},
		},
		{
			name:                  "try delete subnet with instance ip on virtual network update",
			fails:                 true,
			expectedHTTPErrorCode: http.StatusConflict,
			testVnData: &testVn{
				isProviderNetwork: true,
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "network_ipam_a",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								virtualNetworkMakeUserDefinedSubnet("5d54b8ca-e5d4-4cac-bdaa-beefbeefbee3"),
							},
						},
					},
				},
				instanceIPBackRefs: []*models.InstanceIP{
					{
						UUID: "instance_ip_uuid",
					},
				},
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "test_vn_uuid",
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualNetworkFieldNetworkIpamRefs,
					},
				},
			},
		},
		{
			name:                  "delete and add subnet on virtual network update",
			expectedHTTPErrorCode: http.StatusBadRequest,
			testVnData: &testVn{
				isProviderNetwork: true,
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						UUID: "network_ipam_a",
						To:   []string{"test_vn_uuid"},
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								virtualNetworkMakeUserDefinedSubnet("5d54b8ca-e5d4-4cac-bdaa-beefbeefbee3"),
							},
						},
					},
				},
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "test_vn_uuid",
					NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
						{
							UUID: "network_ipam_a",
							To:   []string{"test_vn_uuid"},
							Attr: &models.VnSubnetsType{
								IpamSubnets: []*models.IpamSubnetType{
									virtualNetworkMakeUserDefinedSubnet("df0fc1fd-9a56-492c-b04b-ae4917c3543d"),
								},
							},
						},
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualNetworkFieldNetworkIpamRefs,
					},
				},
			},
			addressManagerSetup: func(s *ContrailTypeLogicService) {
				virtualNetworkMustCreateSubnet(s)
				virtualNetworkMustDeleteSubnet(s)
			},
		},
		{
			name: "update VxLAN ID",
			testVnData: &testVn{
				vxlanID: 10,
			},
			updateRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "test_vn_uuid",
					VirtualNetworkProperties: &models.VirtualNetworkType{
						VxlanNetworkIdentifier: 11,
					},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						basemodels.JoinPath(
							models.VirtualNetworkFieldVirtualNetworkProperties,
							models.VirtualNetworkTypeFieldVxlanNetworkIdentifier,
						),
					},
				},
			},
			intPoolAllocatorSetup: func(s *ContrailTypeLogicService) {
				intPoolAllocator := s.IntPoolAllocator.(*typesmock.MockIntPoolAllocator) //nolint: errcheck
				intPoolAllocator.EXPECT().DeallocateInt(
					gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), int64(10),
				).Return(nil).Times(1)
				intPoolAllocator.EXPECT().SetInt(
					gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), int64(11), gomock.Not(gomock.Nil()),
				).Return(nil).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			vn := createTestVn(tt.testVnData)
			mockedReadServiceAddVirtualNetwork(service, vn)
			virtualNetworkSetupReadServiceMocks(service)
			virtualNetworkSetupNetworkIpam(service, models.UserDefinedSubnet)

			ctx := context.Background()
			// In case of successful flow UpdateVirtualNetwork should be called once on next service
			if !tt.fails {
				nextService := service.Next().(*servicesmock.MockService) //nolint: errcheck
				nextService.EXPECT().UpdateVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
					func(ctx context.Context, request *services.UpdateVirtualNetworkRequest) ( //nolint: unparam
						*services.UpdateVirtualNetworkResponse, error) {
						return &services.UpdateVirtualNetworkResponse{
							VirtualNetwork: request.VirtualNetwork,
						}, nil
					}).Times(1)
			}

			if tt.addressManagerSetup != nil {
				tt.addressManagerSetup(service)
			}

			if tt.intPoolAllocatorSetup != nil {
				tt.intPoolAllocatorSetup(service)
			}

			res, err := service.UpdateVirtualNetwork(ctx, tt.updateRequest)

			if tt.fails {
				assert.Error(t, err)
				if tt.expectedHTTPErrorCode != 0 {
					httpError, ok := errutil.ToHTTPError(err).(*echo.HTTPError)
					assert.True(t, ok, "Expected http error")
					assert.Equal(t, tt.expectedHTTPErrorCode, httpError.Code, "Expected different http status")
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.updateRequest.GetVirtualNetwork(), res.GetVirtualNetwork())
			}
		})
	}
}

func TestDeleteVirtualNetwork(t *testing.T) {
	ipamSubnetUserDefined := virtualNetworkMakeUserDefinedSubnet("5d54b8ca-e5d4-4cac-bdaa-beefbeefbee3")
	var tests = []struct {
		name          string
		UUID          string
		testVnData    *testVn
		fails         bool
		deletesSubnet bool
	}{
		{
			name:  "check missing VirtualNetwork in DB",
			fails: true,
			UUID:  "nonexistent_uuid",
		},
		{
			name: "check DeleteVirtualNetwork",
			UUID: "test_vn_uuid",
			testVnData: &testVn{
				networkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
					{
						Attr: &models.VnSubnetsType{
							IpamSubnets: []*models.IpamSubnetType{
								ipamSubnetUserDefined,
							},
						},
					},
				},
				virtualNetworkNetworkID: 20,
			},
			deletesSubnet: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			if tt.testVnData != nil {
				virtualNetwork := createTestVn(tt.testVnData)
				virtualNetwork.UUID = tt.UUID
				mockedReadServiceAddVirtualNetwork(service, virtualNetwork)
			}

			virtualNetworkSetupReadServiceMocks(service)
			virtualNetworkSetupIntPoolAllocatorMocksForDeletion(service)

			if tt.deletesSubnet {
				virtualNetworkMustDeleteSubnet(service)
			}

			ctx := context.Background()
			// In case of successful flow DeleteVirtualNetwork should be called once on next service
			if !tt.fails {
				nextService := service.Next().(*servicesmock.MockService) //nolint: errcheck
				nextService.EXPECT().DeleteVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
					func(
						_ context.Context, request *services.DeleteVirtualNetworkRequest,
					) (*services.DeleteVirtualNetworkResponse, error) {
						return &services.DeleteVirtualNetworkResponse{
							ID: request.ID,
						}, nil
					}).Times(1)
			}

			res, err := service.DeleteVirtualNetwork(ctx,
				&services.DeleteVirtualNetworkRequest{
					ID: tt.UUID,
				})

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, res.GetID(), tt.UUID)
			}
		})
	}
}

func TestDeleteDefaultRoutingInstance(t *testing.T) {
	tests := []struct {
		name                     string
		routingInstances         []*models.RoutingInstance
		shouldDelete             bool
		defaultRoutingInstanceID string
	}{
		{
			name: "check if default routing instance is deleted",
			routingInstances: []*models.RoutingInstance{
				{UUID: "ri_def_uuid", RoutingInstanceIsDefault: true},
				{UUID: "ri_some_uuid"},
			},
			shouldDelete:             true,
			defaultRoutingInstanceID: "ri_def_uuid",
		},
		{
			name: "check if no routing instances are deleted other than default",
			routingInstances: []*models.RoutingInstance{
				{UUID: "ri_some_uuid"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			virtualNetwork := models.MakeVirtualNetwork()
			virtualNetwork.UUID = "test_vn_red_uuid"
			virtualNetwork.VirtualNetworkNetworkID = 20
			virtualNetwork.RoutingInstances = tt.routingInstances
			service.ReadService.(*servicesmock.MockReadService).EXPECT( //nolint: errcheck
			).GetVirtualNetwork(gomock.Not(gomock.Nil()),
				&services.GetVirtualNetworkRequest{
					ID: virtualNetwork.UUID,
				}).Return(&services.GetVirtualNetworkResponse{VirtualNetwork: virtualNetwork}, nil).Times(1)

			virtualNetworkSetupIntPoolAllocatorMocksForDeletion(service)

			ctx := context.Background()

			expectedDeletes := 0
			if tt.shouldDelete {
				expectedDeletes = 1
			}
			service.WriteService.(*servicesmock.MockWriteService). //nolint: errcheck
										EXPECT().DeleteRoutingInstance(gomock.Not(gomock.Nil()),
				&services.DeleteRoutingInstanceRequest{
					ID: tt.defaultRoutingInstanceID,
				}).Return(&services.DeleteRoutingInstanceResponse{}, nil).Times(expectedDeletes)

			service.Next().(*servicesmock.MockService).
				EXPECT().DeleteVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
				Return(&services.DeleteVirtualNetworkResponse{}, nil).Times(1)

			_, err := service.DeleteVirtualNetwork(ctx,
				&services.DeleteVirtualNetworkRequest{
					ID: virtualNetwork.UUID,
				})

			assert.NoError(t, err)
		})
	}
}

func createTestVn(testVnData *testVn) *models.VirtualNetwork {
	vn := models.MakeVirtualNetwork()
	vn.IsProviderNetwork = testVnData.isProviderNetwork
	vn.MultiPolicyServiceChainsEnabled = testVnData.multiPolicyServiceChainsEnabled
	vn.ImportRouteTargetList = &models.RouteTargetList{RouteTarget: []string{testVnData.importRouteTargetList}}
	vn.ExportRouteTargetList = &models.RouteTargetList{RouteTarget: []string{testVnData.exportRouteTargetList}}
	vn.VirtualNetworkNetworkID = testVnData.virtualNetworkNetworkID
	if len(testVnData.networkIpamRefs) > 0 {
		vn.NetworkIpamRefs = testVnData.networkIpamRefs
	}

	if len(testVnData.bgpVPNRefs) > 0 {
		vn.BGPVPNRefs = testVnData.bgpVPNRefs
	}

	if len(testVnData.logicalRouterRefs) > 0 {
		vn.LogicalRouterRefs = testVnData.logicalRouterRefs
	}

	if len(testVnData.virtualNetworkRefs) > 0 {
		vn.VirtualNetworkRefs = testVnData.virtualNetworkRefs
	}

	if len(testVnData.instanceIPBackRefs) > 0 {
		vn.InstanceIPBackRefs = testVnData.instanceIPBackRefs
	}

	vn.UUID = "test_vn_uuid"
	vn.FQName = []string{"test_vn_uuid"}
	vn.VirtualNetworkProperties.ForwardingMode = testVnData.forwardingMode
	vn.VirtualNetworkProperties.VxlanNetworkIdentifier = testVnData.vxlanID

	return vn
}

func virtualNetworkSetupReadServiceMocks(s *ContrailTypeLogicService) {
	readServiceMock := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	virtualNetwork := models.MakeVirtualNetwork()
	virtualNetwork.UUID = "test_provider_vn_uuid"
	virtualNetwork.FQName = []string{"test_provider_vn"}
	virtualNetwork.IsProviderNetwork = true
	mockedReadServiceAddVirtualNetwork(s, virtualNetwork)

	virtualNetwork = models.MakeVirtualNetwork()
	virtualNetwork.UUID = "test_non_provider_vn_uuid"
	virtualNetwork.FQName = []string{"test_non_provider_vn"}
	mockedReadServiceAddVirtualNetwork(s, virtualNetwork)

	readServiceMock.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, errutil.ErrorNotFound).AnyTimes()

	// BGPVPN
	bgpVPNL3 := models.MakeBGPVPN()
	bgpVPNL3.BGPVPNType = models.L3Mode
	readServiceMock.EXPECT().GetBGPVPN(gomock.Not(gomock.Nil()),
		&services.GetBGPVPNRequest{
			ID: "bgpvpn_uuid_l3",
		}).Return(
		&services.GetBGPVPNResponse{
			BGPVPN: bgpVPNL3,
		}, nil).AnyTimes()

	bgpVPNAny := models.MakeBGPVPN()
	bgpVPNAny.BGPVPNType = models.L2L3Mode
	readServiceMock.EXPECT().GetBGPVPN(gomock.Not(gomock.Nil()),
		&services.GetBGPVPNRequest{
			ID: "bgpvpn_uuid_any",
		}).Return(
		&services.GetBGPVPNResponse{
			BGPVPN: bgpVPNAny,
		}, nil).AnyTimes()

	readServiceMock.EXPECT().GetBGPVPN(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(nil,
		errutil.ErrorNotFound).AnyTimes()

	// Logical Routers
	logicalRouter := models.MakeLogicalRouter()
	readServiceMock.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()),
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
	readServiceMock.EXPECT().GetLogicalRouter(gomock.Not(gomock.Nil()),
		&services.GetLogicalRouterRequest{
			ID: "logical_router_with_bgpvpn_uuid",
		}).Return(
		&services.GetLogicalRouterResponse{
			LogicalRouter: logicalRouterWithBGPVPN,
		}, nil).AnyTimes()

	instanceIP := models.MakeInstanceIP()
	instanceIP.InstanceIPAddress = "10.0.0.6"
	readServiceMock.EXPECT().GetInstanceIP(gomock.Any(),
		&services.GetInstanceIPRequest{
			ID: "instance_ip_uuid",
		}).Return(
		&services.GetInstanceIPResponse{
			InstanceIP: instanceIP,
		}, nil).AnyTimes()
}

func virtualNetworkSetupNetworkIpam(s *ContrailTypeLogicService, ipamSubnetMethod string) {
	readServiceMock := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	ipamSubnetA := models.MakeIpamSubnetType()
	ipamSubnetA.Subnet = &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}
	ipamSubnetA.AllocationPools = []*models.AllocationPoolType{{Start: "10.0.0.0", End: "10.0.0.10"}}
	ipamSubnetA.SubnetUUID = "5d54b8ca-e5d4-4cac-bdaa-3acc8caac4b1"
	ipamSubnetB := models.MakeIpamSubnetType()
	ipamSubnetB.Subnet = &models.SubnetType{IPPrefix: "10.0.1.0", IPPrefixLen: 24}
	ipamSubnetB.AllocationPools = []*models.AllocationPoolType{{Start: "10.0.1.0", End: "10.0.1.10"}}
	ipamSubnetB.SubnetUUID = "5d54b8ca-e5d4-4cac-bdaa-beefbeefbee2"

	networkIpamA := models.MakeNetworkIpam()
	networkIpamA.IpamSubnetMethod = ipamSubnetMethod
	networkIpamA.IpamSubnets.Subnets = []*models.IpamSubnetType{
		ipamSubnetA,
		ipamSubnetB,
	}

	readServiceMock.EXPECT().GetNetworkIpam(gomock.Not(gomock.Nil()),
		&services.GetNetworkIpamRequest{
			ID: "network_ipam_a",
		}).Return(
		&services.GetNetworkIpamResponse{
			NetworkIpam: networkIpamA,
		}, nil).AnyTimes()

	networkIpamB := models.MakeNetworkIpam()
	networkIpamB.IpamSubnetMethod = ipamSubnetMethod
	networkIpamB.IpamSubnets.Subnets = []*models.IpamSubnetType{
		ipamSubnetA,
	}

	readServiceMock.EXPECT().GetNetworkIpam(gomock.Not(gomock.Nil()),
		&services.GetNetworkIpamRequest{
			ID: "network_ipam_b",
		}).Return(
		&services.GetNetworkIpamResponse{
			NetworkIpam: networkIpamB,
		}, nil).AnyTimes()

	readServiceMock.EXPECT().GetNetworkIpam(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, errutil.ErrorNotFound).AnyTimes()
}

func virtualNetworkSetupIntPoolAllocatorMocks(s *ContrailTypeLogicService) {
	intPoolAllocator := s.IntPoolAllocator.(*typesmock.MockIntPoolAllocator) //nolint: errcheck
	intPoolAllocator.EXPECT().AllocateInt(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).Return(int64(13), nil).AnyTimes()
}

func virtualNetworkSetupIntPoolAllocatorMocksForDeletion(s *ContrailTypeLogicService) {
	intPoolAllocator := s.IntPoolAllocator.(*typesmock.MockIntPoolAllocator) //nolint: errcheck
	intPoolAllocator.EXPECT().AllocateInt(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).Return(int64(13), nil).AnyTimes()
	intPoolAllocator.EXPECT().DeallocateInt(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), int64(20),
	).Return(nil).AnyTimes()
}

func virtualNetworkMustCreateSubnet(s *ContrailTypeLogicService) {
	addressManager := s.AddressManager.(*ipammock.MockAddressManager) //nolint: errcheck
	addressManager.EXPECT().CheckIfIpamSubnetExists(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		false, nil,
	).AnyTimes()

	addressManager.EXPECT().CreateIpamSubnet(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(_ context.Context, request *ipam.CreateIpamSubnetRequest) (subnetUUID string, err error) {
			return request.IpamSubnet.GetSubnetUUID(), nil
		}).Times(1)
}

func virtualNetworkMustDeleteSubnet(s *ContrailTypeLogicService) {
	addressManager := s.AddressManager.(*ipammock.MockAddressManager) //nolint: errcheck
	addressManager.EXPECT().DeleteIpamSubnet(gomock.Any(), gomock.Any()).Return(nil).Times(1)
}

func virtualNetworkMakeUserDefinedSubnet(subnetUUID string) *models.IpamSubnetType {
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
	ipamSubnetUserDefined.SubnetUUID = subnetUUID
	return ipamSubnetUserDefined
}
