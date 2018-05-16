package models

const (
	l3Mode   = "l3"
	l2l3Mode = "l2_l3"
)

//MakeNeutronCompatible makes this resource data neutron compatible.
func (m *VirtualNetwork) MakeNeutronCompatible() {
	//  neutorn <-> vnc sharing
	if m.Perms2.GlobalAccess == PermsRWX {
		m.IsShared = true
	}
	if m.IsShared == true {
		m.Perms2.GlobalAccess = PermsRWX
	}
}

//HasVirtualNetworkNetworkID check if the resource has virtual network ID.
func (m *VirtualNetwork) HasVirtualNetworkNetworkID() bool {
	return m.VirtualNetworkNetworkID != 0
}

//IsSupportingAnyVPNType check if the resource has virtual network ID.
func (m *VirtualNetwork) IsSupportingAnyVPNType() bool {
	return m.GetVirtualNetworkProperties().GetForwardingMode() == l2l3Mode
}

//IsL3Mode checks if this network is l3 mode
func (m *VirtualNetwork) IsL3Mode() bool {
	return m.GetVirtualNetworkProperties().GetForwardingMode() == l3Mode
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
