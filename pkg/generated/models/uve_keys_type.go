package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeUveKeysType makes UveKeysType
func MakeUveKeysType() *UveKeysType{
    return &UveKeysType{
    //TODO(nati): Apply default
    UveKey: []string{},
        
    }
}

// MakeUveKeysType makes UveKeysType
func InterfaceToUveKeysType(i interface{}) *UveKeysType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &UveKeysType{
    //TODO(nati): Apply default
    UveKey: schema.InterfaceToStringList(m["uve_key"]),
        
    }
}

// MakeUveKeysTypeSlice() makes a slice of UveKeysType
func MakeUveKeysTypeSlice() []*UveKeysType {
    return []*UveKeysType{}
}

// InterfaceToUveKeysTypeSlice() makes a slice of UveKeysType
func InterfaceToUveKeysTypeSlice(i interface{}) []*UveKeysType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*UveKeysType{}
    for _, item := range list {
        result = append(result, InterfaceToUveKeysType(item) )
    }
    return result
}



