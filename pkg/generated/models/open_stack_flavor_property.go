package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeOpenStackFlavorProperty makes OpenStackFlavorProperty
func MakeOpenStackFlavorProperty() *OpenStackFlavorProperty{
    return &OpenStackFlavorProperty{
    //TODO(nati): Apply default
    ID: "",
        Links: MakeOpenStackLink(),
        
    }
}

// MakeOpenStackFlavorProperty makes OpenStackFlavorProperty
func InterfaceToOpenStackFlavorProperty(i interface{}) *OpenStackFlavorProperty{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &OpenStackFlavorProperty{
    //TODO(nati): Apply default
    ID: schema.InterfaceToString(m["id"]),
        Links: InterfaceToOpenStackLink(m["links"]),
        
    }
}

// MakeOpenStackFlavorPropertySlice() makes a slice of OpenStackFlavorProperty
func MakeOpenStackFlavorPropertySlice() []*OpenStackFlavorProperty {
    return []*OpenStackFlavorProperty{}
}

// InterfaceToOpenStackFlavorPropertySlice() makes a slice of OpenStackFlavorProperty
func InterfaceToOpenStackFlavorPropertySlice(i interface{}) []*OpenStackFlavorProperty {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*OpenStackFlavorProperty{}
    for _, item := range list {
        result = append(result, InterfaceToOpenStackFlavorProperty(item) )
    }
    return result
}



