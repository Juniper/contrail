package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeDatabaseNode makes DatabaseNode
func MakeDatabaseNode() *DatabaseNode {
	return &DatabaseNode{
		//TODO(nati): Apply default
		UUID:                  "",
		ParentUUID:            "",
		ParentType:            "",
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		DatabaseNodeIPAddress: "",
	}
}

// MakeDatabaseNode makes DatabaseNode
func InterfaceToDatabaseNode(i interface{}) *DatabaseNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DatabaseNode{
		//TODO(nati): Apply default
		UUID:                  schema.InterfaceToString(m["uuid"]),
		ParentUUID:            schema.InterfaceToString(m["parent_uuid"]),
		ParentType:            schema.InterfaceToString(m["parent_type"]),
		FQName:                schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:               InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:           schema.InterfaceToString(m["display_name"]),
		Annotations:           InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                InterfaceToPermType2(m["perms2"]),
		DatabaseNodeIPAddress: schema.InterfaceToString(m["database_node_ip_address"]),
	}
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}

// InterfaceToDatabaseNodeSlice() makes a slice of DatabaseNode
func InterfaceToDatabaseNodeSlice(i interface{}) []*DatabaseNode {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DatabaseNode{}
	for _, item := range list {
		result = append(result, InterfaceToDatabaseNode(item))
	}
	return result
}
