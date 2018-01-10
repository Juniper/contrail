package models

// ApplicationPolicySet

import "encoding/json"

// ApplicationPolicySet
type ApplicationPolicySet struct {
	FQName          []string       `json:"fq_name"`
	IDPerms         *IdPermsType   `json:"id_perms"`
	AllApplications bool           `json:"all_applications"`
	ParentUUID      string         `json:"parent_uuid"`
	ParentType      string         `json:"parent_type"`
	DisplayName     string         `json:"display_name"`
	Annotations     *KeyValuePairs `json:"annotations"`
	Perms2          *PermType2     `json:"perms2"`
	UUID            string         `json:"uuid"`

	FirewallPolicyRefs      []*ApplicationPolicySetFirewallPolicyRef      `json:"firewall_policy_refs"`
	GlobalVrouterConfigRefs []*ApplicationPolicySetGlobalVrouterConfigRef `json:"global_vrouter_config_refs"`
}

// ApplicationPolicySetGlobalVrouterConfigRef references each other
type ApplicationPolicySetGlobalVrouterConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// ApplicationPolicySetFirewallPolicyRef references each other
type ApplicationPolicySetFirewallPolicyRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *FirewallSequence
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
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		UUID:            "",
		AllApplications: false,
		ParentUUID:      "",
		ParentType:      "",
		DisplayName:     "",
	}
}

// InterfaceToApplicationPolicySet makes ApplicationPolicySet from interface
func InterfaceToApplicationPolicySet(iData interface{}) *ApplicationPolicySet {
	data := iData.(map[string]interface{})
	return &ApplicationPolicySet{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		AllApplications: data["all_applications"].(bool),

		//{"description":"If set, indicates application policy set to be applied to all application tags","default":false,"type":"boolean"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}

	}
}

// InterfaceToApplicationPolicySetSlice makes a slice of ApplicationPolicySet from interface
func InterfaceToApplicationPolicySetSlice(data interface{}) []*ApplicationPolicySet {
	list := data.([]interface{})
	result := MakeApplicationPolicySetSlice()
	for _, item := range list {
		result = append(result, InterfaceToApplicationPolicySet(item))
	}
	return result
}

// MakeApplicationPolicySetSlice() makes a slice of ApplicationPolicySet
func MakeApplicationPolicySetSlice() []*ApplicationPolicySet {
	return []*ApplicationPolicySet{}
}
