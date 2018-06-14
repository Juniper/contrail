package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/unittest"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

//Structure testVn is used to pass vn parameters during VrirualNetwork object creation
type testVn struct {
	MultiPolicyServiceChainsEnabled bool
	ImportRouteTargetList           string
	ExportRouteTargetList           string
	VirtualNetworkNetworkID         int64
}

func getService() *ContrailTypeLogicService {
	service := &ContrailTypeLogicService{
		DB:               unittest.TestDbService,
		IntPoolAllocator: unittest.TestDbService,
	}

	services.Chain(service, unittest.TestDbService)
	return service
}
func TestMain(m *testing.M) {
	unittest.CreateTestDbService(m)
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
			res, err := service.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{VirtualNetwork: vn})
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
	_, err := service.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: "nonexistent_uuid"})
	assert.Errorf(t, err, "MissigetVirtualNetworkID check for rt incorrect")

	//Check DeleteVirtualNetwork (positive)
	vn := createTestVn(&testVn{})
	intPool := ipam.IntPool{Key: VirtualNetworkIDPoolKey, Start: 0, End: 2}
	err = db.DoInTransaction(ctx, service.DB.DB(), func(ctx context.Context) error {
		err = unittest.TestDbService.CreateIntPool(ctx, &intPool)
		assert.NoError(t, err)
		vn.VirtualNetworkNetworkID, err = service.IntPoolAllocator.AllocateInt(ctx, VirtualNetworkIDPoolKey)
		assert.NoError(t, err)
		return nil
	})
	assert.NoError(t, err)
	vnReq := &services.CreateVirtualNetworkRequest{VirtualNetwork: vn}
	service.DB.CreateVirtualNetwork(ctx, vnReq) // nolint: errcheck
	_, err = service.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: vn.UUID})
	assert.NoErrorf(t, err, "DeleteVirtualNetwork Failed %v", err)
	err = db.DoInTransaction(ctx, service.DB.DB(), func(ctx context.Context) error {
		err = service.IntPoolAllocator.DeleteIntPools(ctx, &intPool)
		assert.NoError(t, err)
		return nil
	})
	assert.NoError(t, err)
}
