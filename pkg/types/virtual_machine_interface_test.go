package types

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/types/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func virtualMachineInterfaceSetupNextServiceMocks(s *ContrailTypeLogicService) {
	nextService := s.Next().(*servicesmock.MockService)
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
	readService := s.ReadService.(*servicesmock.MockReadService)
	virtualNetwork := models.MakeVirtualNetwork()
	virtualNetwork.UUID = "virtual-network-uuid"
	virtualNetwork.FQName = []string{"default", "test-virtual-network"}

	readService.EXPECT().GetVirtualNetwork(
		gomock.Not(gomock.Nil()),
		&services.GetVirtualNetworkRequest{
			ID: "virtual-network-uuid",
		},
	).Return(
		&services.GetVirtualNetworkResponse{VirtualNetwork: virtualNetwork}, nil,
	).AnyTimes()

	readService.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, common.ErrorNotFound,
	).AnyTimes()
}

func virtualMachineInterfacePrepareMetadata(s *ContrailTypeLogicService, mockCtrl *gomock.Controller) {
	fqNameUUIDTranslator := s.FQNameUUIDTranslator.(*typesmock.MockFQNameUUIDTranslator)

	fqNameUUIDTranslator.EXPECT().TranslateBetweenFQNameUUID(
		gomock.Not(gomock.Nil()),
		"",
		[]string{"default", "test-virtual-network", "test-virtual-network"},
	).Return(
		common.ErrorNotFound,
	).AnyTimes()

	fqNameUUIDTranslator.EXPECT().TranslateBetweenFQNameUUID(
		gomock.Not(gomock.Nil()),
		"",
		gomock.Not(gomock.Nil()),
	).Return(
		&models.MetaData{
			UUID:   "routing-instance-uuid",
			FQName: []string{"default", "test-virtual-network", "test-virtual-network"},
		},
		nil,
	).AnyTimes()
}

func virtualMachineInterfacePrepareRoutingInstance(s *ContrailTypeLogicService) {
	readService := s.ReadService.(*servicesmock.MockReadService)
	routingInstance := models.MakeRoutingInstance()
	routingInstance.UUID = "routing-instance-uuid"
	routingInstance.FQName = []string{"default", "test-virtual-network", "test-virtual-network"}

	readService.EXPECT().GetRoutingInstance(
		gomock.Not(gomock.Nil()),
		&services.GetRoutingInstanceRequest{
			ID: "routing-instance-uuid",
		},
	).Return(
		&services.GetRoutingInstanceResponse{RoutingInstance: routingInstance}, nil,
	).AnyTimes()

	readService.EXPECT().GetRoutingInstance(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		nil, common.ErrorNotFound,
	).AnyTimes()
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
			name: "Create virtual machine interface with no mac address",
			paramVMI: models.VirtualMachineInterface{
				UUID: "bbee32a1-1ccc-4bac-a006-646993303e67",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
				},
			},
			expectedVMI: models.VirtualMachineInterface{
				UUID: "bbee32a1-1ccc-4bac-a006-646993303e67",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
				},
				VirtualMachineInterfaceMacAddresses: &models.MacAddressesType{
					MacAddress: []string{"02:bb:ee:32:a1:1c"},
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			virtualMachineInterfaceSetupNextServiceMocks(service)
			virtualMachineInterfacePrepareNetwork(service)
			virtualMachineInterfacePrepareMetadata(service, mockCtrl)
			virtualMachineInterfacePrepareRoutingInstance(service)

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
