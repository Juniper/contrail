package logic

import (
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models"
)

func makePortResponse(
	vn *models.VirtualNetwork, vmi *models.VirtualMachineInterface, iips []*models.InstanceIP,
) *PortResponse {
	pr := &PortResponse{
		ID:                  vmi.GetUUID(),
		NetworkID:           vn.GetUUID(),
		PortSecurityEnabled: vmi.GetPortSecurityEnabled(),
		DeviceOwner:         vmi.GetVirtualMachineInterfaceDeviceOwner(),
	}

	pr.setName(vmi)
	pr.setTenantID(vmi)
	pr.setMacAddress(vmi)
	pr.setBindings(vmi)
	//TODO handle dhcp options
	//TODO handle allowed address pairs
	pr.setFixedIP(iips)
	pr.setSecurityGroups(vmi)
	pr.setDeviceID(vmi)
	pr.setStatus()
	pr.Description = vmi.GetIDPerms().GetDescription()
	pr.setTimeStamps(vmi.GetIDPerms())
	//TODO make contrail extensions customizable
	pr.setContrailExtensions(vmi)

	return pr
}

func (pr *PortResponse) setSecurityGroups(vmi *models.VirtualMachineInterface) {
	for _, scRef := range vmi.GetSecurityGroupRefs() {
		pr.SecurityGroups = append(pr.SecurityGroups, scRef.GetUUID())
	}
}

func (pr *PortResponse) setDeviceID(vmi *models.VirtualMachineInterface) {
	switch {
	case len(vmi.GetLogicalRouterBackRefs()) > 0:
		pr.DeviceID = vmi.GetLogicalRouterBackRefs()[0].GetUUID()
		return
	case vmi.GetParentType() == models.KindVirtualMachine:
		// Get parent name
		pr.DeviceID = vmi.FQName[len(vmi.FQName)-2]
		return
	case len(vmi.GetVirtualMachineRefs()) > 0:
		// Handle neutron router gateway interface
		to := vmi.VirtualMachineRefs[0].GetTo()
		pr.DeviceID = to[len(to)-1]
		return
	}
}

func (pr *PortResponse) setMacAddress(vmi *models.VirtualMachineInterface) {
	macAddresses := vmi.GetVirtualMachineInterfaceMacAddresses().GetMacAddress()
	if len(macAddresses) != 0 {
		pr.MacAddress = macAddresses[0]
	}
}

func (pr *PortResponse) setBindings(vmi *models.VirtualMachineInterface) {
	for _, kvp := range vmi.GetVirtualMachineInterfaceBindings().GetKeyValuePair() {
		switch kvp.GetKey() {
		case "host_id":
			pr.BindingHostID = kvp.GetValue()
		case "vnic_type":
			pr.BindingVnicType = kvp.GetValue()
		case "vif_type":
			pr.BindingVifType = kvp.GetValue()
		default:
			logrus.Warningf("Unsupported vmi binding: %+v", kvp)
		}
	}

	//TODO load vif details from VMI
	pr.BindingVifDetails = BindingVifDetails{
		PortFilter: true,
	}

	// Set defaults
	if pr.BindingVifType == "" {
		pr.BindingVifType = "vrouter"
	}

	if pr.BindingVnicType == "" {
		pr.BindingVnicType = "normal"
	}
}

func (pr *PortResponse) setTimeStamps(idPerms *models.IdPermsType) {
	pr.CreatedAt = idPerms.GetCreated()
	pr.UpdatedAt = idPerms.GetLastModified()
}

func (pr *PortResponse) setStatus() {
	if pr.DeviceID != "" {
		pr.Status = "ACTIVE"
	} else {
		pr.Status = "DOWN"
	}
}

func (pr *PortResponse) setFixedIP(iips []*models.InstanceIP) {
	for _, iip := range iips {
		//TODO handle contrail extensions
		if iip.GetInstanceIPSecondary() {
			continue
		}

		if iip.GetServiceInstanceIP() {
			continue
		}

		if iip.GetServiceHealthCheckIP() {
			continue
		}

		pr.FixedIps = append(pr.FixedIps, &FixedIp{
			IPAddress: iip.GetInstanceIPAddress(),
			SubnetID:  iip.GetSubnetUUID(),
		})
	}
}

func (pr *PortResponse) setContrailExtensions(vmi *models.VirtualMachineInterface) {
	pr.FQName = vmi.GetFQName()
}

func (pr *PortResponse) setName(vmi *models.VirtualMachineInterface) {
	if len(vmi.DisplayName) != 0 {
		pr.Name = vmi.DisplayName
	} else if len(vmi.GetFQName()) >= 1 {
		pr.Name = vmi.FQName[len(vmi.FQName)-1]
	}
}

func (pr *PortResponse) setTenantID(vmi *models.VirtualMachineInterface) {
	pr.TenantID = vmi.GetPerms2().GetOwner()
}
