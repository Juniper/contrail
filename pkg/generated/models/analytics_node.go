package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeAnalyticsNode makes AnalyticsNode
func MakeAnalyticsNode() *AnalyticsNode{
    return &AnalyticsNode{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        AnalyticsNodeIPAddress: "",
        
    }
}

// MakeAnalyticsNode makes AnalyticsNode
func InterfaceToAnalyticsNode(i interface{}) *AnalyticsNode{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &AnalyticsNode{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        AnalyticsNodeIPAddress: schema.InterfaceToString(m["analytics_node_ip_address"]),
        
    }
}

// MakeAnalyticsNodeSlice() makes a slice of AnalyticsNode
func MakeAnalyticsNodeSlice() []*AnalyticsNode {
    return []*AnalyticsNode{}
}

// InterfaceToAnalyticsNodeSlice() makes a slice of AnalyticsNode
func InterfaceToAnalyticsNodeSlice(i interface{}) []*AnalyticsNode {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*AnalyticsNode{}
    for _, item := range list {
        result = append(result, InterfaceToAnalyticsNode(item) )
    }
    return result
}



