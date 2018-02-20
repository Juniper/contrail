package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeVirtualNetwork makes VirtualNetwork
func MakeVirtualNetwork() *VirtualNetwork{
    return &VirtualNetwork{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        VirtualNetworkProperties: MakeVirtualNetworkType(),
        EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
        VirtualNetworkNetworkID: 0,
        AddressAllocationMode: "",
        PBBEvpnEnable: false,
        RouterExternal: false,
        ImportRouteTargetList: MakeRouteTargetList(),
        MacAgingTime: 0,
        ProviderProperties: MakeProviderDetails(),
        RouteTargetList: MakeRouteTargetList(),
        MacLearningEnabled: false,
        ExportRouteTargetList: MakeRouteTargetList(),
        FloodUnknownUnicast: false,
        PBBEtreeEnable: false,
        Layer2ControlWord: false,
        ExternalIpam: false,
        PortSecurityEnabled: false,
        MacMoveControl: MakeMACMoveLimitControlType(),
        MultiPolicyServiceChainsEnabled: false,
        MacLimitControl: MakeMACLimitControlType(),
        IsShared: false,
        
    }
}

// MakeVirtualNetwork makes VirtualNetwork
func InterfaceToVirtualNetwork(i interface{}) *VirtualNetwork{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &VirtualNetwork{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        VirtualNetworkProperties: InterfaceToVirtualNetworkType(m["virtual_network_properties"]),
        EcmpHashingIncludeFields: InterfaceToEcmpHashingIncludeFields(m["ecmp_hashing_include_fields"]),
        VirtualNetworkNetworkID: schema.InterfaceToInt64(m["virtual_network_network_id"]),
        AddressAllocationMode: schema.InterfaceToString(m["address_allocation_mode"]),
        PBBEvpnEnable: schema.InterfaceToBool(m["pbb_evpn_enable"]),
        RouterExternal: schema.InterfaceToBool(m["router_external"]),
        ImportRouteTargetList: InterfaceToRouteTargetList(m["import_route_target_list"]),
        MacAgingTime: schema.InterfaceToInt64(m["mac_aging_time"]),
        ProviderProperties: InterfaceToProviderDetails(m["provider_properties"]),
        RouteTargetList: InterfaceToRouteTargetList(m["route_target_list"]),
        MacLearningEnabled: schema.InterfaceToBool(m["mac_learning_enabled"]),
        ExportRouteTargetList: InterfaceToRouteTargetList(m["export_route_target_list"]),
        FloodUnknownUnicast: schema.InterfaceToBool(m["flood_unknown_unicast"]),
        PBBEtreeEnable: schema.InterfaceToBool(m["pbb_etree_enable"]),
        Layer2ControlWord: schema.InterfaceToBool(m["layer2_control_word"]),
        ExternalIpam: schema.InterfaceToBool(m["external_ipam"]),
        PortSecurityEnabled: schema.InterfaceToBool(m["port_security_enabled"]),
        MacMoveControl: InterfaceToMACMoveLimitControlType(m["mac_move_control"]),
        MultiPolicyServiceChainsEnabled: schema.InterfaceToBool(m["multi_policy_service_chains_enabled"]),
        MacLimitControl: InterfaceToMACLimitControlType(m["mac_limit_control"]),
        IsShared: schema.InterfaceToBool(m["is_shared"]),
        
    }
}

// MakeVirtualNetworkSlice() makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
    return []*VirtualNetwork{}
}

// InterfaceToVirtualNetworkSlice() makes a slice of VirtualNetwork
func InterfaceToVirtualNetworkSlice(i interface{}) []*VirtualNetwork {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*VirtualNetwork{}
    for _, item := range list {
        result = append(result, InterfaceToVirtualNetwork(item) )
    }
    return result
}



