package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualIP makes VirtualIP
// nolint
func MakeVirtualIP() *VirtualIP {
	return &VirtualIP{
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
		VirtualIPProperties:  MakeVirtualIpType(),
	}
}

// MakeVirtualIP makes VirtualIP
// nolint
func InterfaceToVirtualIP(i interface{}) *VirtualIP {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualIP{
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
		VirtualIPProperties:  InterfaceToVirtualIpType(m["virtual_ip_properties"]),

		LoadbalancerPoolRefs: InterfaceToVirtualIPLoadbalancerPoolRefs(m["loadbalancer_pool_refs"]),

		VirtualMachineInterfaceRefs: InterfaceToVirtualIPVirtualMachineInterfaceRefs(m["virtual_machine_interface_refs"]),
	}
}

func InterfaceToVirtualIPLoadbalancerPoolRefs(i interface{}) []*VirtualIPLoadbalancerPoolRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualIPLoadbalancerPoolRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualIPLoadbalancerPoolRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualIPVirtualMachineInterfaceRefs(i interface{}) []*VirtualIPVirtualMachineInterfaceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualIPVirtualMachineInterfaceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualIPVirtualMachineInterfaceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeVirtualIPSlice() makes a slice of VirtualIP
// nolint
func MakeVirtualIPSlice() []*VirtualIP {
	return []*VirtualIP{}
}

// InterfaceToVirtualIPSlice() makes a slice of VirtualIP
// nolint
func InterfaceToVirtualIPSlice(i interface{}) []*VirtualIP {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualIP{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualIP(item))
	}
	return result
}
