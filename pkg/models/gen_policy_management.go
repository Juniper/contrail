package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePolicyManagement makes PolicyManagement
// nolint
func MakePolicyManagement() *PolicyManagement {
	return &PolicyManagement{
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

// MakePolicyManagement makes PolicyManagement
// nolint
func InterfaceToPolicyManagement(i interface{}) *PolicyManagement {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PolicyManagement{
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
	}
}

// MakePolicyManagementSlice() makes a slice of PolicyManagement
// nolint
func MakePolicyManagementSlice() []*PolicyManagement {
	return []*PolicyManagement{}
}

// InterfaceToPolicyManagementSlice() makes a slice of PolicyManagement
// nolint
func InterfaceToPolicyManagementSlice(i interface{}) []*PolicyManagement {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PolicyManagement{}
	for _, item := range list {
		result = append(result, InterfaceToPolicyManagement(item))
	}
	return result
}
