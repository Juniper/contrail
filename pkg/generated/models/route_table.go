package models

// RouteTable

import "encoding/json"

// RouteTable
type RouteTable struct {
	ParentType  string          `json:"parent_type"`
	IDPerms     *IdPermsType    `json:"id_perms"`
	DisplayName string          `json:"display_name"`
	Annotations *KeyValuePairs  `json:"annotations"`
	Routes      *RouteTableType `json:"routes"`
	ParentUUID  string          `json:"parent_uuid"`
	Perms2      *PermType2      `json:"perms2"`
	UUID        string          `json:"uuid"`
	FQName      []string        `json:"fq_name"`
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
		ParentUUID:  "",
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		FQName:      []string{},
		Perms2:      MakePermType2(),
	}
}

// MakeRouteTableSlice() makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
	return []*RouteTable{}
}
