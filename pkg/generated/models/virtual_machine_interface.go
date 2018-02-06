package models

// VirtualMachineInterface

import "encoding/json"

// VirtualMachineInterface
type VirtualMachineInterface struct {
	EcmpHashingIncludeFields                   *EcmpHashingIncludeFields              `json:"ecmp_hashing_include_fields,omitempty"`
	VirtualMachineInterfaceAllowedAddressPairs *AllowedAddressPairs                   `json:"virtual_machine_interface_allowed_address_pairs,omitempty"`
	Annotations                                *KeyValuePairs                         `json:"annotations,omitempty"`
	ParentType                                 string                                 `json:"parent_type,omitempty"`
	DisplayName                                string                                 `json:"display_name,omitempty"`
	VirtualMachineInterfaceHostRoutes          *RouteTableType                        `json:"virtual_machine_interface_host_routes,omitempty"`
	VirtualMachineInterfaceDHCPOptionList      *DhcpOptionsListType                   `json:"virtual_machine_interface_dhcp_option_list,omitempty"`
	VirtualMachineInterfaceDisablePolicy       bool                                   `json:"virtual_machine_interface_disable_policy"`
	VirtualMachineInterfaceFatFlowProtocols    *FatFlowProtocols                      `json:"virtual_machine_interface_fat_flow_protocols,omitempty"`
	VRFAssignTable                             *VrfAssignTableType                    `json:"vrf_assign_table,omitempty"`
	UUID                                       string                                 `json:"uuid,omitempty"`
	ParentUUID                                 string                                 `json:"parent_uuid,omitempty"`
	VirtualMachineInterfaceProperties          *VirtualMachineInterfacePropertiesType `json:"virtual_machine_interface_properties,omitempty"`
	IDPerms                                    *IdPermsType                           `json:"id_perms,omitempty"`
	VirtualMachineInterfaceMacAddresses        *MacAddressesType                      `json:"virtual_machine_interface_mac_addresses,omitempty"`
	VirtualMachineInterfaceBindings            *KeyValuePairs                         `json:"virtual_machine_interface_bindings,omitempty"`
	VlanTagBasedBridgeDomain                   bool                                   `json:"vlan_tag_based_bridge_domain"`
	VirtualMachineInterfaceDeviceOwner         string                                 `json:"virtual_machine_interface_device_owner,omitempty"`
	PortSecurityEnabled                        bool                                   `json:"port_security_enabled"`
	Perms2                                     *PermType2                             `json:"perms2,omitempty"`
	FQName                                     []string                               `json:"fq_name,omitempty"`

	ServiceEndpointRefs         []*VirtualMachineInterfaceServiceEndpointRef         `json:"service_endpoint_refs,omitempty"`
	RoutingInstanceRefs         []*VirtualMachineInterfaceRoutingInstanceRef         `json:"routing_instance_refs,omitempty"`
	ServiceHealthCheckRefs      []*VirtualMachineInterfaceServiceHealthCheckRef      `json:"service_health_check_refs,omitempty"`
	SecurityGroupRefs           []*VirtualMachineInterfaceSecurityGroupRef           `json:"security_group_refs,omitempty"`
	VirtualNetworkRefs          []*VirtualMachineInterfaceVirtualNetworkRef          `json:"virtual_network_refs,omitempty"`
	BridgeDomainRefs            []*VirtualMachineInterfaceBridgeDomainRef            `json:"bridge_domain_refs,omitempty"`
	VirtualMachineRefs          []*VirtualMachineInterfaceVirtualMachineRef          `json:"virtual_machine_refs,omitempty"`
	SecurityLoggingObjectRefs   []*VirtualMachineInterfaceSecurityLoggingObjectRef   `json:"security_logging_object_refs,omitempty"`
	VirtualMachineInterfaceRefs []*VirtualMachineInterfaceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	InterfaceRouteTableRefs     []*VirtualMachineInterfaceInterfaceRouteTableRef     `json:"interface_route_table_refs,omitempty"`
	PhysicalInterfaceRefs       []*VirtualMachineInterfacePhysicalInterfaceRef       `json:"physical_interface_refs,omitempty"`
	BGPRouterRefs               []*VirtualMachineInterfaceBGPRouterRef               `json:"bgp_router_refs,omitempty"`
	QosConfigRefs               []*VirtualMachineInterfaceQosConfigRef               `json:"qos_config_refs,omitempty"`
	PortTupleRefs               []*VirtualMachineInterfacePortTupleRef               `json:"port_tuple_refs,omitempty"`
}

// VirtualMachineInterfaceVirtualMachineRef references each other
type VirtualMachineInterfaceVirtualMachineRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceSecurityLoggingObjectRef references each other
type VirtualMachineInterfaceSecurityLoggingObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfacePhysicalInterfaceRef references each other
type VirtualMachineInterfacePhysicalInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceVirtualMachineInterfaceRef references each other
type VirtualMachineInterfaceVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceInterfaceRouteTableRef references each other
type VirtualMachineInterfaceInterfaceRouteTableRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfacePortTupleRef references each other
type VirtualMachineInterfacePortTupleRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceBGPRouterRef references each other
type VirtualMachineInterfaceBGPRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceQosConfigRef references each other
type VirtualMachineInterfaceQosConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceSecurityGroupRef references each other
type VirtualMachineInterfaceSecurityGroupRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceVirtualNetworkRef references each other
type VirtualMachineInterfaceVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceBridgeDomainRef references each other
type VirtualMachineInterfaceBridgeDomainRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *BridgeDomainMembershipType
}

// VirtualMachineInterfaceServiceEndpointRef references each other
type VirtualMachineInterfaceServiceEndpointRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceRoutingInstanceRef references each other
type VirtualMachineInterfaceRoutingInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *PolicyBasedForwardingRuleType
}

// VirtualMachineInterfaceServiceHealthCheckRef references each other
type VirtualMachineInterfaceServiceHealthCheckRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *VirtualMachineInterface) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualMachineInterface makes VirtualMachineInterface
func MakeVirtualMachineInterface() *VirtualMachineInterface {
	return &VirtualMachineInterface{
		//TODO(nati): Apply default
		VirtualMachineInterfaceBindings:    MakeKeyValuePairs(),
		VlanTagBasedBridgeDomain:           false,
		VirtualMachineInterfaceDeviceOwner: "",
		PortSecurityEnabled:                false,
		Perms2:                             MakePermType2(),
		FQName:                             []string{},
		IDPerms:                            MakeIdPermsType(),
		VirtualMachineInterfaceMacAddresses:        MakeMacAddressesType(),
		VirtualMachineInterfaceAllowedAddressPairs: MakeAllowedAddressPairs(),
		Annotations:                                MakeKeyValuePairs(),
		ParentType:                                 "",
		DisplayName:                                "",
		EcmpHashingIncludeFields:                   MakeEcmpHashingIncludeFields(),
		VirtualMachineInterfaceDHCPOptionList:      MakeDhcpOptionsListType(),
		VirtualMachineInterfaceDisablePolicy:       false,
		VirtualMachineInterfaceFatFlowProtocols:    MakeFatFlowProtocols(),
		VRFAssignTable:                             MakeVrfAssignTableType(),
		UUID:                                       "",
		ParentUUID:                                 "",
		VirtualMachineInterfaceHostRoutes:          MakeRouteTableType(),
		VirtualMachineInterfaceProperties:          MakeVirtualMachineInterfacePropertiesType(),
	}
}

// MakeVirtualMachineInterfaceSlice() makes a slice of VirtualMachineInterface
func MakeVirtualMachineInterfaceSlice() []*VirtualMachineInterface {
	return []*VirtualMachineInterface{}
}
