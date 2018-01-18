package models

// RoutingInstance

import "encoding/json"

// RoutingInstance
type RoutingInstance struct {
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
}

// String returns json representation of the object
func (model *RoutingInstance) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRoutingInstance makes RoutingInstance
func MakeRoutingInstance() *RoutingInstance {
	return &RoutingInstance{
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

// MakeRoutingInstanceSlice() makes a slice of RoutingInstance
func MakeRoutingInstanceSlice() []*RoutingInstance {
	return []*RoutingInstance{}
}
