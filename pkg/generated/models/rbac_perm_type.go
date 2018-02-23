package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeRbacPermType makes RbacPermType
func MakeRbacPermType() *RbacPermType {
	return &RbacPermType{
		//TODO(nati): Apply default
		RoleCrud: "",
		RoleName: "",
	}
}

// MakeRbacPermType makes RbacPermType
func InterfaceToRbacPermType(i interface{}) *RbacPermType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RbacPermType{
		//TODO(nati): Apply default
		RoleCrud: schema.InterfaceToString(m["role_crud"]),
		RoleName: schema.InterfaceToString(m["role_name"]),
	}
}

// MakeRbacPermTypeSlice() makes a slice of RbacPermType
func MakeRbacPermTypeSlice() []*RbacPermType {
	return []*RbacPermType{}
}

// InterfaceToRbacPermTypeSlice() makes a slice of RbacPermType
func InterfaceToRbacPermTypeSlice(i interface{}) []*RbacPermType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RbacPermType{}
	for _, item := range list {
		result = append(result, InterfaceToRbacPermType(item))
	}
	return result
}
