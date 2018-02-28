package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeSubnetType makes SubnetType
// nolint
func MakeSubnetType() *SubnetType {
	return &SubnetType{
		//TODO(nati): Apply default
		IPPrefix:    "",
		IPPrefixLen: 0,
	}
}

// MakeSubnetType makes SubnetType
// nolint
func InterfaceToSubnetType(i interface{}) *SubnetType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SubnetType{
		//TODO(nati): Apply default
		IPPrefix:    common.InterfaceToString(m["ip_prefix"]),
		IPPrefixLen: common.InterfaceToInt64(m["ip_prefix_len"]),
	}
}

// MakeSubnetTypeSlice() makes a slice of SubnetType
// nolint
func MakeSubnetTypeSlice() []*SubnetType {
	return []*SubnetType{}
}

// InterfaceToSubnetTypeSlice() makes a slice of SubnetType
// nolint
func InterfaceToSubnetTypeSlice(i interface{}) []*SubnetType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SubnetType{}
	for _, item := range list {
		result = append(result, InterfaceToSubnetType(item))
	}
	return result
}
