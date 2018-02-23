package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakePeeringPolicy makes PeeringPolicy
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
func InterfaceToPeeringPolicy(i interface{}) *PeeringPolicy {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PeeringPolicy{
		//TODO(nati): Apply default
		UUID:           schema.InterfaceToString(m["uuid"]),
		ParentUUID:     schema.InterfaceToString(m["parent_uuid"]),
		ParentType:     schema.InterfaceToString(m["parent_type"]),
		FQName:         schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:        InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:    schema.InterfaceToString(m["display_name"]),
		Annotations:    InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:         InterfaceToPermType2(m["perms2"]),
		PeeringService: schema.InterfaceToString(m["peering_service"]),
	}
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}

// InterfaceToPeeringPolicySlice() makes a slice of PeeringPolicy
func InterfaceToPeeringPolicySlice(i interface{}) []*PeeringPolicy {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PeeringPolicy{}
	for _, item := range list {
		result = append(result, InterfaceToPeeringPolicy(item))
	}
	return result
}
