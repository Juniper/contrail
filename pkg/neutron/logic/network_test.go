package logic

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

func TestNetworkReadAll(t *testing.T) {
	type readData struct {
		vnRes *services.ListVirtualNetworkResponse
	}

	tests := []struct {
		name     string
		n        *Network
		filters  Filters
		fail     bool
		readData *readData
		expRes   Response
	}{
		{
			name: "read all vns",
			n:    &Network{},
			filters: Filters{
				"shared":    []string{"false"},
				"tenant_id": []string{"uuidp"},
			},
			readData: &readData{
				vnRes: &services.ListVirtualNetworkResponse{
					VirtualNetworks: []*models.VirtualNetwork{
						{
							UUID:        "uuid-vn",
							ParentType:  "project",
							ParentUUID:  "uuid-p",
							DisplayName: "vn-test",
							IDPerms: &models.IdPermsType{
								Enable:       true,
								Created:      "data1",
								LastModified: "data2",
							},
							NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
								{
									Attr: &models.VnSubnetsType{
										IpamSubnets: []*models.IpamSubnetType{
											{
												SubnetUUID: "uuid-sn",
												Subnet: &models.SubnetType{
													IPPrefix:    "10.0.1.0",
													IPPrefixLen: 24,
												},
											},
										},
									},
									To: []string{"fqname", "ipam", "test"},
								},
							},
							PortSecurityEnabled: true,
							FQName:              []string{"fqname", "vn", "test"},
						},
					},
				},
			},
			expRes: []*NetworkResponse{
				{
					Status:              netStatusActive,
					RouterExternal:      false,
					Subnets:             []string{"uuid-sn"},
					FQName:              []string{"fqname", "vn", "test"},
					Name:                "vn-test",
					AdminStateUp:        true,
					TenantID:            "uuidp",
					ProjectID:           "uuidp",
					CreatedAt:           "data1",
					UpdatedAt:           "data2",
					PortSecurityEnabled: true,
					Shared:              false,
					ID:                  "uuid-vn",
					SubnetIpam: []*SubnetIpam{
						{
							SubnetCidr: "10.0.1.0/24",
							IpamFQName: []string{"fqname", "ipam", "test"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockServ := servicesmock.NewMockService(mockCtrl)
			ctx := context.Background()

			if tt.readData.vnRes != nil {
				mockServ.EXPECT().ListVirtualNetwork(gomock.Any(), gomock.Any()).Return(tt.readData.vnRes, nil)
			}

			rp := RequestParameters{
				ReadService:  mockServ,
				WriteService: mockServ,
			}

			r, err := tt.n.ReadAll(ctx, rp, nil, nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.expRes, r)

			if tt.fail {
				assert.Error(t, err)
			}
		})
	}
}

func TestNetworkDelete(t *testing.T) {
	type readData struct {
		vnRes   *services.GetVirtualNetworkResponse
		fippRes *services.ListFloatingIPPoolResponse
		fipsRes *services.ListFloatingIPResponse
	}

	type deleteData struct {
		vnReq   *services.DeleteVirtualNetworkRequest
		fippReq *services.DeleteFloatingIPPoolRequest
		fipReq  *services.DeleteFloatingIPRequest
	}

	tests := []struct {
		name       string
		n          *Network
		expected   *NetworkResponse
		fail       bool
		id         string
		readData   *readData
		deleteData *deleteData
	}{
		{
			name:     "Delete network succeeds",
			n:        &Network{},
			expected: &NetworkResponse{},
			id:       "test-vn",
			readData: &readData{
				vnRes: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{
						UUID:           "test-vn",
						RouterExternal: true,
						FloatingIPPools: []*models.FloatingIPPool{
							{UUID: "test-fipp"},
						},
					},
				},
				fippRes: &services.ListFloatingIPPoolResponse{
					FloatingIPPools: []*models.FloatingIPPool{
						{
							UUID: "test-fipp",
							FloatingIPs: []*models.FloatingIP{
								{UUID: "test-fip"},
							},
						},
					},
				},
				fipsRes: &services.ListFloatingIPResponse{
					FloatingIPs: []*models.FloatingIP{
						{
							UUID: "test-fip",
						},
					},
				},
			},
			deleteData: &deleteData{
				vnReq:   &services.DeleteVirtualNetworkRequest{ID: "test-vn"},
				fippReq: &services.DeleteFloatingIPPoolRequest{ID: "test-fipp"},
				fipReq:  &services.DeleteFloatingIPRequest{ID: "test-fip"},
			},
		},
		{
			name:     "Delete network fails due to network in use",
			n:        &Network{},
			expected: &NetworkResponse{},
			id:       "test-vn",
			fail:     true,
			readData: &readData{
				vnRes: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{
						UUID:           "test-vn",
						RouterExternal: true,
						FloatingIPPools: []*models.FloatingIPPool{
							{UUID: "test-fipp"},
						},
					},
				},
				fippRes: &services.ListFloatingIPPoolResponse{
					FloatingIPPools: []*models.FloatingIPPool{
						{
							UUID: "test-fipp",
							FloatingIPs: []*models.FloatingIP{
								{UUID: "test-fip"},
							},
						},
					},
				},
				fipsRes: &services.ListFloatingIPResponse{
					FloatingIPs: []*models.FloatingIP{
						{
							UUID: "test-fip",
							VirtualMachineInterfaceRefs: []*models.FloatingIPVirtualMachineInterfaceRef{
								{UUID: "test-vmi"},
							},
						},
					},
				},
			},
			deleteData: &deleteData{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockServ := servicesmock.NewMockService(mockCtrl)
			ctx := context.Background()

			if tt.readData.vnRes != nil {
				mockServ.EXPECT().GetVirtualNetwork(gomock.Any(), gomock.Any()).Return(
					tt.readData.vnRes, nil,
				)
			}
			if tt.readData.fippRes != nil {
				mockServ.EXPECT().ListFloatingIPPool(gomock.Any(), gomock.Any()).Return(
					tt.readData.fippRes, nil,
				)
			}
			if tt.readData.fipsRes != nil {
				mockServ.EXPECT().ListFloatingIP(gomock.Any(), gomock.Any()).Return(
					tt.readData.fipsRes, nil,
				)
			}

			if tt.deleteData.vnReq != nil {
				mockServ.EXPECT().DeleteVirtualNetwork(gomock.Any(), tt.deleteData.vnReq)
			}
			if tt.deleteData.fippReq != nil {
				mockServ.EXPECT().DeleteFloatingIPPool(gomock.Any(), tt.deleteData.fippReq)
			}
			if tt.deleteData.fipReq != nil {
				mockServ.EXPECT().DeleteFloatingIP(gomock.Any(), tt.deleteData.fipReq)
			}

			rp := RequestParameters{
				ReadService:  mockServ,
				WriteService: mockServ,
			}

			res, err := tt.n.Delete(ctx, rp, tt.id)

			if tt.fail {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tt.expected, res)
			}
		})
	}
}
