package models

// PluginProperty

import "encoding/json"

// PluginProperty
type PluginProperty struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

// String returns json representation of the object
func (model *PluginProperty) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePluginProperty makes PluginProperty
func MakePluginProperty() *PluginProperty {
	return &PluginProperty{
		//TODO(nati): Apply default
		Value:    "",
		Property: "",
	}
}

// InterfaceToPluginProperty makes PluginProperty from interface
func InterfaceToPluginProperty(iData interface{}) *PluginProperty {
	data := iData.(map[string]interface{})
	return &PluginProperty{
		Property: data["property"].(string),

		//{"type":"string"}
		Value: data["value"].(string),

		//{"type":"string"}

	}
}

// InterfaceToPluginPropertySlice makes a slice of PluginProperty from interface
func InterfaceToPluginPropertySlice(data interface{}) []*PluginProperty {
	list := data.([]interface{})
	result := MakePluginPropertySlice()
	for _, item := range list {
		result = append(result, InterfaceToPluginProperty(item))
	}
	return result
}

// MakePluginPropertySlice() makes a slice of PluginProperty
func MakePluginPropertySlice() []*PluginProperty {
	return []*PluginProperty{}
}
