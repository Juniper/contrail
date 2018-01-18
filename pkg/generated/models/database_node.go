package models

// DatabaseNode

import "encoding/json"

// DatabaseNode
type DatabaseNode struct {
	Perms2                *PermType2     `json:"perms2,omitempty"`
	UUID                  string         `json:"uuid,omitempty"`
	ParentType            string         `json:"parent_type,omitempty"`
	FQName                []string       `json:"fq_name,omitempty"`
	DisplayName           string         `json:"display_name,omitempty"`
	Annotations           *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID            string         `json:"parent_uuid,omitempty"`
	IDPerms               *IdPermsType   `json:"id_perms,omitempty"`
	DatabaseNodeIPAddress IpAddressType  `json:"database_node_ip_address,omitempty"`
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
		ParentType:            "",
		FQName:                []string{},
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		IDPerms:               MakeIdPermsType(),
		DatabaseNodeIPAddress: MakeIpAddressType(),
		ParentUUID:            "",
	}
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}
