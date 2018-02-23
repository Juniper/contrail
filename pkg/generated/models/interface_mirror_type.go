package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeInterfaceMirrorType makes InterfaceMirrorType
func MakeInterfaceMirrorType() *InterfaceMirrorType{
    return &InterfaceMirrorType{
    //TODO(nati): Apply default
    TrafficDirection: "",
        MirrorTo: MakeMirrorActionType(),
        
    }
}

// MakeInterfaceMirrorType makes InterfaceMirrorType
func InterfaceToInterfaceMirrorType(i interface{}) *InterfaceMirrorType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &InterfaceMirrorType{
    //TODO(nati): Apply default
    TrafficDirection: schema.InterfaceToString(m["traffic_direction"]),
        MirrorTo: InterfaceToMirrorActionType(m["mirror_to"]),
        
    }
}

// MakeInterfaceMirrorTypeSlice() makes a slice of InterfaceMirrorType
func MakeInterfaceMirrorTypeSlice() []*InterfaceMirrorType {
    return []*InterfaceMirrorType{}
}

// InterfaceToInterfaceMirrorTypeSlice() makes a slice of InterfaceMirrorType
func InterfaceToInterfaceMirrorTypeSlice(i interface{}) []*InterfaceMirrorType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*InterfaceMirrorType{}
    for _, item := range list {
        result = append(result, InterfaceToInterfaceMirrorType(item) )
    }
    return result
}



