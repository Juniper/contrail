package models

// ConfigNode

import "encoding/json"

// ConfigNode
type ConfigNode struct {
	ParentUUID          string         `json:"parent_uuid,omitempty"`
	ParentType          string         `json:"parent_type,omitempty"`
	FQName              []string       `json:"fq_name,omitempty"`
	ConfigNodeIPAddress IpAddressType  `json:"config_node_ip_address,omitempty"`
	Annotations         *KeyValuePairs `json:"annotations,omitempty"`
	UUID                string         `json:"uuid,omitempty"`
	Perms2              *PermType2     `json:"perms2,omitempty"`
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
		FQName:              []string{},
		ConfigNodeIPAddress: MakeIpAddressType(),
		Annotations:         MakeKeyValuePairs(),
		UUID:                "",
		ParentUUID:          "",
		ParentType:          "",
		Perms2:              MakePermType2(),
		IDPerms:             MakeIdPermsType(),
		DisplayName:         "",
	}
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
func MakeConfigNodeSlice() []*ConfigNode {
	return []*ConfigNode{}
}
