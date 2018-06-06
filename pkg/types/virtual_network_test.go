package types

import (
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/testutils"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
	"time"
)

//structure testVn to pass vn parameters during VrirualNetwork object creation
type testVn struct {
	MultiPolicyServiceChainsEnabled bool
	ImportRouteTargetList           string
	ExportRouteTargetList           string
	VirtualNetworkNetworkID         int64
	globalAccess                    int64
	isShared                        bool
}

func getService() *ContrailTypeLogicService {
	var serviceChain []serviceif.Service
	service := &ContrailTypeLogicService{
		BaseService: serviceif.BaseService{},
		DB:          testutils.TestDbService,
	}
	serviceChain = append(serviceChain, service)
	serviceChain = append(serviceChain, testutils.TestDbService)

	serviceif.Chain(serviceChain)
	return service
}
func TestMain(m *testing.M) {
	testutils.CreateTestDbService(m)
}

func createTestVn(testVnData *testVn) *models.VirtualNetwork {
	vn := models.MakeVirtualNetwork()
	vn.MultiPolicyServiceChainsEnabled = testVnData.MultiPolicyServiceChainsEnabled
	vn.ImportRouteTargetList = &models.RouteTargetList{RouteTarget: []string{testVnData.ImportRouteTargetList}}
	vn.ExportRouteTargetList = &models.RouteTargetList{RouteTarget: []string{testVnData.ExportRouteTargetList}}
	vn.VirtualNetworkNetworkID = testVnData.VirtualNetworkNetworkID
	vn.UUID = "test_vn_uuid"

	return vn
}

func TestIsValidMultiPolicyServiceChainConfig(t *testing.T) {
	var tests = []struct {
		name       string
		testVnData *testVn
		fails      bool
	}{
		{"check for rt",
			&testVn{MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList: "100:101",
				ExportRouteTargetList: "100:102"}, false},
		{"check for rt",
			&testVn{MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList: "100:101",
				ExportRouteTargetList: "100:101"}, true},
		{"check for MultiPolicyServiceChainsEnabled",
			&testVn{MultiPolicyServiceChainsEnabled: false}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vn := createTestVn(tt.testVnData)

			res := vn.IsValidMultiPolicyServiceChainConfig()

			if tt.fails {
				assert.Falsef(t, res, tt.name)
			} else {
				assert.Truef(t, res, tt.name)
			}
		})
	}
}

func TestCreateVirtualNetwork(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Create service chain, add DB service to the end
	service := getService()

	var tests = []struct {
		name       string
		testVnData *testVn
		fails      bool
	}{
		{"check for rt",
			&testVn{MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList: "100:101",
				ExportRouteTargetList: "100:101"}, true},
		{"check for virtual network id",
			&testVn{MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:   "100:101",
				ExportRouteTargetList:   "100:102",
				VirtualNetworkNetworkID: 9999}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vn := createTestVn(tt.testVnData)
			res, err := service.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{VirtualNetwork: vn})
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, res, tt.name)
			}
		})
	}

	//TODO Remaining tests
}

func TestDeleteVirtualNetwork(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Create service chain, add DB service to the end
	service := getService()

	//Check missing VirtualNetwork in DB (negative)
	_, err := service.DeleteVirtualNetwork(ctx, &models.DeleteVirtualNetworkRequest{ID: " "})
	assert.Errorf(t, err, "MissigetVirtualNetworkID check for rt incorrect")

	//Check DeleteVirtualNetwork (positive)
	vn := createTestVn(&testVn{})
	db.DoInTransaction(ctx, service.DB.DB(), func(ctx context.Context) error {
		service.DB.CreateIntPool(ctx, &ipam.IntPool{Key: VirtualNetworkIDPoolKey, Start: 0, End: 2})
		vn.VirtualNetworkNetworkID, _ = service.DB.AllocateInt(ctx, VirtualNetworkIDPoolKey)
		return nil
	})
	vnReq := &models.CreateVirtualNetworkRequest{VirtualNetwork: vn}
	service.DB.CreateVirtualNetwork(ctx, vnReq)

	_, err = service.DeleteVirtualNetwork(ctx, &models.DeleteVirtualNetworkRequest{ID: vn.UUID})
	assert.NoErrorf(t, err, "DeleteVirtualNetwork Failed %v", err)
}
