package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRouteTable makes RouteTable
// nolint
func MakeRouteTable() *RouteTable {
	return &RouteTable{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ConfigurationVersion: 0,
		Routes:               MakeRouteTableType(),
	}
}

// MakeRouteTable makes RouteTable
// nolint
func InterfaceToRouteTable(i interface{}) *RouteTable {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RouteTable{
		//TODO(nati): Apply default
		UUID:                 common.InterfaceToString(m["uuid"]),
		ParentUUID:           common.InterfaceToString(m["parent_uuid"]),
		ParentType:           common.InterfaceToString(m["parent_type"]),
		FQName:               common.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          common.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion: common.InterfaceToInt64(m["configuration_version"]),
		Routes:               InterfaceToRouteTableType(m["routes"]),
	}
}

// MakeRouteTableSlice() makes a slice of RouteTable
// nolint
func MakeRouteTableSlice() []*RouteTable {
	return []*RouteTable{}
}

// InterfaceToRouteTableSlice() makes a slice of RouteTable
// nolint
func InterfaceToRouteTableSlice(i interface{}) []*RouteTable {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RouteTable{}
	for _, item := range list {
		result = append(result, InterfaceToRouteTable(item))
	}
	return result
}
