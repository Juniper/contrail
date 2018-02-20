package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeOpenStackImageProperty makes OpenStackImageProperty
func MakeOpenStackImageProperty() *OpenStackImageProperty{
    return &OpenStackImageProperty{
    //TODO(nati): Apply default
    ID: "",
        Links: MakeOpenStackLink(),
        
    }
}

// MakeOpenStackImageProperty makes OpenStackImageProperty
func InterfaceToOpenStackImageProperty(i interface{}) *OpenStackImageProperty{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &OpenStackImageProperty{
    //TODO(nati): Apply default
    ID: schema.InterfaceToString(m["id"]),
        Links: InterfaceToOpenStackLink(m["links"]),
        
    }
}

// MakeOpenStackImagePropertySlice() makes a slice of OpenStackImageProperty
func MakeOpenStackImagePropertySlice() []*OpenStackImageProperty {
    return []*OpenStackImageProperty{}
}

// InterfaceToOpenStackImagePropertySlice() makes a slice of OpenStackImageProperty
func InterfaceToOpenStackImagePropertySlice(i interface{}) []*OpenStackImageProperty {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*OpenStackImageProperty{}
    for _, item := range list {
        result = append(result, InterfaceToOpenStackImageProperty(item) )
    }
    return result
}



