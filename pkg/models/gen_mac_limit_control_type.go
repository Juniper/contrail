package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeMACLimitControlType makes MACLimitControlType
// nolint
func MakeMACLimitControlType() *MACLimitControlType {
	return &MACLimitControlType{
		//TODO(nati): Apply default
		MacLimit:       0,
		MacLimitAction: "",
	}
}

// MakeMACLimitControlType makes MACLimitControlType
// nolint
func InterfaceToMACLimitControlType(i interface{}) *MACLimitControlType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &MACLimitControlType{
		//TODO(nati): Apply default
		MacLimit:       common.InterfaceToInt64(m["mac_limit"]),
		MacLimitAction: common.InterfaceToString(m["mac_limit_action"]),
	}
}

// MakeMACLimitControlTypeSlice() makes a slice of MACLimitControlType
// nolint
func MakeMACLimitControlTypeSlice() []*MACLimitControlType {
	return []*MACLimitControlType{}
}

// InterfaceToMACLimitControlTypeSlice() makes a slice of MACLimitControlType
// nolint
func InterfaceToMACLimitControlTypeSlice(i interface{}) []*MACLimitControlType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*MACLimitControlType{}
	for _, item := range list {
		result = append(result, InterfaceToMACLimitControlType(item))
	}
	return result
}
