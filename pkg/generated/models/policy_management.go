package models

// PolicyManagement

import "encoding/json"

// PolicyManagement
type PolicyManagement struct {
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`

	AddressGroups         []*AddressGroup         `json:"address_groups"`
	ApplicationPolicySets []*ApplicationPolicySet `json:"application_policy_sets"`
	FirewallPolicys       []*FirewallPolicy       `json:"firewall_policys"`
	FirewallRules         []*FirewallRule         `json:"firewall_rules"`
	ServiceGroups         []*ServiceGroup         `json:"service_groups"`
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
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
	}
}

// MakePolicyManagementSlice() makes a slice of PolicyManagement
func MakePolicyManagementSlice() []*PolicyManagement {
	return []*PolicyManagement{}
}
