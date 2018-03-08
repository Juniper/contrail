package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeConfigRoot makes ConfigRoot
// nolint
func MakeConfigRoot() *ConfigRoot {
	return &ConfigRoot{
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

// MakeConfigRoot makes ConfigRoot
// nolint
func InterfaceToConfigRoot(i interface{}) *ConfigRoot {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ConfigRoot{
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

		TagRefs: InterfaceToConfigRootTagRefs(m["tag_refs"]),
	}
}

func InterfaceToConfigRootTagRefs(i interface{}) []*ConfigRootTagRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ConfigRootTagRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ConfigRootTagRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeConfigRootSlice() makes a slice of ConfigRoot
// nolint
func MakeConfigRootSlice() []*ConfigRoot {
	return []*ConfigRoot{}
}

// InterfaceToConfigRootSlice() makes a slice of ConfigRoot
// nolint
func InterfaceToConfigRootSlice(i interface{}) []*ConfigRoot {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ConfigRoot{}
	for _, item := range list {
		result = append(result, InterfaceToConfigRoot(item))
	}
	return result
}
