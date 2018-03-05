package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeTag makes Tag
// nolint
func MakeTag() *Tag {
	return &Tag{
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
		TagTypeName:          "",
		TagID:                "",
		TagValue:             "",
	}
}

// MakeTag makes Tag
// nolint
func InterfaceToTag(i interface{}) *Tag {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Tag{
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
		TagTypeName:          common.InterfaceToString(m["tag_type_name"]),
		TagID:                common.InterfaceToString(m["tag_id"]),
		TagValue:             common.InterfaceToString(m["tag_value"]),
	}
}

// MakeTagSlice() makes a slice of Tag
// nolint
func MakeTagSlice() []*Tag {
	return []*Tag{}
}

// InterfaceToTagSlice() makes a slice of Tag
// nolint
func InterfaceToTagSlice(i interface{}) []*Tag {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Tag{}
	for _, item := range list {
		result = append(result, InterfaceToTag(item))
	}
	return result
}
