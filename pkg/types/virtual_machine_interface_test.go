package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
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

func TestCreateVirtualMachineInterface(t *testing.T) {
	tests := []struct {
		name        string
		paramVMI    models.VirtualMachineInterface
		expectedVMI models.VirtualMachineInterface
		fails       bool
		errorCode   codes.Code
	}{
		{
			name: "Try to create virtual machine interface without virtual network refs",
			paramVMI: models.VirtualMachineInterface{
				UUID: "uuid",
			},
			fails:     true,
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
			fails:     true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Create virtual machine interface with valid virtual network ref",
			paramVMI: models.VirtualMachineInterface{
				UUID: "uuid",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
				},
			},
			expectedVMI: models.VirtualMachineInterface{
				UUID: "uuid",
				VirtualNetworkRefs: []*models.VirtualMachineInterfaceVirtualNetworkRef{
					{
						UUID: "virtual-network-uuid",
					},
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

			ctx := context.Background()

			paramRequest := services.CreateVirtualMachineInterfaceRequest{VirtualMachineInterface: &tt.paramVMI}
			expectedResponse := services.CreateVirtualMachineInterfaceResponse{VirtualMachineInterface: &tt.expectedVMI}
			createVirtualMachineInterfaceResponse, err := service.CreateVirtualMachineInterface(ctx, &paramRequest)

			if tt.fails {
				assert.Error(t, err)
				if tt.errorCode != codes.OK {
					status, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, tt.errorCode, status.Code())
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createVirtualMachineInterfaceResponse)
			}
		})
	}
}
