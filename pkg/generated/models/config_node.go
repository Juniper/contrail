package models

// ConfigNode

import "encoding/json"

// ConfigNode
type ConfigNode struct {
	ConfigNodeIPAddress IpAddressType  `json:"config_node_ip_address"`
	Annotations         *KeyValuePairs `json:"annotations"`
	Perms2              *PermType2     `json:"perms2"`
	ParentType          string         `json:"parent_type"`
	UUID                string         `json:"uuid"`
	ParentUUID          string         `json:"parent_uuid"`
	FQName              []string       `json:"fq_name"`
	IDPerms             *IdPermsType   `json:"id_perms"`
	DisplayName         string         `json:"display_name"`
}

// String returns json representation of the object
func (model *ConfigNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeConfigNode makes ConfigNode
func MakeConfigNode() *ConfigNode {
	return &ConfigNode{
		//TODO(nati): Apply default
		IDPerms:             MakeIdPermsType(),
		DisplayName:         "",
		UUID:                "",
		ParentUUID:          "",
		FQName:              []string{},
		ParentType:          "",
		ConfigNodeIPAddress: MakeIpAddressType(),
		Annotations:         MakeKeyValuePairs(),
		Perms2:              MakePermType2(),
	}
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
func MakeConfigNodeSlice() []*ConfigNode {
	return []*ConfigNode{}
}
