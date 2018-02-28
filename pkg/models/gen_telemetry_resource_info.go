package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeTelemetryResourceInfo makes TelemetryResourceInfo
// nolint
func MakeTelemetryResourceInfo() *TelemetryResourceInfo {
	return &TelemetryResourceInfo{
		//TODO(nati): Apply default
		Path: "",
		Rate: "",
		Name: "",
	}
}

// MakeTelemetryResourceInfo makes TelemetryResourceInfo
// nolint
func InterfaceToTelemetryResourceInfo(i interface{}) *TelemetryResourceInfo {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &TelemetryResourceInfo{
		//TODO(nati): Apply default
		Path: common.InterfaceToString(m["path"]),
		Rate: common.InterfaceToString(m["rate"]),
		Name: common.InterfaceToString(m["name"]),
	}
}

// MakeTelemetryResourceInfoSlice() makes a slice of TelemetryResourceInfo
// nolint
func MakeTelemetryResourceInfoSlice() []*TelemetryResourceInfo {
	return []*TelemetryResourceInfo{}
}

// InterfaceToTelemetryResourceInfoSlice() makes a slice of TelemetryResourceInfo
// nolint
func InterfaceToTelemetryResourceInfoSlice(i interface{}) []*TelemetryResourceInfo {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*TelemetryResourceInfo{}
	for _, item := range list {
		result = append(result, InterfaceToTelemetryResourceInfo(item))
	}
	return result
}
