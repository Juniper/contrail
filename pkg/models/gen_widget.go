package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeWidget makes Widget
// nolint
func MakeWidget() *Widget {
	return &Widget{
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
		ContainerConfig:      "",
		ContentConfig:        "",
		LayoutConfig:         "",
	}
}

// MakeWidget makes Widget
// nolint
func InterfaceToWidget(i interface{}) *Widget {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Widget{
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
		ContainerConfig:      common.InterfaceToString(m["container_config"]),
		ContentConfig:        common.InterfaceToString(m["content_config"]),
		LayoutConfig:         common.InterfaceToString(m["layout_config"]),
	}
}

// MakeWidgetSlice() makes a slice of Widget
// nolint
func MakeWidgetSlice() []*Widget {
	return []*Widget{}
}

// InterfaceToWidgetSlice() makes a slice of Widget
// nolint
func InterfaceToWidgetSlice(i interface{}) []*Widget {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Widget{}
	for _, item := range list {
		result = append(result, InterfaceToWidget(item))
	}
	return result
}
