package models

// RouteTable

import "encoding/json"

// RouteTable
type RouteTable struct {
	ParentUUID  string          `json:"parent_uuid,omitempty"`
	ParentType  string          `json:"parent_type,omitempty"`
	IDPerms     *IdPermsType    `json:"id_perms,omitempty"`
	DisplayName string          `json:"display_name,omitempty"`
	Annotations *KeyValuePairs  `json:"annotations,omitempty"`
	UUID        string          `json:"uuid,omitempty"`
	Routes      *RouteTableType `json:"routes,omitempty"`
	FQName      []string        `json:"fq_name,omitempty"`
	Perms2      *PermType2      `json:"perms2,omitempty"`
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
		Routes:      MakeRouteTableType(),
		FQName:      []string{},
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
	}
}

// MakeRouteTableSlice() makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
	return []*RouteTable{}
}
