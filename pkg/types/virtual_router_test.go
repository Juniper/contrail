package types

import (
	"context"
	"net/http"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestCreateVirtualRouter(t *testing.T) {
	tests := []struct {
		name               string
		requestedVR        *models.VirtualRouter
		expectedVR         *models.VirtualRouter
		expectedStatusCode int
		networkIpams       []*models.NetworkIpam
		setupDBMocks       func(sv *ContrailTypeLogicService)
	}{
		{
			name: "create provision like virtual router",
			requestedVR: &models.VirtualRouter{
				UUID:                   "vr_uuid",
				ParentType:             "global-system-config",
				FQName:                 []string{"default-global-system-config", "hoge.hoge.novalocal"},
				VirtualRouterIPAddress: "192.168.0.14",
			},
			expectedVR: &models.VirtualRouter{
				UUID:                   "vr_uuid",
				ParentType:             "global-system-config",
				FQName:                 []string{"default-global-system-config", "hoge.hoge.novalocal"},
				VirtualRouterIPAddress: "192.168.0.14",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Try to create virtual router with a reference to user-defined ipam",
			requestedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{
					{
						UUID: "network-ipam-a",
					},
				},
			},
			expectedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
			},
			networkIpams: []*models.NetworkIpam{
				{
					UUID:             "network-ipam-a",
					IpamSubnetMethod: models.UserDefinedSubnet,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create virtual router with a reference to flat-subnet ipam",
			requestedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{
					{
						UUID: "network-ipam-a",
					},
				},
			},
			expectedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
			},
			networkIpams: []*models.NetworkIpam{
				{
					UUID:             "network-ipam-a",
					IpamSubnetMethod: models.UserDefinedSubnet,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		runTest(t, tt.name, func(t *testing.T, sv *ContrailTypeLogicService) {

			readServiceListNetworkIpamCall := sv.ReadService.(*servicesmock.MockReadService).EXPECT().ListNetworkIpam(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
				&services.ListNetworkIpamResponse{
					NetworkIpams: tt.networkIpams,
				}, nil)

			nextServiceCreateCall := sv.Next().(*servicesmock.MockService).
				EXPECT().CreateVirtualRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
				DoAndReturn(func(_ context.Context, request *services.CreateVirtualRouterRequest) (
					*services.CreateVirtualRouterResponse, error) {
					return &services.CreateVirtualRouterResponse{
						VirtualRouter: request.VirtualRouter,
					}, nil
				})

			readServiceListNetworkIpamCall.MaxTimes(1)
			if tt.expectedStatusCode != http.StatusOK {
				nextServiceCreateCall.MaxTimes(1)
			} else {
				nextServiceCreateCall.Times(1)
			}

			res, err := sv.CreateVirtualRouter(context.Background(), &services.CreateVirtualRouterRequest{
				VirtualRouter: tt.requestedVR,
			})

			if tt.expectedStatusCode != http.StatusOK {
				assert.Error(t, err)
				httpError, ok := common.ToHTTPError(err).(*echo.HTTPError)
				assert.True(t, ok, "Expected http error")
				assert.Equal(t, tt.expectedStatusCode, httpError.Code, "Expected different http status")
			} else if assert.NoError(t, err) {
				assert.Equal(t, tt.expectedVR, res.GetVirtualRouter())
			}
		})
	}
}

func TestUpdateVirtualRouter(t *testing.T) {
	tests := []struct {
		name               string
		dataBaseVR         *models.VirtualRouter
		vrUpdateRequest    *services.UpdateVirtualRouterRequest
		expectedVR         *models.VirtualRouter
		expectedStatusCode int
	}{
		{
			name: "update provision like virtual router",
			dataBaseVR: &models.VirtualRouter{
				UUID:                   "vr_uuid",
				ParentType:             "global-system-config",
				FQName:                 []string{"default-global-system-config", "hoge.hoge.novalocal"},
				VirtualRouterIPAddress: "192.168.0.14",
			},
			vrUpdateRequest: &services.UpdateVirtualRouterRequest{
				VirtualRouter: &models.VirtualRouter{
					UUID:                   "vr_uuid",
					ParentType:             "global-system-config",
					FQName:                 []string{"default-global-system-config", "hoge.hoge.novalocal"},
					VirtualRouterIPAddress: "192.168.0.14",
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualNetworkFieldVirtualNetworkNetworkID,
					},
				},
			},
			expectedVR: &models.VirtualRouter{
				UUID:                   "vr_uuid",
				ParentType:             "global-system-config",
				FQName:                 []string{"default-global-system-config", "hoge.hoge.novalocal"},
				VirtualRouterIPAddress: "192.168.0.14",
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		runTest(t, tt.name, func(t *testing.T, sv *ContrailTypeLogicService) {
			nextServiceUpdateCall := sv.Next().(*servicesmock.MockService).
				EXPECT().UpdateVirtualRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
				DoAndReturn(func(_ context.Context, request *services.UpdateVirtualRouterRequest) (
					*services.UpdateVirtualRouterResponse, error) {
					return &services.UpdateVirtualRouterResponse{
						VirtualRouter: request.VirtualRouter,
					}, nil
				})

			if tt.expectedStatusCode != http.StatusOK {
				nextServiceUpdateCall.MaxTimes(1)
			} else {
				nextServiceUpdateCall.Times(1)
			}

			res, err := sv.UpdateVirtualRouter(context.Background(), tt.vrUpdateRequest)
			if tt.expectedStatusCode != http.StatusOK {
				if assert.Error(t, err) {
					httpError, ok := common.ToHTTPError(err).(*echo.HTTPError)
					assert.True(t, ok, "Expected http error")
					assert.Equal(t, tt.expectedStatusCode, httpError.Code, "Expected different http status")
				}
			} else if assert.NoError(t, err) {
				assert.Equal(t, tt.expectedVR, res.VirtualRouter)
			}
		})
	}
}

func TestDeleteVirtualRouter(t *testing.T) {
	tests := []struct {
		name               string
		dataBaseVR         *models.VirtualRouter
		id                 string
		expectedStatusCode int
	}{
		{
			name: "delete provision like virtual router",
			dataBaseVR: &models.VirtualRouter{
				UUID:                   "vr_uuid",
				ParentType:             "global-system-config",
				FQName:                 []string{"default-global-system-config", "hoge.hoge.novalocal"},
				VirtualRouterIPAddress: "192.168.0.14",
			},
			id:                 "vr_uuid",
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		runTest(t, tt.name, func(t *testing.T, sv *ContrailTypeLogicService) {
			nextServiceCreateCall := sv.Next().(*servicesmock.MockService).
				EXPECT().DeleteVirtualRouter(gomock.Not(gomock.Nil()),
				&services.DeleteVirtualRouterRequest{
					ID: tt.id,
				}).Return(&services.DeleteVirtualRouterResponse{
				ID: tt.id,
			}, nil)

			if tt.expectedStatusCode != http.StatusOK {
				nextServiceCreateCall.MaxTimes(1)
			} else {
				nextServiceCreateCall.Times(1)
			}

			res, err := sv.DeleteVirtualRouter(context.Background(), &services.DeleteVirtualRouterRequest{
				ID: tt.id,
			})

			if tt.expectedStatusCode != http.StatusOK {
				assert.Error(t, err)
				httpError, ok := common.ToHTTPError(err).(*echo.HTTPError)
				assert.True(t, ok, "Expected http error")
				assert.Equal(t, tt.expectedStatusCode, httpError.Code, "Expected different http status")
			} else if assert.NoError(t, err) {
				assert.Equal(t, tt.id, res.ID)
			}
		})
	}
}
