package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeOpenStackAddress makes OpenStackAddress
// nolint
func MakeOpenStackAddress() *OpenStackAddress {
	return &OpenStackAddress{
		//TODO(nati): Apply default
		Addr: "",
	}
}

// MakeOpenStackAddress makes OpenStackAddress
// nolint
func InterfaceToOpenStackAddress(i interface{}) *OpenStackAddress {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OpenStackAddress{
		//TODO(nati): Apply default
		Addr: common.InterfaceToString(m["addr"]),
	}
}

// MakeOpenStackAddressSlice() makes a slice of OpenStackAddress
// nolint
func MakeOpenStackAddressSlice() []*OpenStackAddress {
	return []*OpenStackAddress{}
}

// InterfaceToOpenStackAddressSlice() makes a slice of OpenStackAddress
// nolint
func InterfaceToOpenStackAddressSlice(i interface{}) []*OpenStackAddress {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OpenStackAddress{}
	for _, item := range list {
		result = append(result, InterfaceToOpenStackAddress(item))
	}
	return result
}
