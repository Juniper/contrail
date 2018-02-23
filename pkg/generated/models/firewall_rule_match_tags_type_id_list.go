package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeFirewallRuleMatchTagsTypeIdList makes FirewallRuleMatchTagsTypeIdList
func MakeFirewallRuleMatchTagsTypeIdList() *FirewallRuleMatchTagsTypeIdList{
    return &FirewallRuleMatchTagsTypeIdList{
    //TODO(nati): Apply default
    
            
                TagType: []int64{},
            
        
    }
}

// MakeFirewallRuleMatchTagsTypeIdList makes FirewallRuleMatchTagsTypeIdList
func InterfaceToFirewallRuleMatchTagsTypeIdList(i interface{}) *FirewallRuleMatchTagsTypeIdList{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &FirewallRuleMatchTagsTypeIdList{
    //TODO(nati): Apply default
    
            
                TagType: schema.InterfaceToInt64List(m["tag_type"]),
            
        
    }
}

// MakeFirewallRuleMatchTagsTypeIdListSlice() makes a slice of FirewallRuleMatchTagsTypeIdList
func MakeFirewallRuleMatchTagsTypeIdListSlice() []*FirewallRuleMatchTagsTypeIdList {
    return []*FirewallRuleMatchTagsTypeIdList{}
}

// InterfaceToFirewallRuleMatchTagsTypeIdListSlice() makes a slice of FirewallRuleMatchTagsTypeIdList
func InterfaceToFirewallRuleMatchTagsTypeIdListSlice(i interface{}) []*FirewallRuleMatchTagsTypeIdList {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*FirewallRuleMatchTagsTypeIdList{}
    for _, item := range list {
        result = append(result, InterfaceToFirewallRuleMatchTagsTypeIdList(item) )
    }
    return result
}



