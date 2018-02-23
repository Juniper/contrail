package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeApplicationPolicySet makes ApplicationPolicySet
func MakeApplicationPolicySet() *ApplicationPolicySet {
	return &ApplicationPolicySet{
		//TODO(nati): Apply default
		UUID:            "",
		ParentUUID:      "",
		ParentType:      "",
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		AllApplications: false,
	}
}

// MakeApplicationPolicySet makes ApplicationPolicySet
func InterfaceToApplicationPolicySet(i interface{}) *ApplicationPolicySet {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ApplicationPolicySet{
		//TODO(nati): Apply default
		UUID:            schema.InterfaceToString(m["uuid"]),
		ParentUUID:      schema.InterfaceToString(m["parent_uuid"]),
		ParentType:      schema.InterfaceToString(m["parent_type"]),
		FQName:          schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:         InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:     schema.InterfaceToString(m["display_name"]),
		Annotations:     InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:          InterfaceToPermType2(m["perms2"]),
		AllApplications: schema.InterfaceToBool(m["all_applications"]),
	}
}

// MakeApplicationPolicySetSlice() makes a slice of ApplicationPolicySet
func MakeApplicationPolicySetSlice() []*ApplicationPolicySet {
	return []*ApplicationPolicySet{}
}

// InterfaceToApplicationPolicySetSlice() makes a slice of ApplicationPolicySet
func InterfaceToApplicationPolicySetSlice(i interface{}) []*ApplicationPolicySet {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ApplicationPolicySet{}
	for _, item := range list {
		result = append(result, InterfaceToApplicationPolicySet(item))
	}
	return result
}
