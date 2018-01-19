package models

// ConfigNode

import "encoding/json"

// ConfigNode
type ConfigNode struct {
	UUID                string         `json:"uuid,omitempty"`
	ParentType          string         `json:"parent_type,omitempty"`
	FQName              []string       `json:"fq_name,omitempty"`
	IDPerms             *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
	Perms2              *PermType2     `json:"perms2,omitempty"`
	ParentUUID          string         `json:"parent_uuid,omitempty"`
	ConfigNodeIPAddress IpAddressType  `json:"config_node_ip_address,omitempty"`
	Annotations         *KeyValuePairs `json:"annotations,omitempty"`
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
		ConfigNodeIPAddress: MakeIpAddressType(),
		Annotations:         MakeKeyValuePairs(),
		ParentUUID:          "",
		DisplayName:         "",
		Perms2:              MakePermType2(),
		UUID:                "",
		ParentType:          "",
		FQName:              []string{},
		IDPerms:             MakeIdPermsType(),
	}
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
func MakeConfigNodeSlice() []*ConfigNode {
	return []*ConfigNode{}
}
