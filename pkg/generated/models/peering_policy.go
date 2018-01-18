package models

// PeeringPolicy

import "encoding/json"

// PeeringPolicy
type PeeringPolicy struct {
	FQName         []string           `json:"fq_name,omitempty"`
	IDPerms        *IdPermsType       `json:"id_perms,omitempty"`
	Annotations    *KeyValuePairs     `json:"annotations,omitempty"`
	Perms2         *PermType2         `json:"perms2,omitempty"`
	PeeringService PeeringServiceType `json:"peering_service,omitempty"`
	DisplayName    string             `json:"display_name,omitempty"`
	UUID           string             `json:"uuid,omitempty"`
	ParentUUID     string             `json:"parent_uuid,omitempty"`
	ParentType     string             `json:"parent_type,omitempty"`
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
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		UUID:           "",
		ParentUUID:     "",
		ParentType:     "",
		PeeringService: MakePeeringServiceType(),
		DisplayName:    "",
	}
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}
