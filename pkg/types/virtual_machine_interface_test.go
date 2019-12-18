package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
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
