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
	fip *models.FloatingIP,
	ports []*models.VirtualMachineInterface,
	routers []*models.LogicalRouter,
) (*FloatingipResponse, error) {
	fqn := fip.GetFQName()
	netFQName := fqn[:len(fqn)-2]
	netIDResp, err := rp.FQNameService.FQNameToID(ctx, &services.FQNameToIDRequest{
		FQName: netFQName,
		Type:   models.KindVirtualNetwork,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "could not find uuid of network with fq_name: %v", netFQName)
	}

	tenantID, err := getTenant(fip)
	if err != nil {
		return nil, err
	}

	resp := FloatingipResponse{
		ID:                contrailUUIDToNeutronID(fip.GetUUID()),
		TenantID:          contrailUUIDToNeutronID(tenantID),
		FloatingIPAddress: fip.GetFloatingIPAddress(),
		FloatingNetworkID: contrailUUIDToNeutronID(netIDResp.UUID),
		FixedIPAddress:    fip.GetFloatingIPFixedIPAddress(),
		Status:            fipStatusDown,
		CreatedAt:         fip.GetIDPerms().GetCreated(),
		UpdatedAt:         fip.GetIDPerms().GetLastModified(),
		Description:       fip.GetIDPerms().GetDescription(),
	}

	port := getPort(ports)
	if port == nil {
		return &resp, nil
	}

	resp.PortID = port.GetUUID()
	resp.Status = fipStatusActive
	vnRefs := port.GetVirtualNetworkRefs()
	if len(vnRefs) == 0 {
		return nil, errors.Errorf("missing VirtualNetworkRefs for port with UUID %v when building "+
			"FloatingipResponse %v", port.GetUUID(), fip.GetUUID())
	}
	resp.FloatingNetworkID = vnRefs[0].UUID

	vmiRouters := makeVMIToRouterMapping(routers)
	vmis, err := makeVMIList(ctx, rp, ports, vmiRouters)
	if err != nil {
		return nil, err
	}
	for _, vmi := range vmis {
		if uuid, ok := isSameNet(vmi, port); ok {
			resp.RouterID = uuid
		}
	}
	return &resp, nil
}

func getTenant(fip *models.FloatingIP) (string, error) {
	if refs := fip.GetProjectRefs(); len(refs) > 0 {
		return refs[0].GetUUID(), nil
	}
	return "", errors.Errorf("could not get tenant from Project refs from FloatingIp with UUID %v", fip.GetUUID())
}

func getPort(ports []*models.VirtualMachineInterface) *models.VirtualMachineInterface {
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
	for _, router := range routers {
		for _, ref := range router.GetVirtualMachineInterfaceRefs() {
			vmiRouters[ref.GetUUID()] = router.GetUUID()
		}
	}
	return vmiRouters
}

func makeVMIList(ctx context.Context,
	rp RequestParameters,
	ports []*models.VirtualMachineInterface,
	vmiRouters map[string]string) ([]*models.VirtualMachineInterface, error) {
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
