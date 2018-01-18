package models

// DatabaseNode

import "encoding/json"

// DatabaseNode
type DatabaseNode struct {
	IDPerms               *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName           string         `json:"display_name,omitempty"`
	UUID                  string         `json:"uuid,omitempty"`
	ParentUUID            string         `json:"parent_uuid,omitempty"`
	DatabaseNodeIPAddress IpAddressType  `json:"database_node_ip_address,omitempty"`
	FQName                []string       `json:"fq_name,omitempty"`
	Perms2                *PermType2     `json:"perms2,omitempty"`
	ParentType            string         `json:"parent_type,omitempty"`
	Annotations           *KeyValuePairs `json:"annotations,omitempty"`
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
		UUID:                  "",
		ParentUUID:            "",
		DatabaseNodeIPAddress: MakeIpAddressType(),
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		ParentType:            "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
	}
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}
