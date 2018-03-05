package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeInterfaceRouteTable makes InterfaceRouteTable
// nolint
func MakeInterfaceRouteTable() *InterfaceRouteTable {
	return &InterfaceRouteTable{
		//TODO(nati): Apply default
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		ConfigurationVersion:      0,
		InterfaceRouteTableRoutes: MakeRouteTableType(),
	}
}

// MakeInterfaceRouteTable makes InterfaceRouteTable
// nolint
func InterfaceToInterfaceRouteTable(i interface{}) *InterfaceRouteTable {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &InterfaceRouteTable{
		//TODO(nati): Apply default
		UUID:                      common.InterfaceToString(m["uuid"]),
		ParentUUID:                common.InterfaceToString(m["parent_uuid"]),
		ParentType:                common.InterfaceToString(m["parent_type"]),
		FQName:                    common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               common.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:      common.InterfaceToInt64(m["configuration_version"]),
		InterfaceRouteTableRoutes: InterfaceToRouteTableType(m["interface_route_table_routes"]),
	}
}

// MakeInterfaceRouteTableSlice() makes a slice of InterfaceRouteTable
// nolint
func MakeInterfaceRouteTableSlice() []*InterfaceRouteTable {
	return []*InterfaceRouteTable{}
}

// InterfaceToInterfaceRouteTableSlice() makes a slice of InterfaceRouteTable
// nolint
func InterfaceToInterfaceRouteTableSlice(i interface{}) []*InterfaceRouteTable {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*InterfaceRouteTable{}
	for _, item := range list {
		result = append(result, InterfaceToInterfaceRouteTable(item))
	}
	return result
}
