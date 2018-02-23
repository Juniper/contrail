package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeAlarmOrList makes AlarmOrList
func MakeAlarmOrList() *AlarmOrList{
    return &AlarmOrList{
    //TODO(nati): Apply default
    
            
                OrList:  MakeAlarmAndListSlice(),
            
        
    }
}

// MakeAlarmOrList makes AlarmOrList
func InterfaceToAlarmOrList(i interface{}) *AlarmOrList{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &AlarmOrList{
    //TODO(nati): Apply default
    
            
                OrList:  InterfaceToAlarmAndListSlice(m["or_list"]),
            
        
    }
}

// MakeAlarmOrListSlice() makes a slice of AlarmOrList
func MakeAlarmOrListSlice() []*AlarmOrList {
    return []*AlarmOrList{}
}

// InterfaceToAlarmOrListSlice() makes a slice of AlarmOrList
func InterfaceToAlarmOrListSlice(i interface{}) []*AlarmOrList {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*AlarmOrList{}
    for _, item := range list {
        result = append(result, InterfaceToAlarmOrList(item) )
    }
    return result
}



