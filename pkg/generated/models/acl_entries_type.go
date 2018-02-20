package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeAclEntriesType makes AclEntriesType
func MakeAclEntriesType() *AclEntriesType{
    return &AclEntriesType{
    //TODO(nati): Apply default
    Dynamic: false,
        
            
                ACLRule:  MakeAclRuleTypeSlice(),
            
        
    }
}

// MakeAclEntriesType makes AclEntriesType
func InterfaceToAclEntriesType(i interface{}) *AclEntriesType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &AclEntriesType{
    //TODO(nati): Apply default
    Dynamic: schema.InterfaceToBool(m["dynamic"]),
        
            
                ACLRule:  InterfaceToAclRuleTypeSlice(m["acl_rule"]),
            
        
    }
}

// MakeAclEntriesTypeSlice() makes a slice of AclEntriesType
func MakeAclEntriesTypeSlice() []*AclEntriesType {
    return []*AclEntriesType{}
}

// InterfaceToAclEntriesTypeSlice() makes a slice of AclEntriesType
func InterfaceToAclEntriesTypeSlice(i interface{}) []*AclEntriesType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*AclEntriesType{}
    for _, item := range list {
        result = append(result, InterfaceToAclEntriesType(item) )
    }
    return result
}



