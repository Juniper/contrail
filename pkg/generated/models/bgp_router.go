package models

// BGPRouter

import "encoding/json"

// BGPRouter
type BGPRouter struct {
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
}

// String returns json representation of the object
func (model *BGPRouter) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBGPRouter makes BGPRouter
func MakeBGPRouter() *BGPRouter {
	return &BGPRouter{
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

// MakeBGPRouterSlice() makes a slice of BGPRouter
func MakeBGPRouterSlice() []*BGPRouter {
	return []*BGPRouter{}
}
