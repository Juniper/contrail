package models

// RoutingInstance

import "encoding/json"

// RoutingInstance
type RoutingInstance struct {
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
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
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeRoutingInstanceSlice() makes a slice of RoutingInstance
func MakeRoutingInstanceSlice() []*RoutingInstance {
	return []*RoutingInstance{}
}
