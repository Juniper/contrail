package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	fipStatusActive = "ACTIVE"
	fipStatusDown   = "DOWN"
)

// TODO: this file shouldn't make DB transaction and throw errors. Make it simpler, the rest of the logic move into
// floating_ip.go
// TODO: Rewrite this file. See python code.

func makeFloatingipResponse(
	ctx context.Context,
	rp RequestParameters,
	fip *models.FloatingIP,
	vmis *models.VirtualMachineInterface,
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

	return &resp, nil
}

// TODO: move this function into floating_ip.go.
func getTenant(fip *models.FloatingIP) (string, error) {
	if refs := fip.GetProjectRefs(); len(refs) > 0 {
		return refs[0].GetUUID(), nil
	}
	return "", errors.Errorf("could not get tenant from Project refs from FloatingIp with UUID %v", fip.GetUUID())
}
