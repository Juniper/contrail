package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeProviderDetails makes ProviderDetails
func MakeProviderDetails() *ProviderDetails{
    return &ProviderDetails{
    //TODO(nati): Apply default
    SegmentationID: 0,
        PhysicalNetwork: "",
        
    }
}

// MakeProviderDetails makes ProviderDetails
func InterfaceToProviderDetails(i interface{}) *ProviderDetails{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ProviderDetails{
    //TODO(nati): Apply default
    SegmentationID: schema.InterfaceToInt64(m["segmentation_id"]),
        PhysicalNetwork: schema.InterfaceToString(m["physical_network"]),
        
    }
}

// MakeProviderDetailsSlice() makes a slice of ProviderDetails
func MakeProviderDetailsSlice() []*ProviderDetails {
    return []*ProviderDetails{}
}

// InterfaceToProviderDetailsSlice() makes a slice of ProviderDetails
func InterfaceToProviderDetailsSlice(i interface{}) []*ProviderDetails {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ProviderDetails{}
    for _, item := range list {
        result = append(result, InterfaceToProviderDetails(item) )
    }
    return result
}



