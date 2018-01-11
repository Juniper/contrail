package models

// LoadbalancerMember

import "encoding/json"

// LoadbalancerMember
type LoadbalancerMember struct {
	DisplayName                  string                  `json:"display_name"`
	Annotations                  *KeyValuePairs          `json:"annotations"`
	ParentUUID                   string                  `json:"parent_uuid"`
	UUID                         string                  `json:"uuid"`
	ParentType                   string                  `json:"parent_type"`
	LoadbalancerMemberProperties *LoadbalancerMemberType `json:"loadbalancer_member_properties"`
	FQName                       []string                `json:"fq_name"`
	IDPerms                      *IdPermsType            `json:"id_perms"`
	Perms2                       *PermType2              `json:"perms2"`
}

// String returns json representation of the object
func (model *LoadbalancerMember) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerMember makes LoadbalancerMember
func MakeLoadbalancerMember() *LoadbalancerMember {
	return &LoadbalancerMember{
		//TODO(nati): Apply default
		ParentUUID:                   "",
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		IDPerms:                      MakeIdPermsType(),
		Perms2:                       MakePermType2(),
		UUID:                         "",
		ParentType:                   "",
		LoadbalancerMemberProperties: MakeLoadbalancerMemberType(),
		FQName: []string{},
	}
}

// MakeLoadbalancerMemberSlice() makes a slice of LoadbalancerMember
func MakeLoadbalancerMemberSlice() []*LoadbalancerMember {
	return []*LoadbalancerMember{}
}
