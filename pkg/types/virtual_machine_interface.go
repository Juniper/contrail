package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateVirtualMachineInterface validates if there is at least one virtual-network
// reference and allocates MAC-address.
func (sv *ContrailTypeLogicService) CreateVirtualMachineInterface(
	ctx context.Context, request *services.CreateVirtualMachineInterfaceRequest,
) (*services.CreateVirtualMachineInterfaceResponse, error) {

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

			vmi.VirtualMachineInterfaceMacAddresses, err = vmi.GetMacAddressesType()
			if err != nil {
				return err
			}

			if err := sv.checkVirtualMachineInterfaceServiceHealthCheckType(nil, vmi); err != nil {
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

// UpdateVirtualMachineInterface validates if there is at least one virtual-network
// reference and allocates MAC-address.
func (sv *ContrailTypeLogicService) UpdateVirtualMachineInterface(
	ctx context.Context, request *services.UpdateVirtualMachineInterfaceRequest,
) (*services.UpdateVirtualMachineInterfaceResponse, error) {

	var response *services.UpdateVirtualMachineInterfaceResponse
	newVMI := request.GetVirtualMachineInterface()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			vmiResponse, err = sv.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
				ID: newVMI.UUID,
			})
			if err != nil {
				return err
			}
			oldVMI := vmiResponse.GetVirtulMachineInterface()

			// TODO further validation

			if err := sv.checkVirtualMachineInterfaceServiceHealthCheckType(oldVMI, newVMI); err != nil {
				return err
			}

			response, err = sv.BaseService.UpdateVirtualMachineInterface(ctx, request)
			if err != nil {
				return err
			}
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getVirtualMachineInterface(
	ctx context.Context, uuid string,
) (*models.VirtualMachineInterface, error) {

}

func (sv *ContrailTypeLogicService) getVirtualNetworkFromVirtualMachineInterface(
	ctx context.Context, vmi *models.VirtualMachineInterface,
) (*models.VirtualNetwork, error) {

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

func (sv *ContrailTypeLogicService) checkVirtualMachineInterfaceServiceHealthCheckType(
	oldVMI, newVMI *models.VirtualMachineInterface,
) error {
	if len(newVMI.GetPortTupleRefs()) > 0 {
		return nil
	}
	if len(oldVMI.GetPortTupleRefs()) > 0 {
		return nil
	}

	for _, shcRef := range newVMI.GetServiceHealthCheckRefs() {
		if err := sv.fillUUIDFieldInRef(ctx, shcRef); err != nil {
			return err
		}
		shcProps, err := sv.getServiceHealthCheck(ctx, shcRef.GetUUID())
		if err != nil {
			return err
		}
		if shcProps.ServiceHealthCheckType != models.ServiceHealthCheckLinkLocalType {
			return errutil.ErrorBadRequestf(
				"Virtual machine interface(%s) of non service vm can only refer link-local type service health check",
				newVMI.GetUUID(),
			)
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) getServiceHealthCheckProperties(
	ctx context.Context, uuid string,
) (*models.ServiceHealthCheckType, error) {

	response, err := sv.ReadService.GetServiceHealthCheck(
		ctx,
		&services.GetServiceHealthCheckRequest{
			ID:     uuid,
			Fields: []string{models.ServiceHealthCheckFieldServiceHealthCheckProperties},
		},
	)
	if err != nil {
		return nil, errutil.ErrorBadRequestf("missing service-health-check with uuid %s: %v", uuid, err)
	}

	return response.GetServiceHealthCheck().GetServiceHealthCheckProperties(), nil
}

func (sv *ContrailTypeLogicService) createRoutingInstanceRefForVirtualMachineInterface(
	ctx context.Context, vmi *models.VirtualMachineInterface, routingInstance *models.RoutingInstance,
) error {

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
