package models

// TelemetryStateInfo

import "encoding/json"

// TelemetryStateInfo
type TelemetryStateInfo struct {
	ServerPort int                      `json:"server_port,omitempty"`
	ServerIP   string                   `json:"server_ip,omitempty"`
	Resource   []*TelemetryResourceInfo `json:"resource,omitempty"`
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
		ServerPort: 0,
		ServerIP:   "",

		Resource: MakeTelemetryResourceInfoSlice(),
	}
}

// MakeTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
	return []*TelemetryStateInfo{}
}
