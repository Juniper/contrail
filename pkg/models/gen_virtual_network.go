package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualNetwork makes VirtualNetwork
// nolint
func MakeVirtualNetwork() *VirtualNetwork {
	return &VirtualNetwork{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		VirtualNetworkProperties:        MakeVirtualNetworkType(),
		EcmpHashingIncludeFields:        MakeEcmpHashingIncludeFields(),
		VirtualNetworkNetworkID:         0,
		AddressAllocationMode:           "",
		PBBEvpnEnable:                   false,
		RouterExternal:                  false,
		ImportRouteTargetList:           MakeRouteTargetList(),
		MacAgingTime:                    0,
		ProviderProperties:              MakeProviderDetails(),
		RouteTargetList:                 MakeRouteTargetList(),
		MacLearningEnabled:              false,
		ExportRouteTargetList:           MakeRouteTargetList(),
		FloodUnknownUnicast:             false,
		PBBEtreeEnable:                  false,
		Layer2ControlWord:               false,
		ExternalIpam:                    false,
		PortSecurityEnabled:             false,
		MacMoveControl:                  MakeMACMoveLimitControlType(),
		MultiPolicyServiceChainsEnabled: false,
		MacLimitControl:                 MakeMACLimitControlType(),
		IsShared:                        false,
	}
}

// MakeVirtualNetwork makes VirtualNetwork
// nolint
func InterfaceToVirtualNetwork(i interface{}) *VirtualNetwork {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualNetwork{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		VirtualNetworkProperties:        InterfaceToVirtualNetworkType(m["virtual_network_properties"]),
		EcmpHashingIncludeFields:        InterfaceToEcmpHashingIncludeFields(m["ecmp_hashing_include_fields"]),
		VirtualNetworkNetworkID:         common.InterfaceToInt64(m["virtual_network_network_id"]),
		AddressAllocationMode:           common.InterfaceToString(m["address_allocation_mode"]),
		PBBEvpnEnable:                   common.InterfaceToBool(m["pbb_evpn_enable"]),
		RouterExternal:                  common.InterfaceToBool(m["router_external"]),
		ImportRouteTargetList:           InterfaceToRouteTargetList(m["import_route_target_list"]),
		MacAgingTime:                    common.InterfaceToInt64(m["mac_aging_time"]),
		ProviderProperties:              InterfaceToProviderDetails(m["provider_properties"]),
		RouteTargetList:                 InterfaceToRouteTargetList(m["route_target_list"]),
		MacLearningEnabled:              common.InterfaceToBool(m["mac_learning_enabled"]),
		ExportRouteTargetList:           InterfaceToRouteTargetList(m["export_route_target_list"]),
		FloodUnknownUnicast:             common.InterfaceToBool(m["flood_unknown_unicast"]),
		PBBEtreeEnable:                  common.InterfaceToBool(m["pbb_etree_enable"]),
		Layer2ControlWord:               common.InterfaceToBool(m["layer2_control_word"]),
		ExternalIpam:                    common.InterfaceToBool(m["external_ipam"]),
		PortSecurityEnabled:             common.InterfaceToBool(m["port_security_enabled"]),
		MacMoveControl:                  InterfaceToMACMoveLimitControlType(m["mac_move_control"]),
		MultiPolicyServiceChainsEnabled: common.InterfaceToBool(m["multi_policy_service_chains_enabled"]),
		MacLimitControl:                 InterfaceToMACLimitControlType(m["mac_limit_control"]),
		IsShared:                        common.InterfaceToBool(m["is_shared"]),
	}
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
// nolint
func MakeVirtualNetworkSlice() []*VirtualNetwork {
	return []*VirtualNetwork{}
}

// InterfaceToVirtualNetworkSlice() makes a slice of VirtualNetwork
// nolint
func InterfaceToVirtualNetworkSlice(i interface{}) []*VirtualNetwork {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualNetwork{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetwork(item))
	}
	return result
}
