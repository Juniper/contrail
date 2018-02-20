package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeTimerType makes TimerType
func MakeTimerType() *TimerType{
    return &TimerType{
    //TODO(nati): Apply default
    StartTime: "",
        OffInterval: "",
        OnInterval: "",
        EndTime: "",
        
    }
}

// MakeTimerType makes TimerType
func InterfaceToTimerType(i interface{}) *TimerType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &TimerType{
    //TODO(nati): Apply default
    StartTime: schema.InterfaceToString(m["start_time"]),
        OffInterval: schema.InterfaceToString(m["off_interval"]),
        OnInterval: schema.InterfaceToString(m["on_interval"]),
        EndTime: schema.InterfaceToString(m["end_time"]),
        
    }
}

// MakeTimerTypeSlice() makes a slice of TimerType
func MakeTimerTypeSlice() []*TimerType {
    return []*TimerType{}
}

// InterfaceToTimerTypeSlice() makes a slice of TimerType
func InterfaceToTimerTypeSlice(i interface{}) []*TimerType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*TimerType{}
    for _, item := range list {
        result = append(result, InterfaceToTimerType(item) )
    }
    return result
}



