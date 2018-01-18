package models

// RouteTable

import "encoding/json"

// RouteTable
type RouteTable struct {
	Perms2      *PermType2      `json:"perms2,omitempty"`
	UUID        string          `json:"uuid,omitempty"`
	ParentUUID  string          `json:"parent_uuid,omitempty"`
	FQName      []string        `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType    `json:"id_perms,omitempty"`
	DisplayName string          `json:"display_name,omitempty"`
	Routes      *RouteTableType `json:"routes,omitempty"`
	ParentType  string          `json:"parent_type,omitempty"`
	Annotations *KeyValuePairs  `json:"annotations,omitempty"`
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
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Routes:      MakeRouteTableType(),
		ParentType:  "",
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeRouteTableSlice() makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
	return []*RouteTable{}
}
