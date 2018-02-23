package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeTelemetryStateInfo makes TelemetryStateInfo
func MakeTelemetryStateInfo() *TelemetryStateInfo{
    return &TelemetryStateInfo{
    //TODO(nati): Apply default
    
            
                Resource:  MakeTelemetryResourceInfoSlice(),
            
        ServerPort: 0,
        ServerIP: "",
        
    }
}

// MakeTelemetryStateInfo makes TelemetryStateInfo
func InterfaceToTelemetryStateInfo(i interface{}) *TelemetryStateInfo{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &TelemetryStateInfo{
    //TODO(nati): Apply default
    
            
                Resource:  InterfaceToTelemetryResourceInfoSlice(m["resource"]),
            
        ServerPort: schema.InterfaceToInt64(m["server_port"]),
        ServerIP: schema.InterfaceToString(m["server_ip"]),
        
    }
}

// MakeTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
    return []*TelemetryStateInfo{}
}

// InterfaceToTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
func InterfaceToTelemetryStateInfoSlice(i interface{}) []*TelemetryStateInfo {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*TelemetryStateInfo{}
    for _, item := range list {
        result = append(result, InterfaceToTelemetryStateInfo(item) )
    }
    return result
}



