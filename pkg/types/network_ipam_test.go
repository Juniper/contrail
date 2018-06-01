package types

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/Juniper/contrail/pkg/types/mock"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/serviceif/mock"
	"database/sql"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
	"github.com/stretchr/testify/assert"
)

func TestCreateNetworkIpam(t *testing.T) {
	var mockCtrl *gomock.Controller
	var dbMock *typesmock.MockDBServiceInterface
	var nextServiceMock *serviceifmock.MockService
	var logicService ContrailTypeLogicService
	var ctx context.Context

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		dbMock = typesmock.NewMockDBServiceInterface(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService: serviceif.BaseService{},
			DB:          dbMock,
		}
		nextServiceMock = serviceifmock.NewMockService(mockCtrl)
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
			func(ctx context.Context, request *models.CreateNetworkIpamRequest) (response *models.CreateNetworkIpamResponse, err error) {
				return &models.CreateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
			}).AnyTimes()

		createNetworkIpamRequest := &models.CreateNetworkIpamRequest{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "test"}}
		createNetworkIpamResponse, err := logicService.CreateNetworkIpam(ctx, createNetworkIpamRequest)

		assert.NoError(t, err)
		assert.NotNil(t, createNetworkIpamResponse)

	})

	t.Run("Create succeeds because of empty flat subnet", func(t *testing.T) {
		testSetup(t)
		defer testClean()

		nextServiceMock.EXPECT().CreateNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *models.CreateNetworkIpamRequest) (response *models.CreateNetworkIpamResponse, err error) {
				return &models.CreateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
			}).AnyTimes()

		createNetworkIpamRequest := &models.CreateNetworkIpamRequest{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "flat-subnet" , IpamSubnets: &models.IpamSubnets{} }}
		createNetworkIpamResponse, err := logicService.CreateNetworkIpam(ctx, createNetworkIpamRequest)

		assert.NoError(t, err)
		assert.NotNil(t, createNetworkIpamResponse)

	})

	t.Run("Create fails because of non-empty ipam_subnets list and not flat subnet", func(t *testing.T) {
		testSetup(t)
		defer testClean()

		nextServiceMock.EXPECT().CreateNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *models.CreateNetworkIpamRequest) (response *models.CreateNetworkIpamResponse, err error) {
				return &models.CreateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
			}).AnyTimes()

		createNetworkIpamRequest := &models.CreateNetworkIpamRequest{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "notFlat", IpamSubnets: &models.IpamSubnets{} }}
		createNetworkIpamResponse, err := logicService.CreateNetworkIpam(ctx, createNetworkIpamRequest)

		assert.Error(t, err)
		assert.Nil(t, createNetworkIpamResponse)

	})

	t.Run("Delete succeeds because of lack of ipam subnets ", func(t *testing.T) {
		testSetup(t)
		defer testClean()

		getNetworkIpamResponse := &models.GetNetworkIpamResponse{NetworkIpam: &models.NetworkIpam{IpamSubnetMethod: "notFlat"}}
		dbMock.EXPECT().GetNetworkIpam(gomock.Any(), gomock.Any()).Return(getNetworkIpamResponse, nil).AnyTimes()

		nextServiceMock.EXPECT().DeleteNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, request *models.DeleteNetworkIpamRequest) (response *models.DeleteNetworkIpamResponse, err error) {
				return &models.DeleteNetworkIpamResponse{ID: request.ID}, nil
			}).AnyTimes()

		deleteNetworkIpamRequest := &models.DeleteNetworkIpamRequest{ID: "forTestOnly"}
		deleteNetworkIpamResponse, err := logicService.DeleteNetworkIpam(ctx, deleteNetworkIpamRequest)

		assert.NoError(t, err)
		assert.NotNil(t, deleteNetworkIpamResponse)
	})

}
