package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control,omitempty"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time,omitempty"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled,omitempty"`
	Layer2ControlWord               bool                      `json:"layer2_control_word,omitempty"`
	RouterExternal                  bool                      `json:"router_external,omitempty"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
	ExternalIpam                    bool                      `json:"external_ipam,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	ParentType                      string                    `json:"parent_type,omitempty"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id,omitempty"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast,omitempty"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable,omitempty"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled,omitempty"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list,omitempty"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list,omitempty"`
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list,omitempty"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties,omitempty"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	FQName                          []string                  `json:"fq_name,omitempty"`
	DisplayName                     string                    `json:"display_name,omitempty"`
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties,omitempty"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode,omitempty"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable,omitempty"`
	IsShared                        bool                      `json:"is_shared,omitempty"`

	QosConfigRefs             []*VirtualNetworkQosConfigRef             `json:"qos_config_refs,omitempty"`
	RouteTableRefs            []*VirtualNetworkRouteTableRef            `json:"route_table_refs,omitempty"`
	VirtualNetworkRefs        []*VirtualNetworkVirtualNetworkRef        `json:"virtual_network_refs,omitempty"`
	BGPVPNRefs                []*VirtualNetworkBGPVPNRef                `json:"bgpvpn_refs,omitempty"`
	NetworkIpamRefs           []*VirtualNetworkNetworkIpamRef           `json:"network_ipam_refs,omitempty"`
	SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
	NetworkPolicyRefs         []*VirtualNetworkNetworkPolicyRef         `json:"network_policy_refs,omitempty"`

	AccessControlLists []*AccessControlList `json:"access_control_lists,omitempty"`
	AliasIPPools       []*AliasIPPool       `json:"alias_ip_pools,omitempty"`
	BridgeDomains      []*BridgeDomain      `json:"bridge_domains,omitempty"`
	FloatingIPPools    []*FloatingIPPool    `json:"floating_ip_pools,omitempty"`
	RoutingInstances   []*RoutingInstance   `json:"routing_instances,omitempty"`
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

// VirtualNetworkRouteTableRef references each other
type VirtualNetworkRouteTableRef struct {
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
		IsShared:                        false,
		MacMoveControl:                  MakeMACMoveLimitControlType(),
		MultiPolicyServiceChainsEnabled: false,
		UUID:                     "",
		MacAgingTime:             MakeMACAgingTime(),
		MacLearningEnabled:       false,
		Layer2ControlWord:        false,
		RouterExternal:           false,
		IDPerms:                  MakeIdPermsType(),
		Annotations:              MakeKeyValuePairs(),
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		ExternalIpam:             false,
		Perms2:                   MakePermType2(),
		ParentType:               "",
		VirtualNetworkNetworkID:  MakeVirtualNetworkIdType(),
		FloodUnknownUnicast:      false,
		PBBEtreeEnable:           false,
		PortSecurityEnabled:      false,
		ImportRouteTargetList:    MakeRouteTargetList(),
		RouteTargetList:          MakeRouteTargetList(),
		ExportRouteTargetList:    MakeRouteTargetList(),
		ProviderProperties:       MakeProviderDetails(),
		MacLimitControl:          MakeMACLimitControlType(),
		ParentUUID:               "",
		FQName:                   []string{},
		DisplayName:              "",
		VirtualNetworkProperties: MakeVirtualNetworkType(),
		AddressAllocationMode:    MakeAddressAllocationModeType(),
		PBBEvpnEnable:            false,
	}
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
