package models

// DatabaseNode

import "encoding/json"

// DatabaseNode
type DatabaseNode struct {
	DatabaseNodeIPAddress IpAddressType  `json:"database_node_ip_address"`
	ParentUUID            string         `json:"parent_uuid"`
	FQName                []string       `json:"fq_name"`
	IDPerms               *IdPermsType   `json:"id_perms"`
	Annotations           *KeyValuePairs `json:"annotations"`
	ParentType            string         `json:"parent_type"`
	DisplayName           string         `json:"display_name"`
	Perms2                *PermType2     `json:"perms2"`
	UUID                  string         `json:"uuid"`
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
		ParentType:            "",
		DisplayName:           "",
		Perms2:                MakePermType2(),
		UUID:                  "",
		Annotations:           MakeKeyValuePairs(),
		DatabaseNodeIPAddress: MakeIpAddressType(),
		ParentUUID:            "",
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
	}
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}
