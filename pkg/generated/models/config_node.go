package models

// ConfigNode

import "encoding/json"

// ConfigNode
type ConfigNode struct {
	ConfigNodeIPAddress IpAddressType  `json:"config_node_ip_address,omitempty"`
	UUID                string         `json:"uuid,omitempty"`
	FQName              []string       `json:"fq_name,omitempty"`
	Annotations         *KeyValuePairs `json:"annotations,omitempty"`
	Perms2              *PermType2     `json:"perms2,omitempty"`
	ParentUUID          string         `json:"parent_uuid,omitempty"`
	ParentType          string         `json:"parent_type,omitempty"`
	IDPerms             *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
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
		DisplayName:         "",
		ParentUUID:          "",
		ParentType:          "",
		IDPerms:             MakeIdPermsType(),
		Annotations:         MakeKeyValuePairs(),
		Perms2:              MakePermType2(),
		ConfigNodeIPAddress: MakeIpAddressType(),
		UUID:                "",
		FQName:              []string{},
	}
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
func MakeConfigNodeSlice() []*ConfigNode {
	return []*ConfigNode{}
}
