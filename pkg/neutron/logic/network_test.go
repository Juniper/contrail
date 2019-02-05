package logic

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	"github.com/gogo/protobuf/types"
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

func TestNetworkUpdate(t *testing.T) {
	type readData struct {
		vnRes   *services.GetVirtualNetworkResponse
		vmiRes  *services.GetVirtualMachineInterfaceResponse
		fipsRes *services.ListFloatingIPResponse
	}

	type deleteData struct {
		fippReq *services.DeleteFloatingIPPoolRequest
		fipReq  *services.DeleteFloatingIPRequest
	}

	type writeData struct {
		fippRes *services.CreateFloatingIPPoolResponse
	}

	tests := []struct {
		name       string
		n          *Network
		fm         types.FieldMask
		expected   *NetworkResponse
		fail       bool
		id         string
		readData   *readData
		deleteData *deleteData
		writeData  *writeData
	}{
		{
			name: "update with share change allowed",
			n: &Network{
				Shared: false,
			},
			fm: types.FieldMask{
				Paths: []string{
					sharedKey,
				},
			},
			expected: &NetworkResponse{
				Status:              netStatusActive,
				RouterExternal:      false,
				FQName:              []string{"test-project", "test-vn"},
				Name:                "test-vn",
				AdminStateUp:        true,
				TenantID:            "testproject",
				ProjectID:           "testproject",
				CreatedAt:           "data1",
				UpdatedAt:           "data2",
				PortSecurityEnabled: true,
				Shared:              false,
				ID:                  "test-vn",
				Subnets:             []string{},
				SubnetIpam:          []*SubnetIpam{},
			},
			fail: false,
			id:   "test-vn",
			readData: &readData{
				vnRes: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{
						UUID:        "test-vn",
						ParentType:  "project",
						FQName:      []string{"test-project", "test-vn"},
						ParentUUID:  "test-project",
						IsShared:    true,
						DisplayName: "test-vn",
						IDPerms: &models.IdPermsType{
							Enable:       true,
							Created:      "data1",
							LastModified: "data2",
						},
						Perms2: models.MakePermType2(),
						VirtualMachineInterfaceBackRefs: []*models.VirtualMachineInterface{
							{UUID: "test-vmi"},
						},
						PortSecurityEnabled: true,
					},
				},
				vmiRes: &services.GetVirtualMachineInterfaceResponse{
					VirtualMachineInterface: &models.VirtualMachineInterface{
						UUID:       "test-vmi",
						ParentType: models.KindProject,
						ParentUUID: "test-project",
					},
				},
			},
			deleteData: &deleteData{},
			writeData:  &writeData{},
		},
		{
			name: "update with share change not allowed",
			n: &Network{
				Shared: false,
			},
			fm: types.FieldMask{
				Paths: []string{
					sharedKey,
				},
			},
			fail: true,
			id:   "test-vn",
			readData: &readData{
				vnRes: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{
						UUID:        "test-vn",
						ParentType:  "project",
						FQName:      []string{"test-project", "test-vn"},
						ParentUUID:  "test-project",
						IsShared:    true,
						DisplayName: "test-vn",
						IDPerms: &models.IdPermsType{
							Enable:       true,
							Created:      "data1",
							LastModified: "data2",
						},
						Perms2: models.MakePermType2(),
						VirtualMachineInterfaceBackRefs: []*models.VirtualMachineInterface{
							{UUID: "test-vmi"},
						},
						PortSecurityEnabled: true,
					},
				},
				vmiRes: &services.GetVirtualMachineInterfaceResponse{
					VirtualMachineInterface: &models.VirtualMachineInterface{
						UUID:       "test-vmi",
						ParentType: models.KindProject,
						ParentUUID: "different-project",
					},
				},
			},
			deleteData: &deleteData{},
			writeData:  &writeData{},
		},
		{
			name: "update network succeeds and cause floating ip and pools deletion",
			n: &Network{
				RouterExternal: false,
			},
			fm: types.FieldMask{
				Paths: []string{
					routerExternalKey,
				},
			},
			expected: &NetworkResponse{
				Status:              netStatusActive,
				RouterExternal:      false,
				FQName:              []string{"test-project", "test-vn"},
				Name:                "test-vn",
				AdminStateUp:        true,
				TenantID:            "testproject",
				ProjectID:           "testproject",
				CreatedAt:           "data1",
				UpdatedAt:           "data2",
				PortSecurityEnabled: true,
				Shared:              false,
				ID:                  "test-vn",
				Subnets:             []string{},
				SubnetIpam:          []*SubnetIpam{},
			},
			fail: false,
			id:   "test-vn",
			readData: &readData{
				vnRes: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{
						UUID:        "test-vn",
						ParentType:  "project",
						FQName:      []string{"test-project", "test-vn"},
						ParentUUID:  "test-project",
						DisplayName: "test-vn",
						IDPerms: &models.IdPermsType{
							Enable:       true,
							Created:      "data1",
							LastModified: "data2",
						},
						Perms2:              models.MakePermType2(),
						PortSecurityEnabled: true,
						RouterExternal:      true,
						FloatingIPPools: []*models.FloatingIPPool{
							{UUID: "test-fipp"},
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
				fippReq: &services.DeleteFloatingIPPoolRequest{ID: "test-fipp"},
				fipReq:  &services.DeleteFloatingIPRequest{ID: "test-fip"},
			},
			writeData: &writeData{},
		},
		{
			name: "update network succeed and cause floating ip pool creation",
			n: &Network{
				RouterExternal: true,
			},
			fm: types.FieldMask{
				Paths: []string{
					routerExternalKey,
				},
			},
			expected: &NetworkResponse{
				Status:              netStatusActive,
				RouterExternal:      true,
				FQName:              []string{"test-project", "test-vn"},
				Name:                "test-vn",
				AdminStateUp:        true,
				TenantID:            "testproject",
				ProjectID:           "testproject",
				CreatedAt:           "data1",
				UpdatedAt:           "data2",
				PortSecurityEnabled: true,
				Shared:              false,
				ID:                  "test-vn",
				Subnets:             []string{},
				SubnetIpam:          []*SubnetIpam{},
			},
			fail: false,
			id:   "test-vn",
			readData: &readData{
				vnRes: &services.GetVirtualNetworkResponse{
					VirtualNetwork: &models.VirtualNetwork{
						UUID:        "test-vn",
						ParentType:  "project",
						FQName:      []string{"test-project", "test-vn"},
						ParentUUID:  "test-project",
						DisplayName: "test-vn",
						IDPerms: &models.IdPermsType{
							Enable:       true,
							Created:      "data1",
							LastModified: "data2",
						},
						Perms2:              models.MakePermType2(),
						PortSecurityEnabled: true,
						RouterExternal:      false,
					},
				},
			},
			writeData: &writeData{
				fippRes: &services.CreateFloatingIPPoolResponse{
					FloatingIPPool: &models.FloatingIPPool{
						UUID: "test-fipp",
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
			if tt.readData.vmiRes != nil {
				mockServ.EXPECT().GetVirtualMachineInterface(gomock.Any(), gomock.Any()).Return(
					tt.readData.vmiRes, nil,
				)
			}
			if tt.readData.fipsRes != nil {
				mockServ.EXPECT().ListFloatingIP(gomock.Any(), gomock.Any()).Return(
					tt.readData.fipsRes, nil,
				)
			}

			if tt.deleteData.fippReq != nil {
				mockServ.EXPECT().DeleteFloatingIPPool(gomock.Any(), tt.deleteData.fippReq)
			}
			if tt.deleteData.fipReq != nil {
				mockServ.EXPECT().DeleteFloatingIP(gomock.Any(), tt.deleteData.fipReq)
			}

			if tt.writeData.fippRes != nil {
				mockServ.EXPECT().CreateFloatingIPPool(gomock.Any(), gomock.Any()).Return(
					tt.writeData.fippRes, nil,
				)
			}

			if !tt.fail {
				mockServ.EXPECT().UpdateVirtualNetwork(gomock.Any(), gomock.Any())
			}

			rp := RequestParameters{
				ReadService:  mockServ,
				WriteService: mockServ,
				FieldMask:    tt.fm,
			}

			res, err := tt.n.Update(ctx, rp, tt.id)

			if tt.fail {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tt.expected, res)
				assert.NoError(t, err)
			}
		})
	}
}
