package types

import (
	"regexp/syntax"
	"testing"

	"fmt"

	"math"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/types/mocks"
	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestGSCUdc(t *testing.T) {
	var mockCtrl *gomock.Controller
	var dbMock *mock_types.MockDBServiceInterface
	var logicService ContrailTypeLogicService

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		dbMock = mock_types.NewMockDBServiceInterface(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService: serviceif.BaseService{},
			DB:          dbMock,
		}
		logicService.SetNext(dbMock)
	}

	testClean := func() {
		mockCtrl.Finish()
	}

	t.Run("UpdateFail - Invalid statistic", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{&originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		stat := models.UserDefinedLogStat{"*foo", "Test invalid"}
		stats := models.UserDefinedLogStatList{[]*models.UserDefinedLogStat{&stat}}
		updateObj := models.GlobalSystemConfig{UserDefinedLogStatistics: &stats}

		_, err := logicService.UpdateGlobalSystemConfig(nil, &models.UpdateGlobalSystemConfigRequest{&updateObj, types.FieldMask{[]string{}}})
		assert.Equal(t, err, &syntax.Error{Code: "missing argument to repetition operator", Expr: "*"})

		testClean()
	})

	t.Run("UpdateSuccess", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{&originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		stat := models.UserDefinedLogStat{".*[ab][0-9]s1.*", "Test valid"}
		stats := models.UserDefinedLogStatList{[]*models.UserDefinedLogStat{&stat}}
		updateObj := models.GlobalSystemConfig{UserDefinedLogStatistics: &stats}

		_, err := logicService.UpdateGlobalSystemConfig(nil, &models.UpdateGlobalSystemConfigRequest{&updateObj, types.FieldMask{[]string{}}})
		assert.Equal(t, err, nil)

		testClean()
	})

}

func TestGSCBgpaasPorts(t *testing.T) {
	var mockCtrl *gomock.Controller
	var dbMock *mock_types.MockDBServiceInterface
	var logicService ContrailTypeLogicService

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		dbMock = mock_types.NewMockDBServiceInterface(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService: serviceif.BaseService{},
			DB:          dbMock,
		}
		logicService.SetNext(dbMock)
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
		_, err := logicService.UpdateGlobalSystemConfig(nil, &models.UpdateGlobalSystemConfigRequest{&updateObj, types.FieldMask{[]string{}}})
		assert.Equal(t, err, fmt.Errorf("No GlobalSystemConfig found to update"))

		testClean()
	})

	t.Run("UpdateSuccess", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{&originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		updateObj := models.GlobalSystemConfig{}
		_, err := logicService.UpdateGlobalSystemConfig(nil, &models.UpdateGlobalSystemConfigRequest{&updateObj, types.FieldMask{[]string{}}})
		assert.Equal(t, err, nil)

		testClean()
	})

	t.Run("UpdateFail - wrong port range, too small value", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{&originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		ports := models.BGPaaServiceParametersType{0, 1}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(nil, &models.UpdateGlobalSystemConfigRequest{&updateObj, types.FieldMask{[]string{}}})
		assert.Equal(t, err, fmt.Errorf("Invalid Port range specified"))

		testClean()
	})

	t.Run("UpdateFail - wrong port range, too big value", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			originalObj := models.GlobalSystemConfig{}
			resp := models.GetGlobalSystemConfigResponse{&originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		ports := models.BGPaaServiceParametersType{1, math.MaxUint16 + 1}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(nil, &models.UpdateGlobalSystemConfigRequest{&updateObj, types.FieldMask{[]string{}}})
		assert.Equal(t, err, fmt.Errorf("Invalid Port range specified"))

		testClean()
	})

	t.Run("UpdateFail - shrinking ports", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			ports := models.BGPaaServiceParametersType{10, 20}
			originalObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
			resp := models.GetGlobalSystemConfigResponse{&originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().ListBGPAsAService(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.ListBGPAsAServiceRequest) (*models.ListBGPAsAServiceResponse, error) {
			bgps := []*models.BGPAsAService{nil}
			resp := models.ListBGPAsAServiceResponse{bgps}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		ports := models.BGPaaServiceParametersType{12, 18}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(nil, &models.UpdateGlobalSystemConfigRequest{&updateObj, types.FieldMask{[]string{}}})
		assert.Equal(t, err, fmt.Errorf("BGP Port range cannot be shrunk"))

		testClean()
	})

	t.Run("UpdateSuccess - ports updated", func(t *testing.T) {
		testSetup(t)
		dbMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.GetGlobalSystemConfigRequest) (*models.GetGlobalSystemConfigResponse, error) {
			ports := models.BGPaaServiceParametersType{10, 20}
			originalObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
			resp := models.GetGlobalSystemConfigResponse{&originalObj}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().ListBGPAsAService(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, request *models.ListBGPAsAServiceRequest) (*models.ListBGPAsAServiceResponse, error) {
			bgps := []*models.BGPAsAService{nil}
			resp := models.ListBGPAsAServiceResponse{bgps}
			return &resp, nil
		}).AnyTimes()
		dbMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()

		ports := models.BGPaaServiceParametersType{8, 22}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(nil, &models.UpdateGlobalSystemConfigRequest{&updateObj, types.FieldMask{[]string{}}})
		assert.Equal(t, err, nil)

		testClean()
	})
}
