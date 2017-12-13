package models

// AnalyticsNode

import "encoding/json"

// AnalyticsNode
type AnalyticsNode struct {
	ParentUUID             string         `json:"parent_uuid"`
	ParentType             string         `json:"parent_type"`
	Annotations            *KeyValuePairs `json:"annotations"`
	Perms2                 *PermType2     `json:"perms2"`
	UUID                   string         `json:"uuid"`
	IDPerms                *IdPermsType   `json:"id_perms"`
	AnalyticsNodeIPAddress IpAddressType  `json:"analytics_node_ip_address"`
	DisplayName            string         `json:"display_name"`
	FQName                 []string       `json:"fq_name"`
}

// String returns json representation of the object
func (model *AnalyticsNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAnalyticsNode makes AnalyticsNode
func MakeAnalyticsNode() *AnalyticsNode {
	return &AnalyticsNode{
		//TODO(nati): Apply default
		AnalyticsNodeIPAddress: MakeIpAddressType(),
		DisplayName:            "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		Annotations:            MakeKeyValuePairs(),
		Perms2:                 MakePermType2(),
		UUID:                   "",
		ParentUUID:             "",
		ParentType:             "",
	}
}

// InterfaceToAnalyticsNode makes AnalyticsNode from interface
func InterfaceToAnalyticsNode(iData interface{}) *AnalyticsNode {
	data := iData.(map[string]interface{})
	return &AnalyticsNode{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		AnalyticsNodeIPAddress: InterfaceToIpAddressType(data["analytics_node_ip_address"]),

		//{"description":"Ip address of the analytics node, set while provisioning.","type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToAnalyticsNodeSlice makes a slice of AnalyticsNode from interface
func InterfaceToAnalyticsNodeSlice(data interface{}) []*AnalyticsNode {
	list := data.([]interface{})
	result := MakeAnalyticsNodeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAnalyticsNode(item))
	}
	return result
}

// MakeAnalyticsNodeSlice() makes a slice of AnalyticsNode
func MakeAnalyticsNodeSlice() []*AnalyticsNode {
	return []*AnalyticsNode{}
}
