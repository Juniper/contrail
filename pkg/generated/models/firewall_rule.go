package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeFirewallRule makes FirewallRule
func MakeFirewallRule() *FirewallRule{
    return &FirewallRule{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Endpoint1: MakeFirewallRuleEndpointType(),
        Endpoint2: MakeFirewallRuleEndpointType(),
        ActionList: MakeActionListType(),
        Service: MakeFirewallServiceType(),
        Direction: "",
        MatchTagTypes: MakeFirewallRuleMatchTagsTypeIdList(),
        MatchTags: MakeFirewallRuleMatchTagsType(),
        
    }
}

// MakeFirewallRule makes FirewallRule
func InterfaceToFirewallRule(i interface{}) *FirewallRule{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &FirewallRule{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        Endpoint1: InterfaceToFirewallRuleEndpointType(m["endpoint_1"]),
        Endpoint2: InterfaceToFirewallRuleEndpointType(m["endpoint_2"]),
        ActionList: InterfaceToActionListType(m["action_list"]),
        Service: InterfaceToFirewallServiceType(m["service"]),
        Direction: schema.InterfaceToString(m["direction"]),
        MatchTagTypes: InterfaceToFirewallRuleMatchTagsTypeIdList(m["match_tag_types"]),
        MatchTags: InterfaceToFirewallRuleMatchTagsType(m["match_tags"]),
        
    }
}

// MakeFirewallRuleSlice() makes a slice of FirewallRule
func MakeFirewallRuleSlice() []*FirewallRule {
    return []*FirewallRule{}
}

// InterfaceToFirewallRuleSlice() makes a slice of FirewallRule
func InterfaceToFirewallRuleSlice(i interface{}) []*FirewallRule {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*FirewallRule{}
    for _, item := range list {
        result = append(result, InterfaceToFirewallRule(item) )
    }
    return result
}



