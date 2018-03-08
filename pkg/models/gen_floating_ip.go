package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFloatingIP makes FloatingIP
// nolint
func MakeFloatingIP() *FloatingIP {
	return &FloatingIP{
		//TODO(nati): Apply default
		UUID:                         "",
		ParentUUID:                   "",
		ParentType:                   "",
		FQName:                       []string{},
		IDPerms:                      MakeIdPermsType(),
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		Perms2:                       MakePermType2(),
		ConfigurationVersion:         0,
		FloatingIPAddressFamily:      "",
		FloatingIPPortMappings:       MakePortMappings(),
		FloatingIPIsVirtualIP:        false,
		FloatingIPAddress:            "",
		FloatingIPPortMappingsEnable: false,
		FloatingIPFixedIPAddress:     "",
		FloatingIPTrafficDirection:   "",
	}
}

// MakeFloatingIP makes FloatingIP
// nolint
func InterfaceToFloatingIP(i interface{}) *FloatingIP {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FloatingIP{
		//TODO(nati): Apply default
		UUID:                         common.InterfaceToString(m["uuid"]),
		ParentUUID:                   common.InterfaceToString(m["parent_uuid"]),
		ParentType:                   common.InterfaceToString(m["parent_type"]),
		FQName:                       common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                      InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                  common.InterfaceToString(m["display_name"]),
		Annotations:                  InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                       InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:         common.InterfaceToInt64(m["configuration_version"]),
		FloatingIPAddressFamily:      common.InterfaceToString(m["floating_ip_address_family"]),
		FloatingIPPortMappings:       InterfaceToPortMappings(m["floating_ip_port_mappings"]),
		FloatingIPIsVirtualIP:        common.InterfaceToBool(m["floating_ip_is_virtual_ip"]),
		FloatingIPAddress:            common.InterfaceToString(m["floating_ip_address"]),
		FloatingIPPortMappingsEnable: common.InterfaceToBool(m["floating_ip_port_mappings_enable"]),
		FloatingIPFixedIPAddress:     common.InterfaceToString(m["floating_ip_fixed_ip_address"]),
		FloatingIPTrafficDirection:   common.InterfaceToString(m["floating_ip_traffic_direction"]),

		ProjectRefs: InterfaceToFloatingIPProjectRefs(m["project_refs"]),

		VirtualMachineInterfaceRefs: InterfaceToFloatingIPVirtualMachineInterfaceRefs(m["virtual_machine_interface_refs"]),
	}
}

func InterfaceToFloatingIPVirtualMachineInterfaceRefs(i interface{}) []*FloatingIPVirtualMachineInterfaceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*FloatingIPVirtualMachineInterfaceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &FloatingIPVirtualMachineInterfaceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToFloatingIPProjectRefs(i interface{}) []*FloatingIPProjectRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*FloatingIPProjectRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &FloatingIPProjectRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
// nolint
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}

// InterfaceToFloatingIPSlice() makes a slice of FloatingIP
// nolint
func InterfaceToFloatingIPSlice(i interface{}) []*FloatingIP {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FloatingIP{}
	for _, item := range list {
		result = append(result, InterfaceToFloatingIP(item))
	}
	return result
}
