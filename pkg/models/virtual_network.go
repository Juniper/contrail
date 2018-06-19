package models

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
