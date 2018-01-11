package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control"`
	Annotations                     *KeyValuePairs            `json:"annotations"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	IsShared                        bool                      `json:"is_shared"`
	Perms2                          *PermType2                `json:"perms2"`
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties"`
	RouterExternal                  bool                      `json:"router_external"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	ParentType                      string                    `json:"parent_type"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control"`
	ParentUUID                      string                    `json:"parent_uuid"`
	FQName                          []string                  `json:"fq_name"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	DisplayName                     string                    `json:"display_name"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`
	UUID                            string                    `json:"uuid"`
	IDPerms                         *IdPermsType              `json:"id_perms"`

	QosConfigRefs             []*VirtualNetworkQosConfigRef             `json:"qos_config_refs"`
	RouteTableRefs            []*VirtualNetworkRouteTableRef            `json:"route_table_refs"`
	VirtualNetworkRefs        []*VirtualNetworkVirtualNetworkRef        `json:"virtual_network_refs"`
	BGPVPNRefs                []*VirtualNetworkBGPVPNRef                `json:"bgpvpn_refs"`
	NetworkIpamRefs           []*VirtualNetworkNetworkIpamRef           `json:"network_ipam_refs"`
	SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef `json:"security_logging_object_refs"`
	NetworkPolicyRefs         []*VirtualNetworkNetworkPolicyRef         `json:"network_policy_refs"`

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
		EcmpHashingIncludeFields:        MakeEcmpHashingIncludeFields(),
		ExternalIpam:                    false,
		ParentType:                      "",
		FQName:                          []string{},
		VirtualNetworkNetworkID:         MakeVirtualNetworkIdType(),
		MacLearningEnabled:              false,
		PBBEtreeEnable:                  false,
		MacMoveControl:                  MakeMACMoveLimitControlType(),
		ParentUUID:                      "",
		MultiPolicyServiceChainsEnabled: false,
		UUID:                     "",
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		MacAgingTime:             MakeMACAgingTime(),
		ProviderProperties:       MakeProviderDetails(),
		Annotations:              MakeKeyValuePairs(),
		AddressAllocationMode:    MakeAddressAllocationModeType(),
		ImportRouteTargetList:    MakeRouteTargetList(),
		Layer2ControlWord:        false,
		PortSecurityEnabled:      false,
		MacLimitControl:          MakeMACLimitControlType(),
		VirtualNetworkProperties: MakeVirtualNetworkType(),
		RouterExternal:           false,
		FloodUnknownUnicast:      false,
		IsShared:                 false,
		Perms2:                   MakePermType2(),
		RouteTargetList:          MakeRouteTargetList(),
		PBBEvpnEnable:            false,
		ExportRouteTargetList:    MakeRouteTargetList(),
	}
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
