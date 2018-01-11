package models

// NetworkPolicy

import "encoding/json"

// NetworkPolicy
type NetworkPolicy struct {
	Annotations          *KeyValuePairs     `json:"annotations"`
	Perms2               *PermType2         `json:"perms2"`
	UUID                 string             `json:"uuid"`
	ParentType           string             `json:"parent_type"`
	FQName               []string           `json:"fq_name"`
	IDPerms              *IdPermsType       `json:"id_perms"`
	DisplayName          string             `json:"display_name"`
	NetworkPolicyEntries *PolicyEntriesType `json:"network_policy_entries"`
	ParentUUID           string             `json:"parent_uuid"`
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
		ParentUUID:           "",
		NetworkPolicyEntries: MakePolicyEntriesType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		UUID:                 "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
	}
}

// MakeNetworkPolicySlice() makes a slice of NetworkPolicy
func MakeNetworkPolicySlice() []*NetworkPolicy {
	return []*NetworkPolicy{}
}
