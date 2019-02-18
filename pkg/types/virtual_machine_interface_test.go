package types

import (
	"context"
	"fmt"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/types/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func virtualMachineInterfaceSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService) //nolint: errcheck
	nextService.EXPECT().CreateVirtualMachineInterface(
		gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(_ context.Context, request *services.CreateVirtualMachineInterfaceRequest) (
			response *services.CreateVirtualMachineInterfaceResponse, err error,
		) {
			return &services.CreateVirtualMachineInterfaceResponse{VirtualMachineInterface: request.VirtualMachineInterface}, nil
		},
	).AnyTimes()
}

func virtualMachineInterfacePrepareNetwork(s *ContrailTypeLogicService) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	defaultRoutingInstance := models.MakeRoutingInstance()
	defaultRoutingInstance.UUID = "routing-instance-uuid"
	defaultRoutingInstance.RoutingInstanceIsDefault = true

	virtualNetwork := models.MakeVirtualNetwork()
	virtualNetwork.UUID = "virtual-network-uuid"
	virtualNetwork.RoutingInstances = []*models.RoutingInstance{defaultRoutingInstance}

	readService.EXPECT().GetVirtualNetwork(
		gomock.Not(gomock.Nil()),
		&services.GetVirtualNetworkRequest{
			ID: "virtual-network-uuid",
		},
	).Return(
		&services.GetVirtualNetworkResponse{VirtualNetwork: virtualNetwork}, nil,
	).AnyTimes()

	readService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, errutil.ErrorNotFound,
	).AnyTimes()
}

func virtualMachineInterfacePrepareRoutingInstanceRef(s *ContrailTypeLogicService, shouldCreate bool, vmiID string) {
	writeService := s.WriteService.(*servicesmock.MockWriteService) //nolint: errcheck

	times := 0
	if shouldCreate {
		times = 1
	}

	writeService.EXPECT().CreateVirtualMachineInterfaceRoutingInstanceRef(
		gomock.Not(gomock.Nil()),
		&services.CreateVirtualMachineInterfaceRoutingInstanceRefRequest{
			ID: vmiID,
			VirtualMachineInterfaceRoutingInstanceRef: &models.VirtualMachineInterfaceRoutingInstanceRef{
				UUID: "routing-instance-uuid",
				Attr: &models.PolicyBasedForwardingRuleType{
					Direction: "both",
				},
			},
		},
	).Return(
		nil, nil,
	).Times(times)
}

func TestCreateVirtualMachineInterface(t *testing.T) {
	tests := []struct {
		name        string
		paramVMI    models.VirtualMachineInterface
		expectedVMI models.VirtualMachineInterface
		errorCode   codes.Code
	}{
		{
			name: "Try to create virtual machine interface without virtual network refs",
			paramVMI: models.VirtualMachineInterface{
				UUID: "uuid",
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create virtual machine interface with non-existing virtual network ref",
			paramVMI: models.VirtualMachineInterface{
				UUID: "uuid",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network",
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try to create virtual machine interface with no mac address and too short uuid",
			paramVMI: models.VirtualMachineInterface{
				UUID: "deadbeefho",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network",
					},
				},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Create virtual machine interface with no mac address",
			paramVMI: models.VirtualMachineInterface{
				UUID: "deadbeefhog",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
				},
			},
			expectedVMI: models.VirtualMachineInterface{
				UUID: "deadbeefhog",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
				},
				VirtualMachineInterfaceMacAddresses: &models.MacAddressesType{
					MacAddress: []string{"02:de:ad:be:ef:og"},
				},
			},
		},
		{
			name: "Create virtual machine interface with valid virtual network ref and mac address",
			paramVMI: models.VirtualMachineInterface{
				UUID: "vmi-uuid",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
				},
				VirtualMachineInterfaceMacAddresses: &models.MacAddressesType{
					MacAddress: []string{"01-23-45-67-89-10"},
				},
			},
			expectedVMI: models.VirtualMachineInterface{
				UUID: "vmi-uuid",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
				},
				VirtualMachineInterfaceMacAddresses: &models.MacAddressesType{
					MacAddress: []string{"01:23:45:67:89:10"},
				},
			},
		},
		{
			name: "Create virtual machine interface with different mac address format",
			paramVMI: models.VirtualMachineInterface{
				UUID: "vmi-uuid",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
				},
				VirtualMachineInterfaceMacAddresses: &models.MacAddressesType{
					MacAddress: []string{"01:23:45:67:89:10"},
				},
			},
			expectedVMI: models.VirtualMachineInterface{
				UUID: "vmi-uuid",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
				},
				VirtualMachineInterfaceMacAddresses: &models.MacAddressesType{
					MacAddress: []string{"01:23:45:67:89:10"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			virtualMachineInterfaceSetupNextServiceMocks(service)
			virtualMachineInterfacePrepareNetwork(service)
			virtualMachineInterfacePrepareRoutingInstanceRef(service, tt.errorCode == codes.OK, tt.paramVMI.UUID)

			ctx := context.Background()

			paramRequest := services.CreateVirtualMachineInterfaceRequest{VirtualMachineInterface: &tt.paramVMI}
			expectedResponse := services.CreateVirtualMachineInterfaceResponse{VirtualMachineInterface: &tt.expectedVMI}
			createVirtualMachineInterfaceResponse, err := service.CreateVirtualMachineInterface(ctx, &paramRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createVirtualMachineInterfaceResponse)
			}
		})
	}
}

func mockNoVMIs(s *ContrailTypeLogicService) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck

	readService.EXPECT().GetVirtualMachineInterface(
		gomock.Not(gomock.Nil()), &services.GetVirtualMachineInterfaceRequest{ID: "311a1e01-b65c-421b-93a9-8de29a08bb66"},
	).Return(
		nil, errutil.ErrorNotFoundf("no VirtualMachineInterface found with uuid 311a1e01-b65c-421b-93a9-8de29a08bb66"),
	).AnyTimes()
}

func mockDeleteVMINextSerive(s *ContrailTypeLogicService, vmiUUID string) {
	nextService := s.Next().(*servicesmock.MockService) //nolint: errcheck
	nextService.EXPECT().DeleteVirtualMachineInterface(
		context.Background(), &services.DeleteVirtualMachineInterfaceRequest{ID: vmiUUID},
	).Return(
		&services.DeleteVirtualMachineInterfaceResponse{ID: vmiUUID}, nil,
	).AnyTimes()
}

func mockVMINoBindings(s *ContrailTypeLogicService) {
	vmiUUID := "2d9ebd36-709e-450a-8825-581997d06090"

	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
	readService.EXPECT().GetVirtualMachineInterface(
		gomock.Not(gomock.Nil()), &services.GetVirtualMachineInterfaceRequest{ID: vmiUUID},
	).Return(
		&services.GetVirtualMachineInterfaceResponse{
			VirtualMachineInterface: &models.VirtualMachineInterface{
				UUID: vmiUUID,
				Name: "VMI mock one",
			},
		}, nil,
	).AnyTimes()

	mockDeleteVMINextSerive(s, vmiUUID)
}

func mockVMIVRouterInBindings(s *ContrailTypeLogicService) {
	vmiUUID := "60fd88d6-2a2f-4421-8732-3e41f57a820d"

	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
	readService.EXPECT().GetVirtualMachineInterface(
		gomock.Not(gomock.Nil()), &services.GetVirtualMachineInterfaceRequest{ID: vmiUUID},
	).Return(
		&services.GetVirtualMachineInterfaceResponse{
			VirtualMachineInterface: &models.VirtualMachineInterface{
				UUID: vmiUUID,
				Name: "VMI mock one",
				VirtualMachineInterfaceBindings: &models.KeyValuePairs{
					KeyValuePair: []*models.KeyValuePair{
						{
							Key:   "host_id",
							Value: "host-id0001-352d23",
						},
					},
				},
				VirtualMachineRefs: []*models.VirtualMachineInterfaceVirtualMachineRef{
					{
						// TODO: make sure what should be here...
						UUID: "4bec6c76-d96c-49f9-95bc-fbe823480750",
						To:   []string{""},
						Href: "",
					},
				},
			},
		}, nil).AnyTimes()

	metadataGetterService := s.MetadataGetter.(*typesmock.MockMetadataGetter) //nolint: errcheck
	metadataGetterService.EXPECT().GetMetadata(
		gomock.Not(gomock.Nil()), basemodels.Metadata{Type: "virtual_router", FQName: []string{defaultGSCName, "host-id0001-352d23"}},
	).Return(&basemodels.Metadata{UUID: "5c40eeea-a8f9-4ef4-98b0-64c5b855222f"}, nil).AnyTimes() // TODO change virtual_router into constant

	writeService := s.WriteService.(*servicesmock.MockWriteService) //nolint: errcheck
	writeService.EXPECT().DeleteVirtualRouterVirtualMachineRef(
		gomock.Not(gomock.Nil()), &services.DeleteVirtualRouterVirtualMachineRefRequest{
			ID: "5c40eeea-a8f9-4ef4-98b0-64c5b855222f",
			VirtualRouterVirtualMachineRef: &models.VirtualRouterVirtualMachineRef{
				UUID: "4bec6c76-d96c-49f9-95bc-fbe823480750",
			},
		}).Times(1)

	mockDeleteVMINextSerive(s, vmiUUID)
}

func TestDeleteVirtualMachineInterface(t *testing.T) {
	tests := []struct {
		name        string
		vmiID       string
		expectedErr string
		mockSetUp   func(service *ContrailTypeLogicService)
		statusCode  codes.Code
	}{
		{
			name:        "Try to delete not existing VMI",
			vmiID:       "311a1e01-b65c-421b-93a9-8de29a08bb66",
			expectedErr: "no VirtualMachineInterface found",
			mockSetUp: func(service *ContrailTypeLogicService) {
				mockNoVMIs(service)
			},
			statusCode: codes.NotFound,
		},
		{
			name:        "Delete VMI with empty VirtualMachineInterfaceBindings",
			vmiID:       "2d9ebd36-709e-450a-8825-581997d06090",
			expectedErr: "",
			mockSetUp: func(service *ContrailTypeLogicService) {
				mockVMINoBindings(service)
			},
			statusCode: codes.OK,
		},
		{
			name:        "Delete VMI with vRouter in VirtualMachineInterfaceBindings",
			vmiID:       "60fd88d6-2a2f-4421-8732-3e41f57a820d",
			expectedErr: "",
			mockSetUp: func(service *ContrailTypeLogicService) {
				mockVMIVRouterInBindings(service)
			},
			statusCode: codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)

			tt.mockSetUp(service)

			ctx := context.Background()
			result, err := service.DeleteVirtualMachineInterface(ctx,
				&services.DeleteVirtualMachineInterfaceRequest{ID: tt.vmiID})

			if statusCode, ok := status.FromError(err); ok {
				assert.EqualValues(t, tt.statusCode, statusCode.Code())
			}

			if tt.expectedErr == "" {
				assert.NotNil(t, result)

			} else {
				assert.Error(t, err)
				assert.Contains(t, fmt.Sprintf("%s", err), tt.expectedErr)
			}
		})
	}
}
