package logic

import (
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"

	"github.com/pkg/errors"
)

const (
	fipStatusActive = "ACTIVE"
	fipStatusDown   = "DOWN"
)

func acquirePort(ports []*models.VirtualMachineInterface) *models.VirtualMachineInterface {
	for port := range ports {
		if port.GetVirtualMachineInterfaceProperties().GetServiceInterfaceType() != serviceInterfaceTypeRight {
			return port
		}
	}
	return nil
}

func makeFloatingipResponse(
	rp RequestParameters,
	mfip *models.FloatingIP,
	ports []*models.VirtualMachineInterface,
	routers []*models.LogicalRouter, // TODO check type! Output of: https://github.com/Juniper/contrail-controller/blob/0b6850b55a63280bfb339113d24bd24c953cf145/src/config/vnc_openstack/vnc_openstack/neutron_plugin_db.py#L727
) (*FloatingipResponse, error) {
	fqn := mfip.GetFQName()
	netFQName := fqn[:len(fqn)-1]
	netIDResp, err := rp.FQNameService.FQNameToID(services.FQNameToIDRequest{
		FQName: netFQName,
		Type:   models.KindVirtualNetwork,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "looking for uuid of network with fq_name: %v", netFQName)
	}
	fip := FloatingipResponse{
		ID:                contrailUUIDToNeutronID(mfip.GetUUID()),
		TenantID:          contrailUUIDToNeutronID(mfip.GetProjectRefs().UUID),
		FloatingIPAddress: mfip.GetFloatingIPAddress(),
		FloatingNetworkID: contrailUUIDToNeutronID(netIDResp.UUID),
		RouterID:          "", // TODO
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
		portVnID := vnRefs[0].UUID

		vmiRouters := makeVMIToRouterMapping(routers)
		vmis := makeVMIsList(rp, ports, vmiRouters)
		for vmi := range vmis {
			if uuid, ok := isSameNet(vmi, port); ok {
				fip.RouterID = uuid
			}
		}
	}
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
	for router := range routers {
		for ref := range router.GetVirtualMachineInterfaceRefs() {
			vmiRouters[ref.GetUUID()] = router.GetUUID()
		}
	}
	return vmiRouters
}

func makeVMIsList(rp RequestParameters, ports []*models.VirtualMachineInterface, vmiRouters map[string]string) ([]*models.VirtualMachineInterface, error) {
	if len(ports) == 0 {
		if len(vmiRouters) == 0 {
			return nil, nil
		}
		req := services.ListVirtualMachineInterfaceRequest{Spec{ObjectUUIDs: getKeys(vmiRouters)}}
		vmis, err := rp.ReadService.ListVirtualMachineInterface(req)
		if err != nil {
			return nil, errors.Errorf("can't list VirtualMachineInterfaces when building FloatingIPResponse %v", port.GetUUID(), mfip.GetUUID())
		}
		return vmis.VirtualMachineInterfaces, nil
	}
	return findPorts(ports, getKeysAsSet(vmiRouters))
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

func findPorts(ports []*model.VirtualMachineInterface, keys map[string]struct{}) []*model.VirtualMachineInterface {
	found := make([]*model.VirtualMachineInterface, 0, len(ports))
	for port := range ports {
		if _, ok := keys[port.GetUUID()]; ok {
			found = append(found, port)
		}
	}
	return found
}
