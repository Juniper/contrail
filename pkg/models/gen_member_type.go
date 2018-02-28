package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeMemberType makes MemberType
// nolint
func MakeMemberType() *MemberType {
	return &MemberType{
		//TODO(nati): Apply default
		Role: "",
	}
}

// MakeMemberType makes MemberType
// nolint
func InterfaceToMemberType(i interface{}) *MemberType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &MemberType{
		//TODO(nati): Apply default
		Role: common.InterfaceToString(m["role"]),
	}
}

// MakeMemberTypeSlice() makes a slice of MemberType
// nolint
func MakeMemberTypeSlice() []*MemberType {
	return []*MemberType{}
}

// InterfaceToMemberTypeSlice() makes a slice of MemberType
// nolint
func InterfaceToMemberTypeSlice(i interface{}) []*MemberType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*MemberType{}
	for _, item := range list {
		result = append(result, InterfaceToMemberType(item))
	}
	return result
}
