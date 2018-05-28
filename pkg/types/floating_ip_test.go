package types

import (
	"database/sql"
	"testing"

	"github.com/Juniper/contrail/pkg/types/ipam"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/serviceif/mock"
	"github.com/Juniper/contrail/pkg/types/ipam/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

var mockCtrl *gomock.Controller
var ipamMock *ipammock.MockAddressManager
var dbServiceMock *typesmock.MockDBServiceInterface
var logicService ContrailTypeLogicService
var nextServiceMock *serviceifmock.MockService
var ctx context.Context

func testSetup(t *testing.T) {
	mockCtrl = gomock.NewController(t)
	ipamMock = ipammock.NewMockAddressManager(mockCtrl)
	nextServiceMock = serviceifmock.NewMockService(mockCtrl)
	dbServiceMock = typesmock.NewMockDBServiceInterface(mockCtrl)
	logicService = ContrailTypeLogicService{
		BaseService:    serviceif.BaseService{},
		AddressManager: ipamMock,
		DB:             dbServiceMock,
	}
	logicService.SetNext(nextServiceMock)

	// Put empty transaction into context so we could call DoInTransaction() without access to the real db
	emptyTx := sql.Tx{}
	ctx = context.WithValue(ctx, db.Transaction, &emptyTx)
	dbServiceMock.EXPECT().DB().AnyTimes()
	setupDBMocks()
	setupIPAMMocks()
	setupNextServiceMocks()
}

func testClean() {
	mockCtrl.Finish()
}

func setupDBMocks() {
	dbServiceMock.EXPECT().DB().AnyTimes()
	getFloatingIPPoolResponse := &models.GetFloatingIPPoolResponse{FloatingIPPool: &models.FloatingIPPool{}}
	dbServiceMock.EXPECT().GetFloatingIPPool(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(getFloatingIPPoolResponse, nil).AnyTimes()
	getVirtualNetworkResponse := &models.GetVirtualNetworkResponse{VirtualNetwork: &models.VirtualNetwork{}}
	dbServiceMock.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(getVirtualNetworkResponse, nil).AnyTimes()
}

func setupIPAMMocks() {
	ipamMock.EXPECT().AllocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return("10.0.0.1", "uuuu-uuuu-iiii-dddd", nil).AnyTimes()
	ipamMock.EXPECT().IsIPAllocated(gomock.Any(), &ipam.IsIPAllocatedRequest{VirtualNetwork: &models.VirtualNetwork{}, IPAddress: "10.0.0.2"}).Return(true, nil).AnyTimes()
}

func setupNextServiceMocks() {
	nextServiceMock.EXPECT().CreateFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(ctx context.Context, request *models.CreateFloatingIPRequest) (response *models.CreateFloatingIPResponse, err error) {
			return &models.CreateFloatingIPResponse{FloatingIP: request.FloatingIP}, nil
		}).AnyTimes()
	deleteFloatingIPResponse := &models.DeleteFloatingIPResponse{ID: "uuuu-uuuu-iiii-dddd"}

	nextServiceMock.EXPECT().DeleteFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(deleteFloatingIPResponse, nil).AnyTimes()
}

func TestCreateFloatingIP(t *testing.T) {
	tests := []struct {
		name             string
		createRequest    models.CreateFloatingIPRequest
		expectedResponse models.CreateFloatingIPResponse
		fails            bool
	}{
		{
			name:             "Create floating ip when parent type is instance-ip",
			createRequest:    models.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{ParentType: "instance-ip"}},
			expectedResponse: models.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{ParentType: "instance-ip"}},
		},
		{
			name:          "Create floating ip when parent type is floating-ip-pool",
			createRequest: models.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{ParentType: "floating-ip-pool"}},
			expectedResponse: models.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.1",
			}},
		},
		{
			name: "Create floating ip when parent type is floating-ip-pool",
			createRequest: models.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.2",
			}},
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testSetup(t)
			defer testClean()

			createFloatingIPResponse, err := logicService.CreateFloatingIP(ctx, &tt.createRequest)

			if tt.fails {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, createFloatingIPResponse)
			assert.EqualValues(t, createFloatingIPResponse, &tt.expectedResponse)
		})
	}
}

func TestDeleteFloatingIP(t *testing.T) {
	t.Run("Delete floating ip when parent type is instance-ip", func(t *testing.T) {
		testSetup(t)
		getFloatingIPResponse := &models.GetFloatingIPResponse{FloatingIP: &models.FloatingIP{ParentType: "instance-ip"}}
		dbServiceMock.EXPECT().GetFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(getFloatingIPResponse, nil).AnyTimes()

		deleteFloatingIPRequest := &models.DeleteFloatingIPRequest{ID: "uuuu-uuuu-iiii-dddd"}
		deleteFloatingIPResponse, err := logicService.DeleteFloatingIP(ctx, deleteFloatingIPRequest)
		assert.NoError(t, err)
		assert.NotNil(t, deleteFloatingIPResponse)
		testClean()
	})
}

//TODO: implement tests
