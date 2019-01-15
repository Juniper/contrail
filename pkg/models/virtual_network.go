package models

import (
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// Virtual network forwarding modes.
const (
	L3Mode   = "l3"
	L2L3Mode = "l2_l3"
)

// TODO: Enums strings should be generated from schema
const (
	UserDefinedSubnetOnly      = "user-defined-subnet-only"
	UserDefinedSubnetPreferred = "user-defined-subnet-preferred"
	FlatSubnetOnly             = "flat-subnet-only"
)

//MakeNeutronCompatible makes this resource data neutron compatible.
func (m *VirtualNetwork) MakeNeutronCompatible() {
	//  neutorn <-> vnc sharing
	if m.Perms2.GlobalAccess == basemodels.PermsRWX {
		m.IsShared = true
	}
	if m.IsShared {
		m.Perms2.GlobalAccess = basemodels.PermsRWX
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

//ValidateMultiPolicyServiceChainConfig checks if multi policy service chain config is valid
//and throws an error when it's not.
func (m *VirtualNetwork) ValidateMultiPolicyServiceChainConfig() error {
	if !m.MultiPolicyServiceChainsEnabled {
		return nil
	}

	err := errutil.ErrorBadRequest(
		"multi policy service chains are not supported, with both import export external route targets")

	if len(m.GetRouteTargetList().GetRouteTarget()) != 0 {
		return err
	}
	for _, importRouteTarget := range m.GetImportRouteTargetList().GetRouteTarget() {
		for _, exportRouteTarget := range m.GetExportRouteTargetList().GetRouteTarget() {
			if importRouteTarget == exportRouteTarget {
				return err
			}
		}
	}
	return nil
}

//ShouldIgnoreAllocation checks if there is ip-fabric or link-local address allocation
func (m *VirtualNetwork) ShouldIgnoreAllocation() bool {
	fqName := m.GetFQName()
	if format.ContainsString(fqName, "ip-fabric") || format.ContainsString(fqName, "__link_local__") {
		return true
	}
	return false
}

// GetSubnetUUIDs returns list of subnetUUIDs for all subnets
func (m *VirtualNetwork) GetSubnetUUIDs() []string {
	var result []string
	for _, subnet := range m.GetIpamSubnets().GetSubnets() {
		result = append(result, subnet.SubnetUUID)
	}

	return result
}

// GetIpamSubnets returns list of subnets
func (m *VirtualNetwork) GetIpamSubnets() *IpamSubnets {
	var subnets []*IpamSubnetType
	// Take attr subnets
	for _, networkIpam := range m.GetNetworkIpamRefs() {
		subnets = append(subnets, networkIpam.GetAttr().GetIpamSubnets()...)
	}
	return &IpamSubnets{
		Subnets: subnets,
	}
}

// GetAddressAllocationMethod returns address allocation method
func (m *VirtualNetwork) GetAddressAllocationMethod() string {
	allocationMethod := UserDefinedSubnetPreferred
	if m.GetAddressAllocationMode() != "" {
		allocationMethod = m.GetAddressAllocationMode()
	}
	return allocationMethod
}

// GetDefaultRoutingInstance returns the default routing instance of VN or nil if it doesn't exist
func (m *VirtualNetwork) GetDefaultRoutingInstance() *RoutingInstance {
	for _, ri := range m.RoutingInstances {
		if ri.GetRoutingInstanceIsDefault() {
			return ri
		}
	}

	return nil
}

// HasNetworkBasedAllocationMethod checks if allocation method is userdefined or flat subnet only
func (m *VirtualNetwork) HasNetworkBasedAllocationMethod() bool {
	return m.GetAddressAllocationMethod() == UserDefinedSubnetOnly || m.GetAddressAllocationMethod() == FlatSubnetOnly
}

// MakeDefaultRoutingInstance returns the default routing instance for the network.
func (m *VirtualNetwork) MakeDefaultRoutingInstance() *RoutingInstance {
	return &RoutingInstance{
		Name:                      m.Name,
		FQName:                    m.DefaultRoutingInstanceFQName(),
		ParentUUID:                m.UUID,
		RoutingInstanceIsDefault:  true,
		RoutingInstanceFabricSnat: m.FabricSnat,
		RouteTargetRefs:           m.MakeRouteTargetRefs(),
	}
}

// DefaultRoutingInstanceFQName returns the FQName of the network's default RoutingInstance.
func (m *VirtualNetwork) DefaultRoutingInstanceFQName() []string {
	return basemodels.ChildFQName(m.FQName, m.FQName[len(m.FQName)-1])
}

// MakeRouteTargetRefs returns refs to RouteTarget's from import, export and both lists.
func (m *VirtualNetwork) MakeRouteTargetRefs() []*RoutingInstanceRouteTargetRef {
	return append(
		m.GetImportRouteTargetList().AsRefs(&InstanceTargetType{ImportExport: "import"}),
		m.GetExportRouteTargetList().AsRefs(&InstanceTargetType{ImportExport: "export"})...
	)
}

// AsRefs returns refs with instanceTargetType from a RoutingInstance to route targets in the list.
func (m *RouteTargetList) AsRefs(instanceTargetType *InstanceTargetType) (refs []*RoutingInstanceRouteTargetRef) {
	for _, rt := range m.GetRouteTarget() {
		refs = append(refs, &RoutingInstanceRouteTargetRef{
			To:   []string{rt},
			Attr: instanceTargetType,
		})
	}
	return refs
}

// IsLinkLocal returns true if virtual network FQName fits Link Local
func (m *VirtualNetwork) IsLinkLocal() bool {
	fq := []string{"default-domain", "default-project", "__link_local__"}
	return basemodels.FQNameEquals(fq, m.GetFQName())
}
