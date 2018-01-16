package models

// NetworkPolicy

import "encoding/json"

// NetworkPolicy
type NetworkPolicy struct {
	Perms2               *PermType2         `json:"perms2,omitempty"`
	FQName               []string           `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType       `json:"id_perms,omitempty"`
	DisplayName          string             `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs     `json:"annotations,omitempty"`
	UUID                 string             `json:"uuid,omitempty"`
	ParentUUID           string             `json:"parent_uuid,omitempty"`
	ParentType           string             `json:"parent_type,omitempty"`
	NetworkPolicyEntries *PolicyEntriesType `json:"network_policy_entries,omitempty"`
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
		IDPerms:              MakeIdPermsType(),
		Perms2:               MakePermType2(),
		FQName:               []string{},
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		NetworkPolicyEntries: MakePolicyEntriesType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
	}
}

// MakeNetworkPolicySlice() makes a slice of NetworkPolicy
func MakeNetworkPolicySlice() []*NetworkPolicy {
	return []*NetworkPolicy{}
}
