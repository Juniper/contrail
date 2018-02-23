package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeOpenStackLink makes OpenStackLink
func MakeOpenStackLink() *OpenStackLink{
    return &OpenStackLink{
    //TODO(nati): Apply default
    Href: "",
        Rel: "",
        Type: "",
        
    }
}

// MakeOpenStackLink makes OpenStackLink
func InterfaceToOpenStackLink(i interface{}) *OpenStackLink{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &OpenStackLink{
    //TODO(nati): Apply default
    Href: schema.InterfaceToString(m["href"]),
        Rel: schema.InterfaceToString(m["rel"]),
        Type: schema.InterfaceToString(m["type"]),
        
    }
}

// MakeOpenStackLinkSlice() makes a slice of OpenStackLink
func MakeOpenStackLinkSlice() []*OpenStackLink {
    return []*OpenStackLink{}
}

// InterfaceToOpenStackLinkSlice() makes a slice of OpenStackLink
func InterfaceToOpenStackLinkSlice(i interface{}) []*OpenStackLink {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*OpenStackLink{}
    for _, item := range list {
        result = append(result, InterfaceToOpenStackLink(item) )
    }
    return result
}



