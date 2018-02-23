package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeAccessControlList makes AccessControlList
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
		AccessControlListHash:    0,
		AccessControlListEntries: MakeAclEntriesType(),
	}
}

// MakeAccessControlList makes AccessControlList
func InterfaceToAccessControlList(i interface{}) *AccessControlList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AccessControlList{
		//TODO(nati): Apply default
		UUID:                     schema.InterfaceToString(m["uuid"]),
		ParentUUID:               schema.InterfaceToString(m["parent_uuid"]),
		ParentType:               schema.InterfaceToString(m["parent_type"]),
		FQName:                   schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                  InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:              schema.InterfaceToString(m["display_name"]),
		Annotations:              InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                   InterfaceToPermType2(m["perms2"]),
		AccessControlListHash:    schema.InterfaceToInt64(m["access_control_list_hash"]),
		AccessControlListEntries: InterfaceToAclEntriesType(m["access_control_list_entries"]),
	}
}

// MakeAccessControlListSlice() makes a slice of AccessControlList
func MakeAccessControlListSlice() []*AccessControlList {
	return []*AccessControlList{}
}

// InterfaceToAccessControlListSlice() makes a slice of AccessControlList
func InterfaceToAccessControlListSlice(i interface{}) []*AccessControlList {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AccessControlList{}
	for _, item := range list {
		result = append(result, InterfaceToAccessControlList(item))
	}
	return result
}
