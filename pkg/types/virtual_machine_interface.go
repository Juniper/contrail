package types

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	vnicTypeDirect = "direct"
	vnicKeyName    = "vnic_type"
	keyNameHostID  = "host_id"
)

// CreateVirtualMachineInterface validates if there is at least one virtual-network
// reference and allocates MAC-address.
func (sv *ContrailTypeLogicService) CreateVirtualMachineInterface(
	ctx context.Context,
	request *services.CreateVirtualMachineInterfaceRequest,
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

			if err = sv.checkVirtualMachineInterfaceServiceHealthCheckType(ctx, nil, vmi); err != nil {
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

// UpdateVirtualMachineInterface validates update request of virtual machine interface.
func (sv *ContrailTypeLogicService) UpdateVirtualMachineInterface(
	ctx context.Context, request *services.UpdateVirtualMachineInterfaceRequest,
) (*services.UpdateVirtualMachineInterfaceResponse, error) {
	var response *services.UpdateVirtualMachineInterfaceResponse
	newVMI := request.GetVirtualMachineInterface()
	fm := request.GetFieldMask()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			vmiResp, err := sv.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
				ID: newVMI.GetUUID(),
			})
			if err != nil {
				return errors.Wrapf(err, "can't read virtual machine interface from database")
			}
			oldVMI := vmiResp.GetVirtualMachineInterface()
			kvps := oldVMI.GetVirtualMachineInterfaceBindings()

			if basemodels.FieldMaskContains(&fm, models.VirtualMachineInterfaceFieldVirtualMachineRefs) &&
				kvps.GetValue(vnicKeyName) == vnicTypeDirect {
				if err = sv.updateVirtualRouterToVirtualMachinesRefLinks(ctx, oldVMI, newVMI); err != nil {
					return err
				}
			}

			if basemodels.FieldMaskContains(&fm, models.VirtualMachineInterfaceFieldServiceHealthCheckRefs) {
				if err = sv.checkVirtualMachineInterfaceServiceHealthCheckType(ctx, oldVMI, newVMI); err != nil {
					return err
				}
			}

			//nolint: lll
			//TODO: implement rest of pre_dbe_update() logic (python code: https://github.com/Juniper/contrail-controller/blob/b8a2231cfd64f7d2898ea5e1e5bbabb52c7c53ff/src/config/api-server/vnc_cfg_api_server/resources/virtual_machine_interface.py#L574)
			response, err = sv.BaseService.UpdateVirtualMachineInterface(ctx, request)
			//nolint: lll
			//TODO: implement post_dbe_update() logic (python code: https://github.com/Juniper/contrail-controller/blob/b8a2231cfd64f7d2898ea5e1e5bbabb52c7c53ff/src/config/api-server/vnc_cfg_api_server/resources/virtual_machine_interface.py#L885)
			return err
		})
	return response, err
}

// DeleteVirtualMachineInterface validates delete request of virtual machine interface.
func (sv *ContrailTypeLogicService) DeleteVirtualMachineInterface(
	ctx context.Context, request *services.DeleteVirtualMachineInterfaceRequest,
) (*services.DeleteVirtualMachineInterfaceResponse, error) {
	var response *services.DeleteVirtualMachineInterfaceResponse
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			vmiResp, err := sv.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
				ID: request.GetID(),
			})
			if err != nil {
				return errors.Wrapf(err, "can't read virtual_machine_interface of id ='%s' from database", request.GetID())
			}
			vmi := vmiResp.GetVirtualMachineInterface()

			if vmi.GetVirtualMachineInterfaceBindings() != nil {
				if err = sv.updateVirtualRouterToVirtualMachinesRefLinks(ctx, vmi, nil); err != nil {
					return err
				}
			}

			//nolint: lll
			//TODO: implement rest of pre_dbe_delete() logic (python code: https://github.com/Juniper/contrail-controller/blob/b8a2231cfd64f7d2898ea5e1e5bbabb52c7c53ff/src/config/api-server/vnc_cfg_api_server/resources/virtual_machine_interface.py#L923)
			response, err = sv.BaseService.DeleteVirtualMachineInterface(ctx, request)
			//nolint: lll
			//TODO: implement post_dbe_delete() logic (python code: https://github.com/Juniper/contrail-controller/blob/b8a2231cfd64f7d2898ea5e1e5bbabb52c7c53ff/src/config/api-server/vnc_cfg_api_server/resources/virtual_machine_interface.py#L941)
			return err
		})
	return response, err
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
	ctx context.Context, oldVMI, newVMI *models.VirtualMachineInterface,
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
		shcProps, err := sv.getServiceHealthCheckProperties(ctx, shcRef.GetUUID())
		if err != nil {
			return err
		}
		if shcProps.GetHealthCheckType() != models.ServiceHealthCheckLinkLocalType {
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

func (sv *ContrailTypeLogicService) updateVirtualRouterToVirtualMachinesRefLinks(
	ctx context.Context, vmi *models.VirtualMachineInterface, newVMI *models.VirtualMachineInterface,
) error {
	vRouterUUID, err := sv.getVRouterUUID(ctx, vmi.GetVirtualMachineInterfaceBindings())
	if err != nil {
		if errutil.IsNotFound(err) || vRouterUUID == "" {
			return nil
		}
		return err
	}

	vmRefs := vmi.GetVirtualMachineRefs()
	updatedVMRefs := newVMI.GetVirtualMachineRefs()

	if len(vmRefs) == 0 && len(updatedVMRefs) == 0 {
		return nil
	}

	if len(updatedVMRefs) == 0 {
		_, err = sv.WriteService.DeleteVirtualRouterVirtualMachineRef(
			ctx, &services.DeleteVirtualRouterVirtualMachineRefRequest{
				ID: vRouterUUID,
				VirtualRouterVirtualMachineRef: &models.VirtualRouterVirtualMachineRef{
					UUID: vmRefs[0].GetUUID(),
				},
			})
	} else {
		_, err = sv.WriteService.CreateVirtualRouterVirtualMachineRef(
			ctx, &services.CreateVirtualRouterVirtualMachineRefRequest{
				ID: vRouterUUID,
				VirtualRouterVirtualMachineRef: &models.VirtualRouterVirtualMachineRef{
					UUID: updatedVMRefs[0].GetUUID(),
				},
			})
	}
	return err
}

func (sv *ContrailTypeLogicService) getVRouterUUID(ctx context.Context, kvps *models.KeyValuePairs) (string, error) {
	hostID := kvps.GetValue(keyNameHostID)
	if hostID == "" {
		return "",
			errutil.ErrorNotFoundf("Can't find virtual_router because host_id is not set in virtual machine bindings")
	}

	vRouterFQName := []string{defaultGSCName, hostID}
	return sv.FQNameToUUID(ctx, vRouterFQName, models.KindVirtualRouter)
}
