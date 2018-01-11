package models

// DiscoveryServiceAssignment

import "encoding/json"

// DiscoveryServiceAssignment
type DiscoveryServiceAssignment struct {
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`

	DsaRules []*DsaRule `json:"dsa_rules"`
}

// String returns json representation of the object
func (model *DiscoveryServiceAssignment) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDiscoveryServiceAssignment makes DiscoveryServiceAssignment
func MakeDiscoveryServiceAssignment() *DiscoveryServiceAssignment {
	return &DiscoveryServiceAssignment{
		//TODO(nati): Apply default
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
	}
}

// MakeDiscoveryServiceAssignmentSlice() makes a slice of DiscoveryServiceAssignment
func MakeDiscoveryServiceAssignmentSlice() []*DiscoveryServiceAssignment {
	return []*DiscoveryServiceAssignment{}
}
