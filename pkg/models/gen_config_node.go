package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeConfigNode makes ConfigNode
// nolint
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
// nolint
func InterfaceToConfigNode(i interface{}) *ConfigNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ConfigNode{
		//TODO(nati): Apply default
		UUID:                common.InterfaceToString(m["uuid"]),
		ParentUUID:          common.InterfaceToString(m["parent_uuid"]),
		ParentType:          common.InterfaceToString(m["parent_type"]),
		FQName:              common.InterfaceToStringList(m["fq_name"]),
		IDPerms:             InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:         common.InterfaceToString(m["display_name"]),
		Annotations:         InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:              InterfaceToPermType2(m["perms2"]),
		ConfigNodeIPAddress: common.InterfaceToString(m["config_node_ip_address"]),
	}
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
// nolint
func MakeConfigNodeSlice() []*ConfigNode {
	return []*ConfigNode{}
}

// InterfaceToConfigNodeSlice() makes a slice of ConfigNode
// nolint
func InterfaceToConfigNodeSlice(i interface{}) []*ConfigNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ConfigNode{}
	for _, item := range list {
		result = append(result, InterfaceToConfigNode(item))
	}
	return result
}
