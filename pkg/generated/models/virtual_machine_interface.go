package models

// VirtualMachineInterface

import "encoding/json"

// VirtualMachineInterface
type VirtualMachineInterface struct {
	VirtualMachineInterfaceFatFlowProtocols    *FatFlowProtocols                      `json:"virtual_machine_interface_fat_flow_protocols"`
	DisplayName                                string                                 `json:"display_name"`
	EcmpHashingIncludeFields                   *EcmpHashingIncludeFields              `json:"ecmp_hashing_include_fields"`
	VirtualMachineInterfaceDHCPOptionList      *DhcpOptionsListType                   `json:"virtual_machine_interface_dhcp_option_list"`
	VirtualMachineInterfaceDisablePolicy       bool                                   `json:"virtual_machine_interface_disable_policy"`
	VlanTagBasedBridgeDomain                   bool                                   `json:"vlan_tag_based_bridge_domain"`
	VirtualMachineInterfaceDeviceOwner         string                                 `json:"virtual_machine_interface_device_owner"`
	ParentUUID                                 string                                 `json:"parent_uuid"`
	IDPerms                                    *IdPermsType                           `json:"id_perms"`
	VirtualMachineInterfaceHostRoutes          *RouteTableType                        `json:"virtual_machine_interface_host_routes"`
	VirtualMachineInterfaceBindings            *KeyValuePairs                         `json:"virtual_machine_interface_bindings"`
	VRFAssignTable                             *VrfAssignTableType                    `json:"vrf_assign_table"`
	PortSecurityEnabled                        bool                                   `json:"port_security_enabled"`
	VirtualMachineInterfaceProperties          *VirtualMachineInterfacePropertiesType `json:"virtual_machine_interface_properties"`
	ParentType                                 string                                 `json:"parent_type"`
	VirtualMachineInterfaceMacAddresses        *MacAddressesType                      `json:"virtual_machine_interface_mac_addresses"`
	VirtualMachineInterfaceAllowedAddressPairs *AllowedAddressPairs                   `json:"virtual_machine_interface_allowed_address_pairs"`
	Annotations                                *KeyValuePairs                         `json:"annotations"`
	Perms2                                     *PermType2                             `json:"perms2"`
	UUID                                       string                                 `json:"uuid"`
	FQName                                     []string                               `json:"fq_name"`

	VirtualMachineInterfaceRefs []*VirtualMachineInterfaceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
	VirtualMachineRefs          []*VirtualMachineInterfaceVirtualMachineRef          `json:"virtual_machine_refs"`
	BGPRouterRefs               []*VirtualMachineInterfaceBGPRouterRef               `json:"bgp_router_refs"`
	SecurityLoggingObjectRefs   []*VirtualMachineInterfaceSecurityLoggingObjectRef   `json:"security_logging_object_refs"`
	BridgeDomainRefs            []*VirtualMachineInterfaceBridgeDomainRef            `json:"bridge_domain_refs"`
	InterfaceRouteTableRefs     []*VirtualMachineInterfaceInterfaceRouteTableRef     `json:"interface_route_table_refs"`
	RoutingInstanceRefs         []*VirtualMachineInterfaceRoutingInstanceRef         `json:"routing_instance_refs"`
	QosConfigRefs               []*VirtualMachineInterfaceQosConfigRef               `json:"qos_config_refs"`
	ServiceHealthCheckRefs      []*VirtualMachineInterfaceServiceHealthCheckRef      `json:"service_health_check_refs"`
	VirtualNetworkRefs          []*VirtualMachineInterfaceVirtualNetworkRef          `json:"virtual_network_refs"`
	PhysicalInterfaceRefs       []*VirtualMachineInterfacePhysicalInterfaceRef       `json:"physical_interface_refs"`
	SecurityGroupRefs           []*VirtualMachineInterfaceSecurityGroupRef           `json:"security_group_refs"`
	PortTupleRefs               []*VirtualMachineInterfacePortTupleRef               `json:"port_tuple_refs"`
	ServiceEndpointRefs         []*VirtualMachineInterfaceServiceEndpointRef         `json:"service_endpoint_refs"`
}

// VirtualMachineInterfaceServiceHealthCheckRef references each other
type VirtualMachineInterfaceServiceHealthCheckRef struct {
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

// VirtualMachineInterfaceInterfaceRouteTableRef references each other
type VirtualMachineInterfaceInterfaceRouteTableRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceRoutingInstanceRef references each other
type VirtualMachineInterfaceRoutingInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *PolicyBasedForwardingRuleType
}

// VirtualMachineInterfaceQosConfigRef references each other
type VirtualMachineInterfaceQosConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

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

// VirtualMachineInterfacePortTupleRef references each other
type VirtualMachineInterfacePortTupleRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceServiceEndpointRef references each other
type VirtualMachineInterfaceServiceEndpointRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualMachineInterfaceSecurityLoggingObjectRef references each other
type VirtualMachineInterfaceSecurityLoggingObjectRef struct {
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

// String returns json representation of the object
func (model *VirtualMachineInterface) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualMachineInterface makes VirtualMachineInterface
func MakeVirtualMachineInterface() *VirtualMachineInterface {
	return &VirtualMachineInterface{
		//TODO(nati): Apply default
		VirtualMachineInterfaceFatFlowProtocols: MakeFatFlowProtocols(),
		EcmpHashingIncludeFields:                MakeEcmpHashingIncludeFields(),
		VirtualMachineInterfaceDHCPOptionList:   MakeDhcpOptionsListType(),
		VirtualMachineInterfaceDisablePolicy:    false,
		VlanTagBasedBridgeDomain:                false,
		VirtualMachineInterfaceDeviceOwner:      "",
		ParentUUID:                              "",
		DisplayName:                             "",
		VirtualMachineInterfaceHostRoutes:       MakeRouteTableType(),
		VirtualMachineInterfaceBindings:         MakeKeyValuePairs(),
		VRFAssignTable:                          MakeVrfAssignTableType(),
		PortSecurityEnabled:                     false,
		VirtualMachineInterfaceProperties:       MakeVirtualMachineInterfacePropertiesType(),
		ParentType:                              "",
		IDPerms:                                 MakeIdPermsType(),
		VirtualMachineInterfaceMacAddresses:        MakeMacAddressesType(),
		VirtualMachineInterfaceAllowedAddressPairs: MakeAllowedAddressPairs(),
		Annotations:                                MakeKeyValuePairs(),
		Perms2:                                     MakePermType2(),
		UUID:                                       "",
		FQName:                                     []string{},
	}
}

// InterfaceToVirtualMachineInterface makes VirtualMachineInterface from interface
func InterfaceToVirtualMachineInterface(iData interface{}) *VirtualMachineInterface {
	data := iData.(map[string]interface{})
	return &VirtualMachineInterface{
		VirtualMachineInterfaceMacAddresses: InterfaceToMacAddressesType(data["virtual_machine_interface_mac_addresses"]),

		//{"description":"MAC address of the virtual machine interface, automatically assigned by system if not provided.","type":"object","properties":{"mac_address":{"type":"array","item":{"type":"string"}}}}
		VirtualMachineInterfaceAllowedAddressPairs: InterfaceToAllowedAddressPairs(data["virtual_machine_interface_allowed_address_pairs"]),

		//{"description":"List of (IP address, MAC) other than instance ip on this interface.","type":"object","properties":{"allowed_address_pair":{"type":"array","item":{"type":"object","properties":{"address_mode":{"type":"string","enum":["active-active","active-standby"]},"ip":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"mac":{"type":"string"}}}}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		VirtualMachineInterfaceFatFlowProtocols: InterfaceToFatFlowProtocols(data["virtual_machine_interface_fat_flow_protocols"]),

		//{"description":"List of (protocol, port number), for flows to interface with (protocol, destination port number), vrouter will ignore source port while setting up flow and ignore it as source port in reverse flow. Hence many flows will map to single flow.","type":"object","properties":{"fat_flow_protocol":{"type":"array","item":{"type":"object","properties":{"port":{"type":"integer"},"protocol":{"type":"string"}}}}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		EcmpHashingIncludeFields: InterfaceToEcmpHashingIncludeFields(data["ecmp_hashing_include_fields"]),

		//{"description":"ECMP hashing config at global level.","type":"object","properties":{"destination_ip":{"type":"boolean"},"destination_port":{"type":"boolean"},"hashing_configured":{"type":"boolean"},"ip_protocol":{"type":"boolean"},"source_ip":{"type":"boolean"},"source_port":{"type":"boolean"}}}
		VirtualMachineInterfaceDHCPOptionList: InterfaceToDhcpOptionsListType(data["virtual_machine_interface_dhcp_option_list"]),

		//{"description":"DHCP options configuration specific to this interface.","type":"object","properties":{"dhcp_option":{"type":"array","item":{"type":"object","properties":{"dhcp_option_name":{"type":"string"},"dhcp_option_value":{"type":"string"},"dhcp_option_value_bytes":{"type":"string"}}}}}}
		VirtualMachineInterfaceDisablePolicy: data["virtual_machine_interface_disable_policy"].(bool),

		//{"description":"When True all policy checks for ingress and egress traffic from this interface are disabled. Flow table entries are not created. Features that require policy will not work on this interface, these include security group, floating IP, service chain, linklocal services.","default":false,"type":"boolean"}
		VlanTagBasedBridgeDomain: data["vlan_tag_based_bridge_domain"].(bool),

		//{"description":"Enable VLAN tag based bridge domain classification on the port","default":false,"type":"boolean"}
		VirtualMachineInterfaceDeviceOwner: data["virtual_machine_interface_device_owner"].(string),

		//{"description":"For openstack compatibility, not used by system.","type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		VirtualMachineInterfaceHostRoutes: InterfaceToRouteTableType(data["virtual_machine_interface_host_routes"]),

		//{"description":"List of host routes(prefixes, nexthop) that are passed to VM via DHCP.","type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}}
		VirtualMachineInterfaceBindings: InterfaceToKeyValuePairs(data["virtual_machine_interface_bindings"]),

		//{"description":"Dictionary of arbitrary (key, value) for this interface. Neutron port bindings use this.","type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		VRFAssignTable: InterfaceToVrfAssignTableType(data["vrf_assign_table"]),

		//{"description":"VRF assignment policy for this interface, automatically generated by system.","type":"object","properties":{"vrf_assign_rule":{"type":"array","item":{"type":"object","properties":{"ignore_acl":{"type":"boolean"},"match_condition":{"type":"object","properties":{"dst_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"dst_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}},"ethertype":{"type":"string","enum":["IPv4","IPv6"]},"protocol":{"type":"string"},"src_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"src_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}},"routing_instance":{"type":"string"},"vlan_tag":{"type":"integer"}}}}}}
		PortSecurityEnabled: data["port_security_enabled"].(bool),

		//{"description":"Port security status on the network","default":true,"type":"boolean"}
		VirtualMachineInterfaceProperties: InterfaceToVirtualMachineInterfacePropertiesType(data["virtual_machine_interface_properties"]),

		//{"description":"Virtual Machine Interface miscellaneous configurations.","type":"object","properties":{"interface_mirror":{"type":"object","properties":{"mirror_to":{"type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}},"traffic_direction":{"default":"both","type":"string","enum":["ingress","egress","both"]}}},"local_preference":{"type":"integer"},"service_interface_type":{"type":"string"},"sub_interface_vlan_tag":{"type":"integer"}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}

	}
}

// InterfaceToVirtualMachineInterfaceSlice makes a slice of VirtualMachineInterface from interface
func InterfaceToVirtualMachineInterfaceSlice(data interface{}) []*VirtualMachineInterface {
	list := data.([]interface{})
	result := MakeVirtualMachineInterfaceSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualMachineInterface(item))
	}
	return result
}

// MakeVirtualMachineInterfaceSlice() makes a slice of VirtualMachineInterface
func MakeVirtualMachineInterfaceSlice() []*VirtualMachineInterface {
	return []*VirtualMachineInterface{}
}
