package types

import (
	"context"
	"math"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
)

func TestGlobalSystemConfigUdc(t *testing.T) {
	var mockCtrl *gomock.Controller
	var inTransationDoerMock *typesmock.MockInTransactionDoer
	var dataServiceMock *servicesmock.MockService
	var logicService ContrailTypeLogicService

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		inTransationDoerMock = typesmock.NewMockInTransactionDoer(mockCtrl)
		dataServiceMock = servicesmock.NewMockService(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService:       services.BaseService{},
			DataService:       dataServiceMock,
			InTransactionDoer: inTransationDoerMock,
		}
		logicService.SetNext(dataServiceMock)

		inTransationDoerMock.EXPECT().DoInTransaction(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
			func(ctx context.Context, do func(context.Context) error) error {
				return do(ctx)
			},
		)

		// Prepare mock expects
		dataServiceMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetGlobalSystemConfigRequest,
			) (*services.GetGlobalSystemConfigResponse, error) {
				originalObj := models.GlobalSystemConfig{}
				resp := services.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
				return &resp, nil
			}).AnyTimes()
		dataServiceMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()
	}

	testClean := func() {
		mockCtrl.Finish()
	}

	ctx := context.Background()
	t.Run("UpdateFail - Invalid statistic", func(t *testing.T) {
		testSetup(t)

		stat := models.UserDefinedLogStat{Pattern: "*foo", Name: "Test invalid"}
		stats := models.UserDefinedLogStatList{Statlist: []*models.UserDefinedLogStat{&stat}}
		updateObj := models.GlobalSystemConfig{UserDefinedLogStatistics: &stats}

		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateSuccess", func(t *testing.T) {
		testSetup(t)

		stat := models.UserDefinedLogStat{Pattern: ".*[ab][0-9]s1.*", Name: "Test valid"}
		stats := models.UserDefinedLogStatList{Statlist: []*models.UserDefinedLogStat{&stat}}
		updateObj := models.GlobalSystemConfig{UserDefinedLogStatistics: &stats}

		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.Nil(t, err)

		testClean()
	})
}

func TestGlobalSystemConfigBgpaasPorts(t *testing.T) {
	var mockCtrl *gomock.Controller
	var inTransationDoerMock *typesmock.MockInTransactionDoer
	var dataServiceMock *servicesmock.MockService
	var logicService ContrailTypeLogicService

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		inTransationDoerMock = typesmock.NewMockInTransactionDoer(mockCtrl)
		dataServiceMock = servicesmock.NewMockService(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService:       services.BaseService{},
			DataService:       dataServiceMock,
			InTransactionDoer: inTransationDoerMock,
		}
		logicService.SetNext(dataServiceMock)

		inTransationDoerMock.EXPECT().DoInTransaction(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
			func(ctx context.Context, do func(context.Context) error) error {
				return do(ctx)
			},
		)

		// Prepare mock expects
		dataServiceMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()
	}

	testClean := func() {
		mockCtrl.Finish()
	}

	ctx := context.Background()
	t.Run("UpdateFail - no global config", func(t *testing.T) {
		testSetup(t)
		dataServiceMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetGlobalSystemConfigRequest,
			) (*services.GetGlobalSystemConfigResponse, error) {
				return nil, nil
			}).AnyTimes()

		updateObj := models.GlobalSystemConfig{}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateSuccess", func(t *testing.T) {
		testSetup(t)
		dataServiceMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetGlobalSystemConfigRequest,
			) (*services.GetGlobalSystemConfigResponse, error) {
				originalObj := models.GlobalSystemConfig{}
				resp := services.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
				return &resp, nil
			}).AnyTimes()

		updateObj := models.GlobalSystemConfig{}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.Nil(t, err)

		testClean()
	})

	t.Run("UpdateFail - wrong port range, too small value", func(t *testing.T) {
		testSetup(t)
		dataServiceMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetGlobalSystemConfigRequest,
			) (*services.GetGlobalSystemConfigResponse, error) {
				originalObj := models.GlobalSystemConfig{}
				resp := services.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
				return &resp, nil
			}).AnyTimes()

		ports := models.BGPaaServiceParametersType{PortStart: 0, PortEnd: 1}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateFail - wrong port range, too big value", func(t *testing.T) {
		testSetup(t)
		dataServiceMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetGlobalSystemConfigRequest,
			) (*services.GetGlobalSystemConfigResponse, error) {
				originalObj := models.GlobalSystemConfig{}
				resp := services.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
				return &resp, nil
			}).AnyTimes()

		ports := models.BGPaaServiceParametersType{PortStart: 1, PortEnd: math.MaxUint16 + 1}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateFail - shrinking ports", func(t *testing.T) {
		testSetup(t)
		dataServiceMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetGlobalSystemConfigRequest,
			) (*services.GetGlobalSystemConfigResponse, error) {
				ports := models.BGPaaServiceParametersType{PortStart: 10, PortEnd: 20}
				originalObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
				resp := services.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
				return &resp, nil
			}).AnyTimes()
		dataServiceMock.EXPECT().ListBGPAsAService(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.ListBGPAsAServiceRequest,
			) (*services.ListBGPAsAServiceResponse, error) {
				bgps := []*models.BGPAsAService{nil}
				resp := services.ListBGPAsAServiceResponse{BGPAsAServices: bgps}
				return &resp, nil
			}).AnyTimes()

		ports := models.BGPaaServiceParametersType{PortStart: 12, PortEnd: 18}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateSuccess - ports updated", func(t *testing.T) {
		testSetup(t)
		dataServiceMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetGlobalSystemConfigRequest,
			) (*services.GetGlobalSystemConfigResponse, error) {
				ports := models.BGPaaServiceParametersType{PortStart: 10, PortEnd: 20}
				originalObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
				resp := services.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
				return &resp, nil
			}).AnyTimes()
		dataServiceMock.EXPECT().ListBGPAsAService(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.ListBGPAsAServiceRequest,
			) (*services.ListBGPAsAServiceResponse, error) {
				bgps := []*models.BGPAsAService{nil}
				resp := services.ListBGPAsAServiceResponse{BGPAsAServices: bgps}
				return &resp, nil
			}).AnyTimes()

		ports := models.BGPaaServiceParametersType{PortStart: 8, PortEnd: 22}
		updateObj := models.GlobalSystemConfig{BgpaasParameters: &ports}
		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.Nil(t, err)

		testClean()
	})
}

func TestGlobalSystemConfigAsn(t *testing.T) {
	var mockCtrl *gomock.Controller
	var inTransationDoerMock *typesmock.MockInTransactionDoer
	var dataServiceMock *servicesmock.MockService
	var logicService ContrailTypeLogicService

	testSetup := func(t *testing.T) {
		mockCtrl = gomock.NewController(t)
		inTransationDoerMock = typesmock.NewMockInTransactionDoer(mockCtrl)
		dataServiceMock = servicesmock.NewMockService(mockCtrl)
		logicService = ContrailTypeLogicService{
			BaseService:       services.BaseService{},
			DataService:       dataServiceMock,
			InTransactionDoer: inTransationDoerMock,
		}
		logicService.SetNext(dataServiceMock)

		inTransationDoerMock.EXPECT().DoInTransaction(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
			func(ctx context.Context, do func(context.Context) error) error {
				return do(ctx)
			},
		)

		// Prepare mock expects
		dataServiceMock.EXPECT().GetGlobalSystemConfig(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetGlobalSystemConfigRequest,
			) (*services.GetGlobalSystemConfigResponse, error) {
				originalObj := models.GlobalSystemConfig{}
				resp := services.GetGlobalSystemConfigResponse{GlobalSystemConfig: &originalObj}
				return &resp, nil
			}).AnyTimes()
		dataServiceMock.EXPECT().UpdateGlobalSystemConfig(gomock.Any(), gomock.Any()).AnyTimes()
	}

	testClean := func() {
		mockCtrl.Finish()
	}

	ctx := context.Background()
	t.Run("UpdateSuccess - Empty VN list", func(t *testing.T) {
		testSetup(t)

		dataServiceMock.EXPECT().ListVirtualNetwork(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.ListVirtualNetworkRequest,
			) (*services.ListVirtualNetworkResponse, error) {
				resp := services.ListVirtualNetworkResponse{}
				return &resp, nil
			}).AnyTimes()

		updateObj := models.GlobalSystemConfig{AutonomousSystem: 1}

		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.Nil(t, err)

		testClean()
	})

	t.Run("UpdateFail - VN has no user defined route targets", func(t *testing.T) {
		testSetup(t)

		dataServiceMock.EXPECT().ListVirtualNetwork(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.ListVirtualNetworkRequest,
			) (*services.ListVirtualNetworkResponse, error) {
				rtList := models.RouteTargetList{RouteTarget: []string{"target:1:8000001"}}
				vn := models.VirtualNetwork{RouteTargetList: &rtList}
				resp := services.ListVirtualNetworkResponse{VirtualNetworks: []*models.VirtualNetwork{&vn}}
				return &resp, nil
			}).AnyTimes()

		updateObj := models.GlobalSystemConfig{AutonomousSystem: 1}

		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.NotNil(t, err)

		testClean()
	})

	t.Run("UpdateSuccess - VN has user defined route targets with ip", func(t *testing.T) {
		testSetup(t)

		dataServiceMock.EXPECT().ListVirtualNetwork(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.ListVirtualNetworkRequest,
			) (*services.ListVirtualNetworkResponse, error) {
				rtList := models.RouteTargetList{RouteTarget: []string{"target:192.168.0.1:8000001"}}
				vn := models.VirtualNetwork{RouteTargetList: &rtList}
				resp := services.ListVirtualNetworkResponse{VirtualNetworks: []*models.VirtualNetwork{&vn}}
				return &resp, nil
			}).AnyTimes()

		updateObj := models.GlobalSystemConfig{AutonomousSystem: 1}

		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.Nil(t, err)

		testClean()
	})

	t.Run("UpdateSuccess - VN has user defined route targets with small target id", func(t *testing.T) {
		testSetup(t)

		dataServiceMock.EXPECT().ListVirtualNetwork(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.ListVirtualNetworkRequest,
			) (*services.ListVirtualNetworkResponse, error) {
				rtList := models.RouteTargetList{RouteTarget: []string{"target:1:1"}}
				vn := models.VirtualNetwork{RouteTargetList: &rtList}
				resp := services.ListVirtualNetworkResponse{VirtualNetworks: []*models.VirtualNetwork{&vn}}
				return &resp, nil
			}).AnyTimes()

		updateObj := models.GlobalSystemConfig{AutonomousSystem: 1}

		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.Nil(t, err)

		testClean()
	})

	t.Run("UpdateFail - invalid Route Target format", func(t *testing.T) {
		testSetup(t)

		dataServiceMock.EXPECT().ListVirtualNetwork(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.ListVirtualNetworkRequest,
			) (*services.ListVirtualNetworkResponse, error) {
				rtList := models.RouteTargetList{RouteTarget: []string{"target:1a:1"}}
				vn := models.VirtualNetwork{RouteTargetList: &rtList}
				resp := services.ListVirtualNetworkResponse{VirtualNetworks: []*models.VirtualNetwork{&vn}}
				return &resp, nil
			}).AnyTimes()

		updateObj := models.GlobalSystemConfig{AutonomousSystem: 1}

		_, err := logicService.UpdateGlobalSystemConfig(ctx, &services.UpdateGlobalSystemConfigRequest{
			GlobalSystemConfig: &updateObj,
			FieldMask:          types.FieldMask{Paths: []string{}},
		})
		assert.NotNil(t, err)

		testClean()
	})
}
