package models

// PluginProperties

import "encoding/json"

// PluginProperties
type PluginProperties struct {
	PluginProperty []*PluginProperty `json:"plugin_property"`
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

// InterfaceToPluginProperties makes PluginProperties from interface
func InterfaceToPluginProperties(iData interface{}) *PluginProperties {
	data := iData.(map[string]interface{})
	return &PluginProperties{

		PluginProperty: InterfaceToPluginPropertySlice(data["plugin_property"]),

		//{"description":"List of plugin specific properties (property, value)","type":"array","item":{"type":"object","properties":{"property":{"type":"string"},"value":{"type":"string"}}}}

	}
}

// InterfaceToPluginPropertiesSlice makes a slice of PluginProperties from interface
func InterfaceToPluginPropertiesSlice(data interface{}) []*PluginProperties {
	list := data.([]interface{})
	result := MakePluginPropertiesSlice()
	for _, item := range list {
		result = append(result, InterfaceToPluginProperties(item))
	}
	return result
}

// MakePluginPropertiesSlice() makes a slice of PluginProperties
func MakePluginPropertiesSlice() []*PluginProperties {
	return []*PluginProperties{}
}
