package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAclEntriesType makes AclEntriesType
// nolint
func MakeAclEntriesType() *AclEntriesType {
	return &AclEntriesType{
		//TODO(nati): Apply default
		Dynamic: false,

		ACLRule: MakeAclRuleTypeSlice(),
	}
}

// MakeAclEntriesType makes AclEntriesType
// nolint
func InterfaceToAclEntriesType(i interface{}) *AclEntriesType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AclEntriesType{
		//TODO(nati): Apply default
		Dynamic: common.InterfaceToBool(m["dynamic"]),

		ACLRule: InterfaceToAclRuleTypeSlice(m["acl_rule"]),
	}
}

// MakeAclEntriesTypeSlice() makes a slice of AclEntriesType
// nolint
func MakeAclEntriesTypeSlice() []*AclEntriesType {
	return []*AclEntriesType{}
}

// InterfaceToAclEntriesTypeSlice() makes a slice of AclEntriesType
// nolint
func InterfaceToAclEntriesTypeSlice(i interface{}) []*AclEntriesType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AclEntriesType{}
	for _, item := range list {
		result = append(result, InterfaceToAclEntriesType(item))
	}
	return result
}
