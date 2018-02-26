package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeMACMoveLimitControlType makes MACMoveLimitControlType
func MakeMACMoveLimitControlType() *MACMoveLimitControlType {
	return &MACMoveLimitControlType{
		//TODO(nati): Apply default
		MacMoveTimeWindow:  0,
		MacMoveLimit:       0,
		MacMoveLimitAction: "",
	}
}

// MakeMACMoveLimitControlType makes MACMoveLimitControlType
func InterfaceToMACMoveLimitControlType(i interface{}) *MACMoveLimitControlType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &MACMoveLimitControlType{
		//TODO(nati): Apply default
		MacMoveTimeWindow:  schema.InterfaceToInt64(m["mac_move_time_window"]),
		MacMoveLimit:       schema.InterfaceToInt64(m["mac_move_limit"]),
		MacMoveLimitAction: schema.InterfaceToString(m["mac_move_limit_action"]),
	}
}

// MakeMACMoveLimitControlTypeSlice() makes a slice of MACMoveLimitControlType
func MakeMACMoveLimitControlTypeSlice() []*MACMoveLimitControlType {
	return []*MACMoveLimitControlType{}
}

// InterfaceToMACMoveLimitControlTypeSlice() makes a slice of MACMoveLimitControlType
func InterfaceToMACMoveLimitControlTypeSlice(i interface{}) []*MACMoveLimitControlType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*MACMoveLimitControlType{}
	for _, item := range list {
		result = append(result, InterfaceToMACMoveLimitControlType(item))
	}
	return result
}
