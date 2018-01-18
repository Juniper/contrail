package models

// PluginProperties

import "encoding/json"

// PluginProperties
type PluginProperties struct {
	PluginProperty []*PluginProperty `json:"plugin_property,omitempty"`
}

// String returns json representation of the object
func (model *PluginProperties) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePluginProperties makes PluginProperties
func MakePluginProperties() *PluginProperties {
	return &PluginProperties{
		//TODO(nati): Apply default

		PluginProperty: MakePluginPropertySlice(),
	}
}

// MakePluginPropertiesSlice() makes a slice of PluginProperties
func MakePluginPropertiesSlice() []*PluginProperties {
	return []*PluginProperties{}
}
