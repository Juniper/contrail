package models

// TelemetryResourceInfo

import "encoding/json"

// TelemetryResourceInfo
type TelemetryResourceInfo struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	Rate string `json:"rate,omitempty"`
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

// MakeTelemetryResourceInfoSlice() makes a slice of TelemetryResourceInfo
func MakeTelemetryResourceInfoSlice() []*TelemetryResourceInfo {
	return []*TelemetryResourceInfo{}
}
