package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRouteTableType makes RouteTableType
// nolint
func MakeRouteTableType() *RouteTableType {
	return &RouteTableType{
		//TODO(nati): Apply default

		Route: MakeRouteTypeSlice(),
	}
}

// MakeRouteTableType makes RouteTableType
// nolint
func InterfaceToRouteTableType(i interface{}) *RouteTableType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RouteTableType{
		//TODO(nati): Apply default

		Route: InterfaceToRouteTypeSlice(m["route"]),
	}
}

// MakeRouteTableTypeSlice() makes a slice of RouteTableType
// nolint
func MakeRouteTableTypeSlice() []*RouteTableType {
	return []*RouteTableType{}
}

// InterfaceToRouteTableTypeSlice() makes a slice of RouteTableType
// nolint
func InterfaceToRouteTableTypeSlice(i interface{}) []*RouteTableType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RouteTableType{}
	for _, item := range list {
		result = append(result, InterfaceToRouteTableType(item))
	}
	return result
}
