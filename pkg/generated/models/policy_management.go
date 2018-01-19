package models

// PolicyManagement

import "encoding/json"

// PolicyManagement
type PolicyManagement struct {
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`

	AddressGroups         []*AddressGroup         `json:"address_groups,omitempty"`
	ApplicationPolicySets []*ApplicationPolicySet `json:"application_policy_sets,omitempty"`
	FirewallPolicys       []*FirewallPolicy       `json:"firewall_policys,omitempty"`
	FirewallRules         []*FirewallRule         `json:"firewall_rules,omitempty"`
	ServiceGroups         []*ServiceGroup         `json:"service_groups,omitempty"`
}

// String returns json representation of the object
func (model *PolicyManagement) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePolicyManagement makes PolicyManagement
func MakePolicyManagement() *PolicyManagement {
	return &PolicyManagement{
		//TODO(nati): Apply default
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
	}
}

// MakePolicyManagementSlice() makes a slice of PolicyManagement
func MakePolicyManagementSlice() []*PolicyManagement {
	return []*PolicyManagement{}
}
