package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateVirtualMachineInterface validates if there is at least one virtual-network
// reference and allocates MAC-address.
func (sv *ContrailTypeLogicService) CreateVirtualMachineInterface(
	ctx context.Context,
	request *services.CreateVirtualMachineInterfaceRequest) (*services.CreateVirtualMachineInterfaceResponse, error) {

	var response *services.CreateVirtualMachineInterfaceResponse
	virtualMachineInterface := request.GetVirtualMachineInterface()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {

			virtualNetwork, err := sv.getVirtualNetworkFromVirtualMachineInterface(ctx, virtualMachineInterface)
			if err != nil {
				return err
			}

			_ = virtualNetwork

			//TODO further validation
			//mac-address allocation

			response, err = sv.BaseService.CreateVirtualMachineInterface(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getVirtualNetworkFromVirtualMachineInterface(
	ctx context.Context, virtualMachineInterface *models.VirtualMachineInterface) (*models.VirtualNetwork, error) {

	if len(virtualMachineInterface.GetVirtualNetworkRefs()) == 0 {
		return nil, common.ErrorBadRequest("virtual_network_refs are not defined")
	}

	uuid := virtualMachineInterface.GetVirtualNetworkRefs()[0].GetUUID()
	virtualNetworkResponse, err := sv.ReadService.GetVirtualNetwork(
		ctx,
		&services.GetVirtualNetworkRequest{
			ID: uuid,
		},
	)
	if err != nil {
		return nil, common.ErrorBadRequestf("missing virtual-network with uuid %s: %v", uuid, err)
	}

	return virtualNetworkResponse.GetVirtualNetwork(), nil
}
