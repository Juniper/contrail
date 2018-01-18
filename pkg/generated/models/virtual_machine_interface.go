package models

// VirtualMachineInterface

import "encoding/json"

// VirtualMachineInterface
type VirtualMachineInterface struct {
	VRFAssignTable                             *VrfAssignTableType                    `json:"vrf_assign_table,omitempty"`
	VirtualMachineInterfaceProperties          *VirtualMachineInterfacePropertiesType `json:"virtual_machine_interface_properties,omitempty"`
	DisplayName                                string                                 `json:"display_name,omitempty"`
	ParentUUID                                 string                                 `json:"parent_uuid,omitempty"`
	EcmpHashingIncludeFields                   *EcmpHashingIncludeFields              `json:"ecmp_hashing_include_fields,omitempty"`
	VirtualMachineInterfaceHostRoutes          *RouteTableType                        `json:"virtual_machine_interface_host_routes,omitempty"`
	VirtualMachineInterfaceDHCPOptionList      *DhcpOptionsListType                   `json:"virtual_machine_interface_dhcp_option_list,omitempty"`
	VirtualMachineInterfaceAllowedAddressPairs *AllowedAddressPairs                   `json:"virtual_machine_interface_allowed_address_pairs,omitempty"`
	VirtualMachineInterfaceMacAddresses        *MacAddressesType                      `json:"virtual_machine_interface_mac_addresses,omitempty"`
	VlanTagBasedBridgeDomain                   bool                                   `json:"vlan_tag_based_bridge_domain"`
	UUID                                       string                                 `json:"uuid,omitempty"`
	Perms2                                     *PermType2                             `json:"perms2,omitempty"`
	FQName                                     []string                               `json:"fq_name,omitempty"`
	VirtualMachineInterfaceDisablePolicy       bool                                   `json:"virtual_machine_interface_disable_policy"`
	VirtualMachineInterfaceFatFlowProtocols    *FatFlowProtocols                      `json:"virtual_machine_interface_fat_flow_protocols,omitempty"`
	VirtualMachineInterfaceDeviceOwner         string                                 `json:"virtual_machine_interface_device_owner,omitempty"`
	Annotations                                *KeyValuePairs                         `json:"annotations,omitempty"`
	VirtualMachineInterfaceBindings            *KeyValuePairs                         `json:"virtual_machine_interface_bindings,omitempty"`
	PortSecurityEnabled                        bool                                   `json:"port_security_enabled"`
	ParentType                                 string                                 `json:"parent_type,omitempty"`
	IDPerms                                    *IdPermsType                           `json:"id_perms,omitempty"`

	PhysicalInterfaceRefs       []*VirtualMachineInterfacePhysicalInterfaceRef       `json:"physical_interface_refs,omitempty"`
	SecurityGroupRefs           []*VirtualMachineInterfaceSecurityGroupRef           `json:"security_group_refs,omitempty"`
	VirtualNetworkRefs          []*VirtualMachineInterfaceVirtualNetworkRef          `json:"virtual_network_refs,omitempty"`
	SecurityLoggingObjectRefs   []*VirtualMachineInterfaceSecurityLoggingObjectRef   `json:"security_logging_object_refs,omitempty"`
	RoutingInstanceRefs         []*VirtualMachineInterfaceRoutingInstanceRef         `json:"routing_instance_refs,omitempty"`
	ServiceHealthCheckRefs      []*VirtualMachineInterfaceServiceHealthCheckRef      `json:"service_health_check_refs,omitempty"`
	ServiceEndpointRefs         []*VirtualMachineInterfaceServiceEndpointRef         `json:"service_endpoint_refs,omitempty"`
	VirtualMachineInterfaceRefs []*VirtualMachineInterfaceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	VirtualMachineRefs          []*VirtualMachineInterfaceVirtualMachineRef          `json:"virtual_machine_refs,omitempty"`
	BGPRouterRefs               []*VirtualMachineInterfaceBGPRouterRef               `json:"bgp_router_refs,omitempty"`
	InterfaceRouteTableRefs     []*VirtualMachineInterfaceInterfaceRouteTableRef     `json:"interface_route_table_refs,omitempty"`
	QosConfigRefs               []*VirtualMachineInterfaceQosConfigRef               `json:"qos_config_refs,omitempty"`
	PortTupleRefs               []*VirtualMachineInterfacePortTupleRef               `json:"port_tuple_refs,omitempty"`
	BridgeDomainRefs            []*VirtualMachineInterfaceBridgeDomainRef            `json:"bridge_domain_refs,omitempty"`
}

// VirtualMachineInterfaceVirtualNetworkRef references each other
type VirtualMachineInterfaceVirtualNetworkRef struct {
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

// VirtualMachineInterfaceServiceHealthCheckRef references each other
type VirtualMachineInterfaceServiceHealthCheckRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceServiceEndpointRef references each other
type VirtualMachineInterfaceServiceEndpointRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceVirtualMachineInterfaceRef references each other
type VirtualMachineInterfaceVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceVirtualMachineRef references each other
type VirtualMachineInterfaceVirtualMachineRef struct {
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

// VirtualMachineInterfaceQosConfigRef references each other
type VirtualMachineInterfaceQosConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfacePortTupleRef references each other
type VirtualMachineInterfacePortTupleRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceBridgeDomainRef references each other
type VirtualMachineInterfaceBridgeDomainRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *BridgeDomainMembershipType
}

// VirtualMachineInterfacePhysicalInterfaceRef references each other
type VirtualMachineInterfacePhysicalInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceSecurityGroupRef references each other
type VirtualMachineInterfaceSecurityGroupRef struct {
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
		VirtualMachineInterfaceDHCPOptionList:      MakeDhcpOptionsListType(),
		VirtualMachineInterfaceAllowedAddressPairs: MakeAllowedAddressPairs(),
		VRFAssignTable:                             MakeVrfAssignTableType(),
		VirtualMachineInterfaceProperties:          MakeVirtualMachineInterfacePropertiesType(),
		DisplayName:                                "",
		ParentUUID:                                 "",
		EcmpHashingIncludeFields:                   MakeEcmpHashingIncludeFields(),
		VirtualMachineInterfaceHostRoutes:          MakeRouteTableType(),
		UUID: "",
		VirtualMachineInterfaceMacAddresses: MakeMacAddressesType(),
		VlanTagBasedBridgeDomain:            false,
		VirtualMachineInterfaceDeviceOwner:  "",
		Annotations:                         MakeKeyValuePairs(),
		Perms2:                              MakePermType2(),
		FQName:                              []string{},
		VirtualMachineInterfaceDisablePolicy:    false,
		VirtualMachineInterfaceFatFlowProtocols: MakeFatFlowProtocols(),
		ParentType:                              "",
		IDPerms:                                 MakeIdPermsType(),
		VirtualMachineInterfaceBindings: MakeKeyValuePairs(),
		PortSecurityEnabled:             false,
	}
}

// MakeVirtualMachineInterfaceSlice() makes a slice of VirtualMachineInterface
func MakeVirtualMachineInterfaceSlice() []*VirtualMachineInterface {
	return []*VirtualMachineInterface{}
}
