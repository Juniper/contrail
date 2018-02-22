package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeRouteAggregate makes RouteAggregate
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
func InterfaceToRouteAggregate(i interface{}) *RouteAggregate {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RouteAggregate{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
	}
}

// MakeRouteAggregateSlice() makes a slice of RouteAggregate
func MakeRouteAggregateSlice() []*RouteAggregate {
	return []*RouteAggregate{}
}

// InterfaceToRouteAggregateSlice() makes a slice of RouteAggregate
func InterfaceToRouteAggregateSlice(i interface{}) []*RouteAggregate {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RouteAggregate{}
	for _, item := range list {
		result = append(result, InterfaceToRouteAggregate(item))
	}
	return result
}
