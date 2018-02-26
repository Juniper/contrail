package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeMACLimitControlType makes MACLimitControlType
func MakeMACLimitControlType() *MACLimitControlType {
	return &MACLimitControlType{
		//TODO(nati): Apply default
		MacLimit:       0,
		MacLimitAction: "",
	}
}

// MakeMACLimitControlType makes MACLimitControlType
func InterfaceToMACLimitControlType(i interface{}) *MACLimitControlType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &MACLimitControlType{
		//TODO(nati): Apply default
		MacLimit:       schema.InterfaceToInt64(m["mac_limit"]),
		MacLimitAction: schema.InterfaceToString(m["mac_limit_action"]),
	}
}

// MakeMACLimitControlTypeSlice() makes a slice of MACLimitControlType
func MakeMACLimitControlTypeSlice() []*MACLimitControlType {
	return []*MACLimitControlType{}
}

// InterfaceToMACLimitControlTypeSlice() makes a slice of MACLimitControlType
func InterfaceToMACLimitControlTypeSlice(i interface{}) []*MACLimitControlType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*MACLimitControlType{}
	for _, item := range list {
		result = append(result, InterfaceToMACLimitControlType(item))
	}
	return result
}
