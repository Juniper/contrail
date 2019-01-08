package logic

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
	floatingIPResourceName = "floatingip"
)

// Read floating_ip by UUID
func (fip *Floatingip) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	contrailFloatingIP, contrailRelatedPorts, err := getContrailFloatingIPWithRelatedResources(ctx, rp, id)
	if err != nil {
		return nil, newFloatingipError("can't read contrail floatingip resource", err)
	}

	resp, err := makeFloatingipResponse(ctx, rp, contrailFloatingIP, contrailRelatedPorts, nil)

	if err != nil {
		return nil, newFloatingipError("can't create neutron response for floatingip resource", err)
	}

	return resp, nil
}

// ReadAll logic
func (fip *Floatingip) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	// TODO implement ReadAll logic
	return []FloatingipResponse{}, nil
}

func newFloatingipError(message string, err error) error {
	if isNeutronError(err) {
		// If that error is already neutron error than do not override it.
		return err
	}

	if err != nil {
		message = fmt.Sprintf(" %+v: %+v", message, err)
	}

	return newNeutronError(badRequest, errorFields{
		"resource": floatingIPResourceName,
		"msg":      message,
	})
}

func getContrailFloatingIPWithRelatedResources(ctx context.Context, rp RequestParameters, id string) (
	*models.FloatingIP,
	[]*models.VirtualMachineInterface,
	error,
) {

	floatingIP, err := getContrailFloatingIP(ctx, rp, id)
	if err != nil {
		return nil, nil, err
	}

	ports, err := getContrailVMIsRelatedToFloatingIP(ctx, rp, floatingIP)
	if err != nil {
		return nil, nil, err
	}

	return floatingIP, ports, nil
}

func getContrailFloatingIP(ctx context.Context, rp RequestParameters, id string) (*models.FloatingIP, error) {
	uuid, err := neutronIDToContrailUUID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid uuid format %v for READ %s", id, floatingIPResourceName)
	}
	floatingIPResponse, err := rp.ReadService.GetFloatingIP(ctx, &services.GetFloatingIPRequest{ID: uuid})
	if errutil.IsNotFound(err) {
		return nil, newNeutronError(floatingIPNotFound, errorFields{
			"floatingip_id": id,
			"msg":           err,
		})
	} else if err != nil {
		return nil, err
	}

	return floatingIPResponse.GetFloatingIP(), nil
}

func getContrailVMIsRelatedToFloatingIP(ctx context.Context, rp RequestParameters, floatingIP *models.FloatingIP) (
	[]*models.VirtualMachineInterface, error) {
	vmisRefs := floatingIP.GetVirtualMachineInterfaceRefs()
	vmis := make([]*models.VirtualMachineInterface, 0)

	if vmisRefs != nil && len(vmisRefs) > 0 {
		vmiUUIDs := make([]string, 0, len(vmisRefs))
		for _, vmiRef := range vmisRefs {
			vmiUUIDs = append(vmiUUIDs, vmiRef.GetUUID())
		}

		vmisResp, err := rp.ReadService.ListVirtualMachineInterface(
			ctx, &services.ListVirtualMachineInterfaceRequest{
				Spec: baseservices.SimpleListSpec(vmiUUIDs, models.VirtualMachineInterfaceFieldVirtualNetworkRefs),
			},
		)
		if err != nil {
			return nil, errors.Wrapf(err, "error while reading virtual machine interfaces of ids %v from database.",
				vmiUUIDs)
		}

		vmis = vmisResp.GetVirtualMachineInterfaces()
	}

	return vmis, nil
}
