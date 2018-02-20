package models

// TelemetryResourceInfo

// TelemetryResourceInfo
//proteus:generate
type TelemetryResourceInfo struct {
	Path string `json:"path,omitempty"`
	Rate string `json:"rate,omitempty"`
	Name string `json:"name,omitempty"`
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
