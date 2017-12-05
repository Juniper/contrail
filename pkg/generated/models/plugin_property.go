package models

// PluginProperty

import "encoding/json"

type PluginProperty struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

func (model *PluginProperty) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakePluginProperty() *PluginProperty {
	return &PluginProperty{
		//TODO(nati): Apply default
		Value:    "",
		Property: "",
	}
}

func InterfaceToPluginProperty(iData interface{}) *PluginProperty {
	data := iData.(map[string]interface{})
	return &PluginProperty{
		Property: data["property"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Property","GoType":"string"}
		Value: data["value"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}

	}
}

func InterfaceToPluginPropertySlice(data interface{}) []*PluginProperty {
	list := data.([]interface{})
	result := MakePluginPropertySlice()
	for _, item := range list {
		result = append(result, InterfaceToPluginProperty(item))
	}
	return result
}

func MakePluginPropertySlice() []*PluginProperty {
	return []*PluginProperty{}
}
