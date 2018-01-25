package models

// TelemetryStateInfo

// TelemetryStateInfo
//proteus:generate
type TelemetryStateInfo struct {
	Resource   []*TelemetryResourceInfo `json:"resource,omitempty"`
	ServerPort int                      `json:"server_port,omitempty"`
	ServerIP   string                   `json:"server_ip,omitempty"`
}

// MakeTelemetryStateInfo makes TelemetryStateInfo
func MakeTelemetryStateInfo() *TelemetryStateInfo {
	return &TelemetryStateInfo{
		//TODO(nati): Apply default

		Resource: MakeTelemetryResourceInfoSlice(),

		ServerPort: 0,
		ServerIP:   "",
	}
}

// MakeTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
	return []*TelemetryStateInfo{}
}
