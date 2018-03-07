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

		RouteTargetRefs: InterfaceToLogicalRouterRouteTargetRefs(m["route_target_refs"]),

		VirtualMachineInterfaceRefs: InterfaceToLogicalRouterVirtualMachineInterfaceRefs(m["virtual_machine_interface_refs"]),

		ServiceInstanceRefs: InterfaceToLogicalRouterServiceInstanceRefs(m["service_instance_refs"]),

		RouteTableRefs: InterfaceToLogicalRouterRouteTableRefs(m["route_table_refs"]),

		VirtualNetworkRefs: InterfaceToLogicalRouterVirtualNetworkRefs(m["virtual_network_refs"]),

		PhysicalRouterRefs: InterfaceToLogicalRouterPhysicalRouterRefs(m["physical_router_refs"]),

		BGPVPNRefs: InterfaceToLogicalRouterBGPVPNRefs(m["bgpvpn_refs"]),
	}
}

func InterfaceToLogicalRouterVirtualNetworkRefs(i interface{}) []*LogicalRouterVirtualNetworkRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LogicalRouterVirtualNetworkRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LogicalRouterVirtualNetworkRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLogicalRouterPhysicalRouterRefs(i interface{}) []*LogicalRouterPhysicalRouterRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LogicalRouterPhysicalRouterRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LogicalRouterPhysicalRouterRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLogicalRouterBGPVPNRefs(i interface{}) []*LogicalRouterBGPVPNRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LogicalRouterBGPVPNRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LogicalRouterBGPVPNRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLogicalRouterRouteTargetRefs(i interface{}) []*LogicalRouterRouteTargetRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LogicalRouterRouteTargetRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LogicalRouterRouteTargetRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLogicalRouterVirtualMachineInterfaceRefs(i interface{}) []*LogicalRouterVirtualMachineInterfaceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LogicalRouterVirtualMachineInterfaceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LogicalRouterVirtualMachineInterfaceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLogicalRouterServiceInstanceRefs(i interface{}) []*LogicalRouterServiceInstanceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LogicalRouterServiceInstanceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LogicalRouterServiceInstanceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLogicalRouterRouteTableRefs(i interface{}) []*LogicalRouterRouteTableRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LogicalRouterRouteTableRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LogicalRouterRouteTableRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
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
