package models

// PeeringPolicy

import "encoding/json"

// PeeringPolicy
type PeeringPolicy struct {
	PeeringService PeeringServiceType `json:"peering_service,omitempty"`
	ParentType     string             `json:"parent_type,omitempty"`
	FQName         []string           `json:"fq_name,omitempty"`
	ParentUUID     string             `json:"parent_uuid,omitempty"`
	IDPerms        *IdPermsType       `json:"id_perms,omitempty"`
	DisplayName    string             `json:"display_name,omitempty"`
	Annotations    *KeyValuePairs     `json:"annotations,omitempty"`
	Perms2         *PermType2         `json:"perms2,omitempty"`
	UUID           string             `json:"uuid,omitempty"`
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
		PeeringService: MakePeeringServiceType(),
		ParentType:     "",
		FQName:         []string{},
		ParentUUID:     "",
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		UUID:           "",
	}
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}
