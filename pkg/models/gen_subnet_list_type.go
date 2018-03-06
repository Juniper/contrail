package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeSubnetListType makes SubnetListType
// nolint
func MakeSubnetListType() *SubnetListType {
	return &SubnetListType{
		//TODO(nati): Apply default

		Subnet: MakeSubnetTypeSlice(),
	}
}

// MakeSubnetListType makes SubnetListType
// nolint
func InterfaceToSubnetListType(i interface{}) *SubnetListType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SubnetListType{
		//TODO(nati): Apply default

		Subnet: InterfaceToSubnetTypeSlice(m["subnet"]),
	}
}

// MakeSubnetListTypeSlice() makes a slice of SubnetListType
// nolint
func MakeSubnetListTypeSlice() []*SubnetListType {
	return []*SubnetListType{}
}

// InterfaceToSubnetListTypeSlice() makes a slice of SubnetListType
// nolint
func InterfaceToSubnetListTypeSlice(i interface{}) []*SubnetListType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SubnetListType{}
	for _, item := range list {
		result = append(result, InterfaceToSubnetListType(item))
	}
	return result
}
