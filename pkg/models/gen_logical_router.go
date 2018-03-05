package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLogicalRouter makes LogicalRouter
// nolint
func MakeLogicalRouter() *LogicalRouter {
	return &LogicalRouter{
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
		VxlanNetworkIdentifier:    "",
		ConfiguredRouteTargetList: MakeRouteTargetList(),
	}
}

// MakeLogicalRouter makes LogicalRouter
// nolint
func InterfaceToLogicalRouter(i interface{}) *LogicalRouter {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LogicalRouter{
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
		VxlanNetworkIdentifier:    common.InterfaceToString(m["vxlan_network_identifier"]),
		ConfiguredRouteTargetList: InterfaceToRouteTargetList(m["configured_route_target_list"]),
	}
}

// MakeLogicalRouterSlice() makes a slice of LogicalRouter
// nolint
func MakeLogicalRouterSlice() []*LogicalRouter {
	return []*LogicalRouter{}
}

// InterfaceToLogicalRouterSlice() makes a slice of LogicalRouter
// nolint
func InterfaceToLogicalRouterSlice(i interface{}) []*LogicalRouter {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LogicalRouter{}
	for _, item := range list {
		result = append(result, InterfaceToLogicalRouter(item))
	}
	return result
}
