package models

// ConfigNode

import "encoding/json"

// ConfigNode
type ConfigNode struct {
	FQName              []string       `json:"fq_name,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
	Annotations         *KeyValuePairs `json:"annotations,omitempty"`
	Perms2              *PermType2     `json:"perms2,omitempty"`
	UUID                string         `json:"uuid,omitempty"`
	ParentUUID          string         `json:"parent_uuid,omitempty"`
	ConfigNodeIPAddress IpAddressType  `json:"config_node_ip_address,omitempty"`
	ParentType          string         `json:"parent_type,omitempty"`
	IDPerms             *IdPermsType   `json:"id_perms,omitempty"`
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
		Annotations:         MakeKeyValuePairs(),
		Perms2:              MakePermType2(),
		UUID:                "",
		ParentUUID:          "",
		FQName:              []string{},
		DisplayName:         "",
		IDPerms:             MakeIdPermsType(),
		ConfigNodeIPAddress: MakeIpAddressType(),
		ParentType:          "",
	}
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
func MakeConfigNodeSlice() []*ConfigNode {
	return []*ConfigNode{}
}
