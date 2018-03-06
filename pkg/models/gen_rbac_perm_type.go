package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRbacPermType makes RbacPermType
// nolint
func MakeRbacPermType() *RbacPermType {
	return &RbacPermType{
		//TODO(nati): Apply default
		RoleCrud: "",
		RoleName: "",
	}
}

// MakeRbacPermType makes RbacPermType
// nolint
func InterfaceToRbacPermType(i interface{}) *RbacPermType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RbacPermType{
		//TODO(nati): Apply default
		RoleCrud: common.InterfaceToString(m["role_crud"]),
		RoleName: common.InterfaceToString(m["role_name"]),
	}
}

// MakeRbacPermTypeSlice() makes a slice of RbacPermType
// nolint
func MakeRbacPermTypeSlice() []*RbacPermType {
	return []*RbacPermType{}
}

// InterfaceToRbacPermTypeSlice() makes a slice of RbacPermType
// nolint
func InterfaceToRbacPermTypeSlice(i interface{}) []*RbacPermType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RbacPermType{}
	for _, item := range list {
		result = append(result, InterfaceToRbacPermType(item))
	}
	return result
}
