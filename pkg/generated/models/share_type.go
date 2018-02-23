package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeShareType makes ShareType
func MakeShareType() *ShareType {
	return &ShareType{
		//TODO(nati): Apply default
		TenantAccess: 0,
		Tenant:       "",
	}
}

// MakeShareType makes ShareType
func InterfaceToShareType(i interface{}) *ShareType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ShareType{
		//TODO(nati): Apply default
		TenantAccess: schema.InterfaceToInt64(m["tenant_access"]),
		Tenant:       schema.InterfaceToString(m["tenant"]),
	}
}

// MakeShareTypeSlice() makes a slice of ShareType
func MakeShareTypeSlice() []*ShareType {
	return []*ShareType{}
}

// InterfaceToShareTypeSlice() makes a slice of ShareType
func InterfaceToShareTypeSlice(i interface{}) []*ShareType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ShareType{}
	for _, item := range list {
		result = append(result, InterfaceToShareType(item))
	}
	return result
}
