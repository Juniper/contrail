package models

// PeeringPolicy

import "encoding/json"

// PeeringPolicy
type PeeringPolicy struct {
	Annotations    *KeyValuePairs     `json:"annotations,omitempty"`
	UUID           string             `json:"uuid,omitempty"`
	ParentType     string             `json:"parent_type,omitempty"`
	PeeringService PeeringServiceType `json:"peering_service,omitempty"`
	ParentUUID     string             `json:"parent_uuid,omitempty"`
	FQName         []string           `json:"fq_name,omitempty"`
	IDPerms        *IdPermsType       `json:"id_perms,omitempty"`
	DisplayName    string             `json:"display_name,omitempty"`
	Perms2         *PermType2         `json:"perms2,omitempty"`
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
		UUID:           "",
		ParentType:     "",
		PeeringService: MakePeeringServiceType(),
		Annotations:    MakeKeyValuePairs(),
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Perms2:         MakePermType2(),
		ParentUUID:     "",
	}
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}
