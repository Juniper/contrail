package types

import (
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/testutils"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"golang.org/x/net/context"
	"testing"
	"time"
)

//structure to pass vn parameters during VrirualNetwork object creation
type testVn struct {
	mPolicyServChainEn bool
	impRtStr           string
	expRtStr           string
	virtNetID          int64
	globalAccess       int64
	isShared           bool
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
	vn.MultiPolicyServiceChainsEnabled = testVnData.mPolicyServChainEn
	vn.ImportRouteTargetList = &models.RouteTargetList{RouteTarget: []string{testVnData.impRtStr}}
	vn.ExportRouteTargetList = &models.RouteTargetList{RouteTarget: []string{testVnData.expRtStr}}
	vn.VirtualNetworkNetworkID = testVnData.virtNetID
	vn.UUID = "test_vn_uuid"

	return vn
}

func TestIsValidMultiPolicyServiceChainConfig(t *testing.T) {
	var tests = []struct {
		testVnData *testVn
		result     bool
	}{
		{&testVn{mPolicyServChainEn: true,
			impRtStr: "100:101",
			expRtStr: "100:102"}, true},
		{&testVn{mPolicyServChainEn: true,
			impRtStr: "100:101",
			expRtStr: "100:101"}, false},
		{&testVn{mPolicyServChainEn: false}, false},
	}

	//Test for various values of route target
	for _, test := range tests[:2] {
		vn := createTestVn(test.testVnData)
		if res := vn.IsValidMultiPolicyServiceChainConfig(); res != test.result {
			t.Errorf("IsValidMultiPolicyServiceChainConfig check for rt incorrect")
		}
	}

	//Test for multiPolicyServiceChainsEnabled
	vn := createTestVn(tests[2].testVnData)
	if !vn.IsValidMultiPolicyServiceChainConfig() {
		t.Errorf("IsValidMultiPolicyServiceChainConfig check for MultiPolicyServiceChainsEnabled incorrect")
	}
}

func TestCreateVirtualNetwork(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Create service chain, add DB service to the end
	service := getService()

	testVnData := &testVn{
		mPolicyServChainEn: true,
		impRtStr:           "100:101",
		expRtStr:           "100:101",
	}

	//Check multiple policy service chain supported (negative)
	vn := createTestVn(testVnData)
	if _, err := service.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{VirtualNetwork: vn}); err == nil {
		t.Errorf("IsValidMultiPolicyServiceChainConfig check for rt incorrect")
	}

	//Check multiple policy service chain supported (negative)
	testVnData.expRtStr = "100:102"
	testVnData.virtNetID = 9999
	vn1 := createTestVn(testVnData)
	if _, err := service.CreateVirtualNetwork(ctx, &models.CreateVirtualNetworkRequest{VirtualNetwork: vn1}); err == nil {
		t.Errorf("HasVirtualNetworkNetworkID check incorrect")
	}

	//TODO Remaining tests
}

func TestDeleteVirtualNetwork(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Create service chain, add DB service to the end
	service := getService()

	//Check missing VirtualNetwork in DB (negative)
	if _, err := service.DeleteVirtualNetwork(ctx, &models.DeleteVirtualNetworkRequest{ID: " "}); err == nil {
		t.Errorf("MissigetVirtualNetworkID check for rt incorrect")
	}

	//Check DeleteVirtualNetwork (positive)
	vn := createTestVn(&testVn{})
	db.DoInTransaction(ctx, service.DB.DB(), func(ctx context.Context) error {
		service.DB.CreateIntPool(ctx, &ipam.IntPool{Key: VirtualNetworkIDPoolKey, Start: 0, End: 2})
		vn.VirtualNetworkNetworkID, _ = service.DB.AllocateInt(ctx, VirtualNetworkIDPoolKey)
		return nil
	})
	vnReq := &models.CreateVirtualNetworkRequest{VirtualNetwork: vn}
	service.DB.CreateVirtualNetwork(ctx, vnReq)

	if _, err := service.DeleteVirtualNetwork(ctx, &models.DeleteVirtualNetworkRequest{ID: vn.UUID}); err != nil {
		t.Errorf("DeleteVirtualNetwork Failed %v", err)
	}
}
