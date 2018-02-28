package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePermType makes PermType
// nolint
func MakePermType() *PermType {
	return &PermType{
		//TODO(nati): Apply default
		Owner:       "",
		OwnerAccess: 0,
		OtherAccess: 0,
		Group:       "",
		GroupAccess: 0,
	}
}

// MakePermType makes PermType
// nolint
func InterfaceToPermType(i interface{}) *PermType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PermType{
		//TODO(nati): Apply default
		Owner:       common.InterfaceToString(m["owner"]),
		OwnerAccess: common.InterfaceToInt64(m["owner_access"]),
		OtherAccess: common.InterfaceToInt64(m["other_access"]),
		Group:       common.InterfaceToString(m["group"]),
		GroupAccess: common.InterfaceToInt64(m["group_access"]),
	}
}

// MakePermTypeSlice() makes a slice of PermType
// nolint
func MakePermTypeSlice() []*PermType {
	return []*PermType{}
}

// InterfaceToPermTypeSlice() makes a slice of PermType
// nolint
func InterfaceToPermTypeSlice(i interface{}) []*PermType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PermType{}
	for _, item := range list {
		result = append(result, InterfaceToPermType(item))
	}
	return result
}
