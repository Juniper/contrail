package models

// PluginProperties

import "encoding/json"

// PluginProperties
type PluginProperties struct {
	PluginProperty []*PluginProperty `json:"plugin_property"`
}

//  parents relation object

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

		//{"Title":"","Description":"List of plugin specific properties (property, value)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"property":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Property","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PluginProperty","CollectionType":"","Column":"","Item":null,"GoName":"PluginProperty","GoType":"PluginProperty","GoPremitive":false},"GoName":"PluginProperty","GoType":"[]*PluginProperty","GoPremitive":true}

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
