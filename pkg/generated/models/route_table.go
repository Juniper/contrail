package models

// RouteTable

import "encoding/json"

// RouteTable
type RouteTable struct {
	Routes      *RouteTableType `json:"routes,omitempty"`
	Annotations *KeyValuePairs  `json:"annotations,omitempty"`
	Perms2      *PermType2      `json:"perms2,omitempty"`
	ParentUUID  string          `json:"parent_uuid,omitempty"`
	FQName      []string        `json:"fq_name,omitempty"`
	DisplayName string          `json:"display_name,omitempty"`
	UUID        string          `json:"uuid,omitempty"`
	ParentType  string          `json:"parent_type,omitempty"`
	IDPerms     *IdPermsType    `json:"id_perms,omitempty"`
}

// String returns json representation of the object
func (model *RouteTable) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteTable makes RouteTable
func MakeRouteTable() *RouteTable {
	return &RouteTable{
		//TODO(nati): Apply default
		DisplayName: "",
		UUID:        "",
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
		Routes:      MakeRouteTableType(),
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		FQName:      []string{},
	}
}

// MakeRouteTableSlice() makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
	return []*RouteTable{}
}
