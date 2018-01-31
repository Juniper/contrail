package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control,omitempty"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode,omitempty"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties,omitempty"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	FQName                          []string                  `json:"fq_name,omitempty"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
	RouterExternal                  bool                      `json:"router_external"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time,omitempty"`
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list,omitempty"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control,omitempty"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	IsShared                        bool                      `json:"is_shared"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list,omitempty"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties,omitempty"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id,omitempty"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	ParentType                      string                    `json:"parent_type,omitempty"`
	DisplayName                     string                    `json:"display_name,omitempty"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`

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

// VirtualNetworkBGPVPNRef references each other
type VirtualNetworkBGPVPNRef struct {
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
		IsShared:          false,
		PBBEvpnEnable:     false,
		Layer2ControlWord: false,
		ParentUUID:        "",
		UUID:              "",
		ImportRouteTargetList:           MakeRouteTargetList(),
		RouteTargetList:                 MakeRouteTargetList(),
		PortSecurityEnabled:             false,
		ParentType:                      "",
		DisplayName:                     "",
		Annotations:                     MakeKeyValuePairs(),
		VirtualNetworkProperties:        MakeVirtualNetworkType(),
		VirtualNetworkNetworkID:         MakeVirtualNetworkIdType(),
		Perms2:                          MakePermType2(),
		MacLimitControl:                 MakeMACLimitControlType(),
		IDPerms:                         MakeIdPermsType(),
		AddressAllocationMode:           MakeAddressAllocationModeType(),
		ExternalIpam:                    false,
		FloodUnknownUnicast:             false,
		MultiPolicyServiceChainsEnabled: false,
		ProviderProperties:              MakeProviderDetails(),
		MacLearningEnabled:              false,
		MacAgingTime:                    MakeMACAgingTime(),
		ExportRouteTargetList:           MakeRouteTargetList(),
		PBBEtreeEnable:                  false,
		MacMoveControl:                  MakeMACMoveLimitControlType(),
		FQName:                          []string{},
		EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
		RouterExternal:           false,
	}
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
