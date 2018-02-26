package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeSubnetType makes SubnetType
func MakeSubnetType() *SubnetType {
	return &SubnetType{
		//TODO(nati): Apply default
		IPPrefix:    "",
		IPPrefixLen: 0,
	}
}

// MakeSubnetType makes SubnetType
func InterfaceToSubnetType(i interface{}) *SubnetType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SubnetType{
		//TODO(nati): Apply default
		IPPrefix:    schema.InterfaceToString(m["ip_prefix"]),
		IPPrefixLen: schema.InterfaceToInt64(m["ip_prefix_len"]),
	}
}

// MakeSubnetTypeSlice() makes a slice of SubnetType
func MakeSubnetTypeSlice() []*SubnetType {
	return []*SubnetType{}
}

// InterfaceToSubnetTypeSlice() makes a slice of SubnetType
func InterfaceToSubnetTypeSlice(i interface{}) []*SubnetType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SubnetType{}
	for _, item := range list {
		result = append(result, InterfaceToSubnetType(item))
	}
	return result
}
