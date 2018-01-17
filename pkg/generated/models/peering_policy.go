package models

// PeeringPolicy

import "encoding/json"

// PeeringPolicy
type PeeringPolicy struct {
	ParentType     string             `json:"parent_type,omitempty"`
	DisplayName    string             `json:"display_name,omitempty"`
	PeeringService PeeringServiceType `json:"peering_service,omitempty"`
	FQName         []string           `json:"fq_name,omitempty"`
	IDPerms        *IdPermsType       `json:"id_perms,omitempty"`
	Annotations    *KeyValuePairs     `json:"annotations,omitempty"`
	Perms2         *PermType2         `json:"perms2,omitempty"`
	UUID           string             `json:"uuid,omitempty"`
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
		ParentType:     "",
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		UUID:           "",
		ParentUUID:     "",
		PeeringService: MakePeeringServiceType(),
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
	}
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}
