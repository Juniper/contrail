package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeMacAddressesType makes MacAddressesType
// nolint
func MakeMacAddressesType() *MacAddressesType {
	return &MacAddressesType{
		//TODO(nati): Apply default
		MacAddress: []string{},
	}
}

// MakeMacAddressesType makes MacAddressesType
// nolint
func InterfaceToMacAddressesType(i interface{}) *MacAddressesType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &MacAddressesType{
		//TODO(nati): Apply default
		MacAddress: common.InterfaceToStringList(m["mac_address"]),
	}
}

// MakeMacAddressesTypeSlice() makes a slice of MacAddressesType
// nolint
func MakeMacAddressesTypeSlice() []*MacAddressesType {
	return []*MacAddressesType{}
}

// InterfaceToMacAddressesTypeSlice() makes a slice of MacAddressesType
// nolint
func InterfaceToMacAddressesTypeSlice(i interface{}) []*MacAddressesType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*MacAddressesType{}
	for _, item := range list {
		result = append(result, InterfaceToMacAddressesType(item))
	}
	return result
}
