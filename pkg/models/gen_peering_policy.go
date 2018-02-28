package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePeeringPolicy makes PeeringPolicy
// nolint
func MakePeeringPolicy() *PeeringPolicy {
	return &PeeringPolicy{
		//TODO(nati): Apply default
		UUID:           "",
		ParentUUID:     "",
		ParentType:     "",
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		PeeringService: "",
	}
}

// MakePeeringPolicy makes PeeringPolicy
// nolint
func InterfaceToPeeringPolicy(i interface{}) *PeeringPolicy {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PeeringPolicy{
		//TODO(nati): Apply default
		UUID:           common.InterfaceToString(m["uuid"]),
		ParentUUID:     common.InterfaceToString(m["parent_uuid"]),
		ParentType:     common.InterfaceToString(m["parent_type"]),
		FQName:         common.InterfaceToStringList(m["fq_name"]),
		IDPerms:        InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:    common.InterfaceToString(m["display_name"]),
		Annotations:    InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:         InterfaceToPermType2(m["perms2"]),
		PeeringService: common.InterfaceToString(m["peering_service"]),
	}
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
// nolint
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}

// InterfaceToPeeringPolicySlice() makes a slice of PeeringPolicy
// nolint
func InterfaceToPeeringPolicySlice(i interface{}) []*PeeringPolicy {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PeeringPolicy{}
	for _, item := range list {
		result = append(result, InterfaceToPeeringPolicy(item))
	}
	return result
}
