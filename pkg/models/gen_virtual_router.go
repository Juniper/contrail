package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualRouter makes VirtualRouter
// nolint
func MakeVirtualRouter() *VirtualRouter {
	return &VirtualRouter{
		//TODO(nati): Apply default
		UUID:                     "",
		ParentUUID:               "",
		ParentType:               "",
		FQName:                   []string{},
		IDPerms:                  MakeIdPermsType(),
		DisplayName:              "",
		Annotations:              MakeKeyValuePairs(),
		Perms2:                   MakePermType2(),
		ConfigurationVersion:     0,
		VirtualRouterDPDKEnabled: false,
		VirtualRouterType:        "",
		VirtualRouterIPAddress:   "",
	}
}

// MakeVirtualRouter makes VirtualRouter
// nolint
func InterfaceToVirtualRouter(i interface{}) *VirtualRouter {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualRouter{
		//TODO(nati): Apply default
		UUID:                     common.InterfaceToString(m["uuid"]),
		ParentUUID:               common.InterfaceToString(m["parent_uuid"]),
		ParentType:               common.InterfaceToString(m["parent_type"]),
		FQName:                   common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                  InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:              common.InterfaceToString(m["display_name"]),
		Annotations:              InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                   InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:     common.InterfaceToInt64(m["configuration_version"]),
		VirtualRouterDPDKEnabled: common.InterfaceToBool(m["virtual_router_dpdk_enabled"]),
		VirtualRouterType:        common.InterfaceToString(m["virtual_router_type"]),
		VirtualRouterIPAddress:   common.InterfaceToString(m["virtual_router_ip_address"]),

		NetworkIpamRefs: InterfaceToVirtualRouterNetworkIpamRefs(m["network_ipam_refs"]),

		VirtualMachineRefs: InterfaceToVirtualRouterVirtualMachineRefs(m["virtual_machine_refs"]),
	}
}

func InterfaceToVirtualRouterNetworkIpamRefs(i interface{}) []*VirtualRouterNetworkIpamRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualRouterNetworkIpamRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualRouterNetworkIpamRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),

			Attr: InterfaceToVirtualRouterNetworkIpamType(m["attr"]),
		})
	}

	return result
}

func InterfaceToVirtualRouterVirtualMachineRefs(i interface{}) []*VirtualRouterVirtualMachineRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualRouterVirtualMachineRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualRouterVirtualMachineRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeVirtualRouterSlice() makes a slice of VirtualRouter
// nolint
func MakeVirtualRouterSlice() []*VirtualRouter {
	return []*VirtualRouter{}
}

// InterfaceToVirtualRouterSlice() makes a slice of VirtualRouter
// nolint
func InterfaceToVirtualRouterSlice(i interface{}) []*VirtualRouter {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualRouter{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualRouter(item))
	}
	return result
}
