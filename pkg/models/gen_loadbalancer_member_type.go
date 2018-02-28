package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerMemberType makes LoadbalancerMemberType
// nolint
func MakeLoadbalancerMemberType() *LoadbalancerMemberType {
	return &LoadbalancerMemberType{
		//TODO(nati): Apply default
		Status:            "",
		StatusDescription: "",
		Weight:            0,
		AdminState:        false,
		Address:           "",
		ProtocolPort:      0,
	}
}

// MakeLoadbalancerMemberType makes LoadbalancerMemberType
// nolint
func InterfaceToLoadbalancerMemberType(i interface{}) *LoadbalancerMemberType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerMemberType{
		//TODO(nati): Apply default
		Status:            common.InterfaceToString(m["status"]),
		StatusDescription: common.InterfaceToString(m["status_description"]),
		Weight:            common.InterfaceToInt64(m["weight"]),
		AdminState:        common.InterfaceToBool(m["admin_state"]),
		Address:           common.InterfaceToString(m["address"]),
		ProtocolPort:      common.InterfaceToInt64(m["protocol_port"]),
	}
}

// MakeLoadbalancerMemberTypeSlice() makes a slice of LoadbalancerMemberType
// nolint
func MakeLoadbalancerMemberTypeSlice() []*LoadbalancerMemberType {
	return []*LoadbalancerMemberType{}
}

// InterfaceToLoadbalancerMemberTypeSlice() makes a slice of LoadbalancerMemberType
// nolint
func InterfaceToLoadbalancerMemberTypeSlice(i interface{}) []*LoadbalancerMemberType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerMemberType{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerMemberType(item))
	}
	return result
}
