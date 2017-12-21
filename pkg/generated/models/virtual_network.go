package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control"`
	UUID                            string                    `json:"uuid"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	Perms2                          *PermType2                `json:"perms2"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`
	FQName                          []string                  `json:"fq_name"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	IDPerms                         *IdPermsType              `json:"id_perms"`
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`
	DisplayName                     string                    `json:"display_name"`
	Annotations                     *KeyValuePairs            `json:"annotations"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control"`
	IsShared                        bool                      `json:"is_shared"`
	ParentUUID                      string                    `json:"parent_uuid"`
	RouterExternal                  bool                      `json:"router_external"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties"`
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list"`
	ParentType                      string                    `json:"parent_type"`

	NetworkIpamRefs           []*VirtualNetworkNetworkIpamRef           `json:"network_ipam_refs"`
	SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef `json:"security_logging_object_refs"`
	NetworkPolicyRefs         []*VirtualNetworkNetworkPolicyRef         `json:"network_policy_refs"`
	QosConfigRefs             []*VirtualNetworkQosConfigRef             `json:"qos_config_refs"`
	RouteTableRefs            []*VirtualNetworkRouteTableRef            `json:"route_table_refs"`
	VirtualNetworkRefs        []*VirtualNetworkVirtualNetworkRef        `json:"virtual_network_refs"`
	BGPVPNRefs                []*VirtualNetworkBGPVPNRef                `json:"bgpvpn_refs"`

	AccessControlLists []*AccessControlList `json:"access_control_lists"`
	AliasIPPools       []*AliasIPPool       `json:"alias_ip_pools"`
	BridgeDomains      []*BridgeDomain      `json:"bridge_domains"`
	FloatingIPPools    []*FloatingIPPool    `json:"floating_ip_pools"`
	RoutingInstances   []*RoutingInstance   `json:"routing_instances"`
}

// VirtualNetworkRouteTableRef references each other
type VirtualNetworkRouteTableRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualNetworkVirtualNetworkRef references each other
type VirtualNetworkVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualNetworkBGPVPNRef references each other
type VirtualNetworkBGPVPNRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualNetworkNetworkIpamRef references each other
type VirtualNetworkNetworkIpamRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *VnSubnetsType
}

// VirtualNetworkSecurityLoggingObjectRef references each other
type VirtualNetworkSecurityLoggingObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualNetworkNetworkPolicyRef references each other
type VirtualNetworkNetworkPolicyRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *VirtualNetworkPolicyType
}

// VirtualNetworkQosConfigRef references each other
type VirtualNetworkQosConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *VirtualNetwork) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualNetwork makes VirtualNetwork
func MakeVirtualNetwork() *VirtualNetwork {
	return &VirtualNetwork{
		//TODO(nati): Apply default
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		MacLimitControl:          MakeMACLimitControlType(),
		UUID:                     "",
		VirtualNetworkNetworkID: MakeVirtualNetworkIdType(),
		PBBEtreeEnable:          false,
		Layer2ControlWord:       false,
		Perms2:                  MakePermType2(),
		MultiPolicyServiceChainsEnabled: false,
		FQName:                   []string{},
		ImportRouteTargetList:    MakeRouteTargetList(),
		PortSecurityEnabled:      false,
		IDPerms:                  MakeIdPermsType(),
		VirtualNetworkProperties: MakeVirtualNetworkType(),
		MacAgingTime:             MakeMACAgingTime(),
		MacLearningEnabled:       false,
		FloodUnknownUnicast:      false,
		Annotations:              MakeKeyValuePairs(),
		ExternalIpam:             false,
		MacMoveControl:           MakeMACMoveLimitControlType(),
		IsShared:                 false,
		ParentUUID:               "",
		DisplayName:              "",
		RouterExternal:           false,
		ProviderProperties:       MakeProviderDetails(),
		ExportRouteTargetList:    MakeRouteTargetList(),
		ParentType:               "",
		AddressAllocationMode:    MakeAddressAllocationModeType(),
		PBBEvpnEnable:            false,
		RouteTargetList:          MakeRouteTargetList(),
	}
}

// InterfaceToVirtualNetwork makes VirtualNetwork from interface
func InterfaceToVirtualNetwork(iData interface{}) *VirtualNetwork {
	data := iData.(map[string]interface{})
	return &VirtualNetwork{
		MultiPolicyServiceChainsEnabled: data["multi_policy_service_chains_enabled"].(bool),

		//{"type":"boolean"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		ImportRouteTargetList: InterfaceToRouteTargetList(data["import_route_target_list"]),

		//{"description":"List of route targets that are used as import for this virtual network.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		PortSecurityEnabled: data["port_security_enabled"].(bool),

		//{"description":"Port security status on the network","default":true,"type":"boolean"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		FloodUnknownUnicast: data["flood_unknown_unicast"].(bool),

		//{"description":"When true, packets with unknown unicast MAC address are flooded within the network. Default they are dropped.","default":false,"type":"boolean"}
		VirtualNetworkProperties: InterfaceToVirtualNetworkType(data["virtual_network_properties"]),

		//{"description":"Virtual network miscellaneous configurations.","type":"object","properties":{"allow_transit":{"type":"boolean"},"forwarding_mode":{"type":"string","enum":["l2_l3","l2","l3"]},"mirror_destination":{"type":"boolean"},"network_id":{"type":"integer"},"rpf":{"type":"string","enum":["enable","disable"]},"vxlan_network_identifier":{"type":"integer","minimum":1,"maximum":16777215}}}
		MacAgingTime: InterfaceToMACAgingTime(data["mac_aging_time"]),

		//{"description":"MAC aging time on the network","default":"300","type":"integer","minimum":0,"maximum":86400}
		MacLearningEnabled: data["mac_learning_enabled"].(bool),

		//{"description":"Enable MAC learning on the network","default":false,"type":"boolean"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ExternalIpam: data["external_ipam"].(bool),

		//{"description":"IP address assignment to VM is done statically, outside of (external to) Contrail Ipam. vCenter only feature.","type":"boolean"}
		MacMoveControl: InterfaceToMACMoveLimitControlType(data["mac_move_control"]),

		//{"description":"MAC move control on the network","type":"object","properties":{"mac_move_limit":{"type":"integer"},"mac_move_limit_action":{"type":"string","enum":["log","alarm","shutdown","drop"]},"mac_move_time_window":{"type":"integer","minimum":1,"maximum":60}}}
		IsShared: data["is_shared"].(bool),

		//{"description":"When true, this virtual network is shared with all tenants.","type":"boolean"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		RouterExternal: data["router_external"].(bool),

		//{"description":"When true, this virtual network is openstack router external network.","type":"boolean"}
		ProviderProperties: InterfaceToProviderDetails(data["provider_properties"]),

		//{"description":"Virtual network is provider network. Specifies VLAN tag and physical network name.","type":"object","properties":{"physical_network":{"type":"string"},"segmentation_id":{"type":"integer","minimum":1,"maximum":4094}}}
		ExportRouteTargetList: InterfaceToRouteTargetList(data["export_route_target_list"]),

		//{"description":"List of route targets that are used as export for this virtual network.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		AddressAllocationMode: InterfaceToAddressAllocationModeType(data["address_allocation_mode"]),

		//{"description":"Address allocation mode for virtual network.","type":"string","enum":["user-defined-subnet-preferred","user-defined-subnet-only","flat-subnet-preferred","flat-subnet-only"]}
		PBBEvpnEnable: data["pbb_evpn_enable"].(bool),

		//{"description":"Enable/Disable PBB EVPN tunneling on the network","default":false,"type":"boolean"}
		RouteTargetList: InterfaceToRouteTargetList(data["route_target_list"]),

		//{"description":"List of route targets that are used as both import and export for this virtual network.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		EcmpHashingIncludeFields: InterfaceToEcmpHashingIncludeFields(data["ecmp_hashing_include_fields"]),

		//{"description":"ECMP hashing config at global level.","type":"object","properties":{"destination_ip":{"type":"boolean"},"destination_port":{"type":"boolean"},"hashing_configured":{"type":"boolean"},"ip_protocol":{"type":"boolean"},"source_ip":{"type":"boolean"},"source_port":{"type":"boolean"}}}
		MacLimitControl: InterfaceToMACLimitControlType(data["mac_limit_control"]),

		//{"description":"MAC limit control on the network","type":"object","properties":{"mac_limit":{"type":"integer"},"mac_limit_action":{"type":"string","enum":["log","alarm","shutdown","drop"]}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		VirtualNetworkNetworkID: InterfaceToVirtualNetworkIdType(data["virtual_network_network_id"]),

		//{"description":"System assigned unique 32 bit ID for every virtual network.","type":"integer","minimum":1,"maximum":4294967296}
		PBBEtreeEnable: data["pbb_etree_enable"].(bool),

		//{"description":"Enable/Disable PBB ETREE mode on the network","default":false,"type":"boolean"}
		Layer2ControlWord: data["layer2_control_word"].(bool),

		//{"description":"Enable/Disable adding control word to the Layer 2 encapsulation","default":false,"type":"boolean"}

	}
}

// InterfaceToVirtualNetworkSlice makes a slice of VirtualNetwork from interface
func InterfaceToVirtualNetworkSlice(data interface{}) []*VirtualNetwork {
	list := data.([]interface{})
	result := MakeVirtualNetworkSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetwork(item))
	}
	return result
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
