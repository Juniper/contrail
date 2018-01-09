package models

// ApplicationPolicySet

import "encoding/json"

// ApplicationPolicySet
type ApplicationPolicySet struct {
	FQName          []string       `json:"fq_name"`
	IDPerms         *IdPermsType   `json:"id_perms"`
	UUID            string         `json:"uuid"`
	ParentUUID      string         `json:"parent_uuid"`
	ParentType      string         `json:"parent_type"`
	AllApplications bool           `json:"all_applications"`
	DisplayName     string         `json:"display_name"`
	Annotations     *KeyValuePairs `json:"annotations"`
	Perms2          *PermType2     `json:"perms2"`

	FirewallPolicyRefs      []*ApplicationPolicySetFirewallPolicyRef      `json:"firewall_policy_refs"`
	GlobalVrouterConfigRefs []*ApplicationPolicySetGlobalVrouterConfigRef `json:"global_vrouter_config_refs"`
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
		AllApplications: false,
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		ParentType:      "",
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		UUID:            "",
		ParentUUID:      "",
	}
}

// InterfaceToApplicationPolicySet makes ApplicationPolicySet from interface
func InterfaceToApplicationPolicySet(iData interface{}) *ApplicationPolicySet {
	data := iData.(map[string]interface{})
	return &ApplicationPolicySet{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		AllApplications: data["all_applications"].(bool),

		//{"description":"If set, indicates application policy set to be applied to all application tags","default":false,"type":"boolean"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

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
