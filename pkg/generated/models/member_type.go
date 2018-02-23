package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeMemberType makes MemberType
func MakeMemberType() *MemberType{
    return &MemberType{
    //TODO(nati): Apply default
    Role: "",
        
    }
}

// MakeMemberType makes MemberType
func InterfaceToMemberType(i interface{}) *MemberType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &MemberType{
    //TODO(nati): Apply default
    Role: schema.InterfaceToString(m["role"]),
        
    }
}

// MakeMemberTypeSlice() makes a slice of MemberType
func MakeMemberTypeSlice() []*MemberType {
    return []*MemberType{}
}

// InterfaceToMemberTypeSlice() makes a slice of MemberType
func InterfaceToMemberTypeSlice(i interface{}) []*MemberType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*MemberType{}
    for _, item := range list {
        result = append(result, InterfaceToMemberType(item) )
    }
    return result
}



