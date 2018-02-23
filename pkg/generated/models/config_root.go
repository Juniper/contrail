package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeConfigRoot makes ConfigRoot
func MakeConfigRoot() *ConfigRoot {
	return &ConfigRoot{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeConfigRoot makes ConfigRoot
func InterfaceToConfigRoot(i interface{}) *ConfigRoot {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ConfigRoot{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
	}
}

// MakeConfigRootSlice() makes a slice of ConfigRoot
func MakeConfigRootSlice() []*ConfigRoot {
	return []*ConfigRoot{}
}

// InterfaceToConfigRootSlice() makes a slice of ConfigRoot
func InterfaceToConfigRootSlice(i interface{}) []*ConfigRoot {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ConfigRoot{}
	for _, item := range list {
		result = append(result, InterfaceToConfigRoot(item))
	}
	return result
}
