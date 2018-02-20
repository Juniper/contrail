package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeLinklocalServiceEntryType makes LinklocalServiceEntryType
func MakeLinklocalServiceEntryType() *LinklocalServiceEntryType{
    return &LinklocalServiceEntryType{
    //TODO(nati): Apply default
    IPFabricServiceIP: []string{},
        LinklocalServiceName: "",
        LinklocalServiceIP: "",
        IPFabricServicePort: 0,
        IPFabricDNSServiceName: "",
        LinklocalServicePort: 0,
        
    }
}

// MakeLinklocalServiceEntryType makes LinklocalServiceEntryType
func InterfaceToLinklocalServiceEntryType(i interface{}) *LinklocalServiceEntryType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &LinklocalServiceEntryType{
    //TODO(nati): Apply default
    IPFabricServiceIP: schema.InterfaceToStringList(m["ip_fabric_service_ip"]),
        LinklocalServiceName: schema.InterfaceToString(m["linklocal_service_name"]),
        LinklocalServiceIP: schema.InterfaceToString(m["linklocal_service_ip"]),
        IPFabricServicePort: schema.InterfaceToInt64(m["ip_fabric_service_port"]),
        IPFabricDNSServiceName: schema.InterfaceToString(m["ip_fabric_DNS_service_name"]),
        LinklocalServicePort: schema.InterfaceToInt64(m["linklocal_service_port"]),
        
    }
}

// MakeLinklocalServiceEntryTypeSlice() makes a slice of LinklocalServiceEntryType
func MakeLinklocalServiceEntryTypeSlice() []*LinklocalServiceEntryType {
    return []*LinklocalServiceEntryType{}
}

// InterfaceToLinklocalServiceEntryTypeSlice() makes a slice of LinklocalServiceEntryType
func InterfaceToLinklocalServiceEntryTypeSlice(i interface{}) []*LinklocalServiceEntryType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*LinklocalServiceEntryType{}
    for _, item := range list {
        result = append(result, InterfaceToLinklocalServiceEntryType(item) )
    }
    return result
}



