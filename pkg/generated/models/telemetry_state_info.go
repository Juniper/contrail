package models

// TelemetryStateInfo

import "encoding/json"

// TelemetryStateInfo
type TelemetryStateInfo struct {
	Resource   []*TelemetryResourceInfo `json:"resource"`
	ServerPort int                      `json:"server_port"`
	ServerIP   string                   `json:"server_ip"`
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

// InterfaceToTelemetryStateInfo makes TelemetryStateInfo from interface
func InterfaceToTelemetryStateInfo(iData interface{}) *TelemetryStateInfo {
	data := iData.(map[string]interface{})
	return &TelemetryStateInfo{

		Resource: InterfaceToTelemetryResourceInfoSlice(data["resource"]),

		//{"type":"array","item":{"type":"object","properties":{"name":{"type":"string"},"path":{"type":"string"},"rate":{"type":"string"}}}}
		ServerPort: data["server_port"].(int),

		//{"type":"integer"}
		ServerIP: data["server_ip"].(string),

		//{"type":"string"}

	}
}

// InterfaceToTelemetryStateInfoSlice makes a slice of TelemetryStateInfo from interface
func InterfaceToTelemetryStateInfoSlice(data interface{}) []*TelemetryStateInfo {
	list := data.([]interface{})
	result := MakeTelemetryStateInfoSlice()
	for _, item := range list {
		result = append(result, InterfaceToTelemetryStateInfo(item))
	}
	return result
}

// MakeTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
	return []*TelemetryStateInfo{}
}
