package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeShareType makes ShareType
// nolint
func MakeShareType() *ShareType {
	return &ShareType{
		//TODO(nati): Apply default
		TenantAccess: 0,
		Tenant:       "",
	}
}

// MakeShareType makes ShareType
// nolint
func InterfaceToShareType(i interface{}) *ShareType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ShareType{
		//TODO(nati): Apply default
		TenantAccess: common.InterfaceToInt64(m["tenant_access"]),
		Tenant:       common.InterfaceToString(m["tenant"]),
	}
}

// MakeShareTypeSlice() makes a slice of ShareType
// nolint
func MakeShareTypeSlice() []*ShareType {
	return []*ShareType{}
}

// InterfaceToShareTypeSlice() makes a slice of ShareType
// nolint
func InterfaceToShareTypeSlice(i interface{}) []*ShareType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ShareType{}
	for _, item := range list {
		result = append(result, InterfaceToShareType(item))
	}
	return result
}
