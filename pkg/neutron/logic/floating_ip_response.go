package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
	fipStatusActive           = "ACTIVE"
	fipStatusDown             = "DOWN"
)
// TODO: this file shouldn't make DB transaction and throw errors. Make it simpler, the rest of the logic move into
// floating_ip.go

func makeFloatingipResponse(
	ctx context.Context,
	rp RequestParameters,
	fip *models.FloatingIP,
	vmis []*models.VirtualMachineInterface,
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

	// TODO this should be moved into floating_ip.go. See already written code in that file because it can implement
	//      same logic.
	//port := getPort(vmis)
	//if port == nil {
	//	return &resp, nil
	//}

	resp.PortID = port.GetUUID()
	resp.Status = fipStatusActive
	vnRefs := port.GetVirtualNetworkRefs()
	if len(vnRefs) == 0 {
		return nil, errors.Errorf("missing VirtualNetworkRefs for port with UUID %v when building "+
			"FloatingipResponse %v", port.GetUUID(), fip.GetUUID())
	}
	resp.FloatingNetworkID = vnRefs[0].UUID

	vmiRouters := makeVMIToRouterMapping(routers)
	vmis, err := makeVMIList(ctx, rp, vmis, vmiRouters)
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

func getPort(vmis []*models.VirtualMachineInterface) *models.VirtualMachineInterface {
	for _, port := range vmis {
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
	vmis []*models.VirtualMachineInterface,
	vmiRouters map[string]string) ([]*models.VirtualMachineInterface, error) {
	if len(vmis) == 0 {
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
	return findPorts(vmis, getKeysAsSet(vmiRouters)), nil
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

func findPorts(vmis []*models.VirtualMachineInterface, keys map[string]struct{}) []*models.VirtualMachineInterface {
	found := make([]*models.VirtualMachineInterface, 0, len(vmis))
	for _, port := range vmis {
		if _, ok := keys[port.GetUUID()]; ok {
			found = append(found, port)
		}
	}
	return found
}
