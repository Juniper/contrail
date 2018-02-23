package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakePluginProperty makes PluginProperty
func MakePluginProperty() *PluginProperty {
	return &PluginProperty{
		//TODO(nati): Apply default
		Property: "",
		Value:    "",
	}
}

// MakePluginProperty makes PluginProperty
func InterfaceToPluginProperty(i interface{}) *PluginProperty {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PluginProperty{
		//TODO(nati): Apply default
		Property: schema.InterfaceToString(m["property"]),
		Value:    schema.InterfaceToString(m["value"]),
	}
}

// MakePluginPropertySlice() makes a slice of PluginProperty
func MakePluginPropertySlice() []*PluginProperty {
	return []*PluginProperty{}
}

// InterfaceToPluginPropertySlice() makes a slice of PluginProperty
func InterfaceToPluginPropertySlice(i interface{}) []*PluginProperty {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PluginProperty{}
	for _, item := range list {
		result = append(result, InterfaceToPluginProperty(item))
	}
	return result
}
