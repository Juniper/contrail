package types

import (
	"testing"

	"database/sql"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"github.com/Juniper/contrail/pkg/types/ipam/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

type testNetIpamParams struct {
	uuid             string
	ipamSubnetMethod string
	ipamSubnets      *models.IpamSubnets
}

func createTestNetworkIpam(testParams *testNetIpamParams) *models.NetworkIpam {
	networkIpam := models.MakeNetworkIpam()
	networkIpam.UUID = testParams.uuid
	networkIpam.IpamSubnetMethod = testParams.ipamSubnetMethod
	networkIpam.IpamSubnets = testParams.ipamSubnets
	return networkIpam
}

func networkIpamNextServMocks(service *ContrailTypeLogicService) {
	ipamMock := service.AddressManager.(*ipammock.MockAddressManager)
	ipamMock.EXPECT().CreateIpamSubnet(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, request *ipam.CreateIpamSubnetRequest) (subnetUUID string, err error) {
			return "uuuu-uuuu-iiii-dddd", nil
		}).AnyTimes()
	ipamMock.EXPECT().DeleteIpamSubnet(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

}

func networkIpamDBMocks(service *ContrailTypeLogicService) {
	dbService := service.DB.(*typesmock.MockDBServiceInterface)
	dbService.EXPECT().DB().AnyTimes()
}

func networkIpamIPAMMocks(service *ContrailTypeLogicService) {
	nextServiceMock := service.Next().(*servicesmock.MockService)
	nextServiceMock.EXPECT().CreateNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, request *services.CreateNetworkIpamRequest) (response *services.CreateNetworkIpamResponse, err error) {
			return &services.CreateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
		}).AnyTimes()
	nextServiceMock.EXPECT().DeleteNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, request *services.DeleteNetworkIpamRequest) (response *services.DeleteNetworkIpamResponse, err error) {
			return &services.DeleteNetworkIpamResponse{ID: request.ID}, nil
		}).AnyTimes()

}

func TestCreateNetworkIpam(t *testing.T) {
	tests := []struct {
		name              string
		testNetIpamParams *testNetIpamParams
		fails             bool
		changedUUID       bool
	}{
		{
			name: "Create succeed because of empty ipam_subnets list",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "notFlat",
			},
			fails:       false,
			changedUUID: false,
		},
		{
			name: "Create succeeds because of empty flat subnet",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets:      &models.IpamSubnets{},
			},
			fails:       false,
			changedUUID: false,
		},
		{
			name: "Create fails because of non-empty ipam_subnets list and not flat subnet",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "notFlat",
				ipamSubnets:      &models.IpamSubnets{},
			},
			fails:       true,
			changedUUID: false,
		},
		{
			name: "Create network ipam with specific ipam_subnets",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets:      &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
			},
			fails:       false,
			changedUUID: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(t, mockCtrl)
			networkIpamDBMocks(service)
			networkIpamIPAMMocks(service)
			networkIpamNextServMocks(service)

			emptyTx := sql.Tx{}
			ctx := context.WithValue(nil, db.Transaction, &emptyTx)

			networkIpam := createTestNetworkIpam(tt.testNetIpamParams)
			createNetworkIpamRequest := &services.CreateNetworkIpamRequest{NetworkIpam: networkIpam}
			createNetworkIpamResponse, err := service.CreateNetworkIpam(ctx, createNetworkIpamRequest)

			if tt.fails {
				assert.Error(t, err)
				assert.Nil(t, createNetworkIpamResponse)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, createNetworkIpamResponse)

			if tt.changedUUID {
				receivedUUID := createNetworkIpamResponse.NetworkIpam.IpamSubnets.Subnets[0].SubnetUUID
				assert.Equal(t, "uuuu-uuuu-iiii-dddd", receivedUUID)
			}
			mockCtrl.Finish()
		})
	}

}

func TestDeleteNetworkIpam(t *testing.T) {
	deleteIpamDBMock := func(service *ContrailTypeLogicService, getNetworkIpamResponse *services.GetNetworkIpamResponse) {
		dbService := service.DB.(*typesmock.MockDBServiceInterface)
		dbService.EXPECT().GetNetworkIpam(gomock.Any(), gomock.Any()).Return(getNetworkIpamResponse, nil)
	}

	tests := []struct {
		name              string
		testNetIpamParams *testNetIpamParams
	}{
		{
			name: "Delete succeeds because of lack of ipam subnets",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "notFlat",
			},
		},
		{
			name: "Delete succeeds for non-empty ipam network",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets:      &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(t, mockCtrl)
			networkIpamDBMocks(service)
			networkIpamIPAMMocks(service)
			networkIpamNextServMocks(service)

			emptyTx := sql.Tx{}
			ctx := context.WithValue(nil, db.Transaction, &emptyTx)

			networkIpam := createTestNetworkIpam(tt.testNetIpamParams)
			getNetworkIpamResponse := &services.GetNetworkIpamResponse{NetworkIpam: networkIpam}
			deleteIpamDBMock(service, getNetworkIpamResponse)
			deleteNetworkIpamRequest := &services.DeleteNetworkIpamRequest{ID: getNetworkIpamResponse.GetNetworkIpam().GetUUID()}
			deleteNetworkIpamResponse, err := service.DeleteNetworkIpam(ctx, deleteNetworkIpamRequest)

			assert.NoError(t, err)
			assert.NotNil(t, deleteNetworkIpamResponse)
			mockCtrl.Finish()
		})
	}
}
