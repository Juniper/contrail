package models

// DatabaseNode

import "encoding/json"

// DatabaseNode
type DatabaseNode struct {
	DisplayName           string         `json:"display_name,omitempty"`
	Perms2                *PermType2     `json:"perms2,omitempty"`
	UUID                  string         `json:"uuid,omitempty"`
	ParentType            string         `json:"parent_type,omitempty"`
	IDPerms               *IdPermsType   `json:"id_perms,omitempty"`
	DatabaseNodeIPAddress IpAddressType  `json:"database_node_ip_address,omitempty"`
	Annotations           *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID            string         `json:"parent_uuid,omitempty"`
	FQName                []string       `json:"fq_name,omitempty"`
}

// String returns json representation of the object
func (model *DatabaseNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDatabaseNode makes DatabaseNode
func MakeDatabaseNode() *DatabaseNode {
	return &DatabaseNode{
		//TODO(nati): Apply default
		Annotations:           MakeKeyValuePairs(),
		ParentUUID:            "",
		FQName:                []string{},
		DatabaseNodeIPAddress: MakeIpAddressType(),
		Perms2:                MakePermType2(),
		UUID:                  "",
		ParentType:            "",
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
	}
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}
