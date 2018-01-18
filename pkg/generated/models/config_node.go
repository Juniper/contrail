package models

// ConfigNode

import "encoding/json"

// ConfigNode
type ConfigNode struct {
	ConfigNodeIPAddress IpAddressType  `json:"config_node_ip_address,omitempty"`
	Perms2              *PermType2     `json:"perms2,omitempty"`
	UUID                string         `json:"uuid,omitempty"`
	ParentType          string         `json:"parent_type,omitempty"`
	IDPerms             *IdPermsType   `json:"id_perms,omitempty"`
	ParentUUID          string         `json:"parent_uuid,omitempty"`
	FQName              []string       `json:"fq_name,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
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
		ParentUUID:          "",
		FQName:              []string{},
		DisplayName:         "",
		Annotations:         MakeKeyValuePairs(),
		ConfigNodeIPAddress: MakeIpAddressType(),
		Perms2:              MakePermType2(),
		UUID:                "",
		ParentType:          "",
		IDPerms:             MakeIdPermsType(),
	}
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
func MakeConfigNodeSlice() []*ConfigNode {
	return []*ConfigNode{}
}
