package models

// ApplicationPolicySet

import "encoding/json"

// ApplicationPolicySet
type ApplicationPolicySet struct {
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	UUID            string         `json:"uuid,omitempty"`
	AllApplications bool           `json:"all_applications,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`

	FirewallPolicyRefs      []*ApplicationPolicySetFirewallPolicyRef      `json:"firewall_policy_refs,omitempty"`
	GlobalVrouterConfigRefs []*ApplicationPolicySetGlobalVrouterConfigRef `json:"global_vrouter_config_refs,omitempty"`
}

// ApplicationPolicySetFirewallPolicyRef references each other
type ApplicationPolicySetFirewallPolicyRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *FirewallSequence
}

// ApplicationPolicySetGlobalVrouterConfigRef references each other
type ApplicationPolicySetGlobalVrouterConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *ApplicationPolicySet) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeApplicationPolicySet makes ApplicationPolicySet
func MakeApplicationPolicySet() *ApplicationPolicySet {
	return &ApplicationPolicySet{
		//TODO(nati): Apply default
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		IDPerms:         MakeIdPermsType(),
		Perms2:          MakePermType2(),
		UUID:            "",
		AllApplications: false,
		ParentUUID:      "",
		ParentType:      "",
		FQName:          []string{},
	}
}

// MakeApplicationPolicySetSlice() makes a slice of ApplicationPolicySet
func MakeApplicationPolicySetSlice() []*ApplicationPolicySet {
	return []*ApplicationPolicySet{}
}
