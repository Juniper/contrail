package models

// LoadbalancerMember

import "encoding/json"

// LoadbalancerMember
type LoadbalancerMember struct {
	UUID                         string                  `json:"uuid,omitempty"`
	ParentUUID                   string                  `json:"parent_uuid,omitempty"`
	ParentType                   string                  `json:"parent_type,omitempty"`
	Perms2                       *PermType2              `json:"perms2,omitempty"`
	FQName                       []string                `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType            `json:"id_perms,omitempty"`
	DisplayName                  string                  `json:"display_name,omitempty"`
	Annotations                  *KeyValuePairs          `json:"annotations,omitempty"`
	LoadbalancerMemberProperties *LoadbalancerMemberType `json:"loadbalancer_member_properties,omitempty"`
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
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		LoadbalancerMemberProperties: MakeLoadbalancerMemberType(),
		FQName:     []string{},
		IDPerms:    MakeIdPermsType(),
		ParentType: "",
		Perms2:     MakePermType2(),
		UUID:       "",
		ParentUUID: "",
	}
}

// MakeLoadbalancerMemberSlice() makes a slice of LoadbalancerMember
func MakeLoadbalancerMemberSlice() []*LoadbalancerMember {
	return []*LoadbalancerMember{}
}
