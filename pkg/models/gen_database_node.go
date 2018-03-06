package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeDatabaseNode makes DatabaseNode
// nolint
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
// nolint
func InterfaceToDatabaseNode(i interface{}) *DatabaseNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &DatabaseNode{
		//TODO(nati): Apply default
		UUID:                  common.InterfaceToString(m["uuid"]),
		ParentUUID:            common.InterfaceToString(m["parent_uuid"]),
		ParentType:            common.InterfaceToString(m["parent_type"]),
		FQName:                common.InterfaceToStringList(m["fq_name"]),
		IDPerms:               InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:           common.InterfaceToString(m["display_name"]),
		Annotations:           InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                InterfaceToPermType2(m["perms2"]),
		DatabaseNodeIPAddress: common.InterfaceToString(m["database_node_ip_address"]),
	}
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
// nolint
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}

// InterfaceToDatabaseNodeSlice() makes a slice of DatabaseNode
// nolint
func InterfaceToDatabaseNodeSlice(i interface{}) []*DatabaseNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*DatabaseNode{}
	for _, item := range list {
		result = append(result, InterfaceToDatabaseNode(item))
	}
	return result
}
