package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeConfigNode makes ConfigNode
func MakeConfigNode() *ConfigNode {
	return &ConfigNode{
		//TODO(nati): Apply default
		UUID:                "",
		ParentUUID:          "",
		ParentType:          "",
		FQName:              []string{},
		IDPerms:             MakeIdPermsType(),
		DisplayName:         "",
		Annotations:         MakeKeyValuePairs(),
		Perms2:              MakePermType2(),
		ConfigNodeIPAddress: "",
	}
}

// MakeConfigNode makes ConfigNode
func InterfaceToConfigNode(i interface{}) *ConfigNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ConfigNode{
		//TODO(nati): Apply default
		UUID:                schema.InterfaceToString(m["uuid"]),
		ParentUUID:          schema.InterfaceToString(m["parent_uuid"]),
		ParentType:          schema.InterfaceToString(m["parent_type"]),
		FQName:              schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:             InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:         schema.InterfaceToString(m["display_name"]),
		Annotations:         InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:              InterfaceToPermType2(m["perms2"]),
		ConfigNodeIPAddress: schema.InterfaceToString(m["config_node_ip_address"]),
	}
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
func MakeConfigNodeSlice() []*ConfigNode {
	return []*ConfigNode{}
}

// InterfaceToConfigNodeSlice() makes a slice of ConfigNode
func InterfaceToConfigNodeSlice(i interface{}) []*ConfigNode {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ConfigNode{}
	for _, item := range list {
		result = append(result, InterfaceToConfigNode(item))
	}
	return result
}
