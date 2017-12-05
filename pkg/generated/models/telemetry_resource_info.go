package models

// TelemetryResourceInfo

import "encoding/json"

type TelemetryResourceInfo struct {
	Path string `json:"path"`
	Rate string `json:"rate"`
	Name string `json:"name"`
}

func (model *TelemetryResourceInfo) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeTelemetryResourceInfo() *TelemetryResourceInfo {
	return &TelemetryResourceInfo{
		//TODO(nati): Apply default
		Path: "",
		Rate: "",
		Name: "",
	}
}

func InterfaceToTelemetryResourceInfo(iData interface{}) *TelemetryResourceInfo {
	data := iData.(map[string]interface{})
	return &TelemetryResourceInfo{
		Path: data["path"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Path","GoType":"string"}
		Rate: data["rate"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Rate","GoType":"string"}
		Name: data["name"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string"}

	}
}

func InterfaceToTelemetryResourceInfoSlice(data interface{}) []*TelemetryResourceInfo {
	list := data.([]interface{})
	result := MakeTelemetryResourceInfoSlice()
	for _, item := range list {
		result = append(result, InterfaceToTelemetryResourceInfo(item))
	}
	return result
}

func MakeTelemetryResourceInfoSlice() []*TelemetryResourceInfo {
	return []*TelemetryResourceInfo{}
}
