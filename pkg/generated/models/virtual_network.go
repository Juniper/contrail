package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	UUID                            string                    `json:"uuid,omitempty"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list,omitempty"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list,omitempty"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	DisplayName                     string                    `json:"display_name,omitempty"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list,omitempty"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time,omitempty"`
	FQName                          []string                  `json:"fq_name,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties,omitempty"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control,omitempty"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control,omitempty"`
	ParentType                      string                    `json:"parent_type,omitempty"`
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties,omitempty"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode,omitempty"`
	IsShared                        bool                      `json:"is_shared"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id,omitempty"`
	RouterExternal                  bool                      `json:"router_external"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`

	NetworkIpamRefs           []*VirtualNetworkNetworkIpamRef           `json:"network_ipam_refs,omitempty"`
	SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
	NetworkPolicyRefs         []*VirtualNetworkNetworkPolicyRef         `json:"network_policy_refs,omitempty"`
	QosConfigRefs             []*VirtualNetworkQosConfigRef             `json:"qos_config_refs,omitempty"`
	RouteTableRefs            []*VirtualNetworkRouteTableRef            `json:"route_table_refs,omitempty"`
	VirtualNetworkRefs        []*VirtualNetworkVirtualNetworkRef        `json:"virtual_network_refs,omitempty"`
	BGPVPNRefs                []*VirtualNetworkBGPVPNRef                `json:"bgpvpn_refs,omitempty"`

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
		MacAgingTime:                    MakeMACAgingTime(),
		FQName:                          []string{},
		DisplayName:                     "",
		ImportRouteTargetList:           MakeRouteTargetList(),
		PBBEtreeEnable:                  false,
		MacLimitControl:                 MakeMACLimitControlType(),
		Perms2:                          MakePermType2(),
		ProviderProperties:              MakeProviderDetails(),
		ExternalIpam:                    false,
		EcmpHashingIncludeFields:        MakeEcmpHashingIncludeFields(),
		AddressAllocationMode:           MakeAddressAllocationModeType(),
		PBBEvpnEnable:                   false,
		PortSecurityEnabled:             false,
		MacMoveControl:                  MakeMACMoveLimitControlType(),
		ParentType:                      "",
		VirtualNetworkProperties:        MakeVirtualNetworkType(),
		RouterExternal:                  false,
		MultiPolicyServiceChainsEnabled: false,
		IsShared:                        false,
		VirtualNetworkNetworkID:         MakeVirtualNetworkIdType(),
		MacLearningEnabled:              false,
		Layer2ControlWord:               false,
		UUID:                            "",
		IDPerms:                         MakeIdPermsType(),
		Annotations:                     MakeKeyValuePairs(),
		RouteTargetList:                 MakeRouteTargetList(),
		FloodUnknownUnicast:             false,
		ParentUUID:                      "",
		ExportRouteTargetList:           MakeRouteTargetList(),
	}
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
