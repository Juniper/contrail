package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeTelemetryResourceInfo makes TelemetryResourceInfo
func MakeTelemetryResourceInfo() *TelemetryResourceInfo{
    return &TelemetryResourceInfo{
    //TODO(nati): Apply default
    Path: "",
        Rate: "",
        Name: "",
        
    }
}

// MakeTelemetryResourceInfo makes TelemetryResourceInfo
func InterfaceToTelemetryResourceInfo(i interface{}) *TelemetryResourceInfo{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &TelemetryResourceInfo{
    //TODO(nati): Apply default
    Path: schema.InterfaceToString(m["path"]),
        Rate: schema.InterfaceToString(m["rate"]),
        Name: schema.InterfaceToString(m["name"]),
        
    }
}

// MakeTelemetryResourceInfoSlice() makes a slice of TelemetryResourceInfo
func MakeTelemetryResourceInfoSlice() []*TelemetryResourceInfo {
    return []*TelemetryResourceInfo{}
}

// InterfaceToTelemetryResourceInfoSlice() makes a slice of TelemetryResourceInfo
func InterfaceToTelemetryResourceInfoSlice(i interface{}) []*TelemetryResourceInfo {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*TelemetryResourceInfo{}
    for _, item := range list {
        result = append(result, InterfaceToTelemetryResourceInfo(item) )
    }
    return result
}



