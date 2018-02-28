package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRouteType makes RouteType
// nolint
func MakeRouteType() *RouteType {
	return &RouteType{
		//TODO(nati): Apply default
		Prefix:              "",
		NextHop:             "",
		CommunityAttributes: MakeCommunityAttributes(),
		NextHopType:         "",
	}
}

// MakeRouteType makes RouteType
// nolint
func InterfaceToRouteType(i interface{}) *RouteType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RouteType{
		//TODO(nati): Apply default
		Prefix:              common.InterfaceToString(m["prefix"]),
		NextHop:             common.InterfaceToString(m["next_hop"]),
		CommunityAttributes: InterfaceToCommunityAttributes(m["community_attributes"]),
		NextHopType:         common.InterfaceToString(m["next_hop_type"]),
	}
}

// MakeRouteTypeSlice() makes a slice of RouteType
// nolint
func MakeRouteTypeSlice() []*RouteType {
	return []*RouteType{}
}

// InterfaceToRouteTypeSlice() makes a slice of RouteType
// nolint
func InterfaceToRouteTypeSlice(i interface{}) []*RouteType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RouteType{}
	for _, item := range list {
		result = append(result, InterfaceToRouteType(item))
	}
	return result
}
