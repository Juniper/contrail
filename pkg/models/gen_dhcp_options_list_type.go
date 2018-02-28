package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDhcpOptionsListType makes DhcpOptionsListType
// nolint
func MakeDhcpOptionsListType() *DhcpOptionsListType {
	return &DhcpOptionsListType{
		//TODO(nati): Apply default

		DHCPOption: MakeDhcpOptionTypeSlice(),
	}
}

// MakeDhcpOptionsListType makes DhcpOptionsListType
// nolint
func InterfaceToDhcpOptionsListType(i interface{}) *DhcpOptionsListType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DhcpOptionsListType{
		//TODO(nati): Apply default

		DHCPOption: InterfaceToDhcpOptionTypeSlice(m["dhcp_option"]),
	}
}

// MakeDhcpOptionsListTypeSlice() makes a slice of DhcpOptionsListType
// nolint
func MakeDhcpOptionsListTypeSlice() []*DhcpOptionsListType {
	return []*DhcpOptionsListType{}
}

// InterfaceToDhcpOptionsListTypeSlice() makes a slice of DhcpOptionsListType
// nolint
func InterfaceToDhcpOptionsListTypeSlice(i interface{}) []*DhcpOptionsListType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DhcpOptionsListType{}
	for _, item := range list {
		result = append(result, InterfaceToDhcpOptionsListType(item))
	}
	return result
}
