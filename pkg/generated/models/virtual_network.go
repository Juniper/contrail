package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties,omitempty"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id,omitempty"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	RouterExternal                  bool                      `json:"router_external"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	ParentType                      string                    `json:"parent_type,omitempty"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time,omitempty"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties,omitempty"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`
	IsShared                        bool                      `json:"is_shared"`
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list,omitempty"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode,omitempty"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list,omitempty"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	FQName                          []string                  `json:"fq_name,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control,omitempty"`
	DisplayName                     string                    `json:"display_name,omitempty"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control,omitempty"`

	VirtualNetworkRefs        []*VirtualNetworkVirtualNetworkRef        `json:"virtual_network_refs,omitempty"`
	BGPVPNRefs                []*VirtualNetworkBGPVPNRef                `json:"bgpvpn_refs,omitempty"`
	NetworkIpamRefs           []*VirtualNetworkNetworkIpamRef           `json:"network_ipam_refs,omitempty"`
	SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
	NetworkPolicyRefs         []*VirtualNetworkNetworkPolicyRef         `json:"network_policy_refs,omitempty"`
	QosConfigRefs             []*VirtualNetworkQosConfigRef             `json:"qos_config_refs,omitempty"`
	RouteTableRefs            []*VirtualNetworkRouteTableRef            `json:"route_table_refs,omitempty"`

	AccessControlLists []*AccessControlList `json:"access_control_lists,omitempty"`
	AliasIPPools       []*AliasIPPool       `json:"alias_ip_pools,omitempty"`
	BridgeDomains      []*BridgeDomain      `json:"bridge_domains,omitempty"`
	FloatingIPPools    []*FloatingIPPool    `json:"floating_ip_pools,omitempty"`
	RoutingInstances   []*RoutingInstance   `json:"routing_instances,omitempty"`
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

// String returns json representation of the object
func (model *VirtualNetwork) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualNetwork makes VirtualNetwork
func MakeVirtualNetwork() *VirtualNetwork {
	return &VirtualNetwork{
		//TODO(nati): Apply default
		ExternalIpam:                    false,
		MacLimitControl:                 MakeMACLimitControlType(),
		RouteTargetList:                 MakeRouteTargetList(),
		VirtualNetworkProperties:        MakeVirtualNetworkType(),
		VirtualNetworkNetworkID:         MakeVirtualNetworkIdType(),
		MultiPolicyServiceChainsEnabled: false,
		IDPerms:                  MakeIdPermsType(),
		Perms2:                   MakePermType2(),
		RouterExternal:           false,
		PBBEtreeEnable:           false,
		Annotations:              MakeKeyValuePairs(),
		Layer2ControlWord:        false,
		ParentType:               "",
		FloodUnknownUnicast:      false,
		IsShared:                 false,
		ParentUUID:               "",
		MacAgingTime:             MakeMACAgingTime(),
		ProviderProperties:       MakeProviderDetails(),
		ImportRouteTargetList:    MakeRouteTargetList(),
		MacLearningEnabled:       false,
		ExportRouteTargetList:    MakeRouteTargetList(),
		PortSecurityEnabled:      false,
		AddressAllocationMode:    MakeAddressAllocationModeType(),
		PBBEvpnEnable:            false,
		FQName:                   []string{},
		UUID:                     "",
		DisplayName:              "",
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		MacMoveControl:           MakeMACMoveLimitControlType(),
	}
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
