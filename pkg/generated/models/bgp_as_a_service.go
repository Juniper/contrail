package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeBGPAsAService makes BGPAsAService
func MakeBGPAsAService() *BGPAsAService{
    return &BGPAsAService{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        BgpaasShared: false,
        BgpaasSessionAttributes: "",
        BgpaasSuppressRouteAdvertisement: false,
        BgpaasIpv4MappedIpv6Nexthop: false,
        BgpaasIPAddress: "",
        AutonomousSystem: 0,
        
    }
}

// MakeBGPAsAService makes BGPAsAService
func InterfaceToBGPAsAService(i interface{}) *BGPAsAService{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &BGPAsAService{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        BgpaasShared: schema.InterfaceToBool(m["bgpaas_shared"]),
        BgpaasSessionAttributes: schema.InterfaceToString(m["bgpaas_session_attributes"]),
        BgpaasSuppressRouteAdvertisement: schema.InterfaceToBool(m["bgpaas_suppress_route_advertisement"]),
        BgpaasIpv4MappedIpv6Nexthop: schema.InterfaceToBool(m["bgpaas_ipv4_mapped_ipv6_nexthop"]),
        BgpaasIPAddress: schema.InterfaceToString(m["bgpaas_ip_address"]),
        AutonomousSystem: schema.InterfaceToInt64(m["autonomous_system"]),
        
    }
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
    return []*BGPAsAService{}
}

// InterfaceToBGPAsAServiceSlice() makes a slice of BGPAsAService
func InterfaceToBGPAsAServiceSlice(i interface{}) []*BGPAsAService {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*BGPAsAService{}
    for _, item := range list {
        result = append(result, InterfaceToBGPAsAService(item) )
    }
    return result
}



