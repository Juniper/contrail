package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakePermType makes PermType
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
func InterfaceToPermType(i interface{}) *PermType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PermType{
		//TODO(nati): Apply default
		Owner:       schema.InterfaceToString(m["owner"]),
		OwnerAccess: schema.InterfaceToInt64(m["owner_access"]),
		OtherAccess: schema.InterfaceToInt64(m["other_access"]),
		Group:       schema.InterfaceToString(m["group"]),
		GroupAccess: schema.InterfaceToInt64(m["group_access"]),
	}
}

// MakePermTypeSlice() makes a slice of PermType
func MakePermTypeSlice() []*PermType {
	return []*PermType{}
}

// InterfaceToPermTypeSlice() makes a slice of PermType
func InterfaceToPermTypeSlice(i interface{}) []*PermType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PermType{}
	for _, item := range list {
		result = append(result, InterfaceToPermType(item))
	}
	return result
}
