package models

// NetworkPolicy

import "encoding/json"

// NetworkPolicy
type NetworkPolicy struct {
	ParentUUID           string             `json:"parent_uuid,omitempty"`
	Annotations          *KeyValuePairs     `json:"annotations,omitempty"`
	Perms2               *PermType2         `json:"perms2,omitempty"`
	UUID                 string             `json:"uuid,omitempty"`
	IDPerms              *IdPermsType       `json:"id_perms,omitempty"`
	DisplayName          string             `json:"display_name,omitempty"`
	NetworkPolicyEntries *PolicyEntriesType `json:"network_policy_entries,omitempty"`
	ParentType           string             `json:"parent_type,omitempty"`
	FQName               []string           `json:"fq_name,omitempty"`
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
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		NetworkPolicyEntries: MakePolicyEntriesType(),
		Perms2:               MakePermType2(),
		UUID:                 "",
		ParentUUID:           "",
		Annotations:          MakeKeyValuePairs(),
	}
}

// MakeNetworkPolicySlice() makes a slice of NetworkPolicy
func MakeNetworkPolicySlice() []*NetworkPolicy {
	return []*NetworkPolicy{}
}
