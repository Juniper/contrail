package types

import (
	"testing"

	"math"

	"database/sql"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/types/mocks"
	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestGlobalSystemConfigUdc(t *testing.T) {
	var mockCtrl *gomock.Controller
	var dbMock *mock_types.MockDBServiceInterface
	var logicService ContrailTypeLogicService
	var ctx context.Context

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		dbMock = mock_types.NewMockDBServiceInterface(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService: serviceif.BaseService{},
			DB:          dbMock,
		}
		logicService.SetNext(dbMock)

		// Put empty transaction into context so we could call DoInTransaction() without access to the real db
		emptyTx := sql.Tx{}
		ctx = context.WithValue(ctx, db.Transaction, &emptyTx)
		dbMock.EXPECT().DB().AnyTimes()
	}

	testClean := func() {
		mockCtrl.Finish()
	}

	t.Run("UpdateFail - Invalid statistic", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		stat := models.UserDefinedLogStat{Pattern: "*foo", Name: "Test invalid"}
		stats := models.UserDefinedLogStatList{Statlist: []*models.UserDefinedLogStat{&stat}}
		updateObj := models.GlobalSystemConfig{UserDefinedLogStatistics: &stats}

		_, err := logicService.UpdateGlobalSystemConfig(ctx, &models.UpdateGlobalSystemConfigRequest{GlobalSystemConfig: &updateObj, FieldMask: types.FieldMask{Paths: []string{}}})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateSuccess", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		stat := models.UserDefinedLogStat{Pattern: ".*[ab][0-9]s1.*", Name: "Test valid"}
		stats := models.UserDefinedLogStatList{Statlist: []*models.UserDefinedLogStat{&stat}}
		updateObj := models.GlobalSystemConfig{UserDefinedLogStatistics: &stats}

		_, err := logicService.UpdateGlobalSystemConfig(ctx, &models.UpdateGlobalSystemConfigRequest{GlobalSystemConfig: &updateObj, FieldMask: types.FieldMask{Paths: []string{}}})
		assert.Nil(t, err)

		testClean()
	})

}

func TestGlobalSystemConfigBgpaasPorts(t *testing.T) {
	var mockCtrl *gomock.Controller
	var dbMock *mock_types.MockDBServiceInterface
	var logicService ContrailTypeLogicService
	var ctx context.Context

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		dbMock = mock_types.NewMockDBServiceInterface(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService: serviceif.BaseService{},
			DB:          dbMock,
		}
		logicService.SetNext(dbMock)

		emptyTx := sql.Tx{}
		ctx = context.WithValue(ctx, db.Transaction, &emptyTx)
		dbMock.EXPECT().DB().AnyTimes()
	}

	testClean := func() {
		mockCtrl.Finish()
	}

	t.Run("UpdateFail - no global config", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {

			return nil, nil
		}).AnyTimes()

		updateObj := models.GlobalSystemConfig{}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &models.UpdateGlobalSystemConfigRequest{GlobalSystemConfig: &updateObj, FieldMask: types.FieldMask{Paths: []string{}}})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateSuccess", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		updateObj := models.GlobalSystemConfig{}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &models.UpdateGlobalSystemConfigRequest{GlobalSystemConfig: &updateObj, FieldMask: types.FieldMask{Paths: []string{}}})
		assert.Nil(t, err)

		testClean()
	})

	t.Run("UpdateFail - wrong port range, too small value", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		ports := models.BGPaaServiceParametersType{PortStart: 0, PortEnd: 1}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &models.UpdateGlobalSystemConfigRequest{GlobalSystemConfig: &updateObj, FieldMask: types.FieldMask{Paths: []string{}}})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateFail - wrong port range, too big value", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		ports := models.BGPaaServiceParametersType{PortStart: 1, PortEnd: math.MaxUint16 + 1}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &models.UpdateGlobalSystemConfigRequest{GlobalSystemConfig: &updateObj, FieldMask: types.FieldMask{Paths: []string{}}})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateFail - shrinking ports", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			ports := models.BGPaaServiceParametersType{PortStart: 10, PortEnd: 20}
			originalObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
			resp := models.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().ListBGPAsAService(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.ListBGPAsAServiceRequest) (*models.ListBGPAsAServiceResponse, error) {
			bgps := []*models.BGPAsAService{nil}
			resp := models.ListBGPAsAServiceResponse{BGPAsAServices: bgps}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		ports := models.BGPaaServiceParametersType{PortStart: 12, PortEnd: 18}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &models.UpdateGlobalSystemConfigRequest{GlobalSystemConfig: &updateObj, FieldMask: types.FieldMask{Paths: []string{}}})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateSuccess - ports updated", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			ports := models.BGPaaServiceParametersType{PortStart: 10, PortEnd: 20}
			originalObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
			resp := models.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().ListBGPAsAService(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.ListBGPAsAServiceRequest) (*models.ListBGPAsAServiceResponse, error) {
			bgps := []*models.BGPAsAService{nil}
			resp := models.ListBGPAsAServiceResponse{BGPAsAServices: bgps}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		ports := models.BGPaaServiceParametersType{PortStart: 8, PortEnd: 22}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &models.UpdateGlobalSystemConfigRequest{GlobalSystemConfig: &updateObj, FieldMask: types.FieldMask{Paths: []string{}}})
		assert.Nil(t, err)

		testClean()
	})
}
