package models

// APIAccessList

import "encoding/json"

// APIAccessList
type APIAccessList struct {
	DisplayName          string               `json:"display_name"`
	ParentType           string               `json:"parent_type"`
	FQName               []string             `json:"fq_name"`
	APIAccessListEntries *RbacRuleEntriesType `json:"api_access_list_entries"`
	Perms2               *PermType2           `json:"perms2"`
	UUID                 string               `json:"uuid"`
	ParentUUID           string               `json:"parent_uuid"`
	IDPerms              *IdPermsType         `json:"id_perms"`
	Annotations          *KeyValuePairs       `json:"annotations"`
}

// String returns json representation of the object
func (model *APIAccessList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAPIAccessList makes APIAccessList
func MakeAPIAccessList() *APIAccessList {
	return &APIAccessList{
		//TODO(nati): Apply default
		ParentType:           "",
		FQName:               []string{},
		APIAccessListEntries: MakeRbacRuleEntriesType(),
		DisplayName:          "",
		UUID:                 "",
		ParentUUID:           "",
		IDPerms:              MakeIdPermsType(),
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
	}
}

// InterfaceToAPIAccessList makes APIAccessList from interface
func InterfaceToAPIAccessList(iData interface{}) *APIAccessList {
	data := iData.(map[string]interface{})
	return &APIAccessList{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		APIAccessListEntries: InterfaceToRbacRuleEntriesType(data["api_access_list_entries"]),

		//{"description":"List of rules e.g network.* =\u003e admin:CRUD (admin can perform all ops on networks).","type":"object","properties":{"rbac_rule":{"type":"array","item":{"type":"object","properties":{"rule_field":{"type":"string"},"rule_object":{"type":"string"},"rule_perms":{"type":"array","item":{"type":"object","properties":{"role_crud":{"type":"string"},"role_name":{"type":"string"}}}}}}}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToAPIAccessListSlice makes a slice of APIAccessList from interface
func InterfaceToAPIAccessListSlice(data interface{}) []*APIAccessList {
	list := data.([]interface{})
	result := MakeAPIAccessListSlice()
	for _, item := range list {
		result = append(result, InterfaceToAPIAccessList(item))
	}
	return result
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
	return []*APIAccessList{}
}
