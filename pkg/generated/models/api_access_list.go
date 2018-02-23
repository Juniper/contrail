package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeAPIAccessList makes APIAccessList
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
		APIAccessListEntries: MakeRbacRuleEntriesType(),
	}
}

// MakeAPIAccessList makes APIAccessList
func InterfaceToAPIAccessList(i interface{}) *APIAccessList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &APIAccessList{
		//TODO(nati): Apply default
		UUID:                 schema.InterfaceToString(m["uuid"]),
		ParentUUID:           schema.InterfaceToString(m["parent_uuid"]),
		ParentType:           schema.InterfaceToString(m["parent_type"]),
		FQName:               schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          schema.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		APIAccessListEntries: InterfaceToRbacRuleEntriesType(m["api_access_list_entries"]),
	}
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
	return []*APIAccessList{}
}

// InterfaceToAPIAccessListSlice() makes a slice of APIAccessList
func InterfaceToAPIAccessListSlice(i interface{}) []*APIAccessList {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*APIAccessList{}
	for _, item := range list {
		result = append(result, InterfaceToAPIAccessList(item))
	}
	return result
}
