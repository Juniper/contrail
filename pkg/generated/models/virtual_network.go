package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	DisplayName                     string                    `json:"display_name"`
	ParentUUID                      string                    `json:"parent_uuid"`
	IDPerms                         *IdPermsType              `json:"id_perms"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control"`
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields"`
	RouterExternal                  bool                      `json:"router_external"`
	Annotations                     *KeyValuePairs            `json:"annotations"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control"`
	UUID                            string                    `json:"uuid"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list"`
	IsShared                        bool                      `json:"is_shared"`
	FQName                          []string                  `json:"fq_name"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list"`
	Perms2                          *PermType2                `json:"perms2"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	ParentType                      string                    `json:"parent_type"`

	BGPVPNRefs                []*VirtualNetworkBGPVPNRef                `json:"bgpvpn_refs"`
	NetworkIpamRefs           []*VirtualNetworkNetworkIpamRef           `json:"network_ipam_refs"`
	SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef `json:"security_logging_object_refs"`
	NetworkPolicyRefs         []*VirtualNetworkNetworkPolicyRef         `json:"network_policy_refs"`
	QosConfigRefs             []*VirtualNetworkQosConfigRef             `json:"qos_config_refs"`
	RouteTableRefs            []*VirtualNetworkRouteTableRef            `json:"route_table_refs"`
	VirtualNetworkRefs        []*VirtualNetworkVirtualNetworkRef        `json:"virtual_network_refs"`

	AccessControlLists []*AccessControlList `json:"access_control_lists"`
	AliasIPPools       []*AliasIPPool       `json:"alias_ip_pools"`
	BridgeDomains      []*BridgeDomain      `json:"bridge_domains"`
	FloatingIPPools    []*FloatingIPPool    `json:"floating_ip_pools"`
	RoutingInstances   []*RoutingInstance   `json:"routing_instances"`
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
		RouterExternal:           false,
		Annotations:              MakeKeyValuePairs(),
		VirtualNetworkNetworkID:  MakeVirtualNetworkIdType(),
		MacLimitControl:          MakeMACLimitControlType(),
		UUID:                     "",
		ImportRouteTargetList:           MakeRouteTargetList(),
		IsShared:                        false,
		FQName:                          []string{},
		MacLearningEnabled:              false,
		ExportRouteTargetList:           MakeRouteTargetList(),
		Perms2:                          MakePermType2(),
		MacAgingTime:                    MakeMACAgingTime(),
		Layer2ControlWord:               false,
		ParentType:                      "",
		ProviderProperties:              MakeProviderDetails(),
		ExternalIpam:                    false,
		PortSecurityEnabled:             false,
		MultiPolicyServiceChainsEnabled: false,
		DisplayName:                     "",
		ParentUUID:                      "",
		IDPerms:                         MakeIdPermsType(),
		FloodUnknownUnicast:             false,
		VirtualNetworkProperties:        MakeVirtualNetworkType(),
		AddressAllocationMode:           MakeAddressAllocationModeType(),
		PBBEvpnEnable:                   false,
		RouteTargetList:                 MakeRouteTargetList(),
		PBBEtreeEnable:                  false,
		MacMoveControl:                  MakeMACMoveLimitControlType(),
	}
}

// InterfaceToVirtualNetwork makes VirtualNetwork from interface
func InterfaceToVirtualNetwork(iData interface{}) *VirtualNetwork {
	data := iData.(map[string]interface{})
	return &VirtualNetwork{
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ProviderProperties: InterfaceToProviderDetails(data["provider_properties"]),

		//{"description":"Virtual network is provider network. Specifies VLAN tag and physical network name.","type":"object","properties":{"physical_network":{"type":"string"},"segmentation_id":{"type":"integer","minimum":1,"maximum":4094}}}
		ExternalIpam: data["external_ipam"].(bool),

		//{"description":"IP address assignment to VM is done statically, outside of (external to) Contrail Ipam. vCenter only feature.","type":"boolean"}
		PortSecurityEnabled: data["port_security_enabled"].(bool),

		//{"description":"Port security status on the network","default":true,"type":"boolean"}
		MultiPolicyServiceChainsEnabled: data["multi_policy_service_chains_enabled"].(bool),

		//{"type":"boolean"}
		FloodUnknownUnicast: data["flood_unknown_unicast"].(bool),

		//{"description":"When true, packets with unknown unicast MAC address are flooded within the network. Default they are dropped.","default":false,"type":"boolean"}
		PBBEtreeEnable: data["pbb_etree_enable"].(bool),

		//{"description":"Enable/Disable PBB ETREE mode on the network","default":false,"type":"boolean"}
		MacMoveControl: InterfaceToMACMoveLimitControlType(data["mac_move_control"]),

		//{"description":"MAC move control on the network","type":"object","properties":{"mac_move_limit":{"type":"integer"},"mac_move_limit_action":{"type":"string","enum":["log","alarm","shutdown","drop"]},"mac_move_time_window":{"type":"integer","minimum":1,"maximum":60}}}
		VirtualNetworkProperties: InterfaceToVirtualNetworkType(data["virtual_network_properties"]),

		//{"description":"Virtual network miscellaneous configurations.","type":"object","properties":{"allow_transit":{"type":"boolean"},"forwarding_mode":{"type":"string","enum":["l2_l3","l2","l3"]},"mirror_destination":{"type":"boolean"},"network_id":{"type":"integer"},"rpf":{"type":"string","enum":["enable","disable"]},"vxlan_network_identifier":{"type":"integer","minimum":1,"maximum":16777215}}}
		AddressAllocationMode: InterfaceToAddressAllocationModeType(data["address_allocation_mode"]),

		//{"description":"Address allocation mode for virtual network.","type":"string","enum":["user-defined-subnet-preferred","user-defined-subnet-only","flat-subnet-preferred","flat-subnet-only"]}
		PBBEvpnEnable: data["pbb_evpn_enable"].(bool),

		//{"description":"Enable/Disable PBB EVPN tunneling on the network","default":false,"type":"boolean"}
		RouteTargetList: InterfaceToRouteTargetList(data["route_target_list"]),

		//{"description":"List of route targets that are used as both import and export for this virtual network.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		EcmpHashingIncludeFields: InterfaceToEcmpHashingIncludeFields(data["ecmp_hashing_include_fields"]),

		//{"description":"ECMP hashing config at global level.","type":"object","properties":{"destination_ip":{"type":"boolean"},"destination_port":{"type":"boolean"},"hashing_configured":{"type":"boolean"},"ip_protocol":{"type":"boolean"},"source_ip":{"type":"boolean"},"source_port":{"type":"boolean"}}}
		RouterExternal: data["router_external"].(bool),

		//{"description":"When true, this virtual network is openstack router external network.","type":"boolean"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		VirtualNetworkNetworkID: InterfaceToVirtualNetworkIdType(data["virtual_network_network_id"]),

		//{"description":"System assigned unique 32 bit ID for every virtual network.","type":"integer","minimum":1,"maximum":4294967296}
		MacLimitControl: InterfaceToMACLimitControlType(data["mac_limit_control"]),

		//{"description":"MAC limit control on the network","type":"object","properties":{"mac_limit":{"type":"integer"},"mac_limit_action":{"type":"string","enum":["log","alarm","shutdown","drop"]}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ImportRouteTargetList: InterfaceToRouteTargetList(data["import_route_target_list"]),

		//{"description":"List of route targets that are used as import for this virtual network.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		IsShared: data["is_shared"].(bool),

		//{"description":"When true, this virtual network is shared with all tenants.","type":"boolean"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		MacLearningEnabled: data["mac_learning_enabled"].(bool),

		//{"description":"Enable MAC learning on the network","default":false,"type":"boolean"}
		ExportRouteTargetList: InterfaceToRouteTargetList(data["export_route_target_list"]),

		//{"description":"List of route targets that are used as export for this virtual network.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		MacAgingTime: InterfaceToMACAgingTime(data["mac_aging_time"]),

		//{"description":"MAC aging time on the network","default":"300","type":"integer","minimum":0,"maximum":86400}
		Layer2ControlWord: data["layer2_control_word"].(bool),

		//{"description":"Enable/Disable adding control word to the Layer 2 encapsulation","default":false,"type":"boolean"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}

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
