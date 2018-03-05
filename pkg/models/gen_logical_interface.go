package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLogicalInterface makes LogicalInterface
// nolint
func MakeLogicalInterface() *LogicalInterface {
	return &LogicalInterface{
		//TODO(nati): Apply default
		UUID:                    "",
		ParentUUID:              "",
		ParentType:              "",
		FQName:                  []string{},
		IDPerms:                 MakeIdPermsType(),
		DisplayName:             "",
		Annotations:             MakeKeyValuePairs(),
		Perms2:                  MakePermType2(),
		ConfigurationVersion:    0,
		LogicalInterfaceVlanTag: 0,
		LogicalInterfaceType:    "",
	}
}

// MakeLogicalInterface makes LogicalInterface
// nolint
func InterfaceToLogicalInterface(i interface{}) *LogicalInterface {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LogicalInterface{
		//TODO(nati): Apply default
		UUID:                    common.InterfaceToString(m["uuid"]),
		ParentUUID:              common.InterfaceToString(m["parent_uuid"]),
		ParentType:              common.InterfaceToString(m["parent_type"]),
		FQName:                  common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                 InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:             common.InterfaceToString(m["display_name"]),
		Annotations:             InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                  InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:    common.InterfaceToInt64(m["configuration_version"]),
		LogicalInterfaceVlanTag: common.InterfaceToInt64(m["logical_interface_vlan_tag"]),
		LogicalInterfaceType:    common.InterfaceToString(m["logical_interface_type"]),
	}
}

// MakeLogicalInterfaceSlice() makes a slice of LogicalInterface
// nolint
func MakeLogicalInterfaceSlice() []*LogicalInterface {
	return []*LogicalInterface{}
}

// InterfaceToLogicalInterfaceSlice() makes a slice of LogicalInterface
// nolint
func InterfaceToLogicalInterfaceSlice(i interface{}) []*LogicalInterface {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LogicalInterface{}
	for _, item := range list {
		result = append(result, InterfaceToLogicalInterface(item))
	}
	return result
}
