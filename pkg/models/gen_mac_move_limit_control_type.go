package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeMACMoveLimitControlType makes MACMoveLimitControlType
// nolint
func MakeMACMoveLimitControlType() *MACMoveLimitControlType {
	return &MACMoveLimitControlType{
		//TODO(nati): Apply default
		MacMoveTimeWindow:  0,
		MacMoveLimit:       0,
		MacMoveLimitAction: "",
	}
}

// MakeMACMoveLimitControlType makes MACMoveLimitControlType
// nolint
func InterfaceToMACMoveLimitControlType(i interface{}) *MACMoveLimitControlType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &MACMoveLimitControlType{
		//TODO(nati): Apply default
		MacMoveTimeWindow:  common.InterfaceToInt64(m["mac_move_time_window"]),
		MacMoveLimit:       common.InterfaceToInt64(m["mac_move_limit"]),
		MacMoveLimitAction: common.InterfaceToString(m["mac_move_limit_action"]),
	}
}

// MakeMACMoveLimitControlTypeSlice() makes a slice of MACMoveLimitControlType
// nolint
func MakeMACMoveLimitControlTypeSlice() []*MACMoveLimitControlType {
	return []*MACMoveLimitControlType{}
}

// InterfaceToMACMoveLimitControlTypeSlice() makes a slice of MACMoveLimitControlType
// nolint
func InterfaceToMACMoveLimitControlTypeSlice(i interface{}) []*MACMoveLimitControlType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*MACMoveLimitControlType{}
	for _, item := range list {
		result = append(result, InterfaceToMACMoveLimitControlType(item))
	}
	return result
}
