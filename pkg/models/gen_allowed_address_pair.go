package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAllowedAddressPair makes AllowedAddressPair
// nolint
func MakeAllowedAddressPair() *AllowedAddressPair {
	return &AllowedAddressPair{
		//TODO(nati): Apply default
		IP:          MakeSubnetType(),
		Mac:         "",
		AddressMode: "",
	}
}

// MakeAllowedAddressPair makes AllowedAddressPair
// nolint
func InterfaceToAllowedAddressPair(i interface{}) *AllowedAddressPair {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AllowedAddressPair{
		//TODO(nati): Apply default
		IP:          InterfaceToSubnetType(m["ip"]),
		Mac:         common.InterfaceToString(m["mac"]),
		AddressMode: common.InterfaceToString(m["address_mode"]),
	}
}

// MakeAllowedAddressPairSlice() makes a slice of AllowedAddressPair
// nolint
func MakeAllowedAddressPairSlice() []*AllowedAddressPair {
	return []*AllowedAddressPair{}
}

// InterfaceToAllowedAddressPairSlice() makes a slice of AllowedAddressPair
// nolint
func InterfaceToAllowedAddressPairSlice(i interface{}) []*AllowedAddressPair {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AllowedAddressPair{}
	for _, item := range list {
		result = append(result, InterfaceToAllowedAddressPair(item))
	}
	return result
}
