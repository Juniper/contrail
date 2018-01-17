package models

// NetworkPolicy

import "encoding/json"

// NetworkPolicy
type NetworkPolicy struct {
	NetworkPolicyEntries *PolicyEntriesType `json:"network_policy_entries,omitempty"`
	ParentUUID           string             `json:"parent_uuid,omitempty"`
	ParentType           string             `json:"parent_type,omitempty"`
	FQName               []string           `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType       `json:"id_perms,omitempty"`
	Annotations          *KeyValuePairs     `json:"annotations,omitempty"`
	Perms2               *PermType2         `json:"perms2,omitempty"`
	DisplayName          string             `json:"display_name,omitempty"`
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
		IDPerms:              MakeIdPermsType(),
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		NetworkPolicyEntries: MakePolicyEntriesType(),
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		DisplayName:          "",
		UUID:                 "",
	}
}

// MakeNetworkPolicySlice() makes a slice of NetworkPolicy
func MakeNetworkPolicySlice() []*NetworkPolicy {
	return []*NetworkPolicy{}
}
