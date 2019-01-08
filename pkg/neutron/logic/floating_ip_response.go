package logic

import (
	"github.com/Juniper/contrail/pkg/models"
)

const (
	fipStatusActive = "ACTIVE"
	fipStatusDown   = "DOWN"
)

func makeFloatingipResponse(
	fip *models.FloatingIP,
	networkID string,
	tenantID string,
	vmi *models.VirtualMachineInterface,
	virtualNetworkRef *models.VirtualMachineInterfaceVirtualNetworkRef,
	routerID string,
) *FloatingipResponse {

	resp := FloatingipResponse{
		ID:                contrailUUIDToNeutronID(fip.GetUUID()),
		TenantID:          contrailUUIDToNeutronID(tenantID),
		FloatingIPAddress: fip.GetFloatingIPAddress(),
		FloatingNetworkID: contrailUUIDToNeutronID(networkID),
		FixedIPAddress:    fip.GetFloatingIPFixedIPAddress(),
		Status:            fipStatusDown,
		CreatedAt:         fip.GetIDPerms().GetCreated(),
		UpdatedAt:         fip.GetIDPerms().GetLastModified(),
		Description:       fip.GetIDPerms().GetDescription(),
	}

	if vmi == nil {
		return &resp
	}

	resp.PortID = vmi.GetUUID()
	resp.Status = fipStatusActive
	resp.FloatingNetworkID = virtualNetworkRef.GetUUID()
	resp.RouterID = routerID

	return &resp

}
