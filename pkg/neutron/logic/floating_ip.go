package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/services"

	"github.com/pkg/errors"
)

// Read floating_ip by UUID
func (fip *Floatingip) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	uuid, err := neutronIDToContrailUUID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid uuid %v for READ FloatingIp", id)
	}
	resp, err := rp.ReadService.GetFloatingIP(ctx, &services.GetFloatingIPRequest{ID: uuid})
	if errutil.IsNotFound(err) {
		return nil, newNeutronError(floatingIPNotFound, errorFields{
			"floatingip_id": id,
		})
	} else if err != nil {
		return nil, err
	}
	return makeFloatingipResponse(ctx, rp, resp.FloatingIP, nil, nil)
}

// ReadAll logic
func (fip *Floatingip) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	// TODO implement ReadAll logic
	return []FloatingipResponse{}, nil
}
