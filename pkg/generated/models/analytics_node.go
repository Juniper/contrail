package models

// AnalyticsNode

import "encoding/json"

// AnalyticsNode
type AnalyticsNode struct {
	AnalyticsNodeIPAddress IpAddressType  `json:"analytics_node_ip_address,omitempty"`
	FQName                 []string       `json:"fq_name,omitempty"`
	Annotations            *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                 *PermType2     `json:"perms2,omitempty"`
	ParentType             string         `json:"parent_type,omitempty"`
	IDPerms                *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName            string         `json:"display_name,omitempty"`
	UUID                   string         `json:"uuid,omitempty"`
	ParentUUID             string         `json:"parent_uuid,omitempty"`
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
		ParentUUID:             "",
		ParentType:             "",
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		UUID:                   "",
		Perms2:                 MakePermType2(),
		AnalyticsNodeIPAddress: MakeIpAddressType(),
		FQName:                 []string{},
		Annotations:            MakeKeyValuePairs(),
	}
}

// MakeAnalyticsNodeSlice() makes a slice of AnalyticsNode
func MakeAnalyticsNodeSlice() []*AnalyticsNode {
	return []*AnalyticsNode{}
}
