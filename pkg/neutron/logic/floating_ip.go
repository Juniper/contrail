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
	serviceInterfaceTypeRight = "right"
)

// Read floating_ip by UUID
func (fip *Floatingip) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	contrailFloatingIP, contrailRelatedVMIs, err := fip.getContrailFloatingIPWithRelatedResources(ctx, rp, id)
	if err != nil {
		return nil, fip.newFloatingipError("can't read contrail floatingip resource", err)
	}

	resp, err := makeFloatingipResponse(ctx, rp, contrailFloatingIP, contrailRelatedVMIs, nil)

	if err != nil {
		return nil, fip.newFloatingipError("can't create neutron response for floatingip resource", err)
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

func (fip *Floatingip) newFloatingipError(message string, err error) error {
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

func (fip *Floatingip) getContrailFloatingIPWithRelatedResources(ctx context.Context, rp RequestParameters, id string) (
	*models.FloatingIP,
	*models.VirtualMachineInterface,
	error,
) {

	floatingIP, err := fip.getContrailFloatingIP(ctx, rp, id)
	if err != nil {
		return nil, nil, err
	}

	vmis, err := fip.getContrailVMIsRelatedToFloatingIP(ctx, rp, floatingIP)
	if err != nil {
		return nil, nil, err
	}

	vmis = fip.filterOutUndesirableVMIs(vmis) // TODO write test to cover it.
	// TODO in python code we take only first elem (see neutron_plugin_db.py:2087)

	var firstVMI *models.VirtualMachineInterface
	if len(vmis) > 0 {
		firstVMI = vmis[0]
	}

	return floatingIP, firstVMI, nil
}

func (fip *Floatingip) getContrailFloatingIP(ctx context.Context, rp RequestParameters, id string) (*models.FloatingIP, error) {
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

func (fip *Floatingip) getContrailVMIsRelatedToFloatingIP(ctx context.Context, rp RequestParameters, floatingIP *models.FloatingIP) (
	[]*models.VirtualMachineInterface, error) {
	vmisRefs := floatingIP.GetVirtualMachineInterfaceRefs()
	vmis := make([]*models.VirtualMachineInterface, 0)

	if len(vmisRefs) > 0 {
		vmiRefsUUIDs := make([]string, 0, len(vmisRefs))
		for _, vmiRef := range vmisRefs {
			vmiRefsUUIDs = append(vmiRefsUUIDs, vmiRef.GetUUID())
		}

		vmisResp, err := rp.ReadService.ListVirtualMachineInterface(
			ctx, &services.ListVirtualMachineInterfaceRequest{
				Spec: baseservices.SimpleListSpec(vmiRefsUUIDs, models.VirtualMachineInterfaceFieldVirtualNetworkRefs),
			},
		)
		if err != nil {
			return nil, errors.Wrapf(err, "error while reading virtual machine interfaces of ids %v from database.",
				vmiRefsUUIDs)
		}

		vmis = vmisResp.GetVirtualMachineInterfaces()
	}

	return vmis, nil
}

func (fip *Floatingip) filterOutUndesirableVMIs(vmis []*models.VirtualMachineInterface) []*models.VirtualMachineInterface {
	// TODO write test to it! Perhaps unit test will be the best option.
	 validVMIs := make([]*models.VirtualMachineInterface, len(vmis))
	 for _, vmi := range vmis {
	   /* In case of floating ip on the Virtual-ip, svc-monitor will *
	 	* link floating ip to "right" interface of service VMs       *
	 	* launched by ha-proxy service instance. Skip them           */
	 	if !fip.isRightInterfaceOfvmHaProxy(vmi) {
	 		validVMIs = append(validVMIs, vmi)
		}
	 }
	 return validVMIs
}

func (fip *Floatingip) isRightInterfaceOfvmHaProxy(vmi *models.VirtualMachineInterface) bool {
	if prop := vmi.GetVirtualMachineInterfaceProperties(); prop != nil {
		if prop.ServiceInterfaceType == serviceInterfaceTypeRight {
			return true
		}
	}
	return false
}
