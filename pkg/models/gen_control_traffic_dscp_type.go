package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeControlTrafficDscpType makes ControlTrafficDscpType
// nolint
func MakeControlTrafficDscpType() *ControlTrafficDscpType {
	return &ControlTrafficDscpType{
		//TODO(nati): Apply default
		Control:   0,
		Analytics: 0,
		DNS:       0,
	}
}

// MakeControlTrafficDscpType makes ControlTrafficDscpType
// nolint
func InterfaceToControlTrafficDscpType(i interface{}) *ControlTrafficDscpType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ControlTrafficDscpType{
		//TODO(nati): Apply default
		Control:   common.InterfaceToInt64(m["control"]),
		Analytics: common.InterfaceToInt64(m["analytics"]),
		DNS:       common.InterfaceToInt64(m["dns"]),
	}
}

// MakeControlTrafficDscpTypeSlice() makes a slice of ControlTrafficDscpType
// nolint
func MakeControlTrafficDscpTypeSlice() []*ControlTrafficDscpType {
	return []*ControlTrafficDscpType{}
}

// InterfaceToControlTrafficDscpTypeSlice() makes a slice of ControlTrafficDscpType
// nolint
func InterfaceToControlTrafficDscpTypeSlice(i interface{}) []*ControlTrafficDscpType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ControlTrafficDscpType{}
	for _, item := range list {
		result = append(result, InterfaceToControlTrafficDscpType(item))
	}
	return result
}
