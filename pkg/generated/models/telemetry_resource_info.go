package models

// TelemetryResourceInfo

import "encoding/json"

// TelemetryResourceInfo
type TelemetryResourceInfo struct {
	Rate string `json:"rate"`
	Name string `json:"name"`
	Path string `json:"path"`
}

//  parents relation object

// String returns json representation of the object
func (model *TelemetryResourceInfo) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeTelemetryResourceInfo makes TelemetryResourceInfo
func MakeTelemetryResourceInfo() *TelemetryResourceInfo {
	return &TelemetryResourceInfo{
		//TODO(nati): Apply default
		Path: "",
		Rate: "",
		Name: "",
	}
}

// InterfaceToTelemetryResourceInfo makes TelemetryResourceInfo from interface
func InterfaceToTelemetryResourceInfo(iData interface{}) *TelemetryResourceInfo {
	data := iData.(map[string]interface{})
	return &TelemetryResourceInfo{
		Path: data["path"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Path","GoType":"string","GoPremitive":true}
		Rate: data["rate"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Rate","GoType":"string","GoPremitive":true}
		Name: data["name"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToTelemetryResourceInfoSlice makes a slice of TelemetryResourceInfo from interface
func InterfaceToTelemetryResourceInfoSlice(data interface{}) []*TelemetryResourceInfo {
	list := data.([]interface{})
	result := MakeTelemetryResourceInfoSlice()
	for _, item := range list {
		result = append(result, InterfaceToTelemetryResourceInfo(item))
	}
	return result
}

// MakeTelemetryResourceInfoSlice() makes a slice of TelemetryResourceInfo
func MakeTelemetryResourceInfoSlice() []*TelemetryResourceInfo {
	return []*TelemetryResourceInfo{}
}
