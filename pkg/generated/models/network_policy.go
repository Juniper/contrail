package models

// NetworkPolicy

import "encoding/json"

// NetworkPolicy
type NetworkPolicy struct {
	Perms2               *PermType2         `json:"perms2,omitempty"`
	UUID                 string             `json:"uuid,omitempty"`
	ParentUUID           string             `json:"parent_uuid,omitempty"`
	ParentType           string             `json:"parent_type,omitempty"`
	Annotations          *KeyValuePairs     `json:"annotations,omitempty"`
	NetworkPolicyEntries *PolicyEntriesType `json:"network_policy_entries,omitempty"`
	FQName               []string           `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType       `json:"id_perms,omitempty"`
	DisplayName          string             `json:"display_name,omitempty"`
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
		Perms2:               MakePermType2(),
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		NetworkPolicyEntries: MakePolicyEntriesType(),
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
	}
}

// MakeNetworkPolicySlice() makes a slice of NetworkPolicy
func MakeNetworkPolicySlice() []*NetworkPolicy {
	return []*NetworkPolicy{}
}
