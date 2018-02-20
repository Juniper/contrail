package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeFlowAgingTimeout makes FlowAgingTimeout
func MakeFlowAgingTimeout() *FlowAgingTimeout{
    return &FlowAgingTimeout{
    //TODO(nati): Apply default
    TimeoutInSeconds: 0,
        Protocol: "",
        Port: 0,
        
    }
}

// MakeFlowAgingTimeout makes FlowAgingTimeout
func InterfaceToFlowAgingTimeout(i interface{}) *FlowAgingTimeout{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &FlowAgingTimeout{
    //TODO(nati): Apply default
    TimeoutInSeconds: schema.InterfaceToInt64(m["timeout_in_seconds"]),
        Protocol: schema.InterfaceToString(m["protocol"]),
        Port: schema.InterfaceToInt64(m["port"]),
        
    }
}

// MakeFlowAgingTimeoutSlice() makes a slice of FlowAgingTimeout
func MakeFlowAgingTimeoutSlice() []*FlowAgingTimeout {
    return []*FlowAgingTimeout{}
}

// InterfaceToFlowAgingTimeoutSlice() makes a slice of FlowAgingTimeout
func InterfaceToFlowAgingTimeoutSlice(i interface{}) []*FlowAgingTimeout {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*FlowAgingTimeout{}
    for _, item := range list {
        result = append(result, InterfaceToFlowAgingTimeout(item) )
    }
    return result
}



