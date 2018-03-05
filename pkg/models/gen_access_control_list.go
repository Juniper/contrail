package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAccessControlList makes AccessControlList
// nolint
func MakeAccessControlList() *AccessControlList {
	return &AccessControlList{
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
		AccessControlListHash:    0,
		AccessControlListEntries: MakeAclEntriesType(),
	}
}

// MakeAccessControlList makes AccessControlList
// nolint
func InterfaceToAccessControlList(i interface{}) *AccessControlList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AccessControlList{
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
		AccessControlListHash:    common.InterfaceToInt64(m["access_control_list_hash"]),
		AccessControlListEntries: InterfaceToAclEntriesType(m["access_control_list_entries"]),
	}
}

// MakeAccessControlListSlice() makes a slice of AccessControlList
// nolint
func MakeAccessControlListSlice() []*AccessControlList {
	return []*AccessControlList{}
}

// InterfaceToAccessControlListSlice() makes a slice of AccessControlList
// nolint
func InterfaceToAccessControlListSlice(i interface{}) []*AccessControlList {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AccessControlList{}
	for _, item := range list {
		result = append(result, InterfaceToAccessControlList(item))
	}
	return result
}
