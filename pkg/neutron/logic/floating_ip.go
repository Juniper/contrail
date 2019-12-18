package logic

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	serviceInterfaceTypeRight = "right"
)

// Create logic.
func (fip *Floatingip) Create(
	ctx context.Context, rp RequestParameters,
) (Response, error) {
	fIPP, err := fip.getFloatingIPPool(ctx, rp)
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": KindFloatingip,
			"msg":      fmt.Sprintf("Network %s doesn't provide a floatingip pool", fip.FloatingNetworkID),
		})
	}

	project, err := getProject(ctx, rp)
	if err != nil {
		return nil, err
	}

	vmiRefs, err := fip.getVMIRefs(ctx, rp)
	if err != nil {
		return nil, err
	}

	fixedIP, err := fip.getFixedIPAddress(ctx, rp)
	if err != nil {
		return nil, err
	}

	response, err := rp.WriteService.CreateFloatingIP(ctx, &services.CreateFloatingIPRequest{
		FloatingIP: fip.newFloatingIP(fIPP, project, vmiRefs, fixedIP),
	})
	if err != nil {
		return nil, newNeutronError(ipAddressGenerationFailure, errorFields{
			"net_id": fip.FloatingNetworkID,
		})
	}

	return floatingipVncToNeutron(ctx, rp, response.GetFloatingIP())
}

func (fip *Floatingip) newFloatingIP(
	fIPP *models.FloatingIPPool,
	project *models.Project,
	vmiRefs []*models.FloatingIPVirtualMachineInterfaceRef,
	fixedIP string,
) *models.FloatingIP {
	return &models.FloatingIP{
		Name:              fip.ID,
		UUID:              fip.ID,
		ParentUUID:        fIPP.GetUUID(),
		ParentType:        fIPP.Kind(),
		FloatingIPAddress: fip.FloatingIPAddress,
		Perms2: &models.PermType2{
			Owner:       project.UUID,
			OwnerAccess: permsRWX,
		},
		ProjectRefs: []*models.FloatingIPProjectRef{
			{
				To:   project.FQName,
				UUID: project.UUID,
			},
		},
		VirtualMachineInterfaceRefs: vmiRefs,
		FloatingIPFixedIPAddress:    fixedIP,
		IDPerms: &models.IdPermsType{
			Enable:      true,
			Description: fip.Description,
		},
	}
}

// Read logic.
func (fip *Floatingip) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	readResponse, err := rp.ReadService.GetFloatingIP(ctx, &services.GetFloatingIPRequest{
		ID: id,
	})
	if err != nil {
		return nil, newNeutronError(floatingIPNotFound, errorFields{
			"floatingip_id": id,
			"msg":           errors.Wrapf(err, "failed to get floating ip with uuid: '%s'", id),
		})
	}
	return floatingipVncToNeutron(ctx, rp, readResponse.GetFloatingIP())
}

// ReadAll logic.
func (fip *Floatingip) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	rf, err := createRefFilters(rp.RequestContext, filters)
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"filters": filters,
			"msg":     errors.Wrapf(err, "failed to create correct filters"),
		})
	}

	listResponse, err := rp.ReadService.ListFloatingIP(ctx, &services.ListFloatingIPRequest{
		Spec: &baseservices.ListSpec{
			ObjectUUIDs: filters["id"],
			Filters:     createVncFilters(filters),
			RefUUIDs:    rf,
			Detail:      true,
		},
	})
	if err != nil {
		return nil, newNeutronError(floatingIPNotFound, errorFields{
			"filters": filters,
			"msg":     errors.Wrapf(err, "failed to list floating ips"),
		})
	}

	response := []*FloatingipResponse{}
	for _, fip := range listResponse.GetFloatingIPs() {
		nfip, err := floatingipVncToNeutron(ctx, rp, fip)
		if err != nil {
			rp.Log.Warnf("Failed to convert floating ip (id: %s) from vnc to neutron: %s", fip.GetUUID(), err)
			continue
		}
		response = append(response, nfip)

	}
	return response, nil
}

// Update logic.
func (fip *Floatingip) Update(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	vmiRefs, err := fip.getVMIRefs(ctx, rp)
	if err != nil {
		return nil, err
	}

	fixedIP, err := fip.getFixedIPAddress(ctx, rp)
	if err != nil {
		return nil, err
	}

	_, err = rp.WriteService.UpdateFloatingIP(ctx, &services.UpdateFloatingIPRequest{
		FloatingIP: &models.FloatingIP{
			UUID:                        id,
			FloatingIPAddress:           fip.FloatingIPAddress,
			VirtualMachineInterfaceRefs: vmiRefs,
			FloatingIPFixedIPAddress:    fixedIP,
		},
		FieldMask: createFieldMaskForUpdate(rp),
	})
	if err != nil {
		return nil, err
	}
	getResponse, err := rp.ReadService.GetFloatingIP(ctx, &services.GetFloatingIPRequest{
		ID: id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get floating ip after updating")
	}
	return floatingipVncToNeutron(ctx, rp, getResponse.GetFloatingIP())
}

// Delete logic.
func (fip *Floatingip) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	_, err := rp.WriteService.DeleteFloatingIP(ctx, &services.DeleteFloatingIPRequest{
		ID: id,
	})
	if err == nil {
		return nil, nil
	}
	switch {
	case errutil.IsNotFound(err):
		return nil, newNeutronError(floatingIPNotFound, errorFields{
			"floatingip_id": id,
		})
	default:
		return nil, err
	}
}

func createRefFilters(rc RequestContext, f Filters) (map[string]*baseservices.UUIDs, error) {
	refs := make(map[string]*baseservices.UUIDs)
	if !rc.IsAdmin {
		refs[models.FloatingIPFieldProjectRefs] = &baseservices.UUIDs{
			UUIDs: []string{rc.TenantID},
		}
	} else if projects, ok := f["tenant_id"]; ok {
		ps, err := neutronIDsToVncUUIDs(projects)
		if err != nil {
			return nil, err
		}
		refs[models.FloatingIPFieldProjectRefs] = &baseservices.UUIDs{
			UUIDs: ps,
		}
	}

	if vmis, ok := f["port_id"]; ok {
		refs[models.FloatingIPFieldVirtualMachineInterfaceRefs] = &baseservices.UUIDs{
			UUIDs: vmis,
		}
	}

	return refs, nil
}

func createVncFilters(f Filters) []*baseservices.Filter {
	var filters []*baseservices.Filter
	if ip, ok := f["floating_ip_address"]; ok {
		filters = append(filters, &baseservices.Filter{
			Key:    models.FloatingIPFieldFloatingIPAddress,
			Values: ip,
		})
	}

	return filters
}

func createFieldMaskForUpdate(rp RequestParameters) types.FieldMask {
	var paths []string
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(FloatingipFieldFloatingIPAddress)) {
		paths = append(paths, models.FloatingIPFieldFloatingIPAddress)
	}
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(FloatingipFieldPortID)) {
		paths = append(paths, models.FloatingIPFieldFloatingIPFixedIPAddress)
		paths = append(paths, models.FloatingIPFieldVirtualMachineInterfaceRefs)
	}
	return types.FieldMask{
		Paths: paths,
	}
}

func (fip *Floatingip) getFloatingIPPool(
	ctx context.Context, rp RequestParameters,
) (*models.FloatingIPPool, error) {
	fIPPListResponse, err := rp.ReadService.ListFloatingIPPool(ctx, &services.ListFloatingIPPoolRequest{
		Spec: &baseservices.ListSpec{
			ParentUUIDs: []string{fip.FloatingNetworkID},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(fIPPListResponse.GetFloatingIPPools()) == 0 {
		return nil, errors.Errorf("no floating-ip-pool with parent_uuid: '%s'", fip.FloatingNetworkID)
	}
	return fIPPListResponse.GetFloatingIPPools()[0], nil
}

func (fip *Floatingip) getVMIRefs(
	ctx context.Context, rp RequestParameters,
) (refs []*models.FloatingIPVirtualMachineInterfaceRef, err error) {
	if fip.PortID == "" {
		return nil, nil
	}
	response, err := rp.IDToFQNameService.IDToFQName(ctx, &services.IDToFQNameRequest{
		UUID: fip.PortID,
	})
	if err != nil {
		return nil, newNeutronError(portNotFound, errorFields{
			"resource": KindFloatingip,
			"port_id":  fip.PortID,
		})
	}
	//TODO: validate vmi
	//TODO: validate if strict_compliance enabled
	return []*models.FloatingIPVirtualMachineInterfaceRef{
		{
			To:   response.GetFQName(),
			UUID: fip.PortID,
		},
	}, nil
}

func (fip *Floatingip) getFixedIPAddress(
	ctx context.Context, rp RequestParameters,
) (fixedIP string, err error) {
	if fip.PortID == "" && fip.FixedIPAddress != "" {
		return "", newNeutronError(badRequest, errorFields{
			"resource": KindFloatingip,
			"msg":      fmt.Sprint("fixed_ip_address cannot be specified without a port_id"),
		})
	}
	if fip.PortID != "" && fip.FixedIPAddress == "" {
		if fip.FixedIPAddress, err = fip.extractFixedIPAddressFromPort(ctx, rp); err != nil {
			return "", err
		}
	}

	//TODO: _check_port_fip_assoc
	return fip.FixedIPAddress, nil
}

func (fip *Floatingip) extractFixedIPAddressFromPort(ctx context.Context, rp RequestParameters) (string, error) {
	vmiResponse, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
		ID: fip.PortID,
	})
	if err != nil {
		return "", errors.Wrapf(err, "failed to fetch port with uuid: '%s'", fip.PortID)
	}
	iipBackRefs := vmiResponse.GetVirtualMachineInterface().GetInstanceIPBackRefs()
	if len(iipBackRefs) > 1 {
		return "", newNeutronError(badRequest, errorFields{
			"resource": KindFloatingip,
			"msg": fmt.Sprintf("Port %s has multiple fixed IP addresses. "+
				"Must provide a specific IP address when assigning a floating IP",
				vmiResponse.GetVirtualMachineInterface().GetUUID()),
		})
	}
	if len(iipBackRefs) == 1 {
		return iipBackRefs[0].GetInstanceIPAddress(), nil
	}
	return "", nil
}

func floatingipVncToNeutron(
	ctx context.Context, rp RequestParameters, fIP *models.FloatingIP,
) (*FloatingipResponse, error) {
	port, id := getPortAndID(ctx, rp, fIP)

	netID, err := getFloatingNetworkID(ctx, rp, fIP)
	if err != nil {
		return nil, err
	}
	return &FloatingipResponse{
		ID:                fIP.UUID,
		TenantID:          getFloatingIPTenantID(fIP),
		FloatingIPAddress: fIP.FloatingIPAddress,
		FloatingNetworkID: netID,
		FixedIPAddress:    fIP.FloatingIPFixedIPAddress,
		CreatedAt:         fIP.GetIDPerms().GetCreated(),
		UpdatedAt:         fIP.GetIDPerms().GetLastModified(),
		Status:            getStatus(port),
		Description:       fIP.GetIDPerms().GetDescription(),
		PortID:            id,
		RouterID:          port.GetRouterID(),
	}, nil
}

func getFloatingIPTenantID(fIP *models.FloatingIP) string {
	if len(fIP.ProjectRefs) > 0 {
		return VncUUIDToNeutronID(fIP.ProjectRefs[0].UUID)
	}
	return ""
}

func getFloatingNetworkID(ctx context.Context, rp RequestParameters, fIP *models.FloatingIP) (string, error) {
	response, err := rp.FQNameToIDService.FQNameToID(ctx, &services.FQNameToIDRequest{
		FQName: getGrandParentFQName(fIP.GetFQName()),
		Type:   models.KindVirtualNetwork,
	})
	if err != nil {
		return "", err
	}
	return response.GetUUID(), nil
}

func getGrandParentFQName(fqName []string) []string {
	if len(fqName) < 2 {
		return nil
	}
	return fqName[:len(fqName)-2]
}

func getPortAndID(
	ctx context.Context, rp RequestParameters, fIP *models.FloatingIP,
) (port *models.VirtualMachineInterface, id string) {
	for _, ref := range fIP.GetVirtualMachineInterfaceRefs() {
		response, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
			ID: ref.GetUUID(),
		})
		if err != nil {
			continue
		}
		//TODO: investigate vmis with service_interface_type "right"
		// in vnc_openstack if vmi has service_interface_type set to "right" then the port_id is set to None
		// even though port is still used to set status and router_id properties.
		// original comment:
		//# In case of floating ip on the Virtual-ip, svc-monitor will
		//# link floating ip to "right" interface of service VMs
		//# launched by ha-proxy service instance. Skip them
		port = response.GetVirtualMachineInterface()
		if props := port.GetVirtualMachineInterfaceProperties(); props != nil {
			if props.GetServiceInterfaceType() == serviceInterfaceTypeRight {
				continue
			}
		}
		id = ref.GetUUID()
	}
	return port, id
}

func getStatus(port *models.VirtualMachineInterface) string {
	if port == nil {
		return netStatusDown
	}
	return netStatusActive
}
