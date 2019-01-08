package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"

	"github.com/pkg/errors"
)

const (
	fipStatusActive           = "ACTIVE"
	fipStatusDown             = "DOWN"
	serviceInterfaceTypeRight = "right"
)

func makeFloatingipResponse(
	ctx context.Context,
	rp RequestParameters,
	mfip *models.FloatingIP,
	ports []*models.VirtualMachineInterface,
	routers []*models.LogicalRouter,
) (*FloatingipResponse, error) {
	fqn := mfip.GetFQName()
	netFQName := fqn[:len(fqn)-2]
	netIDResp, err := rp.FQNameService.FQNameToID(ctx, &services.FQNameToIDRequest{
		FQName: netFQName,
		Type:   models.KindVirtualNetwork,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "looking for uuid of network with fq_name: %v", netFQName)
	}
	tenantID := acquireTenant(mfip)
	if tenantID == "" {
		return nil, errors.Wrapf(err, "could not read tenant from FloatingIp with uuid %v refs to project", mfip.GetUUID())
	}
	fip := FloatingipResponse{
		ID:                contrailUUIDToNeutronID(mfip.GetUUID()),
		TenantID:          contrailUUIDToNeutronID(acquireTenant(mfip)),
		FloatingIPAddress: mfip.GetFloatingIPAddress(),
		FloatingNetworkID: contrailUUIDToNeutronID(netIDResp.UUID),
		FixedIPAddress:    mfip.GetFloatingIPFixedIPAddress(),
		Status:            fipStatusDown,
		CreatedAt:         mfip.GetIDPerms().GetCreated(),
		UpdatedAt:         mfip.GetIDPerms().GetLastModified(),
		Description:       mfip.GetIDPerms().GetDescription(),
	}

	port := acquirePort(ports)
	if port != nil {
		fip.PortID = port.GetUUID()
		fip.Status = fipStatusActive
		vnRefs := port.GetVirtualNetworkRefs()
		if len(vnRefs) == 0 {
			return nil, errors.Errorf("missing VirtualNetworkRefs for port with UUID %v when building FloatingIPResponse %v", port.GetUUID(), mfip.GetUUID())
		}
		fip.FloatingNetworkID = vnRefs[0].UUID

		vmiRouters := makeVMIToRouterMapping(routers)
		vmis, err := makeVMIsList(ctx, rp, ports, vmiRouters)
		if err != nil {
			return nil, err
		}
		for _, vmi := range vmis {
			if uuid, ok := isSameNet(vmi, port); ok {
				fip.RouterID = uuid
			}
		}
	}
	return &fip, nil
}

func acquireTenant(mfip *models.FloatingIP) string {
	if refs := mfip.GetProjectRefs(); len(refs) > 0 {
		return refs[0].GetUUID()
	}
	return ""
}

func acquirePort(ports []*models.VirtualMachineInterface) *models.VirtualMachineInterface {
	for _, port := range ports {
		if port.GetVirtualMachineInterfaceProperties().GetServiceInterfaceType() != serviceInterfaceTypeRight {
			return port
		}
	}
	return nil
}

func isSameNet(port1, port2 *models.VirtualMachineInterface) (string, bool) {
	refs1 := port1.GetVirtualNetworkRefs()
	refs2 := port2.GetVirtualNetworkRefs()
	if len(refs1) == 0 || len(refs2) == 0 {
		return "", false
	}
	return refs1[0].GetUUID(), refs1[0].GetUUID() == refs2[0].GetUUID()
}

func makeVMIToRouterMapping(routers []*models.LogicalRouter) map[string]string {
	vmiRouters := map[string]string{}
	// ajaj: vmi_routers.update(dict((vmi_ref['uuid'], router_obj.uuid) for vmi_ref in (router_obj.get_virtual_machine_interface_refs() or [])))
	for _, router := range routers {
		for _, ref := range router.GetVirtualMachineInterfaceRefs() {
			vmiRouters[ref.GetUUID()] = router.GetUUID()
		}
	}
	return vmiRouters
}

func makeVMIsList(ctx context.Context, rp RequestParameters, ports []*models.VirtualMachineInterface, vmiRouters map[string]string) ([]*models.VirtualMachineInterface, error) {
	if len(ports) == 0 {
		if len(vmiRouters) == 0 {
			return nil, nil
		}
		req := services.ListVirtualMachineInterfaceRequest{
			Spec: &baseservices.ListSpec{ObjectUUIDs: getKeys(vmiRouters)},
		}
		vmis, err := rp.ReadService.ListVirtualMachineInterface(ctx, &req)
		if err != nil {
			return nil, errors.Wrap(err, "can't list VirtualMachineInterfaces when building FloatingIPResponse")
		}
		return vmis.VirtualMachineInterfaces, nil
	}
	return findPorts(ports, getKeysAsSet(vmiRouters)), nil
}

func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getKeysAsSet(m map[string]string) map[string]struct{} {
	keys := map[string]struct{}{}
	for k := range m {
		keys[k] = struct{}{}
	}
	return keys
}

func findPorts(ports []*models.VirtualMachineInterface, keys map[string]struct{}) []*models.VirtualMachineInterface {
	found := make([]*models.VirtualMachineInterface, 0, len(ports))
	for _, port := range ports {
		if _, ok := keys[port.GetUUID()]; ok {
			found = append(found, port)
		}
	}
	return found
}
