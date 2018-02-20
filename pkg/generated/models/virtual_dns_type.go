package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeVirtualDnsType makes VirtualDnsType
func MakeVirtualDnsType() *VirtualDnsType{
    return &VirtualDnsType{
    //TODO(nati): Apply default
    FloatingIPRecord: "",
        DomainName: "",
        ExternalVisible: false,
        NextVirtualDNS: "",
        DynamicRecordsFromClient: false,
        ReverseResolution: false,
        DefaultTTLSeconds: 0,
        RecordOrder: "",
        
    }
}

// MakeVirtualDnsType makes VirtualDnsType
func InterfaceToVirtualDnsType(i interface{}) *VirtualDnsType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &VirtualDnsType{
    //TODO(nati): Apply default
    FloatingIPRecord: schema.InterfaceToString(m["floating_ip_record"]),
        DomainName: schema.InterfaceToString(m["domain_name"]),
        ExternalVisible: schema.InterfaceToBool(m["external_visible"]),
        NextVirtualDNS: schema.InterfaceToString(m["next_virtual_DNS"]),
        DynamicRecordsFromClient: schema.InterfaceToBool(m["dynamic_records_from_client"]),
        ReverseResolution: schema.InterfaceToBool(m["reverse_resolution"]),
        DefaultTTLSeconds: schema.InterfaceToInt64(m["default_ttl_seconds"]),
        RecordOrder: schema.InterfaceToString(m["record_order"]),
        
    }
}

// MakeVirtualDnsTypeSlice() makes a slice of VirtualDnsType
func MakeVirtualDnsTypeSlice() []*VirtualDnsType {
    return []*VirtualDnsType{}
}

// InterfaceToVirtualDnsTypeSlice() makes a slice of VirtualDnsType
func InterfaceToVirtualDnsTypeSlice(i interface{}) []*VirtualDnsType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*VirtualDnsType{}
    for _, item := range list {
        result = append(result, InterfaceToVirtualDnsType(item) )
    }
    return result
}



