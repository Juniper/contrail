package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAnalyticsNode makes AnalyticsNode
// nolint
func MakeAnalyticsNode() *AnalyticsNode {
	return &AnalyticsNode{
		//TODO(nati): Apply default
		UUID:                   "",
		ParentUUID:             "",
		ParentType:             "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Annotations:            MakeKeyValuePairs(),
		Perms2:                 MakePermType2(),
		AnalyticsNodeIPAddress: "",
	}
}

// MakeAnalyticsNode makes AnalyticsNode
// nolint
func InterfaceToAnalyticsNode(i interface{}) *AnalyticsNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AnalyticsNode{
		//TODO(nati): Apply default
		UUID:                   common.InterfaceToString(m["uuid"]),
		ParentUUID:             common.InterfaceToString(m["parent_uuid"]),
		ParentType:             common.InterfaceToString(m["parent_type"]),
		FQName:                 common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:            common.InterfaceToString(m["display_name"]),
		Annotations:            InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                 InterfaceToPermType2(m["perms2"]),
		AnalyticsNodeIPAddress: common.InterfaceToString(m["analytics_node_ip_address"]),
	}
}

// MakeAnalyticsNodeSlice() makes a slice of AnalyticsNode
// nolint
func MakeAnalyticsNodeSlice() []*AnalyticsNode {
	return []*AnalyticsNode{}
}

// InterfaceToAnalyticsNodeSlice() makes a slice of AnalyticsNode
// nolint
func InterfaceToAnalyticsNodeSlice(i interface{}) []*AnalyticsNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AnalyticsNode{}
	for _, item := range list {
		result = append(result, InterfaceToAnalyticsNode(item))
	}
	return result
}
