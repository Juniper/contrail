package models

// PolicyManagement

import "encoding/json"

// PolicyManagement
type PolicyManagement struct {
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`

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

// InterfaceToPolicyManagement makes PolicyManagement from interface
func InterfaceToPolicyManagement(iData interface{}) *PolicyManagement {
	data := iData.(map[string]interface{})
	return &PolicyManagement{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToPolicyManagementSlice makes a slice of PolicyManagement from interface
func InterfaceToPolicyManagementSlice(data interface{}) []*PolicyManagement {
	list := data.([]interface{})
	result := MakePolicyManagementSlice()
	for _, item := range list {
		result = append(result, InterfaceToPolicyManagement(item))
	}
	return result
}

// MakePolicyManagementSlice() makes a slice of PolicyManagement
func MakePolicyManagementSlice() []*PolicyManagement {
	return []*PolicyManagement{}
}
