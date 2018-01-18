package models

// VirtualNetwork

import "encoding/json"

// VirtualNetwork
type VirtualNetwork struct {
	ExportRouteTargetList           *RouteTargetList          `json:"export_route_target_list,omitempty"`
	MacMoveControl                  *MACMoveLimitControlType  `json:"mac_move_control,omitempty"`
	IsShared                        bool                      `json:"is_shared"`
	EcmpHashingIncludeFields        *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	DisplayName                     string                    `json:"display_name,omitempty"`
	PBBEtreeEnable                  bool                      `json:"pbb_etree_enable"`
	Layer2ControlWord               bool                      `json:"layer2_control_word"`
	PortSecurityEnabled             bool                      `json:"port_security_enabled"`
	MultiPolicyServiceChainsEnabled bool                      `json:"multi_policy_service_chains_enabled"`
	MacAgingTime                    MACAgingTime              `json:"mac_aging_time,omitempty"`
	ProviderProperties              *ProviderDetails          `json:"provider_properties,omitempty"`
	MacLearningEnabled              bool                      `json:"mac_learning_enabled"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	FQName                          []string                  `json:"fq_name,omitempty"`
	VirtualNetworkProperties        *VirtualNetworkType       `json:"virtual_network_properties,omitempty"`
	VirtualNetworkNetworkID         VirtualNetworkIdType      `json:"virtual_network_network_id,omitempty"`
	RouteTargetList                 *RouteTargetList          `json:"route_target_list,omitempty"`
	MacLimitControl                 *MACLimitControlType      `json:"mac_limit_control,omitempty"`
	ParentType                      string                    `json:"parent_type,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	PBBEvpnEnable                   bool                      `json:"pbb_evpn_enable"`
	RouterExternal                  bool                      `json:"router_external"`
	FloodUnknownUnicast             bool                      `json:"flood_unknown_unicast"`
	ExternalIpam                    bool                      `json:"external_ipam"`
	AddressAllocationMode           AddressAllocationModeType `json:"address_allocation_mode,omitempty"`
	ImportRouteTargetList           *RouteTargetList          `json:"import_route_target_list,omitempty"`

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
		PBBEvpnEnable:                   false,
		RouterExternal:                  false,
		FloodUnknownUnicast:             false,
		ExternalIpam:                    false,
		ParentUUID:                      "",
		AddressAllocationMode:           MakeAddressAllocationModeType(),
		ImportRouteTargetList:           MakeRouteTargetList(),
		ExportRouteTargetList:           MakeRouteTargetList(),
		MacMoveControl:                  MakeMACMoveLimitControlType(),
		IsShared:                        false,
		EcmpHashingIncludeFields:        MakeEcmpHashingIncludeFields(),
		PBBEtreeEnable:                  false,
		Layer2ControlWord:               false,
		PortSecurityEnabled:             false,
		MultiPolicyServiceChainsEnabled: false,
		Perms2:                   MakePermType2(),
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		MacAgingTime:             MakeMACAgingTime(),
		ProviderProperties:       MakeProviderDetails(),
		MacLearningEnabled:       false,
		Annotations:              MakeKeyValuePairs(),
		VirtualNetworkProperties: MakeVirtualNetworkType(),
		VirtualNetworkNetworkID:  MakeVirtualNetworkIdType(),
		RouteTargetList:          MakeRouteTargetList(),
		MacLimitControl:          MakeMACLimitControlType(),
		UUID:                     "",
		FQName:                   []string{},
		ParentType:               "",
	}
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}
