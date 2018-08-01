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

			virtualNetwork, err := sv.getVirtualNetworkFromVirtualMachineInterface(ctx, virtualMachineInterface)
			if err != nil {
				return err
			}

			//TODO further validation

			calculateMacAddresses(virtualMachineInterface)

			response, err = sv.BaseService.CreateVirtualMachineInterface(ctx, request)
			if err != nil {
				return err
			}

			routingInstance, err := sv.getRoutingInstanceFromVirtualNetwork(ctx, virtualNetwork)
			if err != nil {
				return err
			}

			return sv.createRoutingInstanceRefForVirtualMachineInterface(ctx, virtualMachineInterface, routingInstance)
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

func (sv *ContrailTypeLogicService) getRoutingInstanceFromVirtualNetwork(
	ctx context.Context, vn *models.VirtualNetwork) (*models.RoutingInstance, error) {

	vnFqName := vn.GetFQName()
	routingInstanceFqName := append(vnFqName, vnFqName[len(vnFqName)-1])

	metadata, err := sv.FQNameUUIDTranslator.TranslateBetweenFQNameUUID(ctx, "", routingInstanceFqName)
	if err != nil {
		return nil, common.ErrorBadRequestf(
			"missing routing-instance with fq-name [%s]: %v", models.FQNameToString(routingInstanceFqName), err)
	}

	routingInstanceResponse, err := sv.ReadService.GetRoutingInstance(
		ctx,
		&services.GetRoutingInstanceRequest{
			ID: metadata.UUID,
		},
	)
	if err != nil {
		return nil, common.ErrorBadRequestf("missing routing-instance with uuid %s: %v", metadata.UUID, err)
	}

	return routingInstanceResponse.GetRoutingInstance(), nil
}

func (sv *ContrailTypeLogicService) createRoutingInstanceRefForVirtualMachineInterface(
	ctx context.Context, vmi *models.VirtualMachineInterface, routingInstance *models.RoutingInstance) error {

	_, err := sv.WriteService.CreateVirtualMachineInterfaceRoutingInstanceRef(
		ctx,
		&services.CreateVirtualMachineInterfaceRoutingInstanceRefRequest{
			ID: vmi.GetUUID(),
			VirtualMachineInterfaceRoutingInstanceRef: &models.VirtualMachineInterfaceRoutingInstanceRef{
				UUID: routingInstance.UUID,
				To:   routingInstance.FQName,
				Attr: &models.PolicyBasedForwardingRuleType{
					Direction: "both",
				},
			},
		},
	)
	if err != nil {
		return common.ErrorBadRequestf("cannot add routing-instance ref to virtual-machine-interface: %v", err)
	}

	return nil
}
