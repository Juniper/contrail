package models

// DatabaseNode

import "encoding/json"

// DatabaseNode
type DatabaseNode struct {
	ParentType            string         `json:"parent_type,omitempty"`
	FQName                []string       `json:"fq_name,omitempty"`
	IDPerms               *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName           string         `json:"display_name,omitempty"`
	Perms2                *PermType2     `json:"perms2,omitempty"`
	ParentUUID            string         `json:"parent_uuid,omitempty"`
	DatabaseNodeIPAddress IpAddressType  `json:"database_node_ip_address,omitempty"`
	Annotations           *KeyValuePairs `json:"annotations,omitempty"`
	UUID                  string         `json:"uuid,omitempty"`
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
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		UUID:        "",
		DatabaseNodeIPAddress: MakeIpAddressType(),
		Annotations:           MakeKeyValuePairs(),
	}
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}
