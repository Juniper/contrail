package models

// FirewallPolicy

import "encoding/json"

// FirewallPolicy
type FirewallPolicy struct {
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`

	FirewallRuleRefs          []*FirewallPolicyFirewallRuleRef          `json:"firewall_rule_refs"`
	SecurityLoggingObjectRefs []*FirewallPolicySecurityLoggingObjectRef `json:"security_logging_object_refs"`
}

// FirewallPolicySecurityLoggingObjectRef references each other
type FirewallPolicySecurityLoggingObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// FirewallPolicyFirewallRuleRef references each other
type FirewallPolicyFirewallRuleRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *FirewallSequence
}

// String returns json representation of the object
func (model *FirewallPolicy) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallPolicy makes FirewallPolicy
func MakeFirewallPolicy() *FirewallPolicy {
	return &FirewallPolicy{
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

// MakeFirewallPolicySlice() makes a slice of FirewallPolicy
func MakeFirewallPolicySlice() []*FirewallPolicy {
	return []*FirewallPolicy{}
}
