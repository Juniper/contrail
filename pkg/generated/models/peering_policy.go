package models

// PeeringPolicy

import "encoding/json"

// PeeringPolicy
type PeeringPolicy struct {
	UUID           string             `json:"uuid"`
	ParentUUID     string             `json:"parent_uuid"`
	ParentType     string             `json:"parent_type"`
	FQName         []string           `json:"fq_name"`
	PeeringService PeeringServiceType `json:"peering_service"`
	IDPerms        *IdPermsType       `json:"id_perms"`
	DisplayName    string             `json:"display_name"`
	Annotations    *KeyValuePairs     `json:"annotations"`
	Perms2         *PermType2         `json:"perms2"`
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
		ParentUUID:     "",
		ParentType:     "",
		FQName:         []string{},
		PeeringService: MakePeeringServiceType(),
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
	}
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}
