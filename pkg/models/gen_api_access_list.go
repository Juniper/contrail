package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAPIAccessList makes APIAccessList
// nolint
func MakeAPIAccessList() *APIAccessList {
	return &APIAccessList{
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
		APIAccessListEntries: MakeRbacRuleEntriesType(),
	}
}

// MakeAPIAccessList makes APIAccessList
// nolint
func InterfaceToAPIAccessList(i interface{}) *APIAccessList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &APIAccessList{
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
		APIAccessListEntries: InterfaceToRbacRuleEntriesType(m["api_access_list_entries"]),
	}
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
// nolint
func MakeAPIAccessListSlice() []*APIAccessList {
	return []*APIAccessList{}
}

// InterfaceToAPIAccessListSlice() makes a slice of APIAccessList
// nolint
func InterfaceToAPIAccessListSlice(i interface{}) []*APIAccessList {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*APIAccessList{}
	for _, item := range list {
		result = append(result, InterfaceToAPIAccessList(item))
	}
	return result
}
