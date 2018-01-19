package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties,omitempty"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control,omitempty"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control,omitempty"`
	IsShared                        bool                      `json:"is_shared"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties,omitempty"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`
	ParentType                      string                    `json:"parent_type,omitempty"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id,omitempty"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list,omitempty"`
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list,omitempty"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode,omitempty"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list,omitempty"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`
	FQName                          []string                  `json:"fq_name,omitempty"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	RouterExternal                  bool                      `json:"router_external"`
	DisplayName                     string                    `json:"display_name,omitempty"`

	NetworkPolicyRefs         []*VirtualNetworkNetworkPolicyRef         `json:"network_policy_refs,omitempty"`
	QosConfigRefs             []*VirtualNetworkQosConfigRef             `json:"qos_config_refs,omitempty"`
	RouteTableRefs            []*VirtualNetworkRouteTableRef            `json:"route_table_refs,omitempty"`
	VirtualNetworkRefs        []*VirtualNetworkVirtualNetworkRef        `json:"virtual_network_refs,omitempty"`
	BGPVPNRefs                []*VirtualNetworkBGPVPNRef                `json:"bgpvpn_refs,omitempty"`
	NetworkIpamRefs           []*VirtualNetworkNetworkIpamRef           `json:"network_ipam_refs,omitempty"`
	SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`

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
		FloodUnknownUnicast:             false,
		ParentType:                      "",
		ProviderProperties:              MakeProviderDetails(),
		MacLearningEnabled:              false,
		ExportRouteTargetList:           MakeRouteTargetList(),
		Layer2ControlWord:               false,
		ExternalIpam:                    false,
		PortSecurityEnabled:             false,
		Perms2:                          MakePermType2(),
		VirtualNetworkNetworkID:         MakeVirtualNetworkIdType(),
		RouteTargetList:                 MakeRouteTargetList(),
		ImportRouteTargetList:           MakeRouteTargetList(),
		PBBEtreeEnable:                  false,
		MultiPolicyServiceChainsEnabled: false,
		FQName: []string{},
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		AddressAllocationMode:    MakeAddressAllocationModeType(),
		MacAgingTime:             MakeMACAgingTime(),
		UUID:                     "",
		ParentUUID:               "",
		RouterExternal:           false,
		DisplayName:              "",
		MacMoveControl:           MakeMACMoveLimitControlType(),
		MacLimitControl:          MakeMACLimitControlType(),
		IsShared:                 false,
		VirtualNetworkProperties: MakeVirtualNetworkType(),
		PBBEvpnEnable:            false,
		Annotations:              MakeKeyValuePairs(),
		IDPerms:                  MakeIdPermsType(),
	}
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
