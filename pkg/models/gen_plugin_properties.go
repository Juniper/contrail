package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePluginProperties makes PluginProperties
// nolint
func MakePluginProperties() *PluginProperties {
	return &PluginProperties{
		//TODO(nati): Apply default

		PluginProperty: MakePluginPropertySlice(),
	}
}

// MakePluginProperties makes PluginProperties
// nolint
func InterfaceToPluginProperties(i interface{}) *PluginProperties {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PluginProperties{
		//TODO(nati): Apply default

		PluginProperty: InterfaceToPluginPropertySlice(m["plugin_property"]),
	}
}

// MakePluginPropertiesSlice() makes a slice of PluginProperties
// nolint
func MakePluginPropertiesSlice() []*PluginProperties {
	return []*PluginProperties{}
}

// InterfaceToPluginPropertiesSlice() makes a slice of PluginProperties
// nolint
func InterfaceToPluginPropertiesSlice(i interface{}) []*PluginProperties {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PluginProperties{}
	for _, item := range list {
		result = append(result, InterfaceToPluginProperties(item))
	}
	return result
}
