package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePluginProperty makes PluginProperty
// nolint
func MakePluginProperty() *PluginProperty {
	return &PluginProperty{
		//TODO(nati): Apply default
		Property: "",
		Value:    "",
	}
}

// MakePluginProperty makes PluginProperty
// nolint
func InterfaceToPluginProperty(i interface{}) *PluginProperty {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PluginProperty{
		//TODO(nati): Apply default
		Property: common.InterfaceToString(m["property"]),
		Value:    common.InterfaceToString(m["value"]),
	}
}

// MakePluginPropertySlice() makes a slice of PluginProperty
// nolint
func MakePluginPropertySlice() []*PluginProperty {
	return []*PluginProperty{}
}

// InterfaceToPluginPropertySlice() makes a slice of PluginProperty
// nolint
func InterfaceToPluginPropertySlice(i interface{}) []*PluginProperty {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PluginProperty{}
	for _, item := range list {
		result = append(result, InterfaceToPluginProperty(item))
	}
	return result
}
