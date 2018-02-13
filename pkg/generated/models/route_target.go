package models

// RouteTarget

import "encoding/json"

// RouteTarget
//proteus:generate
type RouteTarget struct {
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
}

// String returns json representation of the object
func (model *RouteTarget) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteTarget makes RouteTarget
func MakeRouteTarget() *RouteTarget {
	return &RouteTarget{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeRouteTargetSlice() makes a slice of RouteTarget
func MakeRouteTargetSlice() []*RouteTarget {
	return []*RouteTarget{}
}
