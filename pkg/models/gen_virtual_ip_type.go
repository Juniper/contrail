package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualIpType makes VirtualIpType
// nolint
func MakeVirtualIpType() *VirtualIpType {
	return &VirtualIpType{
		//TODO(nati): Apply default
		Status:                "",
		StatusDescription:     "",
		Protocol:              "",
		SubnetID:              "",
		PersistenceCookieName: "",
		ConnectionLimit:       0,
		PersistenceType:       "",
		AdminState:            false,
		Address:               "",
		ProtocolPort:          0,
	}
}

// MakeVirtualIpType makes VirtualIpType
// nolint
func InterfaceToVirtualIpType(i interface{}) *VirtualIpType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualIpType{
		//TODO(nati): Apply default
		Status:                common.InterfaceToString(m["status"]),
		StatusDescription:     common.InterfaceToString(m["status_description"]),
		Protocol:              common.InterfaceToString(m["protocol"]),
		SubnetID:              common.InterfaceToString(m["subnet_id"]),
		PersistenceCookieName: common.InterfaceToString(m["persistence_cookie_name"]),
		ConnectionLimit:       common.InterfaceToInt64(m["connection_limit"]),
		PersistenceType:       common.InterfaceToString(m["persistence_type"]),
		AdminState:            common.InterfaceToBool(m["admin_state"]),
		Address:               common.InterfaceToString(m["address"]),
		ProtocolPort:          common.InterfaceToInt64(m["protocol_port"]),
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
// nolint
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}

// InterfaceToVirtualIpTypeSlice() makes a slice of VirtualIpType
// nolint
func InterfaceToVirtualIpTypeSlice(i interface{}) []*VirtualIpType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualIpType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualIpType(item))
	}
	return result
}
