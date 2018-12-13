package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "github.com/golang/mock/gomock"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services/mock"
	"context"
)

const (
	testDataPath = "./test_data/"

	createNetworkRequestPath = testDataPath + "create_network.json"
	readallNetworkRequestPath = testDataPath + "readall_network.json"
)
//
//func TestNetworkCreate(t *testing.T) {
//	runTest(t, func(t *testing.T, client *integration.HTTPAPIClient) {
//		response, err := client.NeutronPost(
//			context.Background(),
//			loadRequestFromJSONFile(t, createNetworkRequestPath),
//			[]int{200})
//		assert.NoError(t, err)
//		assertEqual(t, logic.NetworkResponse{Name: "ctest-vn-49391908"}, response)
//	})
//}

func TestNetworkReadAll(t *testing.T) {
	type readData struct {
		vnRes *services.ListVirtualNetworkResponse
	}

	tests := []struct {
		name string
		n *Network
		filters Filters
		fail bool
		readData *readData
		expRes Response
	}{
		{
			name: "read all vns",
			n: &Network{},
			filters: Filters{
				"shared": []string{"false"},
				"tenant_id": []string{"uuidp"},
			},
			readData:  &readData{
				vnRes: &services.ListVirtualNetworkResponse{
					VirtualNetworks: []*models.VirtualNetwork{
						{
							UUID: "uuid-vn",
							ParentType: "project",
							ParentUUID: "uuid-p",
							DisplayName: "vn-test",
							IDPerms: &models.IdPermsType{
								Enable: true,
								Created: "data1",
								LastModified: "data2",
							},
							NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
								{
									Attr: &models.VnSubnetsType{
										IpamSubnets: []*models.IpamSubnetType{
											{
												SubnetUUID: "uuid-sn",
											},
										},
									},
								},
							},
							PortSecurityEnabled: true,
							FQName: []string{"fqname", "vn", "test"},
						},
					},
				},
			},
			expRes: []*NetworkResponse{{
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
			}},
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
