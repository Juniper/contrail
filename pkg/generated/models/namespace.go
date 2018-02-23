package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeNamespace makes Namespace
func MakeNamespace() *Namespace {
	return &Namespace{
		//TODO(nati): Apply default
		UUID:          "",
		ParentUUID:    "",
		ParentType:    "",
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
		NamespaceCidr: MakeSubnetType(),
	}
}

// MakeNamespace makes Namespace
func InterfaceToNamespace(i interface{}) *Namespace {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Namespace{
		//TODO(nati): Apply default
		UUID:          schema.InterfaceToString(m["uuid"]),
		ParentUUID:    schema.InterfaceToString(m["parent_uuid"]),
		ParentType:    schema.InterfaceToString(m["parent_type"]),
		FQName:        schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:       InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:   schema.InterfaceToString(m["display_name"]),
		Annotations:   InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:        InterfaceToPermType2(m["perms2"]),
		NamespaceCidr: InterfaceToSubnetType(m["namespace_cidr"]),
	}
}

// MakeNamespaceSlice() makes a slice of Namespace
func MakeNamespaceSlice() []*Namespace {
	return []*Namespace{}
}

// InterfaceToNamespaceSlice() makes a slice of Namespace
func InterfaceToNamespaceSlice(i interface{}) []*Namespace {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Namespace{}
	for _, item := range list {
		result = append(result, InterfaceToNamespace(item))
	}
	return result
}
