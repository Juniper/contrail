package models

// SecurityLoggingObject

import "encoding/json"

// SecurityLoggingObject
type SecurityLoggingObject struct {
	SecurityLoggingObjectRules *SecurityLoggingObjectRuleListType `json:"security_logging_object_rules"`
	SecurityLoggingObjectRate  int                                `json:"security_logging_object_rate"`
	FQName                     []string                           `json:"fq_name"`
	DisplayName                string                             `json:"display_name"`
	Annotations                *KeyValuePairs                     `json:"annotations"`
	UUID                       string                             `json:"uuid"`
	ParentUUID                 string                             `json:"parent_uuid"`
	ParentType                 string                             `json:"parent_type"`
	IDPerms                    *IdPermsType                       `json:"id_perms"`
	Perms2                     *PermType2                         `json:"perms2"`

	SecurityGroupRefs []*SecurityLoggingObjectSecurityGroupRef `json:"security_group_refs"`
	NetworkPolicyRefs []*SecurityLoggingObjectNetworkPolicyRef `json:"network_policy_refs"`
}

// SecurityLoggingObjectSecurityGroupRef references each other
type SecurityLoggingObjectSecurityGroupRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *SecurityLoggingObjectRuleListType
}

// SecurityLoggingObjectNetworkPolicyRef references each other
type SecurityLoggingObjectNetworkPolicyRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *SecurityLoggingObjectRuleListType
}

// String returns json representation of the object
func (model *SecurityLoggingObject) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSecurityLoggingObject makes SecurityLoggingObject
func MakeSecurityLoggingObject() *SecurityLoggingObject {
	return &SecurityLoggingObject{
		//TODO(nati): Apply default
		ParentType: "",
		IDPerms:    MakeIdPermsType(),
		Perms2:     MakePermType2(),
		SecurityLoggingObjectRules: MakeSecurityLoggingObjectRuleListType(),
		SecurityLoggingObjectRate:  0,
		FQName:      []string{},
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ParentUUID:  "",
	}
}

// InterfaceToSecurityLoggingObject makes SecurityLoggingObject from interface
func InterfaceToSecurityLoggingObject(iData interface{}) *SecurityLoggingObject {
	data := iData.(map[string]interface{})
	return &SecurityLoggingObject{
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		SecurityLoggingObjectRules: InterfaceToSecurityLoggingObjectRuleListType(data["security_logging_object_rules"]),

		//{"description":"Security logging object rules derived internally.","type":"object","properties":{"rule":{"type":"array","item":{"type":"object","properties":{"rate":{"type":"integer"},"rule_uuid":{"type":"string"}}}}}}
		SecurityLoggingObjectRate: data["security_logging_object_rate"].(int),

		//{"description":"Security logging object rate defining rate of session logging","default":"100","type":"integer"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToSecurityLoggingObjectSlice makes a slice of SecurityLoggingObject from interface
func InterfaceToSecurityLoggingObjectSlice(data interface{}) []*SecurityLoggingObject {
	list := data.([]interface{})
	result := MakeSecurityLoggingObjectSlice()
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObject(item))
	}
	return result
}

// MakeSecurityLoggingObjectSlice() makes a slice of SecurityLoggingObject
func MakeSecurityLoggingObjectSlice() []*SecurityLoggingObject {
	return []*SecurityLoggingObject{}
}
