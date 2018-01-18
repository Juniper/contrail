package models

// TelemetryResourceInfo

import "encoding/json"

// TelemetryResourceInfo
type TelemetryResourceInfo struct {
	Rate string `json:"rate,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
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
		Name: "",
		Path: "",
		Rate: "",
	}
}

// MakeTelemetryResourceInfoSlice() makes a slice of TelemetryResourceInfo
func MakeTelemetryResourceInfoSlice() []*TelemetryResourceInfo {
	return []*TelemetryResourceInfo{}
}
