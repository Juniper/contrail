package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRouteTargetList makes RouteTargetList
// nolint
func MakeRouteTargetList() *RouteTargetList {
	return &RouteTargetList{
		//TODO(nati): Apply default
		RouteTarget: []string{},
	}
}

// MakeRouteTargetList makes RouteTargetList
// nolint
func InterfaceToRouteTargetList(i interface{}) *RouteTargetList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RouteTargetList{
		//TODO(nati): Apply default
		RouteTarget: common.InterfaceToStringList(m["route_target"]),
	}
}

// MakeRouteTargetListSlice() makes a slice of RouteTargetList
// nolint
func MakeRouteTargetListSlice() []*RouteTargetList {
	return []*RouteTargetList{}
}

// InterfaceToRouteTargetListSlice() makes a slice of RouteTargetList
// nolint
func InterfaceToRouteTargetListSlice(i interface{}) []*RouteTargetList {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RouteTargetList{}
	for _, item := range list {
		result = append(result, InterfaceToRouteTargetList(item))
	}
	return result
}
