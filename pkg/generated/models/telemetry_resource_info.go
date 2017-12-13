package models

// TelemetryResourceInfo

import "encoding/json"

// TelemetryResourceInfo
type TelemetryResourceInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Rate string `json:"rate"`
}

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

		//{"type":"string"}
		Rate: data["rate"].(string),

		//{"type":"string"}
		Name: data["name"].(string),

		//{"type":"string"}

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
