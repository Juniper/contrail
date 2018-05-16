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
	VirtualNetworkNetworkID         int64
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
	var tests = []struct {
		name       string
		testVnData *testVn
		fails      bool
	}{
		{
			"check for rt",
			&testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:101",
			},
			true,
		},
		{
			"check for virtual network id",
			&testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:102",
				VirtualNetworkNetworkID:         9999,
			},
			true,
		},
		{
			"check for virtual network id",
			&testVn{
				MultiPolicyServiceChainsEnabled: true,
				ImportRouteTargetList:           "100:101",
				ExportRouteTargetList:           "100:102",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(t, mockCtrl)

			virtualNetworkSetupDBMocks(service)
			virtualNetworkSetupIntPoolAllocatorMocks(service)
			virtualNetworkSetupNextServiceMocks(service)

			vn := createTestVn(tt.testVnData)
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
				assert.NotNil(t, res.GetVirtualNetwork())
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
