package models

// VirtualMachineInterface

// VirtualMachineInterface
//proteus:generate
type VirtualMachineInterface struct {
	UUID                                       string                                 `json:"uuid,omitempty"`
	ParentUUID                                 string                                 `json:"parent_uuid,omitempty"`
	ParentType                                 string                                 `json:"parent_type,omitempty"`
	FQName                                     []string                               `json:"fq_name,omitempty"`
	IDPerms                                    *IdPermsType                           `json:"id_perms,omitempty"`
	DisplayName                                string                                 `json:"display_name,omitempty"`
	Annotations                                *KeyValuePairs                         `json:"annotations,omitempty"`
	Perms2                                     *PermType2                             `json:"perms2,omitempty"`
	EcmpHashingIncludeFields                   *EcmpHashingIncludeFields              `json:"ecmp_hashing_include_fields,omitempty"`
	VirtualMachineInterfaceHostRoutes          *RouteTableType                        `json:"virtual_machine_interface_host_routes,omitempty"`
	VirtualMachineInterfaceMacAddresses        *MacAddressesType                      `json:"virtual_machine_interface_mac_addresses,omitempty"`
	VirtualMachineInterfaceDHCPOptionList      *DhcpOptionsListType                   `json:"virtual_machine_interface_dhcp_option_list,omitempty"`
	VirtualMachineInterfaceBindings            *KeyValuePairs                         `json:"virtual_machine_interface_bindings,omitempty"`
	VirtualMachineInterfaceDisablePolicy       bool                                   `json:"virtual_machine_interface_disable_policy"`
	VirtualMachineInterfaceAllowedAddressPairs *AllowedAddressPairs                   `json:"virtual_machine_interface_allowed_address_pairs,omitempty"`
	VirtualMachineInterfaceFatFlowProtocols    *FatFlowProtocols                      `json:"virtual_machine_interface_fat_flow_protocols,omitempty"`
	VlanTagBasedBridgeDomain                   bool                                   `json:"vlan_tag_based_bridge_domain"`
	VirtualMachineInterfaceDeviceOwner         string                                 `json:"virtual_machine_interface_device_owner,omitempty"`
	VRFAssignTable                             *VrfAssignTableType                    `json:"vrf_assign_table,omitempty"`
	PortSecurityEnabled                        bool                                   `json:"port_security_enabled"`
	VirtualMachineInterfaceProperties          *VirtualMachineInterfacePropertiesType `json:"virtual_machine_interface_properties,omitempty"`

	VirtualMachineRefs          []*VirtualMachineInterfaceVirtualMachineRef          `json:"virtual_machine_refs,omitempty"`
	VirtualNetworkRefs          []*VirtualMachineInterfaceVirtualNetworkRef          `json:"virtual_network_refs,omitempty"`
	InterfaceRouteTableRefs     []*VirtualMachineInterfaceInterfaceRouteTableRef     `json:"interface_route_table_refs,omitempty"`
	BridgeDomainRefs            []*VirtualMachineInterfaceBridgeDomainRef            `json:"bridge_domain_refs,omitempty"`
	VirtualMachineInterfaceRefs []*VirtualMachineInterfaceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	BGPRouterRefs               []*VirtualMachineInterfaceBGPRouterRef               `json:"bgp_router_refs,omitempty"`
	QosConfigRefs               []*VirtualMachineInterfaceQosConfigRef               `json:"qos_config_refs,omitempty"`
	PortTupleRefs               []*VirtualMachineInterfacePortTupleRef               `json:"port_tuple_refs,omitempty"`
	PhysicalInterfaceRefs       []*VirtualMachineInterfacePhysicalInterfaceRef       `json:"physical_interface_refs,omitempty"`
	ServiceHealthCheckRefs      []*VirtualMachineInterfaceServiceHealthCheckRef      `json:"service_health_check_refs,omitempty"`
	SecurityGroupRefs           []*VirtualMachineInterfaceSecurityGroupRef           `json:"security_group_refs,omitempty"`
	ServiceEndpointRefs         []*VirtualMachineInterfaceServiceEndpointRef         `json:"service_endpoint_refs,omitempty"`
	SecurityLoggingObjectRefs   []*VirtualMachineInterfaceSecurityLoggingObjectRef   `json:"security_logging_object_refs,omitempty"`
	RoutingInstanceRefs         []*VirtualMachineInterfaceRoutingInstanceRef         `json:"routing_instance_refs,omitempty"`
}

// VirtualMachineInterfaceQosConfigRef references each other
type VirtualMachineInterfaceQosConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceSecurityLoggingObjectRef references each other
type VirtualMachineInterfaceSecurityLoggingObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceRoutingInstanceRef references each other
type VirtualMachineInterfaceRoutingInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *PolicyBasedForwardingRuleType
}

// VirtualMachineInterfacePortTupleRef references each other
type VirtualMachineInterfacePortTupleRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfacePhysicalInterfaceRef references each other
type VirtualMachineInterfacePhysicalInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceServiceHealthCheckRef references each other
type VirtualMachineInterfaceServiceHealthCheckRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceSecurityGroupRef references each other
type VirtualMachineInterfaceSecurityGroupRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceServiceEndpointRef references each other
type VirtualMachineInterfaceServiceEndpointRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceVirtualMachineRef references each other
type VirtualMachineInterfaceVirtualMachineRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceVirtualNetworkRef references each other
type VirtualMachineInterfaceVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceVirtualMachineInterfaceRef references each other
type VirtualMachineInterfaceVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceBGPRouterRef references each other
type VirtualMachineInterfaceBGPRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceInterfaceRouteTableRef references each other
type VirtualMachineInterfaceInterfaceRouteTableRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceBridgeDomainRef references each other
type VirtualMachineInterfaceBridgeDomainRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *BridgeDomainMembershipType
}

// MakeVirtualMachineInterface makes VirtualMachineInterface
func MakeVirtualMachineInterface() *VirtualMachineInterface {
	return &VirtualMachineInterface{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		EcmpHashingIncludeFields:                   MakeEcmpHashingIncludeFields(),
		VirtualMachineInterfaceHostRoutes:          MakeRouteTableType(),
		VirtualMachineInterfaceMacAddresses:        MakeMacAddressesType(),
		VirtualMachineInterfaceDHCPOptionList:      MakeDhcpOptionsListType(),
		VirtualMachineInterfaceBindings:            MakeKeyValuePairs(),
		VirtualMachineInterfaceDisablePolicy:       false,
		VirtualMachineInterfaceAllowedAddressPairs: MakeAllowedAddressPairs(),
		VirtualMachineInterfaceFatFlowProtocols:    MakeFatFlowProtocols(),
		VlanTagBasedBridgeDomain:                   false,
		VirtualMachineInterfaceDeviceOwner:         "",
		VRFAssignTable:                             MakeVrfAssignTableType(),
		PortSecurityEnabled:                        false,
		VirtualMachineInterfaceProperties:          MakeVirtualMachineInterfacePropertiesType(),
	}
}

// MakeVirtualMachineInterfaceSlice() makes a slice of VirtualMachineInterface
func MakeVirtualMachineInterfaceSlice() []*VirtualMachineInterface {
	return []*VirtualMachineInterface{}
}
