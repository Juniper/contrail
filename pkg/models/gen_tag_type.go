package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeTagType makes TagType
// nolint
func MakeTagType() *TagType {
	return &TagType{
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
		TagTypeID:            "",
	}
}

// MakeTagType makes TagType
// nolint
func InterfaceToTagType(i interface{}) *TagType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &TagType{
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
		TagTypeID:            common.InterfaceToString(m["tag_type_id"]),
	}
}

// MakeTagTypeSlice() makes a slice of TagType
// nolint
func MakeTagTypeSlice() []*TagType {
	return []*TagType{}
}

// InterfaceToTagTypeSlice() makes a slice of TagType
// nolint
func InterfaceToTagTypeSlice(i interface{}) []*TagType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*TagType{}
	for _, item := range list {
		result = append(result, InterfaceToTagType(item))
	}
	return result
}
