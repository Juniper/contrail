package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeAllowedAddressPairs makes AllowedAddressPairs
func MakeAllowedAddressPairs() *AllowedAddressPairs{
    return &AllowedAddressPairs{
    //TODO(nati): Apply default
    
            
                AllowedAddressPair:  MakeAllowedAddressPairSlice(),
            
        
    }
}

// MakeAllowedAddressPairs makes AllowedAddressPairs
func InterfaceToAllowedAddressPairs(i interface{}) *AllowedAddressPairs{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &AllowedAddressPairs{
    //TODO(nati): Apply default
    
            
                AllowedAddressPair:  InterfaceToAllowedAddressPairSlice(m["allowed_address_pair"]),
            
        
    }
}

// MakeAllowedAddressPairsSlice() makes a slice of AllowedAddressPairs
func MakeAllowedAddressPairsSlice() []*AllowedAddressPairs {
    return []*AllowedAddressPairs{}
}

// InterfaceToAllowedAddressPairsSlice() makes a slice of AllowedAddressPairs
func InterfaceToAllowedAddressPairsSlice(i interface{}) []*AllowedAddressPairs {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*AllowedAddressPairs{}
    for _, item := range list {
        result = append(result, InterfaceToAllowedAddressPairs(item) )
    }
    return result
}



