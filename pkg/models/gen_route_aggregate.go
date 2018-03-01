package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRouteAggregate makes RouteAggregate
// nolint
func MakeRouteAggregate() *RouteAggregate {
	return &RouteAggregate{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeRouteAggregate makes RouteAggregate
// nolint
func InterfaceToRouteAggregate(i interface{}) *RouteAggregate {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RouteAggregate{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
	}
}

// MakeRouteAggregateSlice() makes a slice of RouteAggregate
// nolint
func MakeRouteAggregateSlice() []*RouteAggregate {
	return []*RouteAggregate{}
}

// InterfaceToRouteAggregateSlice() makes a slice of RouteAggregate
// nolint
func InterfaceToRouteAggregateSlice(i interface{}) []*RouteAggregate {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RouteAggregate{}
	for _, item := range list {
		result = append(result, InterfaceToRouteAggregate(item))
	}
	return result
}
