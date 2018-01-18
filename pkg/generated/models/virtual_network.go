package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list,omitempty"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties,omitempty"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	RouterExternal                  bool                      `json:"router_external"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list,omitempty"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties,omitempty"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`
	IsShared                        bool                      `json:"is_shared"`
	ParentType                      string                    `json:"parent_type,omitempty"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id,omitempty"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode,omitempty"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	DisplayName                     string                    `json:"display_name,omitempty"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control,omitempty"`
	FQName                          []string                  `json:"fq_name,omitempty"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list,omitempty"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`

	SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
	NetworkPolicyRefs         []*VirtualNetworkNetworkPolicyRef         `json:"network_policy_refs,omitempty"`
	QosConfigRefs             []*VirtualNetworkQosConfigRef             `json:"qos_config_refs,omitempty"`
	RouteTableRefs            []*VirtualNetworkRouteTableRef            `json:"route_table_refs,omitempty"`
	VirtualNetworkRefs        []*VirtualNetworkVirtualNetworkRef        `json:"virtual_network_refs,omitempty"`
	BGPVPNRefs                []*VirtualNetworkBGPVPNRef                `json:"bgpvpn_refs,omitempty"`
	NetworkIpamRefs           []*VirtualNetworkNetworkIpamRef           `json:"network_ipam_refs,omitempty"`

	AccessControlLists []*AccessControlList `json:"access_control_lists,omitempty"`
	AliasIPPools       []*AliasIPPool       `json:"alias_ip_pools,omitempty"`
	BridgeDomains      []*BridgeDomain      `json:"bridge_domains,omitempty"`
	FloatingIPPools    []*FloatingIPPool    `json:"floating_ip_pools,omitempty"`
	RoutingInstances   []*RoutingInstance   `json:"routing_instances,omitempty"`
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

// String returns json representation of the object
func (model *VirtualNetwork) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualNetwork makes VirtualNetwork
func MakeVirtualNetwork() *VirtualNetwork {
	return &VirtualNetwork{
		//TODO(nati): Apply default
		ExportRouteTargetList: MakeRouteTargetList(),
		MacMoveControl:        MakeMACMoveLimitControlType(),
		Perms2:                MakePermType2(),
		VirtualNetworkProperties: MakeVirtualNetworkType(),
		PBBEvpnEnable:            false,
		RouterExternal:           false,
		RouteTargetList:          MakeRouteTargetList(),
		MacLearningEnabled:       false,
		PortSecurityEnabled:      false,
		ProviderProperties:       MakeProviderDetails(),
		PBBEtreeEnable:           false,
		IDPerms:                  MakeIdPermsType(),
		MultiPolicyServiceChainsEnabled: false,
		IsShared:                        false,
		ParentType:                      "",
		VirtualNetworkNetworkID:         MakeVirtualNetworkIdType(),
		AddressAllocationMode:           MakeAddressAllocationModeType(),
		Layer2ControlWord:               false,
		DisplayName:                     "",
		Annotations:                     MakeKeyValuePairs(),
		MacAgingTime:                    MakeMACAgingTime(),
		ParentUUID:                      "",
		EcmpHashingIncludeFields:        MakeEcmpHashingIncludeFields(),
		ExternalIpam:                    false,
		MacLimitControl:                 MakeMACLimitControlType(),
		FQName:                          []string{},
		UUID:                            "",
		ImportRouteTargetList: MakeRouteTargetList(),
		FloodUnknownUnicast:   false,
	}
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
