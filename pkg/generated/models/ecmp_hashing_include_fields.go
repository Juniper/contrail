package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeEcmpHashingIncludeFields makes EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFields() *EcmpHashingIncludeFields{
    return &EcmpHashingIncludeFields{
    //TODO(nati): Apply default
    DestinationIP: false,
        IPProtocol: false,
        SourceIP: false,
        HashingConfigured: false,
        SourcePort: false,
        DestinationPort: false,
        
    }
}

// MakeEcmpHashingIncludeFields makes EcmpHashingIncludeFields
func InterfaceToEcmpHashingIncludeFields(i interface{}) *EcmpHashingIncludeFields{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &EcmpHashingIncludeFields{
    //TODO(nati): Apply default
    DestinationIP: schema.InterfaceToBool(m["destination_ip"]),
        IPProtocol: schema.InterfaceToBool(m["ip_protocol"]),
        SourceIP: schema.InterfaceToBool(m["source_ip"]),
        HashingConfigured: schema.InterfaceToBool(m["hashing_configured"]),
        SourcePort: schema.InterfaceToBool(m["source_port"]),
        DestinationPort: schema.InterfaceToBool(m["destination_port"]),
        
    }
}

// MakeEcmpHashingIncludeFieldsSlice() makes a slice of EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFieldsSlice() []*EcmpHashingIncludeFields {
    return []*EcmpHashingIncludeFields{}
}

// InterfaceToEcmpHashingIncludeFieldsSlice() makes a slice of EcmpHashingIncludeFields
func InterfaceToEcmpHashingIncludeFieldsSlice(i interface{}) []*EcmpHashingIncludeFields {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*EcmpHashingIncludeFields{}
    for _, item := range list {
        result = append(result, InterfaceToEcmpHashingIncludeFields(item) )
    }
    return result
}



