package models

// NetworkPolicy

import "encoding/json"

// NetworkPolicy
type NetworkPolicy struct {
	NetworkPolicyEntries *PolicyEntriesType `json:"network_policy_entries,omitempty"`
	FQName               []string           `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType       `json:"id_perms,omitempty"`
	ParentUUID           string             `json:"parent_uuid,omitempty"`
	ParentType           string             `json:"parent_type,omitempty"`
	DisplayName          string             `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs     `json:"annotations,omitempty"`
	Perms2               *PermType2         `json:"perms2,omitempty"`
	UUID                 string             `json:"uuid,omitempty"`
}

// String returns json representation of the object
func (model *NetworkPolicy) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeNetworkPolicy makes NetworkPolicy
func MakeNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{
		//TODO(nati): Apply default
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		DisplayName:          "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		NetworkPolicyEntries: MakePolicyEntriesType(),
	}
}

// MakeNetworkPolicySlice() makes a slice of NetworkPolicy
func MakeNetworkPolicySlice() []*NetworkPolicy {
	return []*NetworkPolicy{}
}
