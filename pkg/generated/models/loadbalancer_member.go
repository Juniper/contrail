package models

// LoadbalancerMember

import "encoding/json"

// LoadbalancerMember
type LoadbalancerMember struct {
	ParentType                   string                  `json:"parent_type,omitempty"`
	Annotations                  *KeyValuePairs          `json:"annotations,omitempty"`
	LoadbalancerMemberProperties *LoadbalancerMemberType `json:"loadbalancer_member_properties,omitempty"`
	ParentUUID                   string                  `json:"parent_uuid,omitempty"`
	FQName                       []string                `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType            `json:"id_perms,omitempty"`
	DisplayName                  string                  `json:"display_name,omitempty"`
	Perms2                       *PermType2              `json:"perms2,omitempty"`
	UUID                         string                  `json:"uuid,omitempty"`
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
		ParentType:  "",
		Annotations: MakeKeyValuePairs(),
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Perms2:      MakePermType2(),
		UUID:        "",
		LoadbalancerMemberProperties: MakeLoadbalancerMemberType(),
		ParentUUID:                   "",
		FQName:                       []string{},
	}
}

// MakeLoadbalancerMemberSlice() makes a slice of LoadbalancerMember
func MakeLoadbalancerMemberSlice() []*LoadbalancerMember {
	return []*LoadbalancerMember{}
}
