package types

import (
	"context"
	"net/http"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
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
			name:               "Create provision like virtual router",
			expectedStatusCode: http.StatusOK,
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
		},
		{
			name:               "Create virtual router with a reference to flat-subnet ipam",
			expectedStatusCode: http.StatusOK,
			requestedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
					UUID: "network-ipam-a",
					Attr: &models.VirtualRouterNetworkIpamType{
						AllocationPools: []*models.AllocationPoolType{
							{
								VrouterSpecificPool: true,
								Start:               "10.0.0.1",
								End:                 "10.0.0.254",
							},
							{
								VrouterSpecificPool: true,
								Start:               "10.1.0.1",
								End:                 "10.1.0.254",
							},
						},
					},
				}},
			},
			networkIpams: []*models.NetworkIpam{
				{
					UUID:             "network-ipam-a",
					IpamSubnetMethod: models.FlatSubnet,
					IpamSubnets: &models.IpamSubnets{
						Subnets: []*models.IpamSubnetType{{
							AllocationPools: []*models.AllocationPoolType{
								{
									VrouterSpecificPool: true,
									Start:               "10.0.0.1",
									End:                 "10.0.0.254",
								},
								{
									VrouterSpecificPool: true,
									Start:               "10.1.0.1",
									End:                 "10.1.0.254",
								},
							}},
						},
					},
				},
			},
		},
		{
			name:               "Try to create virtual router with a reference to flat-subnet ipam without alloc pools",
			expectedStatusCode: http.StatusBadRequest,
			requestedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
					UUID: "network-ipam-a",
				}},
			},
			networkIpams: []*models.NetworkIpam{{
				UUID:             "network-ipam-a",
				IpamSubnetMethod: models.FlatSubnet,
			}},
		},
		{
			name:               "Try to create virtual router with a reference to flat-subnet ipam with no subnet prefix",
			expectedStatusCode: http.StatusBadRequest,
			requestedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
					UUID: "network-ipam-a",
				}},
			},
			networkIpams: []*models.NetworkIpam{{
				UUID:             "network-ipam-a",
				IpamSubnetMethod: models.FlatSubnet,
			}},
		},
		{
			name:               "Try to create virtual router with a reference to ipam with incorrect ip prefix",
			expectedStatusCode: http.StatusBadRequest,
			requestedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
					UUID: "network-ipam-a",
					Attr: &models.VirtualRouterNetworkIpamType{
						Subnet: []*models.SubnetType{{
							IPPrefix:    "hogohoge",
							IPPrefixLen: 24,
						}},
						AllocationPools: []*models.AllocationPoolType{{
							VrouterSpecificPool: true,
							Start:               "10.0.0.1",
							End:                 "10.0.0.254",
						}},
					},
				}},
			},
			networkIpams: []*models.NetworkIpam{
				{
					UUID:             "network-ipam-a",
					IpamSubnetMethod: models.FlatSubnet,
					IpamSubnets: &models.IpamSubnets{
						Subnets: []*models.IpamSubnetType{{
							AllocationPools: []*models.AllocationPoolType{{
								VrouterSpecificPool: true,
								Start:               "10.0.0.1",
								End:                 "10.0.0.254",
							}},
						}},
					},
				},
			},
		},
		{
			name:               "Try to create virtual router with a reference to user-defined ipam",
			expectedStatusCode: http.StatusBadRequest,
			requestedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
					UUID: "network-ipam-a",
				}},
			},
			networkIpams: []*models.NetworkIpam{{
				UUID:             "network-ipam-a",
				IpamSubnetMethod: models.UserDefinedSubnet,
			}},
		},
		{
			name:               "Try to create virtual router with a reference to ipam with different alloc pool",
			expectedStatusCode: http.StatusBadRequest,
			requestedVR: &models.VirtualRouter{
				UUID: "vr_uuid",
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
					UUID: "network-ipam-a",
					Attr: &models.VirtualRouterNetworkIpamType{
						AllocationPools: []*models.AllocationPoolType{{
							VrouterSpecificPool: true,
							Start:               "10.0.0.1",
							End:                 "10.0.0.254",
						}},
					},
				}},
			},
			networkIpams: []*models.NetworkIpam{
				{
					UUID:             "network-ipam-a",
					IpamSubnetMethod: models.FlatSubnet,
					IpamSubnets: &models.IpamSubnets{
						Subnets: []*models.IpamSubnetType{{
							AllocationPools: []*models.AllocationPoolType{{
								VrouterSpecificPool: true,
								Start:               "10.1.0.1",
								End:                 "10.1.0.254",
							}},
						}},
					},
				},
			},
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
				httpError, ok := errutil.ToHTTPError(err).(*echo.HTTPError)
				assert.True(t, ok, "Expected http error")
				assert.Equal(t, tt.expectedStatusCode, httpError.Code, "Expected different http status")
			} else if assert.NoError(t, err) && tt.expectedVR != nil {
				assert.Equal(t, tt.expectedVR, res.GetVirtualRouter())
			}
		})
	}
}

func TestUpdateVirtualRouter(t *testing.T) {
	tests := []struct {
		name               string
		dataBaseVR         *models.VirtualRouter
		networkIpams       []*models.NetworkIpam
		vrUpdateRequest    *services.UpdateVirtualRouterRequest
		expectedVR         *models.VirtualRouter
		expectedStatusCode int
	}{
		{
			name: "Update ip address of provision like virtual router",
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
					VirtualRouterIPAddress: "10.0.0.1",
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualRouterFieldVirtualRouterIPAddress,
					},
				},
			},
			expectedVR: &models.VirtualRouter{
				UUID:                   "vr_uuid",
				ParentType:             "global-system-config",
				FQName:                 []string{"default-global-system-config", "hoge.hoge.novalocal"},
				VirtualRouterIPAddress: "10.0.0.1",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Add allocation pool",
			dataBaseVR: &models.VirtualRouter{
				InstanceIPBackRefs: []*models.InstanceIP{{InstanceIPAddress: "10.1.0.2"}},
				NetworkIpamRefs:    []*models.VirtualRouterNetworkIpamRef{},
			},
			networkIpams: []*models.NetworkIpam{
				{
					UUID:             "network-ipam-a",
					IpamSubnetMethod: models.FlatSubnet,
					IpamSubnets: &models.IpamSubnets{
						Subnets: []*models.IpamSubnetType{{
							AllocationPools: []*models.AllocationPoolType{{
								VrouterSpecificPool: true,
								Start:               "10.1.0.1",
								End:                 "10.1.0.254",
							}},
						}},
					},
				},
			},
			vrUpdateRequest: &services.UpdateVirtualRouterRequest{
				VirtualRouter: &models.VirtualRouter{
					NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
						UUID: "network-ipam-a",
						Attr: &models.VirtualRouterNetworkIpamType{
							AllocationPools: []*models.AllocationPoolType{
								{
									VrouterSpecificPool: true,
									Start:               "10.1.0.1",
									End:                 "10.1.0.254",
								},
							},
						},
					}},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualRouterFieldNetworkIpamRefs,
					},
				},
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Remove unused allocation pool",
			dataBaseVR: &models.VirtualRouter{
				InstanceIPBackRefs: []*models.InstanceIP{{InstanceIPAddress: "10.1.0.2"}},
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
					UUID: "network-ipam-a",
					Attr: &models.VirtualRouterNetworkIpamType{
						AllocationPools: []*models.AllocationPoolType{
							{
								VrouterSpecificPool: true,
								Start:               "10.0.0.1",
								End:                 "10.0.0.254",
							},
							{
								VrouterSpecificPool: true,
								Start:               "10.1.0.1",
								End:                 "10.1.0.254",
							},
						},
					},
				}},
			},
			networkIpams: []*models.NetworkIpam{
				{
					UUID:             "network-ipam-a",
					IpamSubnetMethod: models.FlatSubnet,
					IpamSubnets: &models.IpamSubnets{
						Subnets: []*models.IpamSubnetType{{
							AllocationPools: []*models.AllocationPoolType{{
								VrouterSpecificPool: true,
								Start:               "10.1.0.1",
								End:                 "10.1.0.254",
							}},
						}},
					},
				},
			},
			vrUpdateRequest: &services.UpdateVirtualRouterRequest{
				VirtualRouter: &models.VirtualRouter{
					NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
						UUID: "network-ipam-a",
						Attr: &models.VirtualRouterNetworkIpamType{
							AllocationPools: []*models.AllocationPoolType{
								{
									VrouterSpecificPool: true,
									Start:               "10.1.0.1",
									End:                 "10.1.0.254",
								},
							},
						},
					}},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualRouterFieldNetworkIpamRefs,
					},
				},
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Try to remove used allocation pool",
			dataBaseVR: &models.VirtualRouter{
				InstanceIPBackRefs: []*models.InstanceIP{{InstanceIPAddress: "10.0.0.2"}},
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
					UUID: "network-ipam-a",
					Attr: &models.VirtualRouterNetworkIpamType{
						AllocationPools: []*models.AllocationPoolType{
							{
								VrouterSpecificPool: true,
								Start:               "10.0.0.1",
								End:                 "10.0.0.254",
							},
							{
								VrouterSpecificPool: true,
								Start:               "10.1.0.1",
								End:                 "10.1.0.254",
							},
						},
					},
				}},
			},
			networkIpams: []*models.NetworkIpam{
				{
					UUID:             "network-ipam-a",
					IpamSubnetMethod: models.FlatSubnet,
					IpamSubnets: &models.IpamSubnets{
						Subnets: []*models.IpamSubnetType{{
							AllocationPools: []*models.AllocationPoolType{{
								VrouterSpecificPool: true,
								Start:               "10.1.0.1",
								End:                 "10.1.0.254",
							}},
						}},
					},
				},
			},
			vrUpdateRequest: &services.UpdateVirtualRouterRequest{
				VirtualRouter: &models.VirtualRouter{
					NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
						UUID: "network-ipam-a",
						Attr: &models.VirtualRouterNetworkIpamType{
							AllocationPools: []*models.AllocationPoolType{
								{
									VrouterSpecificPool: true,
									Start:               "10.1.0.1",
									End:                 "10.1.0.254",
								},
							},
						},
					}},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualRouterFieldNetworkIpamRefs,
					},
				},
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Try to remove all allocation pools",
			dataBaseVR: &models.VirtualRouter{
				InstanceIPBackRefs: []*models.InstanceIP{{InstanceIPAddress: "10.0.0.2"}},
				NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{{
					UUID: "network-ipam-a",
					Attr: &models.VirtualRouterNetworkIpamType{
						AllocationPools: []*models.AllocationPoolType{{
							VrouterSpecificPool: true,
							Start:               "10.0.0.1",
							End:                 "10.0.0.254",
						}},
					},
				}},
			},
			vrUpdateRequest: &services.UpdateVirtualRouterRequest{
				VirtualRouter: &models.VirtualRouter{
					NetworkIpamRefs: []*models.VirtualRouterNetworkIpamRef{},
				},
				FieldMask: types.FieldMask{
					Paths: []string{
						models.VirtualRouterFieldNetworkIpamRefs,
					},
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

			readServiceGetVirtualRouterCall := sv.ReadService.(*servicesmock.MockReadService).
				EXPECT().GetVirtualRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
				Return(&services.GetVirtualRouterResponse{
					VirtualRouter: tt.dataBaseVR,
				}, nil)

			nextServiceUpdateCall := sv.Next().(*servicesmock.MockService).
				EXPECT().UpdateVirtualRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
				DoAndReturn(func(_ context.Context, request *services.UpdateVirtualRouterRequest) (
					*services.UpdateVirtualRouterResponse, error) {
					return &services.UpdateVirtualRouterResponse{
						VirtualRouter: request.VirtualRouter,
					}, nil
				})

			readServiceListNetworkIpamCall.MaxTimes(1)
			if tt.expectedStatusCode != http.StatusOK {
				nextServiceUpdateCall.MaxTimes(1)
				readServiceGetVirtualRouterCall.MaxTimes(1)
			} else {
				nextServiceUpdateCall.Times(1)
				readServiceGetVirtualRouterCall.Times(1)
			}

			res, err := sv.UpdateVirtualRouter(context.Background(), tt.vrUpdateRequest)
			if tt.expectedStatusCode != http.StatusOK {
				require.Error(t, err)
				httpError, ok := errutil.ToHTTPError(err).(*echo.HTTPError)
				assert.True(t, ok, "Expected http error")
				assert.Equal(t, tt.expectedStatusCode, httpError.Code, "Expected different http status")
			} else {
				assert.NoError(t, err)
				if tt.expectedVR != nil {
					assert.Equal(t, tt.expectedVR, res.VirtualRouter)
				}
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
				require.Error(t, err)
				httpError, ok := errutil.ToHTTPError(err).(*echo.HTTPError)
				assert.True(t, ok, "Expected http error")
				assert.Equal(t, tt.expectedStatusCode, httpError.Code, "Expected different http status")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.id, res.ID)
			}
		})
	}
}
