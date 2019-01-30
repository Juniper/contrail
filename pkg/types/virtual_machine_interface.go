package types

import (
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateVirtualMachineInterface validates if there is at least one virtual-network
// reference and allocates MAC-address.
func (sv *ContrailTypeLogicService) CreateVirtualMachineInterface(
	ctx context.Context,
	request *services.CreateVirtualMachineInterfaceRequest) (*services.CreateVirtualMachineInterfaceResponse, error) {

	var response *services.CreateVirtualMachineInterfaceResponse
	vmi := request.GetVirtualMachineInterface()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {

			vn, err := sv.getVirtualNetworkFromVirtualMachineInterface(ctx, vmi)
			if err != nil {
				return err
			}

			//TODO further validation

			vmi.VirtualMachineInterfaceMacAddresses, err = calculateMacAddresses(vmi)
			if err != nil {
				return err
			}

			response, err = sv.BaseService.CreateVirtualMachineInterface(ctx, request)
			if err != nil {
				return err
			}

			ri := vn.GetDefaultRoutingInstance()
			if ri == nil {
				return errutil.ErrorBadRequestf("could not get default routing instance for VN (%v)", vn.UUID)
			}

			return sv.createRoutingInstanceRefForVirtualMachineInterface(ctx, vmi, ri)
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getVirtualNetworkFromVirtualMachineInterface(
	ctx context.Context, vmi *models.VirtualMachineInterface) (*models.VirtualNetwork, error) {

	if len(vmi.GetVirtualNetworkRefs()) == 0 {
		return nil, errutil.ErrorBadRequest("virtual_network_refs are not defined")
	}

	uuid := vmi.GetVirtualNetworkRefs()[0].GetUUID()
	response, err := sv.ReadService.GetVirtualNetwork(
		ctx,
		&services.GetVirtualNetworkRequest{
			ID: uuid,
		},
	)
	if err != nil {
		return nil, errutil.ErrorBadRequestf("missing virtual-network with uuid %s: %v", uuid, err)
	}

	return response.GetVirtualNetwork(), nil
}

func calculateMacAddresses(vmi *models.VirtualMachineInterface) (*models.MacAddressesType, error) {
	addrs := len(vmi.GetVirtualMachineInterfaceMacAddresses().GetMacAddress())

	if addrs == 1 {
		oldMacAddress := vmi.VirtualMachineInterfaceMacAddresses.GetMacAddress()[0]
		newMacAddress := strings.Replace(oldMacAddress, "-", ":", -1)
		return &models.MacAddressesType{
			MacAddress: []string{newMacAddress},
		}, nil
	}

	uuid := vmi.GetUUID()
	if len(uuid) < 11 {
		return nil, errutil.ErrorBadRequestf("could not generate mac address: vn uuid (%v) too short", uuid)
	}

	macAddress := fmt.Sprintf("02:%s:%s:%s:%s:%s", uuid[0:2], uuid[2:4], uuid[4:6], uuid[6:8], uuid[9:11])
	return &models.MacAddressesType{
		MacAddress: []string{macAddress},
	}, nil
}

func (sv *ContrailTypeLogicService) createRoutingInstanceRefForVirtualMachineInterface(
	ctx context.Context, vmi *models.VirtualMachineInterface, routingInstance *models.RoutingInstance) error {

	_, err := sv.WriteService.CreateVirtualMachineInterfaceRoutingInstanceRef(
		ctx, &services.CreateVirtualMachineInterfaceRoutingInstanceRefRequest{
			ID: vmi.GetUUID(),
			VirtualMachineInterfaceRoutingInstanceRef: &models.VirtualMachineInterfaceRoutingInstanceRef{
				UUID: routingInstance.UUID,
				Attr: &models.PolicyBasedForwardingRuleType{
					Direction: "both",
				},
			},
		},
	)

	if err != nil {
		return errutil.ErrorBadRequestf("cannot add routing-instance ref to virtual-machine-interface: %v", err)
	}

	return nil
}
