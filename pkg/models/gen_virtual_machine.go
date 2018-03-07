package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualMachine makes VirtualMachine
// nolint
func MakeVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
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
	}
}

// MakeVirtualMachine makes VirtualMachine
// nolint
func InterfaceToVirtualMachine(i interface{}) *VirtualMachine {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualMachine{
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

		ServiceInstanceRefs: InterfaceToVirtualMachineServiceInstanceRefs(m["service_instance_refs"]),
	}
}

func InterfaceToVirtualMachineServiceInstanceRefs(i interface{}) []*VirtualMachineServiceInstanceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineServiceInstanceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineServiceInstanceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeVirtualMachineSlice() makes a slice of VirtualMachine
// nolint
func MakeVirtualMachineSlice() []*VirtualMachine {
	return []*VirtualMachine{}
}

// InterfaceToVirtualMachineSlice() makes a slice of VirtualMachine
// nolint
func InterfaceToVirtualMachineSlice(i interface{}) []*VirtualMachine {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualMachine{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualMachine(item))
	}
	return result
}
