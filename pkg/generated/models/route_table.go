package models

// RouteTable

import "encoding/json"

// RouteTable
type RouteTable struct {
	ParentType  string          `json:"parent_type,omitempty"`
	FQName      []string        `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType    `json:"id_perms,omitempty"`
	Annotations *KeyValuePairs  `json:"annotations,omitempty"`
	Perms2      *PermType2      `json:"perms2,omitempty"`
	Routes      *RouteTableType `json:"routes,omitempty"`
	UUID        string          `json:"uuid,omitempty"`
	ParentUUID  string          `json:"parent_uuid,omitempty"`
	DisplayName string          `json:"display_name,omitempty"`
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
		UUID:        "",
		ParentUUID:  "",
		DisplayName: "",
		Perms2:      MakePermType2(),
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeRouteTableSlice() makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
	return []*RouteTable{}
}
