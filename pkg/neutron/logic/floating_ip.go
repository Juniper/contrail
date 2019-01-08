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
	floatingIPResourceName    = "floatingip"
	serviceInterfaceTypeRight = "right"
)

// Read floating_ip by UUID
func (fip *Floatingip) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	cFloatingIP, networkID, tenantID, cRelatedVMIs, cVnRef, routerId, err := fip.getContrailFloatingIPWithRelatedResources(ctx, rp, id)
	if err != nil {
		return nil, fip.newFloatingipError("can't read contrail floatingip resource", err)
	}

	return makeFloatingipResponse(cFloatingIP, networkID, tenantID, cRelatedVMIs, cVnRef, routerId), nil
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
	string,
	string,
	*models.VirtualMachineInterface,
	*models.VirtualMachineInterfaceVirtualNetworkRef,
	string,
	error,
) {

	floatingIP, err := fip.getContrailFloatingIP(ctx, rp, id)
	if err != nil {
		return nil, "", "", nil, nil, "", err
	}

	floatingNetworkID, err := fip.getFloatingNetworkID(ctx, rp, floatingIP)
	if err != nil {
		return nil, "", "", nil, nil, "", err
	}

	tenantID, err := getTenant(floatingIP)
	if err != nil {
		return nil, "", "", nil, nil, "", err
	}

	vmis, err := fip.getContrailVMIsRelatedToFloatingIP(ctx, rp, floatingIP)
	if err != nil {
		return nil, "", "", nil, nil, "", err
	}

	vmis = fip.filterOutUndesirableVMIs(vmis)
	var vmi *models.VirtualMachineInterface
	if len(vmis) > 0 {
		vmi = vmis[0]
	}

	if vmi == nil {
		return floatingIP, floatingNetworkID, tenantID, nil, nil, "", err
	}

	vnRef, err := fip.getVirtualNetworkRef(vmi)
	if err != nil {
		return nil, "", "", nil, nil, "", err
	}

	routerID, err := fip.findRouterID(ctx, rp, vnRef)
	if err != nil {
		return nil, "", "", nil, nil, "", err
	}

	return floatingIP, floatingNetworkID, tenantID, vmi, vnRef, routerID, nil
}

func (fip *Floatingip) getContrailFloatingIP(ctx context.Context, rp RequestParameters, id string) (
	*models.FloatingIP,
	error,
) {
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

func (fip *Floatingip) getFloatingNetworkID(ctx context.Context, rp RequestParameters, contrailFip *models.FloatingIP) (
	string,
	error,
) {
	fqn := contrailFip.GetFQName()
	if len(fqn) < 2 {
		return "", errors.New(
			fmt.Sprintf("invalid FQName of %s too few number of elements", floatingIPResourceName),
		)
	}
	netFQName := fqn[:len(fqn)-2]
	netIDResp, err := rp.FQNameService.FQNameToID(ctx, &services.FQNameToIDRequest{
		FQName: netFQName,
		Type:   models.KindVirtualNetwork,
	})
	if err != nil {
		return "", errors.Wrapf(err, "could not find uuid of network with fq_name: %v", netFQName)
	}

	return netIDResp.GetUUID(), nil
}

func getTenant(fip *models.FloatingIP) (string, error) {
	if refs := fip.GetProjectRefs(); len(refs) > 0 {
		return refs[0].GetUUID(), nil
	}
	return "", errors.Errorf("could not get tenant from Project refs from FloatingIp with UUID %v", fip.GetUUID())
}

func (fip *Floatingip) getContrailVMIsRelatedToFloatingIP(
	ctx context.Context,
	rp RequestParameters,
	floatingIP *models.FloatingIP) (
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

func (fip *Floatingip) filterOutUndesirableVMIs(
	vmis []*models.VirtualMachineInterface) []*models.VirtualMachineInterface {
	// TODO write test to it. Perhaps unit test will be the best option.
	validVMIs := make([]*models.VirtualMachineInterface, 0, len(vmis))
	for _, vmi := range vmis {
		/* In case of floating ip on the Virtual-ip, svc-monitor will *
		* link floating ip to "right" interface of service VMs       *
		* launched by ha-proxy service instance. Skip them           */
		if fip.isRightInterfaceOfvmHaProxy(vmi) {
			continue // omit this vmi
		}
		validVMIs = append(validVMIs, vmi)
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

func (fip *Floatingip) getVirtualNetworkRef(vmi *models.VirtualMachineInterface) (
	*models.VirtualMachineInterfaceVirtualNetworkRef, error) {
	vnRefs := vmi.GetVirtualNetworkRefs()
	if len(vnRefs) == 0 {
		return nil, errors.Errorf("missing VirtualNetworkRefs for vmi with UUID %v", vmi.GetUUID())
	}
	return vnRefs[0], nil
}

func (fip *Floatingip) findRouterID(ctx context.Context,
	rp RequestParameters,
	ref *models.VirtualMachineInterfaceVirtualNetworkRef) (string, error) {
	linkedVirtualNetwork, err := fip.getVirtualNetwork(ctx, rp, ref)
	if err != nil {
		return "", err
	}

	linkedRouter, err := fip.getLinkedRouter(ctx, rp, linkedVirtualNetwork)
	if err != nil {
		return "", err
	}
	if linkedRouter != nil {
		return linkedRouter.GetUUID(), nil
	}

	return "", errors.New(
		fmt.Sprintf("can't find linked router in assigned virtual network (id = %s)", ref.GetUUID()),
	)
}

func (fip *Floatingip) getVirtualNetwork(ctx context.Context,
	rp RequestParameters,
	ref *models.VirtualMachineInterfaceVirtualNetworkRef) (*models.VirtualNetwork, error) {
	vnResp, err := rp.ReadService.GetVirtualNetwork(
		ctx, &services.GetVirtualNetworkRequest{
			ID: ref.GetUUID(),
		},
	)
	if err != nil {
		return nil, errors.Wrapf(err, "can't read virtual network (id = %s) from database", ref.GetUUID())
	}
	return vnResp.GetVirtualNetwork(), nil
}

func (fip *Floatingip) getLinkedRouter(ctx context.Context,
	rp RequestParameters,
	virtualNetwork *models.VirtualNetwork,
) (*models.LogicalRouter, error) {
	for _, vmi := range virtualNetwork.GetVirtualMachineInterfaceBackRefs() {
		// We need to refresh vmi resource because BackRefs Slice is flat.
		var err error
		vmi, err = fip.reReadVirtualMachineInterface(ctx, rp, vmi)
		if err != nil {
			return nil, err
		}
		linkedLogicalRouters := vmi.GetLogicalRouterBackRefs()
		if len(linkedLogicalRouters) > 0 {
			return linkedLogicalRouters[0], nil
		}
	}
	return nil
}

func (fip *Floatingip) reReadVirtualMachineInterface(
	ctx context.Context,
	rp RequestParameters,
	vmi *models.VirtualMachineInterface,
) (*models.VirtualMachineInterface, error) {
	vmiResp, err := rp.ReadService.GetVirtualMachineInterface(
		ctx, &services.GetVirtualMachineInterfaceRequest{
			ID: vmi.GetUUID(),
		},
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can't refresh virtual machine interface of id=%s", vmi.GetUUID()))
	}
	return vmiResp.GetVirtualMachineInterface(), nil
}
