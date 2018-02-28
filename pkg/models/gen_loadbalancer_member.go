package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerMember makes LoadbalancerMember
// nolint
func MakeLoadbalancerMember() *LoadbalancerMember {
	return &LoadbalancerMember{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		LoadbalancerMemberProperties: MakeLoadbalancerMemberType(),
	}
}

// MakeLoadbalancerMember makes LoadbalancerMember
// nolint
func InterfaceToLoadbalancerMember(i interface{}) *LoadbalancerMember {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerMember{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		LoadbalancerMemberProperties: InterfaceToLoadbalancerMemberType(m["loadbalancer_member_properties"]),
	}
}

// MakeLoadbalancerMemberSlice() makes a slice of LoadbalancerMember
// nolint
func MakeLoadbalancerMemberSlice() []*LoadbalancerMember {
	return []*LoadbalancerMember{}
}

// InterfaceToLoadbalancerMemberSlice() makes a slice of LoadbalancerMember
// nolint
func InterfaceToLoadbalancerMemberSlice(i interface{}) []*LoadbalancerMember {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerMember{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerMember(item))
	}
	return result
}
