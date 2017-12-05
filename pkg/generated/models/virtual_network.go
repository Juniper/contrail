package models

// VirtualNetwork

import "encoding/json"

type VirtualNetwork struct {
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	DisplayName                     string                    `json:"display_name"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	RouterExternal                  bool                      `json:"router_external"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control"`
	FQName                          []string                  `json:"fq_name"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time"`
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list"`
	IsShared                        bool                      `json:"is_shared"`
	Perms2                          *PermType2                `json:"perms2"`
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	IDPerms                         *IdPermsType              `json:"id_perms"`
	UUID                            string                    `json:"uuid"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list"`
	Annotations                     *KeyValuePairs            `json:"annotations"`

	// security_logging_object <utils.Reference Value>
	SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef
	// network_policy <utils.Reference Value>
	NetworkPolicyRefs []*VirtualNetworkNetworkPolicyRef
	// qos_config <utils.Reference Value>
	QosConfigRefs []*VirtualNetworkQosConfigRef
	// route_table <utils.Reference Value>
	RouteTableRefs []*VirtualNetworkRouteTableRef
	// virtual_network <utils.Reference Value>
	VirtualNetworkRefs []*VirtualNetworkVirtualNetworkRef
	// bgpvpn <utils.Reference Value>
	BGPVPNRefs []*VirtualNetworkBGPVPNRef
	// network_ipam <utils.Reference Value>
	NetworkIpamRefs []*VirtualNetworkNetworkIpamRef

	Projects []*VirtualNetworkProject
}

// <utils.Reference Value>
type VirtualNetworkRouteTableRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// <utils.Reference Value>
type VirtualNetworkVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// <utils.Reference Value>
type VirtualNetworkBGPVPNRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// <utils.Reference Value>
type VirtualNetworkNetworkIpamRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *VnSubnetsType
}

// <utils.Reference Value>
type VirtualNetworkSecurityLoggingObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// <utils.Reference Value>
type VirtualNetworkNetworkPolicyRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *VirtualNetworkPolicyType
}

// <utils.Reference Value>
type VirtualNetworkQosConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

type VirtualNetworkProject struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

func (model *VirtualNetwork) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeVirtualNetwork() *VirtualNetwork {
	return &VirtualNetwork{
		//TODO(nati): Apply default
		RouterExternal:                  false,
		FloodUnknownUnicast:             false,
		ExternalIpam:                    false,
		MultiPolicyServiceChainsEnabled: false,
		MacLimitControl:                 MakeMACLimitControlType(),
		VirtualNetworkNetworkID:         MakeVirtualNetworkIdType(),
		MacMoveControl:                  MakeMACMoveLimitControlType(),
		FQName:                          []string{},
		PBBEvpnEnable:                   false,
		MacAgingTime:                    MakeMACAgingTime(),
		ExportRouteTargetList:           MakeRouteTargetList(),
		IsShared:                        false,
		Perms2:                          MakePermType2(),
		VirtualNetworkProperties: MakeVirtualNetworkType(),
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		AddressAllocationMode:    MakeAddressAllocationModeType(),
		ImportRouteTargetList:    MakeRouteTargetList(),
		Layer2ControlWord:        false,
		ProviderProperties:       MakeProviderDetails(),
		MacLearningEnabled:       false,
		IDPerms:                  MakeIdPermsType(),
		UUID:                     "",
		RouteTargetList:          MakeRouteTargetList(),
		Annotations:              MakeKeyValuePairs(),
		PBBEtreeEnable:           false,
		DisplayName:              "",
		PortSecurityEnabled:      false,
	}
}

func InterfaceToVirtualNetwork(iData interface{}) *VirtualNetwork {
	data := iData.(map[string]interface{})
	return &VirtualNetwork{
		MacAgingTime: InterfaceToMACAgingTime(data["mac_aging_time"]),

		//{"Title":"","Description":"MAC aging time on the network","SQL":"int","Default":"300","Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":86400,"Ref":"types.json#/definitions/MACAgingTime","CollectionType":"","Column":"mac_aging_time","Item":null,"GoName":"MacAgingTime","GoType":"MACAgingTime"}
		ExportRouteTargetList: InterfaceToRouteTargetList(data["export_route_target_list"]),

		//{"Title":"","Description":"List of route targets that are used as export for this virtual network.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"route_target":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"export_route_target_list_route_target","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RouteTarget","GoType":"string"},"GoName":"RouteTarget","GoType":"[]string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteTargetList","CollectionType":"","Column":"","Item":null,"GoName":"ExportRouteTargetList","GoType":"RouteTargetList"}
		IsShared: data["is_shared"].(bool),

		//{"Title":"","Description":"When true, this virtual network is shared with all tenants.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"is_shared","Item":null,"GoName":"IsShared","GoType":"bool"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string"},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType"},"GoName":"Share","GoType":"[]*ShareType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2"}
		PBBEvpnEnable: data["pbb_evpn_enable"].(bool),

		//{"Title":"","Description":"Enable/Disable PBB EVPN tunneling on the network","SQL":"bool","Default":false,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"pbb_evpn_enable","Item":null,"GoName":"PBBEvpnEnable","GoType":"bool"}
		EcmpHashingIncludeFields: InterfaceToEcmpHashingIncludeFields(data["ecmp_hashing_include_fields"]),

		//{"Title":"","Description":"ECMP hashing config at global level.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"destination_ip":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"destination_ip","Item":null,"GoName":"DestinationIP","GoType":"bool"},"destination_port":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"destination_port","Item":null,"GoName":"DestinationPort","GoType":"bool"},"hashing_configured":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"hashing_configured","Item":null,"GoName":"HashingConfigured","GoType":"bool"},"ip_protocol":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_protocol","Item":null,"GoName":"IPProtocol","GoType":"bool"},"source_ip":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"source_ip","Item":null,"GoName":"SourceIP","GoType":"bool"},"source_port":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"source_port","Item":null,"GoName":"SourcePort","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/EcmpHashingIncludeFields","CollectionType":"","Column":"","Item":null,"GoName":"EcmpHashingIncludeFields","GoType":"EcmpHashingIncludeFields"}
		AddressAllocationMode: InterfaceToAddressAllocationModeType(data["address_allocation_mode"]),

		//{"Title":"","Description":"Address allocation mode for virtual network.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["user-defined-subnet-preferred","user-defined-subnet-only","flat-subnet-preferred","flat-subnet-only"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressAllocationModeType","CollectionType":"","Column":"address_allocation_mode","Item":null,"GoName":"AddressAllocationMode","GoType":"AddressAllocationModeType"}
		ImportRouteTargetList: InterfaceToRouteTargetList(data["import_route_target_list"]),

		//{"Title":"","Description":"List of route targets that are used as import for this virtual network.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"route_target":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"import_route_target_list_route_target","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RouteTarget","GoType":"string"},"GoName":"RouteTarget","GoType":"[]string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteTargetList","CollectionType":"","Column":"","Item":null,"GoName":"ImportRouteTargetList","GoType":"RouteTargetList"}
		Layer2ControlWord: data["layer2_control_word"].(bool),

		//{"Title":"","Description":"Enable/Disable adding control word to the Layer 2 encapsulation","SQL":"bool","Default":false,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"layer2_control_word","Item":null,"GoName":"Layer2ControlWord","GoType":"bool"}
		VirtualNetworkProperties: InterfaceToVirtualNetworkType(data["virtual_network_properties"]),

		//{"Title":"","Description":"Virtual network miscellaneous configurations.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"allow_transit":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"allow_transit","Item":null,"GoName":"AllowTransit","GoType":"bool"},"forwarding_mode":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["l2_l3","l2","l3"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ForwardingModeType","CollectionType":"","Column":"forwarding_mode","Item":null,"GoName":"ForwardingMode","GoType":"ForwardingModeType"},"mirror_destination":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mirror_destination","Item":null,"GoName":"MirrorDestination","GoType":"bool"},"network_id":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"system-only","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"network_id","Item":null,"GoName":"NetworkID","GoType":"int"},"rpf":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["enable","disable"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RpfModeType","CollectionType":"","Column":"rpf","Item":null,"GoName":"RPF","GoType":"RpfModeType"},"vxlan_network_identifier":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"vxlan_network_identifier","Item":null,"GoName":"VxlanNetworkIdentifier","GoType":"VxlanNetworkIdentifierType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/VirtualNetworkType","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetworkProperties","GoType":"VirtualNetworkType"}
		MacLearningEnabled: data["mac_learning_enabled"].(bool),

		//{"Title":"","Description":"Enable MAC learning on the network","SQL":"bool","Default":false,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mac_learning_enabled","Item":null,"GoName":"MacLearningEnabled","GoType":"bool"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string"},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string"},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string"},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool"},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string"},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string"},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType"},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType"},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType"}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string"}
		ProviderProperties: InterfaceToProviderDetails(data["provider_properties"]),

		//{"Title":"","Description":"Virtual network is provider network. Specifies VLAN tag and physical network name.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"physical_network":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_network","Item":null,"GoName":"PhysicalNetwork","GoType":"string"},"segmentation_id":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"segmentation_id","Item":null,"GoName":"SegmentationID","GoType":"VlanIdType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ProviderDetails","CollectionType":"","Column":"","Item":null,"GoName":"ProviderProperties","GoType":"ProviderDetails"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string"},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair"},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs"}
		RouteTargetList: InterfaceToRouteTargetList(data["route_target_list"]),

		//{"Title":"","Description":"List of route targets that are used as both import and export for this virtual network.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"route_target":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"route_target","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RouteTarget","GoType":"string"},"GoName":"RouteTarget","GoType":"[]string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteTargetList","CollectionType":"","Column":"","Item":null,"GoName":"RouteTargetList","GoType":"RouteTargetList"}
		PBBEtreeEnable: data["pbb_etree_enable"].(bool),

		//{"Title":"","Description":"Enable/Disable PBB ETREE mode on the network","SQL":"bool","Default":false,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"pbb_etree_enable","Item":null,"GoName":"PBBEtreeEnable","GoType":"bool"}
		PortSecurityEnabled: data["port_security_enabled"].(bool),

		//{"Title":"","Description":"Port security status on the network","SQL":"bool","Default":true,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"port_security_enabled","Item":null,"GoName":"PortSecurityEnabled","GoType":"bool"}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string"}
		FloodUnknownUnicast: data["flood_unknown_unicast"].(bool),

		//{"Title":"","Description":"When true, packets with unknown unicast MAC address are flooded within the network. Default they are dropped.","SQL":"bool","Default":false,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"flood_unknown_unicast","Item":null,"GoName":"FloodUnknownUnicast","GoType":"bool"}
		ExternalIpam: data["external_ipam"].(bool),

		//{"Title":"","Description":"IP address assignment to VM is done statically, outside of (external to) Contrail Ipam. vCenter only feature.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"external_ipam","Item":null,"GoName":"ExternalIpam","GoType":"bool"}
		MultiPolicyServiceChainsEnabled: data["multi_policy_service_chains_enabled"].(bool),

		//{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"multi_policy_service_chains_enabled","Item":null,"GoName":"MultiPolicyServiceChainsEnabled","GoType":"bool"}
		MacLimitControl: InterfaceToMACLimitControlType(data["mac_limit_control"]),

		//{"Title":"","Description":"MAC limit control on the network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"mac_limit":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mac_limit","Item":null,"GoName":"MacLimit","GoType":"int"},"mac_limit_action":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["log","alarm","shutdown","drop"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitExceedActionType","CollectionType":"","Column":"mac_limit_action","Item":null,"GoName":"MacLimitAction","GoType":"MACLimitExceedActionType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitControlType","CollectionType":"","Column":"","Item":null,"GoName":"MacLimitControl","GoType":"MACLimitControlType"}
		RouterExternal: data["router_external"].(bool),

		//{"Title":"","Description":"When true, this virtual network is openstack router external network.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"router_external","Item":null,"GoName":"RouterExternal","GoType":"bool"}
		MacMoveControl: InterfaceToMACMoveLimitControlType(data["mac_move_control"]),

		//{"Title":"","Description":"MAC move control on the network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"mac_move_limit":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mac_move_limit","Item":null,"GoName":"MacMoveLimit","GoType":"int"},"mac_move_limit_action":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["log","alarm","shutdown","drop"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitExceedActionType","CollectionType":"","Column":"mac_move_limit_action","Item":null,"GoName":"MacMoveLimitAction","GoType":"MACLimitExceedActionType"},"mac_move_time_window":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":60,"Ref":"types.json#/definitions/MACMoveTimeWindow","CollectionType":"","Column":"mac_move_time_window","Item":null,"GoName":"MacMoveTimeWindow","GoType":"MACMoveTimeWindow"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACMoveLimitControlType","CollectionType":"","Column":"","Item":null,"GoName":"MacMoveControl","GoType":"MACMoveLimitControlType"}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string"},"GoName":"FQName","GoType":"[]string"}
		VirtualNetworkNetworkID: InterfaceToVirtualNetworkIdType(data["virtual_network_network_id"]),

		//{"Title":"","Description":"System assigned unique 32 bit ID for every virtual network.","SQL":"int","Default":null,"Operation":"","Presence":"system-only","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4294967296,"Ref":"types.json#/definitions/VirtualNetworkIdType","CollectionType":"","Column":"virtual_network_network_id","Item":null,"GoName":"VirtualNetworkNetworkID","GoType":"VirtualNetworkIdType"}

	}
}

func InterfaceToVirtualNetworkSlice(data interface{}) []*VirtualNetwork {
	list := data.([]interface{})
	result := MakeVirtualNetworkSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetwork(item))
	}
	return result
}

func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
