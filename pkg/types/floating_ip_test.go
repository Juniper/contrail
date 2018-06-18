package types

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"

	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	ipammock "github.com/Juniper/contrail/pkg/types/ipam/mock"
	typesmock "github.com/Juniper/contrail/pkg/types/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

type addrMgrSubnetExhausted int

func (e addrMgrSubnetExhausted) SubnetExhausted() {
}

func (e addrMgrSubnetExhausted) Error() string {
	return ""
}

var mockCtrl *gomock.Controller
var ipamMock *ipammock.MockAddressManager
var dbServiceMock *servicesmock.MockService
var logicService ContrailTypeLogicService
var nextServiceMock *servicesmock.MockService
var dbMock *typesmock.MockDBInterface

func testSetup(t *testing.T) {
	mockCtrl = gomock.NewController(t)
	ipamMock = ipammock.NewMockAddressManager(mockCtrl)
	nextServiceMock = servicesmock.NewMockService(mockCtrl)
	dbServiceMock = servicesmock.NewMockService(mockCtrl)
	dbMock = typesmock.NewMockDBInterface(mockCtrl)
	logicService = ContrailTypeLogicService{
		BaseService:    services.BaseService{},
		AddressManager: ipamMock,
		DBService:      dbServiceMock,
		DB:             dbMock,
	}
	logicService.SetNext(nextServiceMock)

	setupDBMocks()
	setupIPAMMocks()
	setupNextServiceMocks()
}

func testClean() {
	mockCtrl.Finish()
}

func setupDBMocks() {
	dbMock.EXPECT().DB().AnyTimes()
	dbServiceMock.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetVirtualNetworkResponse{
			VirtualNetwork: &models.VirtualNetwork{},
		}, nil).AnyTimes()
}

func setupIPAMMocks() {
	ipamMock.EXPECT().AllocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(ctx context.Context, request *ipam.AllocateIPRequest) (address string, subnetUUID string, err error) {

			if request.SubnetUUID == "uuid-1" {
				return "10.0.0.1", "uuid-1", nil
			}

			exhaustedError := addrMgrSubnetExhausted(0)
			if request.SubnetUUID == "uuid-2" {
				return "", "", &exhaustedError
			}
			if request.SubnetUUID == "uuid-3" {
				return "", "", &exhaustedError
			}
			if request.SubnetUUID == "uuid-4" {
				return "", "", fmt.Errorf("Generic error")
			}

			return "10.0.0.1", "uuid-1", nil
		}).AnyTimes()

	ipamMock.EXPECT().IsIPAllocated(gomock.Not(gomock.Nil()),
		&ipam.IsIPAllocatedRequest{
			VirtualNetwork: &models.VirtualNetwork{},
			IPAddress:      "10.0.0.2",
		}).Return(true, nil).AnyTimes()
}

func setupNextServiceMocks() {
	// CreateFloatingIP - response
	nextServiceMock.EXPECT().CreateFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(ctx context.Context, request *services.CreateFloatingIPRequest) (response *services.CreateFloatingIPResponse, err error) {
			return &services.CreateFloatingIPResponse{
				FloatingIP: request.FloatingIP,
			}, nil
		}).AnyTimes()

	// DeleteFloatingIP - response
	nextServiceMock.EXPECT().DeleteFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(ctx context.Context, request *services.DeleteFloatingIPRequest) (response *services.DeleteFloatingIPResponse, err error) {
			return &services.DeleteFloatingIPResponse{
				ID: request.ID,
			}, nil
		}).AnyTimes()
}

func prepareParent(floatingIPPool *models.FloatingIPPool) {
	if floatingIPPool != nil {
		dbServiceMock.EXPECT().GetFloatingIPPool(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			&services.GetFloatingIPPoolResponse{
				FloatingIPPool: floatingIPPool,
			}, nil).AnyTimes()
	} else {
		dbServiceMock.EXPECT().GetFloatingIPPool(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
			nil, fmt.Errorf("DB error")).AnyTimes()
	}
}

func TestCreateFloatingIP(t *testing.T) {
	tests := []struct {
		name             string
		floatingIPParent *models.FloatingIPPool
		createRequest    services.CreateFloatingIPRequest
		expectedResponse services.CreateFloatingIPResponse
		fails            bool
		errorCode        codes.Code
	}{
		{
			name:             "Create floating ip when parent type is instance-ip",
			createRequest:    services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{ParentType: "instance-ip"}},
			floatingIPParent: &models.FloatingIPPool{},
			expectedResponse: services.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{ParentType: "instance-ip"}},
		},
		{
			name:             "Create floating ip with a free ip address",
			createRequest:    services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{ParentType: "floating-ip-pool"}},
			floatingIPParent: &models.FloatingIPPool{},
			expectedResponse: services.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.1",
			}},
		},
		{
			name: "Try to create floating ip with IP address which is already allocated",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.2",
			}},
			floatingIPParent: &models.FloatingIPPool{},
			fails:            true,
			errorCode:        codes.AlreadyExists,
		},
		{
			name: "Create floating ip without IP address",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			floatingIPParent: &models.FloatingIPPool{},
			expectedResponse: services.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.1",
			}},
		},
		{
			name: "Create floating ip with subnets from floating ip pool",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			floatingIPParent: &models.FloatingIPPool{
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"uuid-2", "uuid-1"},
				},
			},
			expectedResponse: services.CreateFloatingIPResponse{FloatingIP: &models.FloatingIP{
				ParentType:        "floating-ip-pool",
				FloatingIPAddress: "10.0.0.1",
			}},
		},
		{
			name: "Try to create floating ip with exhausted subnets from floating ip pool",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			floatingIPParent: &models.FloatingIPPool{
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"uuid-3", "uuid-2"},
				},
			},
			fails:     true,
			errorCode: codes.ResourceExhausted,
		},
		{
			name: "Try to create floating ip with subnets from floating ip pool with generic error",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			floatingIPParent: &models.FloatingIPPool{
				FloatingIPPoolSubnets: &models.FloatingIpPoolSubnetType{
					SubnetUUID: []string{"uuid-3", "uuid-4"},
				},
			},
			fails: true,
		},
		{
			name: "Try to create floating ip when parent can't be get from DB ",
			createRequest: services.CreateFloatingIPRequest{FloatingIP: &models.FloatingIP{
				ParentType: "floating-ip-pool",
			}},
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testSetup(t)
			defer testClean()
			prepareParent(tt.floatingIPParent)
			// Put an empty transaction into context so we could call DoInTransaction() without access to the real db
			ctx := context.WithValue(nil, db.Transaction, &sql.Tx{})
			createFloatingIPResponse, err := logicService.CreateFloatingIP(ctx, &tt.createRequest)

			if tt.fails {
				assert.Error(t, err)
				if tt.errorCode != codes.OK {
					status, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, tt.errorCode, status.Code())
				}
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, createFloatingIPResponse)
			assert.EqualValues(t, &tt.expectedResponse, createFloatingIPResponse)
		})
	}
}

func TestDeleteFloatingIP(t *testing.T) {

	tests := []struct {
		name             string
		floatingIPParent *models.FloatingIPPool
		floatingIP       *models.FloatingIP
		deleteRequest    *services.DeleteFloatingIPRequest
		expectedResponse *services.DeleteFloatingIPResponse
		fails            bool
		deallocatesIP    bool
		errorCode        codes.Code
	}{
		{
			name:             "Delete floating ip when parent type is instance-ip",
			deleteRequest:    &services.DeleteFloatingIPRequest{ID: "uuid-1"},
			floatingIPParent: &models.FloatingIPPool{},
			floatingIP:       &models.FloatingIP{UUID: "uuid-1", ParentType: "instance-ip"},
			expectedResponse: &services.DeleteFloatingIPResponse{ID: "uuid-1"},
			deallocatesIP:    false,
		},
		{
			name:             "Delete floating ip when parent type is floating-ip-pool",
			deleteRequest:    &services.DeleteFloatingIPRequest{ID: "uuid-1"},
			floatingIPParent: &models.FloatingIPPool{},
			floatingIP:       &models.FloatingIP{UUID: "uuid-1", ParentType: "floating-ip-pool"},
			expectedResponse: &services.DeleteFloatingIPResponse{ID: "uuid-1"},
			deallocatesIP:    true,
		},
		{
			name:          "Try to delete floating ip if it doesn't exist in DB",
			deleteRequest: &services.DeleteFloatingIPRequest{ID: "uuid-1"},
			deallocatesIP: false,
			fails:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testSetup(t)
			defer testClean()
			prepareParent(tt.floatingIPParent)

			if tt.floatingIP != nil {
				dbServiceMock.EXPECT().GetFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					&services.GetFloatingIPResponse{
						FloatingIP: tt.floatingIP,
					}, nil).AnyTimes()
			} else {
				dbServiceMock.EXPECT().GetFloatingIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
					nil, fmt.Errorf("Not found")).AnyTimes()
			}

			if tt.deallocatesIP {
				ipamMock.EXPECT().DeallocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(nil)
			}

			// Put an empty transaction into context so we could call DoInTransaction() without access to the real db
			ctx := context.WithValue(nil, db.Transaction, &sql.Tx{})
			deleteFloatingIPResponse, err := logicService.DeleteFloatingIP(ctx, tt.deleteRequest)

			if tt.fails {
				assert.Error(t, err)
				if tt.errorCode != codes.OK {
					status, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, tt.errorCode, status.Code())
				}
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, deleteFloatingIPResponse)
			assert.EqualValues(t, tt.expectedResponse, deleteFloatingIPResponse)
		})
	}
}
