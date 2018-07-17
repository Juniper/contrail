package models

// Virtual network forwarding modes.
const (
	L3Mode   = "l3"
	L2L3Mode = "l2_l3"
)

//MakeNeutronCompatible makes this resource data neutron compatible.
func (m *VirtualNetwork) MakeNeutronCompatible() {
	//  neutorn <-> vnc sharing
	if m.Perms2.GlobalAccess == PermsRWX {
		m.IsShared = true
	}
	if m.IsShared {
		m.Perms2.GlobalAccess = PermsRWX
	}
}

//HasVirtualNetworkNetworkID check if the resource has virtual network ID.
func (m *VirtualNetwork) HasVirtualNetworkNetworkID() bool {
	return m.VirtualNetworkNetworkID != 0
}

//IsSupportingAnyVPNType checks if this network is l2 and l3 mode
func (m *VirtualNetwork) IsSupportingAnyVPNType() bool {
	return m.GetVirtualNetworkProperties().GetForwardingMode() == L2L3Mode
}

//IsSupportingL3VPNType checks if this network is l3 mode
func (m *VirtualNetwork) IsSupportingL3VPNType() bool {
	return m.GetVirtualNetworkProperties().GetForwardingMode() == L3Mode
}

//IsValidMultiPolicyServiceChainConfig checks if multi policy service chain config is valid or not.
func (m *VirtualNetwork) IsValidMultiPolicyServiceChainConfig() bool {
	if !m.MultiPolicyServiceChainsEnabled {
		return true
	}
	if len(m.GetRouteTargetList().GetRouteTarget()) != 0 {
		return false
	}
	for _, importRouteTarget := range m.GetImportRouteTargetList().GetRouteTarget() {
		for _, exportRouteTarget := range m.GetExportRouteTargetList().GetRouteTarget() {
			if importRouteTarget == exportRouteTarget {
				return false
			}
		}
	}
	return true
}

// GetSubnetUUIDs returns list of subnetUUIDs for all subnets
func (m *VirtualNetwork) GetSubnetUUIDs() ([]string, error) {
	subnets, err := m.GetSubnets()
	if err != nil {
		return nil, err
	}

	var result []string
	for _, subnet := range subnets {
		result = append(result, subnet.SubnetUUID)
	}

	return result, nil
}

// GetSubnets returns list of subnets
func (m *VirtualNetwork) GetSubnets() ([]*IpamSubnetType, error) {
	var result []*IpamSubnetType
	// Take attr subnets
	for _, networkIpam := range m.GetNetworkIpamRefs() {
		result = append(result, networkIpam.GetAttr().GetIpamSubnets()...)
	}
	return result, nil
}

// GetAddressAllocationMethod returns address allocation method
func (m *VirtualNetwork) GetAddressAllocationMethod() string {
	// TODO: Enums strings should be generated from schema
	allocationMethod := "user-defined-subnet-preferred"
	if m.GetAddressAllocationMode() != "" {
		allocationMethod = m.GetAddressAllocationMode()
	}
	return allocationMethod
}
