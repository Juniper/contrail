package models

// PeeringPolicy

import "encoding/json"

// PeeringPolicy
type PeeringPolicy struct {
	UUID           string             `json:"uuid,omitempty"`
	IDPerms        *IdPermsType       `json:"id_perms,omitempty"`
	DisplayName    string             `json:"display_name,omitempty"`
	PeeringService PeeringServiceType `json:"peering_service,omitempty"`
	Perms2         *PermType2         `json:"perms2,omitempty"`
	ParentType     string             `json:"parent_type,omitempty"`
	FQName         []string           `json:"fq_name,omitempty"`
	Annotations    *KeyValuePairs     `json:"annotations,omitempty"`
	ParentUUID     string             `json:"parent_uuid,omitempty"`
}

// String returns json representation of the object
func (model *PeeringPolicy) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePeeringPolicy makes PeeringPolicy
func MakePeeringPolicy() *PeeringPolicy {
	return &PeeringPolicy{
		//TODO(nati): Apply default
		Perms2:         MakePermType2(),
		UUID:           "",
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		PeeringService: MakePeeringServiceType(),
		ParentUUID:     "",
		ParentType:     "",
		FQName:         []string{},
		Annotations:    MakeKeyValuePairs(),
	}
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}
