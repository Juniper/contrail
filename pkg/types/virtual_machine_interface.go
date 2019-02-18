package types

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	vnicTypeDirect            = "direct"
	keyNameHostID             = "host_id"
	resourceNameVirtualRouter = "virtual_router"
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

// UpdateVirtualMachineInterface validates update request of virtual machine interface.
func (sv *ContrailTypeLogicService) UpdateVirtualMachineInterface(
	ctx context.Context, request *services.UpdateVirtualMachineInterfaceRequest,
) (*services.UpdateVirtualMachineInterfaceResponse, error) {

	var response *services.UpdateVirtualMachineInterfaceResponse
	updatedVmi := request.GetVirtualMachineInterface()
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			vmiResp, err := sv.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
				ID: updatedVmi.GetUUID(),
			})
			if err != nil {
				return err
			}
			vmi := vmiResp.GetVirtualMachineInterface()
			kvps := vmi.GetVirtualMachineInterfaceBindings()
			if value, err := kvps.GetValue("vnic_type"); value == vnicTypeDirect && err == nil { //nolint: govet
				if err = sv.updateVRouterLinks(ctx, vmi, updatedVmi); err != nil {
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
				if err = sv.updateVRouterLinks(ctx, vmi, nil); err != nil {
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

func (sv *ContrailTypeLogicService) updateVRouterLinks(
	ctx context.Context, vmi *models.VirtualMachineInterface, updatedVmi *models.VirtualMachineInterface) error {
	vRouterUUID, err := sv.getVRouterUUID(ctx, vmi.GetVirtualMachineInterfaceBindings())
	if err != nil {
		return nil
	}

	var vmiRefs, updatedVmiRefs []*models.VirtualMachineInterfaceVirtualMachineRef
	if vmi != nil {
		vmiRefs = vmi.GetVirtualMachineRefs()
	}
	if updatedVmi != nil {
		updatedVmiRefs = updatedVmi.GetVirtualMachineRefs()
	}

	if len(vmiRefs) == 0 && len(updatedVmiRefs) == 0 {
		return nil
	}

	if len(updatedVmiRefs) == 0 {
		_, err = sv.WriteService.DeleteVirtualRouterVirtualMachineRef(
			ctx, &services.DeleteVirtualRouterVirtualMachineRefRequest{
				ID: vRouterUUID,
				VirtualRouterVirtualMachineRef: &models.VirtualRouterVirtualMachineRef{
					UUID: vmiRefs[0].GetUUID(),
				},
			})
	} else {
		_, err = sv.WriteService.CreateVirtualRouterVirtualMachineRef(
			ctx, &services.CreateVirtualRouterVirtualMachineRefRequest{
				ID: vRouterUUID,
				VirtualRouterVirtualMachineRef: &models.VirtualRouterVirtualMachineRef{
					UUID: updatedVmiRefs[0].GetUUID(),
				},
			})
	}
	return err
}

func (sv *ContrailTypeLogicService) getVRouterUUID(ctx context.Context, kvps *models.KeyValuePairs) (string, error) {
	hostID, err := kvps.GetValue(keyNameHostID)
	if err != nil {
		return "", err
	}

	vRouterFQName := []string{defaultGSCName, hostID}
	return sv.FQNameToUUID(ctx, vRouterFQName, resourceNameVirtualRouter)
}
