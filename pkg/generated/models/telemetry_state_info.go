package models

// TelemetryStateInfo

import "encoding/json"

// TelemetryStateInfo
type TelemetryStateInfo struct {
	Resource   []*TelemetryResourceInfo `json:"resource,omitempty"`
	ServerPort int                      `json:"server_port,omitempty"`
	ServerIP   string                   `json:"server_ip,omitempty"`
}

// String returns json representation of the object
func (model *TelemetryStateInfo) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeTelemetryStateInfo makes TelemetryStateInfo
func MakeTelemetryStateInfo() *TelemetryStateInfo {
	return &TelemetryStateInfo{
		//TODO(nati): Apply default
		ServerIP: "",

		Resource: MakeTelemetryResourceInfoSlice(),

		ServerPort: 0,
	}
}

// MakeTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
	return []*TelemetryStateInfo{}
}
