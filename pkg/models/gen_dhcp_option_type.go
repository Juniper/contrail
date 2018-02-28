package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDhcpOptionType makes DhcpOptionType
// nolint
func MakeDhcpOptionType() *DhcpOptionType {
	return &DhcpOptionType{
		//TODO(nati): Apply default
		DHCPOptionValue:      "",
		DHCPOptionValueBytes: "",
		DHCPOptionName:       "",
	}
}

// MakeDhcpOptionType makes DhcpOptionType
// nolint
func InterfaceToDhcpOptionType(i interface{}) *DhcpOptionType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DhcpOptionType{
		//TODO(nati): Apply default
		DHCPOptionValue:      common.InterfaceToString(m["dhcp_option_value"]),
		DHCPOptionValueBytes: common.InterfaceToString(m["dhcp_option_value_bytes"]),
		DHCPOptionName:       common.InterfaceToString(m["dhcp_option_name"]),
	}
}

// MakeDhcpOptionTypeSlice() makes a slice of DhcpOptionType
// nolint
func MakeDhcpOptionTypeSlice() []*DhcpOptionType {
	return []*DhcpOptionType{}
}

// InterfaceToDhcpOptionTypeSlice() makes a slice of DhcpOptionType
// nolint
func InterfaceToDhcpOptionTypeSlice(i interface{}) []*DhcpOptionType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DhcpOptionType{}
	for _, item := range list {
		result = append(result, InterfaceToDhcpOptionType(item))
	}
	return result
}
