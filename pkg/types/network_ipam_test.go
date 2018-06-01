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

func TestCreateNetworkIpam(t *testing.T) {
	var mockCtrl *gomock.Controller
	var dbMock *typesmock.MockDBServiceInterface
	var nextServiceMock *servicesmock.MockService
	var ipamMock *ipammock.MockAddressManager
	var logicService ContrailTypeLogicService
	var ctx context.Context

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		dbMock = typesmock.NewMockDBServiceInterface(mockCtrl)
		ipamMock = ipammock.NewMockAddressManager(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService:    services.BaseService{},
			AddressManager: ipamMock,
			DB:             dbMock,
		}
		nextServiceMock = servicesmock.NewMockService(mockCtrl)
		logicService.SetNext(nextServiceMock)

		// Put empty transaction into context so we could call DoInTransaction() without access to the real db
		emptyTx := sql.Tx{}
		ctx = context.WithValue(ctx, db.Transaction, &emptyTx)
		dbMock.EXPECT().DB().AnyTimes()
	}

	testClean := func() {
		mockCtrl.Finish()
	}

	t.Run("Create succeed because of empty ipam_subnets list", func(t *testing.T) {
		testSetup(t)
		defer testClean()

		nextServiceMock.EXPECT().CreateNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *services.CreateNetworkIpamRequest) (response *services.CreateNetworkIpamResponse, err error) {
				return &services.CreateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
			}).AnyTimes()

		createNetworkIpamRequest := &services.CreateNetworkIpamRequest{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "test"}}
		createNetworkIpamResponse, err := logicService.CreateNetworkIpam(ctx, createNetworkIpamRequest)

		assert.NoError(t, err)
		assert.NotNil(t, createNetworkIpamResponse)

	})

	t.Run("Create succeeds because of empty flat subnet", func(t *testing.T) {
		testSetup(t)
		defer testClean()

		nextServiceMock.EXPECT().CreateNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *services.CreateNetworkIpamRequest) (response *services.CreateNetworkIpamResponse, err error) {
				return &services.CreateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
			}).AnyTimes()

		createNetworkIpamRequest := &services.CreateNetworkIpamRequest{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "flat-subnet", IpamSubnets: &models.IpamSubnets{}}}
		createNetworkIpamResponse, err := logicService.CreateNetworkIpam(ctx, createNetworkIpamRequest)

		assert.NoError(t, err)
		assert.NotNil(t, createNetworkIpamResponse)

	})

	t.Run("Create fails because of non-empty ipam_subnets list and not flat subnet", func(t *testing.T) {
		testSetup(t)
		defer testClean()

		nextServiceMock.EXPECT().CreateNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *services.CreateNetworkIpamRequest) (response *services.CreateNetworkIpamResponse, err error) {
				return &services.CreateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
			}).AnyTimes()

		createNetworkIpamRequest := &services.CreateNetworkIpamRequest{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "notFlat", IpamSubnets: &models.IpamSubnets{}}}
		createNetworkIpamResponse, err := logicService.CreateNetworkIpam(ctx, createNetworkIpamRequest)

		assert.Error(t, err)
		assert.Nil(t, createNetworkIpamResponse)

	})

	t.Run("Create network ipam with specific ipam_subnets", func(t *testing.T) {
		testSetup(t)
		defer testClean()

		nextServiceMock.EXPECT().CreateNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *services.CreateNetworkIpamRequest) (response *services.CreateNetworkIpamResponse, err error) {
				return &services.CreateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
			}).AnyTimes()

		ipamMock.EXPECT().CreateIpamSubnet(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *ipam.CreateIpamSubnetRequest) (subnetUUID string, err error) {
				return "uuuu-uuuu-iiii-dddd", nil
			}).AnyTimes()

		subnet := &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}
		ipamSubnet := models.IpamSubnetType{Subnet: subnet}
		createNetworkIpamRequest := &services.CreateNetworkIpamRequest{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "flat-subnet", IpamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{&ipamSubnet}}}}
		createNetworkIpamResponse, err := logicService.CreateNetworkIpam(ctx, createNetworkIpamRequest)

		receivedUUID := createNetworkIpamResponse.NetworkIpam.IpamSubnets.Subnets[0].SubnetUUID
		assert.Equal(t, "uuuu-uuuu-iiii-dddd", receivedUUID)
		assert.NoError(t, err)
		assert.NotNil(t, createNetworkIpamResponse)

	})
}

func TestDeleteNetworkIpam(t *testing.T) {
	var mockCtrl *gomock.Controller
	var dbMock *typesmock.MockDBServiceInterface
	var nextServiceMock *servicesmock.MockService
	var ipamMock *ipammock.MockAddressManager
	var logicService ContrailTypeLogicService
	var ctx context.Context

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		dbMock = typesmock.NewMockDBServiceInterface(mockCtrl)
		ipamMock = ipammock.NewMockAddressManager(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService:    services.BaseService{},
			AddressManager: ipamMock,
			DB:             dbMock,
		}
		nextServiceMock = servicesmock.NewMockService(mockCtrl)
		logicService.SetNext(nextServiceMock)

		// Put empty transaction into context so we could call DoInTransaction() without access to the real db
		emptyTx := sql.Tx{}
		ctx = context.WithValue(ctx, db.Transaction, &emptyTx)
		dbMock.EXPECT().DB().AnyTimes()
	}

	testClean := func() {
		mockCtrl.Finish()
	}

	t.Run("Delete succeeds because of lack of ipam subnets ", func(t *testing.T) {
		testSetup(t)
		defer testClean()

		getNetworkIpamResponse := &services.GetNetworkIpamResponse{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "notFlat"}}
		dbMock.EXPECT().GetNetworkIpam(gomock.Any(), gomock.Any()).Return(getNetworkIpamResponse, nil).AnyTimes()

		nextServiceMock.EXPECT().DeleteNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *services.DeleteNetworkIpamRequest) (response *services.DeleteNetworkIpamResponse, err error) {
				return &services.DeleteNetworkIpamResponse{ID: request.ID}, nil
			}).AnyTimes()

		deleteNetworkIpamRequest := &services.DeleteNetworkIpamRequest{ID: "forTestOnly"}
		deleteNetworkIpamResponse, err := logicService.DeleteNetworkIpam(ctx, deleteNetworkIpamRequest)

		assert.NoError(t, err)
		assert.NotNil(t, deleteNetworkIpamResponse)
	})

	t.Run("Delete succeeds for non-empty ipam network ", func(t *testing.T) {
		testSetup(t)
		defer testClean()

		getNetworkIpamResponse := &services.GetNetworkIpamResponse{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "notFlat"}}
		dbMock.EXPECT().GetNetworkIpam(gomock.Any(), gomock.Any()).Return(getNetworkIpamResponse, nil).AnyTimes()

		nextServiceMock.EXPECT().DeleteNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *services.DeleteNetworkIpamRequest) (response *services.DeleteNetworkIpamResponse, err error) {
				return &services.DeleteNetworkIpamResponse{ID: request.ID}, nil
			}).AnyTimes()

		ipamMock.EXPECT().DeleteIpamSubnet(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

		subnet := &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}
		ipamSubnet := models.IpamSubnetType{Subnet: subnet}
		networkIpam := &models.NetworkIpam{UUID: "uuid", IpamSubnetMethod: "flat-subnet", IpamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{&ipamSubnet}}}
		deleteNetworkIpamRequest := &services.DeleteNetworkIpamRequest{ID: networkIpam.GetUUID()}
		deleteNetworkIpamResponse, err := logicService.DeleteNetworkIpam(ctx, deleteNetworkIpamRequest)

		assert.NoError(t, err)
		assert.NotNil(t, deleteNetworkIpamResponse)
	})
}
