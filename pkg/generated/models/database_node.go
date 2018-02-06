package models

// DatabaseNode

import "encoding/json"

// DatabaseNode
type DatabaseNode struct {
	UUID                  string         `json:"uuid,omitempty"`
	ParentUUID            string         `json:"parent_uuid,omitempty"`
	FQName                []string       `json:"fq_name,omitempty"`
	IDPerms               *IdPermsType   `json:"id_perms,omitempty"`
	DatabaseNodeIPAddress IpAddressType  `json:"database_node_ip_address,omitempty"`
	ParentType            string         `json:"parent_type,omitempty"`
	DisplayName           string         `json:"display_name,omitempty"`
	Annotations           *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                *PermType2     `json:"perms2,omitempty"`
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
		DatabaseNodeIPAddress: MakeIpAddressType(),
		UUID:        "",
		ParentUUID:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		Perms2:      MakePermType2(),
		ParentType:  "",
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}
