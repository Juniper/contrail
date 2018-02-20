package models

// DatabaseNode

// DatabaseNode
//proteus:generate
type DatabaseNode struct {
	UUID                  string         `json:"uuid,omitempty"`
	ParentUUID            string         `json:"parent_uuid,omitempty"`
	ParentType            string         `json:"parent_type,omitempty"`
	FQName                []string       `json:"fq_name,omitempty"`
	IDPerms               *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName           string         `json:"display_name,omitempty"`
	Annotations           *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                *PermType2     `json:"perms2,omitempty"`
	DatabaseNodeIPAddress IpAddressType  `json:"database_node_ip_address,omitempty"`
}

// MakeDatabaseNode makes DatabaseNode
func MakeDatabaseNode() *DatabaseNode {
	return &DatabaseNode{
		//TODO(nati): Apply default
		UUID:                  "",
		ParentUUID:            "",
		ParentType:            "",
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		DatabaseNodeIPAddress: MakeIpAddressType(),
	}
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
	return []*DatabaseNode{}
}
