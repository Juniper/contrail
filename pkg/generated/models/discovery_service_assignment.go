package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeDiscoveryServiceAssignment makes DiscoveryServiceAssignment
func MakeDiscoveryServiceAssignment() *DiscoveryServiceAssignment {
	return &DiscoveryServiceAssignment{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeDiscoveryServiceAssignment makes DiscoveryServiceAssignment
func InterfaceToDiscoveryServiceAssignment(i interface{}) *DiscoveryServiceAssignment {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DiscoveryServiceAssignment{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
	}
}

// MakeDiscoveryServiceAssignmentSlice() makes a slice of DiscoveryServiceAssignment
func MakeDiscoveryServiceAssignmentSlice() []*DiscoveryServiceAssignment {
	return []*DiscoveryServiceAssignment{}
}

// InterfaceToDiscoveryServiceAssignmentSlice() makes a slice of DiscoveryServiceAssignment
func InterfaceToDiscoveryServiceAssignmentSlice(i interface{}) []*DiscoveryServiceAssignment {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DiscoveryServiceAssignment{}
	for _, item := range list {
		result = append(result, InterfaceToDiscoveryServiceAssignment(item))
	}
	return result
}
