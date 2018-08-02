package types

import (
	"context"
	"fmt"
	"strings"

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

			_, err := sv.getVirtualNetworkFromVirtualMachineInterface(ctx, virtualMachineInterface)
			if err != nil {
				return err
			}

			//TODO further validation

			calculateMacAddresses(virtualMachineInterface)

			response, err = sv.BaseService.CreateVirtualMachineInterface(ctx, request)
			return err
		})

	return response, err
}

//nolint TODO has to be removed when vn is used
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

func calculateMacAddresses(vmi *models.VirtualMachineInterface) {
	switch len(vmi.GetVirtualMachineInterfaceMacAddresses().GetMacAddress()) {
	case 0:
		uuid := vmi.GetUUID()
		macAddress := fmt.Sprintf("02:%s:%s:%s:%s:%s", uuid[0:2], uuid[2:4], uuid[4:6], uuid[6:8], uuid[9:11])
		vmi.VirtualMachineInterfaceMacAddresses = &models.MacAddressesType{
			MacAddress: []string{macAddress},
		}
	case 1:
		oldMacAddress := vmi.VirtualMachineInterfaceMacAddresses.GetMacAddress()[0]
		newMacAddress := strings.Replace(oldMacAddress, "-", ":", -1)
		vmi.VirtualMachineInterfaceMacAddresses = &models.MacAddressesType{
			MacAddress: []string{newMacAddress},
		}
	}
}
