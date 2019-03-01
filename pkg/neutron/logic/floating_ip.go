package logic

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
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
			"resource": "floatingip",
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
		FloatingIP: &models.FloatingIP{
			Name:              fip.ID,
			ParentUUID:        fIPP.UUID,
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
		},
	})
	if err != nil {
		return nil, newNeutronError(ipAddressGenerationFailure, errorFields{
			"net_id": fip.FloatingNetworkID,
		})
	}

	readResponse, err := rp.ReadService.GetFloatingIP(ctx, &services.GetFloatingIPRequest{
		ID: response.GetFloatingIP().GetUUID(),
	})

	if err != nil {
		return nil, err
	}

	return floatingipVncToNeutron(ctx, rp, readResponse.GetFloatingIP())
}

// Read default implementation
func (fip *Floatingip) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	readResponse, err := rp.ReadService.GetFloatingIP(ctx, &services.GetFloatingIPRequest{
		ID: id,
	})
	if err != nil {
		return nil, newNeutronError("FloatingIPNotFound", errorFields{
			"floatingip_id": id,
		})
	}
	return floatingipVncToNeutron(ctx, rp, readResponse.GetFloatingIP())
}

// ReadAll logic
func (fip *Floatingip) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	// TODO implement ReadAll logic
	return []FloatingipResponse{}, nil
}

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
		FieldMask: fip.createFieldMaskForUpdate(),
	})
	if err != nil {
		return nil, err
	}
	getResponse, err := rp.ReadService.GetFloatingIP(ctx, &services.GetFloatingIPRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return floatingipVncToNeutron(ctx, rp, getResponse.GetFloatingIP())
}

func (fip *Floatingip) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	_, err := rp.WriteService.DeleteFloatingIP(ctx, &services.DeleteFloatingIPRequest{
		ID: id,
	})
	if err != nil {
		switch {
		case errutil.IsNotFound(err):
			return nil, newNeutronError("FloatingIPNotFound", errorFields{
				"floatingip_id": id,
			})
		default:
			return nil, err
		}
	}
	return nil, nil
}

func (fip *Floatingip) createFieldMaskForUpdate() types.FieldMask {
	var paths []string
	if fip.FloatingIPAddress != "" {
		paths = append(paths, models.FloatingIPFieldFloatingIPAddress)
	}
	if fip.PortID != "" {
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
	if len(fIPPListResponse.FloatingIPPools) == 0 {
		return nil, errors.Errorf("no floating-ip-pool with parent_uuid: '%s'", fip.FloatingNetworkID)
	}
	fIPPResponse, err := rp.ReadService.GetFloatingIPPool(ctx, &services.GetFloatingIPPoolRequest{
		ID: fIPPListResponse.FloatingIPPools[0].UUID,
	})
	if err != nil {
		return nil, err
	}
	return fIPPResponse.FloatingIPPool, nil
}

func getProject(ctx context.Context, rp RequestParameters) (*models.Project, error) {
	projectID, err := neutronIDToVncUUID(rp.RequestContext.TenantID)
	if err != nil {
		return nil, err
	}

	projectResponse, err := rp.ReadService.GetProject(
		ctx,
		&services.GetProjectRequest{
			ID: projectID,
		},
	)
	if err != nil {
		return nil, err

	}

	return projectResponse.GetProject(), nil
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
			"resource": "floatingip",
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
) (string, error) {
	if fip.FixedIPAddress != "" && fip.PortID == "" {
		return "", newNeutronError(badRequest, errorFields{
			"resource": "floatingip",
			"msg":      fmt.Sprint("fixed_ip_address cannot be specified without a port_id"),
		})
	}
	if fip.FixedIPAddress != "" {
		//TODO: _check_port_fip_assoc
		return fip.FixedIPAddress, nil
	}
	if fip.PortID == "" {
		return "", nil
	}
	vmiResponse, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
		ID: fip.PortID,
		Fields: Fields{
			"instance_ip_back_refs",
			"floating_ip_back_refs",
		},
	})
	if err != nil {
		return "", err
	}
	iipBackRefs := vmiResponse.GetVirtualMachineInterface().GetInstanceIPBackRefs()
	if len(iipBackRefs) > 1 {
		return "", newNeutronError(badRequest, errorFields{
			"resource": "floatingip",
			"msg": fmt.Sprintf("Port %s has multiple fixed IP addresses. "+
				"Must provide a specific IP address when assigning a floating IP",
				vmiResponse.GetVirtualMachineInterface().GetUUID()),
		})
	}
	if len(iipBackRefs) == 1 {
		//TODO: _check_port_fip_assoc
		return iipBackRefs[0].GetInstanceIPAddress(), nil
	}
	return "", nil
}

// review: help with name
func floatingipVncToNeutron(ctx context.Context, rp RequestParameters, fIP *models.FloatingIP) (*FloatingipResponse, error) {
	port, id := getPortAndID(ctx, rp, fIP)

	routerID, err := getRouterID(port)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get router id")
	}
	netID, err := getFloatingNetworkID(ctx, rp, fIP)
	if err != nil {
		return nil, err
	}
	return &FloatingipResponse{
		ID:                fIP.UUID,
		TenantID:          getTenantID(fIP),
		FloatingIPAddress: fIP.FloatingIPAddress,
		FloatingNetworkID: netID,
		FixedIPAddress:    fIP.FloatingIPFixedIPAddress,
		CreatedAt:         fIP.GetIDPerms().GetCreated(),
		UpdatedAt:         fIP.GetIDPerms().GetLastModified(),
		Status:            getStatus(port),
		Description:       fIP.GetIDPerms().GetDescription(),
		PortID:            id,
		RouterID:          routerID,
	}, nil
}

func getTenantID(fIP *models.FloatingIP) string {
	if len(fIP.ProjectRefs) > 0 {
		return VncUUIDToNeutronID(fIP.ProjectRefs[0].UUID)
	}
	return ""
}

func getFloatingNetworkID(ctx context.Context, rp RequestParameters, fIP *models.FloatingIP) (string, error) {
	response, err := rp.FQNameToIDService.FQNameToID(ctx, &services.FQNameToIDRequest{
		FQName: fIP.FQName[:len(fIP.FQName)-2],
		Type:   models.KindVirtualNetwork,
	})
	if err != nil {
		return "", err
	}
	return response.GetUUID(), nil
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

func getRouterID(
	port *models.VirtualMachineInterface,
) (string, error) {
	if port == nil {
		return "", nil
	}
	if len(port.LogicalRouterBackRefs) == 0 {
		return "", nil
	}
	return port.LogicalRouterBackRefs[0].GetUUID(), nil
}
