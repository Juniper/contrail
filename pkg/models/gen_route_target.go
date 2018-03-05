package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRouteTarget makes RouteTarget
// nolint
func MakeRouteTarget() *RouteTarget {
	return &RouteTarget{
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
	}
}

// MakeRouteTarget makes RouteTarget
// nolint
func InterfaceToRouteTarget(i interface{}) *RouteTarget {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RouteTarget{
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
	}
}

// MakeRouteTargetSlice() makes a slice of RouteTarget
// nolint
func MakeRouteTargetSlice() []*RouteTarget {
	return []*RouteTarget{}
}

// InterfaceToRouteTargetSlice() makes a slice of RouteTarget
// nolint
func InterfaceToRouteTargetSlice(i interface{}) []*RouteTarget {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RouteTarget{}
	for _, item := range list {
		result = append(result, InterfaceToRouteTarget(item))
	}
	return result
}
