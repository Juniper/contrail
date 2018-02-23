package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeRouteTargetList makes RouteTargetList
func MakeRouteTargetList() *RouteTargetList {
	return &RouteTargetList{
		//TODO(nati): Apply default
		RouteTarget: []string{},
	}
}

// MakeRouteTargetList makes RouteTargetList
func InterfaceToRouteTargetList(i interface{}) *RouteTargetList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RouteTargetList{
		//TODO(nati): Apply default
		RouteTarget: schema.InterfaceToStringList(m["route_target"]),
	}
}

// MakeRouteTargetListSlice() makes a slice of RouteTargetList
func MakeRouteTargetListSlice() []*RouteTargetList {
	return []*RouteTargetList{}
}

// InterfaceToRouteTargetListSlice() makes a slice of RouteTargetList
func InterfaceToRouteTargetListSlice(i interface{}) []*RouteTargetList {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RouteTargetList{}
	for _, item := range list {
		result = append(result, InterfaceToRouteTargetList(item))
	}
	return result
}
