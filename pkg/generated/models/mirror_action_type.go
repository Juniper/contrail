package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeMirrorActionType makes MirrorActionType
func MakeMirrorActionType() *MirrorActionType{
    return &MirrorActionType{
    //TODO(nati): Apply default
    NicAssistedMirroringVlan: 0,
        AnalyzerName: "",
        NHMode: "",
        JuniperHeader: false,
        UDPPort: 0,
        RoutingInstance: "",
        StaticNHHeader: MakeStaticMirrorNhType(),
        AnalyzerIPAddress: "",
        Encapsulation: "",
        AnalyzerMacAddress: "",
        NicAssistedMirroring: false,
        
    }
}

// MakeMirrorActionType makes MirrorActionType
func InterfaceToMirrorActionType(i interface{}) *MirrorActionType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &MirrorActionType{
    //TODO(nati): Apply default
    NicAssistedMirroringVlan: schema.InterfaceToInt64(m["nic_assisted_mirroring_vlan"]),
        AnalyzerName: schema.InterfaceToString(m["analyzer_name"]),
        NHMode: schema.InterfaceToString(m["nh_mode"]),
        JuniperHeader: schema.InterfaceToBool(m["juniper_header"]),
        UDPPort: schema.InterfaceToInt64(m["udp_port"]),
        RoutingInstance: schema.InterfaceToString(m["routing_instance"]),
        StaticNHHeader: InterfaceToStaticMirrorNhType(m["static_nh_header"]),
        AnalyzerIPAddress: schema.InterfaceToString(m["analyzer_ip_address"]),
        Encapsulation: schema.InterfaceToString(m["encapsulation"]),
        AnalyzerMacAddress: schema.InterfaceToString(m["analyzer_mac_address"]),
        NicAssistedMirroring: schema.InterfaceToBool(m["nic_assisted_mirroring"]),
        
    }
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
    return []*MirrorActionType{}
}

// InterfaceToMirrorActionTypeSlice() makes a slice of MirrorActionType
func InterfaceToMirrorActionTypeSlice(i interface{}) []*MirrorActionType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*MirrorActionType{}
    for _, item := range list {
        result = append(result, InterfaceToMirrorActionType(item) )
    }
    return result
}



