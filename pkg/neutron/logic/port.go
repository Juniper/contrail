package logic

import (
	"context"
	"fmt"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	uuid "github.com/satori/go.uuid"
)

const (
	sizeIPv6AddressBits = 128
)

// Create handles port create request.
func (port *Port) Create(ctx context.Context, rp RequestParameters) (Response, error) {
	if port.ID == "" {
		port.ID = uuid.NewV4().String()
	}

	if err := port.checkIfMacAddressIsAvailable(ctx, rp); err != nil {
		return nil, err
	}

	vn, err := port.getVirtualNetwork(ctx, rp)
	if err != nil {
		return nil, err
	}

	vmi, err := port.createVirtualMachineInterface(ctx, rp, vn)
	if err != nil {
		return nil, err
	}

	iip, err := port.allocateIPAddress(ctx, rp, vn, vmi)
	if err != nil {
		return nil, err
	}

	// TODO:
	// create interface route table for the port if
	// subnet has a host route for this port ip.

	return makePortResponse(vn, vmi, []*models.InstanceIP{iip}), nil
}

// Update handles port update request.
func (port *Port) Update(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	vmi, vn, err := port.readVNCPort(ctx, rp, id)
	if err != nil {
		return nil, newNeutronError(portNotFound, errorFields{
			"port_id": id,
			"msg":     err.Error(),
		})
	}

	fm := types.FieldMask{}

	if err = port.handleUpdateOfFields(ctx, rp, vmi, vn, &fm); err != nil {
		return nil, err
	}

	if err = port.handleUpdateOfReferences(ctx, rp, vmi, vn, &fm); err != nil {
		return nil, err
	}

	if _, err = rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
		FieldMask:               fm,
	}); err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      fmt.Sprintf("failed to update virtual machine interface: %s", err.Error()),
		})
	}

	return makePortResponse(vn, vmi, vmi.GetInstanceIPBackRefs()), nil
}

// Delete handles port delete requests.
func (port *Port) Delete(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	vmi, _, err := port.readVNCPort(ctx, rp, id)
	if err != nil {
		return nil, newNeutronError(portNotFound, errorFields{
			"port_id": id,
			"msg":     err.Error(),
		})
	}

	if len(vmi.GetLogicalRouterBackRefs()) > 0 {
		return nil, newNeutronError(l3PortInUse, errorFields{
			"port_id":      id,
			"device_owner": "network:router_interface",
		})
	}

	if err = releaseIPAddresses(ctx, rp, vmi); err != nil {
		return nil, err
	}
	//TODO disassociate any floating IP used by instance

	if _, err = rp.WriteService.DeleteVirtualMachineInterface(ctx, &services.DeleteVirtualMachineInterfaceRequest{
		ID: vmi.GetUUID(),
	}); err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	vmID := port.getAsssociatedVirtualMachineID(vmi)

	if err = removeVirtualMachineIfNoPortsLeft(ctx, rp, vmID); err != nil {
		return nil, err
	}

	//TODO delete any interface route table associated with the port to handle
	// subnet host route Neutron extension, un-reference others

	return &PortResponse{}, nil
}

func releaseIPAddresses(ctx context.Context, rp RequestParameters, vmi *models.VirtualMachineInterface) error {
	for _, iip := range vmi.GetInstanceIPBackRefs() {
		// TODO handle shared ip case

		response, err := rp.ReadService.ListInstanceIP(ctx, &services.ListInstanceIPRequest{
			Spec: &baseservices.ListSpec{
				ObjectUUIDs: []string{iip.GetUUID()},
			},
		})
		if err != nil {
			return err
		}

		if len(response.InstanceIPs) == 0 {
			// shared ip can be removed by svc monitor so neutron does not need to do that
			continue
		}

		_, err = rp.WriteService.DeleteInstanceIP(ctx, &services.DeleteInstanceIPRequest{
			ID: iip.GetUUID(),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func removeVirtualMachineIfNoPortsLeft(ctx context.Context, rp RequestParameters, id string) error {
	if id == "" {
		return nil
	}

	vmRes, err := rp.ReadService.GetVirtualMachine(ctx, &services.GetVirtualMachineRequest{ID: id})
	if err != nil {
		return err
	}

	vm := vmRes.GetVirtualMachine()
	if len(vm.GetVirtualMachineInterfaceBackRefs()) != 0 {
		return nil
	}

	return deleteVM(ctx, rp, id)
}

func deleteVM(ctx context.Context, rp RequestParameters, id string) error {
	_, err := rp.WriteService.DeleteVirtualMachine(ctx, &services.DeleteVirtualMachineRequest{
		ID: id,
	})
	if err != nil {
		return newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}
	return nil
}

// Read handles port read requests.
func (port *Port) Read(ctx context.Context, rp RequestParameters, id string) (Response, error) {
	vmi, vn, err := port.readVNCPort(ctx, rp, id)
	if err != nil {
		return nil, newNeutronError(portNotFound, errorFields{
			"port_id": id,
			"msg":     err.Error(),
		})
	}

	return makePortResponse(vn, vmi, vmi.GetInstanceIPBackRefs()), nil
}

// ReadAll handles port read all requests.
func (port *Port) ReadAll(
	ctx context.Context, rp RequestParameters, filters Filters, fields Fields,
) (Response, error) {
	// TODO implement rest of ReadAll logic
	ps := []*PortResponse{}
	if filters.HaveKeys("device_id") {
		deviceUUIDs := filters["device_id"]
		if len(deviceUUIDs) != 1 {
			return ps, nil
		}

		idToFQNameRes, err := rp.IDToFQNameService.IDToFQName(ctx, &services.IDToFQNameRequest{
			UUID: deviceUUIDs[0],
		})

		if errutil.IsNotFound(err) {
			return ps, nil
		}

		if err != nil {
			return nil, newNeutronError(badRequest, errorFields{
				"resource": "port",
				"msg":      err.Error(),
			})
		}

		// TODO handle another resources associated with port using device_id field in filters
		if idToFQNameRes.GetType() == models.KindVirtualMachine {
			return port.readPortsAssociatedWithVM(ctx, rp, filters, deviceUUIDs[0])
		}
	}
	return ps, nil
}

func (port *Port) handleUpdateOfFields(
	ctx context.Context,
	rp RequestParameters,
	vmi *models.VirtualMachineInterface,
	vn *models.VirtualNetwork,
	fm *types.FieldMask,
) error {
	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(PortFieldName)) {
		vmi.DisplayName = port.Name
		basemodels.FieldMaskAppend(fm, models.VirtualMachineInterfaceFieldDisplayName)
	}

	if port.setBindings(vmi) {
		basemodels.FieldMaskAppend(fm, models.VirtualMachineInterfaceFieldVirtualMachineInterfaceBindings)
	}

	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(PortFieldMacAddress)) {
		// TODO Verify if mac address change allowed
		vmi.VirtualMachineInterfaceMacAddresses = &models.MacAddressesType{
			MacAddress: []string{port.MacAddress},
		}
		basemodels.FieldMaskAppend(fm, models.VirtualMachineInterfaceFieldVirtualMachineInterfaceMacAddresses)
	}

	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(PortFieldPortSecurityEnabled)) {
		vmi.PortSecurityEnabled = port.PortSecurityEnabled
		basemodels.FieldMaskAppend(fm, models.VirtualMachineInterfaceFieldPortSecurityEnabled)
	}

	if err := port.handleAllowedAddressPairs(rp, vmi, fm); err != nil {
		return err
	}

	if err := port.checkFixedIPs(ctx, rp, vmi, true); err != nil {
		return err
	}

	//TODO id perms update (???)

	return nil
}

func (port *Port) handleUpdateOfReferences(
	ctx context.Context,
	rp RequestParameters,
	vmi *models.VirtualMachineInterface,
	vn *models.VirtualNetwork,
	fm *types.FieldMask,
) error {
	if err := port.handleDeviceUpdate(ctx, rp, vmi, fm); err != nil {
		return errors.Wrap(err, "failed to handle device update")
	}

	if err := port.handleSecurityGroupUpdate(ctx, rp, vmi, vn, fm); err != nil {
		return err
	}
	return nil
}

func (port *Port) handleAllowedAddressPairs(
	rp RequestParameters,
	vmi *models.VirtualMachineInterface,
	fm *types.FieldMask,
) error {
	if !basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(PortFieldAllowedAddressPairs)) {
		return nil
	}
	pairs := models.AllowedAddressPairs{}
	for i, pair := range port.AllowedAddressPairs {
		addrInfo, err := decodeIP(pair.IPAddress)
		if err != nil {
			return newNeutronError(badRequest, errorFields{
				"resource": "port",
				"msg":      "Invalid address pair argument",
				"pairIdx":  i,
				"address":  pair.IPAddress,
				"err":      err.Error(),
			})
		}
		subnet := models.SubnetType{IPPrefix: addrInfo.IP}
		if addrInfo.ver == ipV4 {
			subnet.IPPrefixLen = addrInfo.prefixLen
		} else {
			subnet.IPPrefixLen = sizeIPv6AddressBits
		}
		pairs.AllowedAddressPair = append(pairs.AllowedAddressPair, &models.AllowedAddressPair{
			AddressMode: "active-standby",
			Mac:         pair.MacAddress,
			IP:          &subnet,
		})
	}
	basemodels.FieldMaskAppend(fm, models.VirtualMachineInterfaceFieldVirtualMachineInterfaceAllowedAddressPairs)
	vmi.VirtualMachineInterfaceAllowedAddressPairs = &pairs
	return nil
}

func (port *Port) checkFixedIPs(
	ctx context.Context,
	rp RequestParameters,
	vmi *models.VirtualMachineInterface,
	isUpdate bool,
) error {
	if !basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(PortFieldFixedIps)) {
		return nil
	}
	if !basemodels.FieldMaskContains(&rp.FieldMask,
		buildDataResourcePath(PortFieldFixedIps, FixedIpFieldIPAddress)) {
		return nil
	}

	IPs := map[string]struct{}{}
	for _, ip := range vmi.GetInstanceIPBackRefs() {
		IPs[ip.GetInstanceIPAddress()] = struct{}{}
	}
	netID, err := port.getNetworkID(ctx, rp)
	if err != nil {
		return err
	}
	instanceIPs := new([]*models.InstanceIP) // Will be initialized on-demand
	for _, fxip := range port.FixedIps {
		if _, ok := IPs[fxip.IPAddress]; ok {
			continue
		}
		if isUpdate {
			return newNeutronError(badRequest, errorFields{
				"resource": "port",
				"msg":      "Fixed IP cannot be updated on a port",
			})
		}
		if err := port.checkUnusedInstanceIP(ctx, rp, instanceIPs, fxip, netID); err != nil {
			return err
		}
	}
	return nil
}

func (port *Port) checkUnusedInstanceIP(
	ctx context.Context,
	rp RequestParameters,
	instanceIPs *[]*models.InstanceIP,
	fxip *FixedIp,
	netID string,
) error {
	var err error
	if instanceIPs == nil {
		if instanceIPs, err = port.listInstanceIPForNetwork(ctx, rp, netID, []string{
			"instance_ip_address",
		}); err != nil {
			return err
		}
	}
	for _, instanceIP := range *instanceIPs {
		if fxip.IPAddress == instanceIP.GetInstanceIPAddress() {
			return newNeutronError(ipAddressInUse, errorFields{
				"net_id":     netID,
				"ip_address": fxip.IPAddress,
			})

		}
	}
	return nil
}

func (port *Port) listInstanceIPForNetwork(
	ctx context.Context,
	rp RequestParameters,
	netID string,
	fields []string,
) (*[]*models.InstanceIP, error) {
	list, err := rp.ReadService.ListInstanceIP(ctx, &services.ListInstanceIPRequest{
		Spec: &baseservices.ListSpec{
			BackRefUUIDs: []string{netID},
			Fields:       fields,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error reading InstanceIP with Network backref UUID: %v", netID)
	}
	return &list.InstanceIPs, nil
}

func (port *Port) getNetworkID(ctx context.Context, rp RequestParameters) (string, error) {
	netID := port.NetworkID // Try net id from request, if none read it
	if !basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(PortFieldNetworkID)) {
		net, err := port.getVirtualNetwork(ctx, rp)
		if err != nil {
			return "", err
		}
		netID = net.GetUUID()
	}
	return netID, nil
}

func (port *Port) handleSecurityGroupUpdate(
	ctx context.Context,
	rp RequestParameters,
	vmi *models.VirtualMachineInterface,
	vn *models.VirtualNetwork,
	fm *types.FieldMask,
) error {
	if !basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(PortFieldSecurityGroups)) {
		return nil
	}

	if err := port.setPortSecurity(ctx, rp, vmi, vn); err != nil {
		return err
	}
	basemodels.FieldMaskAppend(fm, models.VirtualMachineInterfaceFieldSecurityGroupRefs)

	return nil
}

func (port *Port) handleDeviceUpdate(
	ctx context.Context,
	rp RequestParameters,
	vmi *models.VirtualMachineInterface,
	fm *types.FieldMask,
) error {
	if shouldChangeDevice(port.DeviceOwner, &rp.FieldMask) {
		if err := port.setVMInstance(ctx, rp, vmi); err != nil {
			return errors.Wrapf(err, "failed to set virtual machine (device) to %q", port.DeviceID)
		}
		basemodels.FieldMaskAppend(fm, models.VirtualMachineInterfaceFieldVirtualMachineRefs)
	}

	if basemodels.FieldMaskContains(&rp.FieldMask, buildDataResourcePath(PortFieldDeviceOwner)) {
		vmi.VirtualMachineInterfaceDeviceOwner = port.DeviceOwner
		basemodels.FieldMaskAppend(fm, models.VirtualMachineInterfaceFieldVirtualMachineInterfaceDeviceOwner)
	}
	return nil
}

func shouldChangeDevice(deviceOwner string, fm *types.FieldMask) bool {
	return basemodels.FieldMaskContains(fm, buildDataResourcePath(PortFieldDeviceID)) &&
		deviceOwner != "network:router_interface" &&
		deviceOwner != "network:router_gateway"
}

// updateMacAddress modify mac address only for baremetal deployments or when port is not attached to any VM
// Here nil is returned - same as original config node
// However this silently discards mac address update request if not allowed
func (port *Port) getAsssociatedVirtualMachineID(vmi *models.VirtualMachineInterface) string {
	if vmi.GetParentType() == models.KindVirtualMachine {
		return vmi.GetParentUUID()
	}

	if len(vmi.GetVirtualMachineRefs()) > 0 {
		return vmi.VirtualMachineRefs[0].GetUUID()
	}

	return ""
}

func (port *Port) readPortsAssociatedWithVM(
	ctx context.Context, rp RequestParameters, filters Filters, vmUUID string,
) ([]*PortResponse, error) {
	vmRes, err := rp.ReadService.GetVirtualMachine(ctx, &services.GetVirtualMachineRequest{
		ID: vmUUID,
	})
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	ps := []*PortResponse{}
	vmiBackRefs := vmRes.GetVirtualMachine().GetVirtualMachineInterfaceBackRefs()
	for _, vmiRef := range vmiBackRefs {
		var vmi *models.VirtualMachineInterface
		var vn *models.VirtualNetwork
		vmi, vn, err = port.readVNCPort(ctx, rp, vmiRef.GetUUID())
		if err != nil {
			return nil, newNeutronError(portNotFound, errorFields{
				"port_id": vmiRef.GetUUID(),
				"msg":     err.Error(),
			})
		}

		ps = append(ps, makePortResponse(vn, vmi, vmi.GetInstanceIPBackRefs()))
	}
	return ps, nil
}

func (port *Port) getAssociatedVirtualNetwork(ctx context.Context, rp RequestParameters,
	vmi *models.VirtualMachineInterface,
) (*models.VirtualNetwork, error) {
	vnRefs := vmi.GetVirtualNetworkRefs()
	if len(vnRefs) == 0 {
		return nil, nil
	}
	vnRes, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: vnRefs[0].GetUUID(),
	})
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}
	return vnRes.GetVirtualNetwork(), nil
}

func (port *Port) readVNCPort(
	ctx context.Context, rp RequestParameters, id string,
) (*models.VirtualMachineInterface, *models.VirtualNetwork, error) {
	vmiRes, err := rp.ReadService.GetVirtualMachineInterface(ctx, &services.GetVirtualMachineInterfaceRequest{
		ID: id,
	})
	if err != nil {
		return nil, nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	vmi := vmiRes.GetVirtualMachineInterface()
	vn, err := port.getAssociatedVirtualNetwork(ctx, rp, vmi)
	if err != nil {
		return nil, nil, err
	}

	return vmi, vn, nil
}

func (port *Port) allocateIPAddress(
	ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork, vmi *models.VirtualMachineInterface,
) (*models.InstanceIP, error) {

	//TODO handle fixed_ips
	if len(vn.NetworkIpamRefs) == 0 {
		return nil, nil
	}

	return port.createInstanceIP(ctx, rp, vn, vmi, "", "")
}

func (port *Port) createInstanceIP(
	ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork, vmi *models.VirtualMachineInterface,
	subnetUUID string, ipAddress string,
) (*models.InstanceIP, error) {

	ipUUID := uuid.NewV4().String()
	iip := &models.InstanceIP{
		Name:              ipUUID,
		UUID:              ipUUID,
		SubnetUUID:        subnetUUID,
		InstanceIPAddress: ipAddress,
		Perms2: &models.PermType2{
			Owner: vmi.GetPerms2().GetOwner(),
		},
	}

	iip.AddVirtualMachineInterfaceRef(&models.InstanceIPVirtualMachineInterfaceRef{
		UUID: vmi.UUID,
	})

	iip.AddVirtualNetworkRef(&models.InstanceIPVirtualNetworkRef{
		UUID: vn.UUID,
	})

	iipRes, err := rp.WriteService.CreateInstanceIP(ctx, &services.CreateInstanceIPRequest{
		InstanceIP: iip,
	})

	if err != nil {
		errorName := ipAddressGenerationFailure
		if errutil.IsBadRequest(err) {
			errorName = badRequest
		}
		return nil, newNeutronError(errorName, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	return iipRes.InstanceIP, nil
}

func (port *Port) getVirtualNetwork(
	ctx context.Context, rp RequestParameters,
) (*models.VirtualNetwork, error) {
	res, err := rp.ReadService.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{
		ID: port.NetworkID,
	})
	if err != nil {
		return nil, newNeutronError(networkNotFound, errorFields{
			"net_id": port.NetworkID,
			"msg":    err.Error(),
		})
	}

	return res.GetVirtualNetwork(), err
}

func (port *Port) setVMInstance(
	ctx context.Context,
	rp RequestParameters,
	vmi *models.VirtualMachineInterface,
) error {
	if port.DeviceID == "" {
		for _, vmRef := range vmi.GetVirtualMachineRefs() {
			vmi.RemoveVirtualMachineRef(vmRef)
			mask := types.FieldMask{}
			basemodels.FieldMaskAppend(&mask, models.VirtualMachineInterfaceFieldVirtualMachineRefs)
			_, err := rp.WriteService.UpdateVirtualMachineInterface(ctx, &services.UpdateVirtualMachineInterfaceRequest{
				VirtualMachineInterface: vmi,
				FieldMask:               mask,
			})
			if err != nil {
				return errors.Wrap(err, "failed to update removed VirtualMachine reference")
			}
			_, err = rp.WriteService.DeleteVirtualMachine(ctx, &services.DeleteVirtualMachineRequest{ID: vmRef.UUID})
			if err != nil {
				return errors.Wrapf(err, "deleting VirtualMachine (as DeviceID) (uuid %v) failed", vmRef.UUID)
			}
		}

		vmi.VirtualMachineRefs = nil
		return nil
	}

	vm, err := port.ensureVMInstanceExists(ctx, rp, vmi.GetPerms2().GetOwner())
	if err != nil {
		return err
	}

	vmi.AddVirtualMachineRef(&models.VirtualMachineInterfaceVirtualMachineRef{
		UUID: vm.GetUUID(),
		To:   vm.GetFQName(),
	})

	return nil
}

func (port *Port) ensureVMInstanceExists(
	ctx context.Context, rp RequestParameters, tenantID string,
) (*models.VirtualMachine, error) {
	// TODO: Handle bare metal
	uuid := parseUUID(port.DeviceID)
	gr, err := rp.ReadService.GetVirtualMachine(ctx, &services.GetVirtualMachineRequest{
		ID: uuid,
	})
	if errutil.IsNotFound(err) {
		rp.Log.WithField("uuid", uuid).Debug("No virtual machine for port - creating new one")
		cr, cErr := rp.WriteService.CreateVirtualMachine(ctx, &services.CreateVirtualMachineRequest{
			VirtualMachine: &models.VirtualMachine{
				UUID: uuid,
				Name: port.DeviceID,
				Perms2: &models.PermType2{
					Owner: tenantID,
				},
				ServerType: "virtual-server", // TODO: Handle bare metal
			},
		})
		if cErr != nil {
			return nil, errors.Wrapf(err, "failed to create new virtual machine (device) %q", uuid)
		}

		return cr.GetVirtualMachine(), nil
	} else if err != nil {
		return nil, errors.Wrapf(err, "failed to get virtual machine (device) %q", uuid)
	}

	return gr.GetVirtualMachine(), nil
}

func parseUUID(rawUUID string) string {
	uuid, err := uuid.FromString(rawUUID)
	if err != nil {
		return ""
	}
	return uuid.String()
}

func (port *Port) setPortSecurity(
	ctx context.Context, rp RequestParameters, vmi *models.VirtualMachineInterface, vn *models.VirtualNetwork,
) error {
	vmi.PortSecurityEnabled = port.PortSecurityEnabled
	if !vmi.PortSecurityEnabled {
		vmi.PortSecurityEnabled = vn.PortSecurityEnabled
	}

	securityGroups, err := port.listSecurityGroups(ctx, rp, port.SecurityGroups, []string{"uuid", "fqname"})
	if err != nil {
		return err
	}

	vmi.SecurityGroupRefs = nil
	for _, sc := range securityGroups {
		vmi.AddSecurityGroupRef(&models.VirtualMachineInterfaceSecurityGroupRef{
			UUID: sc.GetUUID(),
		})
	}

	if len(vmi.SecurityGroupRefs) == 0 && vmi.PortSecurityEnabled {
		// When there is no security group for a port, the internal no_rule group should be used
		vmi.AddSecurityGroupRef(&models.VirtualMachineInterfaceSecurityGroupRef{
			To: sgNoRuleFQName,
		})
	}

	//TODO Handle default security group

	return nil
}

func (port *Port) listSecurityGroups(
	ctx context.Context, rp RequestParameters, uuids []string, fields []string,
) ([]*models.SecurityGroup, error) {
	if len(uuids) == 0 {
		return nil, nil
	}

	res, err := rp.ReadService.ListSecurityGroup(ctx, &services.ListSecurityGroupRequest{
		Spec: &baseservices.ListSpec{
			ObjectUUIDs: uuids,
			Fields:      fields,
		},
	})
	if errutil.IsNotFound(err) {
		// TODO add information which group is missing
		return nil, newNeutronError(securityGroupNotFound, errorFields{
			"device_owner": "network:router_interface",
			"msg":          err.Error(),
		})
	} else if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}
	return res.GetSecurityGroups(), nil
}

func (port *Port) createVirtualMachineInterface(
	ctx context.Context, rp RequestParameters, vn *models.VirtualNetwork,
) (*models.VirtualMachineInterface, error) {

	vmi := &models.VirtualMachineInterface{
		UUID:       port.ID,
		ParentType: models.KindProject,
		ParentUUID: parseUUID(port.TenantID),
		IDPerms: &models.IdPermsType{
			Enable: true,
		},
		Perms2: &models.PermType2{
			Owner: port.TenantID,
		},
	}

	if port.Name == "" {
		vmi.Name = port.ID
	} else {
		vmi.Name = port.Name
		vmi.DisplayName = port.Name
	}

	if port.MacAddress != "" {
		vmi.VirtualMachineInterfaceMacAddresses = &models.MacAddressesType{
			MacAddress: []string{port.MacAddress},
		}
	}

	vmi.AddVirtualNetworkRef(&models.VirtualMachineInterfaceVirtualNetworkRef{
		UUID: vn.GetUUID(),
	})

	if port.DeviceOwner != "network:router_interface" &&
		port.DeviceOwner != "network:router_gateway" && port.DeviceID != "" {
		if err := port.setVMInstance(ctx, rp, vmi); err != nil {
			return nil, err
		}
	}

	vmi.VirtualMachineInterfaceDeviceOwner = port.DeviceOwner
	port.setBindings(vmi)

	if err := port.setPortSecurity(ctx, rp, vmi, vn); err != nil {
		return nil, err
	}

	//TODO Handle allowed address pair
	//TODO Handle fixed ips

	vmiRes, err := rp.WriteService.CreateVirtualMachineInterface(ctx, &services.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: vmi,
	})
	if err != nil {
		return nil, newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	return vmiRes.GetVirtualMachineInterface(), nil
}

func (port *Port) setBinding(vmi *models.VirtualMachineInterface, key string, value string) {
	if vmi.VirtualMachineInterfaceBindings == nil {
		vmi.VirtualMachineInterfaceBindings = models.MakeKeyValuePairs()
	}

	//TODO SetInMap shouldn't return an error
	_ = vmi.GetVirtualMachineInterfaceBindings().SetInMap(&models.KeyValuePair{ // nolint: errcheck
		Key:   key,
		Value: value,
	})
}

func (port *Port) setBindings(vmi *models.VirtualMachineInterface) bool {
	modified := false
	if port.BindingVnicType != "" {
		port.setBinding(vmi, "vnic_type", port.BindingVnicType)
		modified = true
	}
	if port.BindingVifType != "" {
		port.setBinding(vmi, "vif_type", port.BindingVifType)
		modified = true
	}
	if port.BindingHostID != "" {
		port.setBinding(vmi, "host_id", port.BindingHostID)
		modified = true
	}
	return modified
}

func (port *Port) checkIfMacAddressIsAvailable(ctx context.Context, rp RequestParameters) error {
	if port.MacAddress == "" {
		return nil
	}

	vmisRes, err := rp.ReadService.ListVirtualMachineInterface(ctx, &services.ListVirtualMachineInterfaceRequest{
		Spec: &baseservices.ListSpec{
			Filters: []*baseservices.Filter{
				{
					Key:    models.VirtualMachineInterfaceFieldVirtualMachineInterfaceMacAddresses,
					Values: []string{port.MacAddress},
				},
			},
		},
	})

	if err != nil {
		return newNeutronError(badRequest, errorFields{
			"resource": "port",
			"msg":      err.Error(),
		})
	}

	if vmisRes.GetVirtualMachineInterfaceCount() != 0 {
		return newNeutronError(macAddressInUse, errorFields{
			"net_id": port.NetworkID,
			"mac":    port.MacAddress,
		})
	}

	return nil
}
