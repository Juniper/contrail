package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeTelemetryStateInfo makes TelemetryStateInfo
// nolint
func MakeTelemetryStateInfo() *TelemetryStateInfo {
	return &TelemetryStateInfo{
		//TODO(nati): Apply default

		Resource: MakeTelemetryResourceInfoSlice(),

		ServerPort: 0,
		ServerIP:   "",
	}
}

// MakeTelemetryStateInfo makes TelemetryStateInfo
// nolint
func InterfaceToTelemetryStateInfo(i interface{}) *TelemetryStateInfo {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &TelemetryStateInfo{
		//TODO(nati): Apply default

		Resource: InterfaceToTelemetryResourceInfoSlice(m["resource"]),

		ServerPort: common.InterfaceToInt64(m["server_port"]),
		ServerIP:   common.InterfaceToString(m["server_ip"]),
	}
}

// MakeTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
// nolint
func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
	return []*TelemetryStateInfo{}
}

// InterfaceToTelemetryStateInfoSlice() makes a slice of TelemetryStateInfo
// nolint
func InterfaceToTelemetryStateInfoSlice(i interface{}) []*TelemetryStateInfo {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*TelemetryStateInfo{}
	for _, item := range list {
		result = append(result, InterfaceToTelemetryStateInfo(item))
	}
	return result
}
