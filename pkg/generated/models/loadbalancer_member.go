package models

// LoadbalancerMember

import "encoding/json"

// LoadbalancerMember
type LoadbalancerMember struct {
	ParentUUID                   string                  `json:"parent_uuid,omitempty"`
	IDPerms                      *IdPermsType            `json:"id_perms,omitempty"`
	DisplayName                  string                  `json:"display_name,omitempty"`
	Perms2                       *PermType2              `json:"perms2,omitempty"`
	LoadbalancerMemberProperties *LoadbalancerMemberType `json:"loadbalancer_member_properties,omitempty"`
	UUID                         string                  `json:"uuid,omitempty"`
	ParentType                   string                  `json:"parent_type,omitempty"`
	FQName                       []string                `json:"fq_name,omitempty"`
	Annotations                  *KeyValuePairs          `json:"annotations,omitempty"`
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
		Perms2:                       MakePermType2(),
		ParentUUID:                   "",
		IDPerms:                      MakeIdPermsType(),
		DisplayName:                  "",
		FQName:                       []string{},
		Annotations:                  MakeKeyValuePairs(),
		LoadbalancerMemberProperties: MakeLoadbalancerMemberType(),
		UUID:       "",
		ParentType: "",
	}
}

// MakeLoadbalancerMemberSlice() makes a slice of LoadbalancerMember
func MakeLoadbalancerMemberSlice() []*LoadbalancerMember {
	return []*LoadbalancerMember{}
}
