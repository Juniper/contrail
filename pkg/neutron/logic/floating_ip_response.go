package logic

import (
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	fipStatusActive = "ACTIVE"
	fipStatusDown   = "DOWN"
)

type memoData struct {
	// TODO: this will be used for READALL operation
	Ports   struct{}
	NetFQN  struct{}
	Routers struct{}
}

func getPortObject(uuid string, memoReq *memoData) {
}

func gatherFloatingIPDetails(fip *models.FloatingIP, memoReq *memoData) {
	vmiRefs := fip.GetVirtualMachineInterfaceRefs()
	for vmi := range vmiRefs {
		if memoData != nil {
			portObj := getPortObject(memoReq, fip)
		}
	}
}

func makeFloatingipResponse(rp RequestParameters, mfip *models.FloatingIP, memoReq *memoData) *FloatingipResponse {
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
		PortID:            "", // TODO\
		FixedIPAddress:    mfip.GetFloatingIPFixedIPAddress(),
		Status:            fipStatusDown, // TODO
		CreatedAt:         mfip.GetIDPerms().GetCreated(),
		UpdatedAt:         mfip.GetIDPerms().GetLastModified(),
		Description:       mfip.GetIDPerms().GetDescription(),
	}

	details := gatherFloatingIPDetails(mfip, memoReq)
}
