package models

// AnalyticsNode

import "encoding/json"

// AnalyticsNode
type AnalyticsNode struct {
	Annotations            *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID             string         `json:"parent_uuid,omitempty"`
	FQName                 []string       `json:"fq_name,omitempty"`
	IDPerms                *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName            string         `json:"display_name,omitempty"`
	Perms2                 *PermType2     `json:"perms2,omitempty"`
	UUID                   string         `json:"uuid,omitempty"`
	ParentType             string         `json:"parent_type,omitempty"`
	AnalyticsNodeIPAddress IpAddressType  `json:"analytics_node_ip_address,omitempty"`
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
		DisplayName:            "",
		Annotations:            MakeKeyValuePairs(),
		ParentUUID:             "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		AnalyticsNodeIPAddress: MakeIpAddressType(),
		Perms2:                 MakePermType2(),
		UUID:                   "",
		ParentType:             "",
	}
}

// MakeAnalyticsNodeSlice() makes a slice of AnalyticsNode
func MakeAnalyticsNodeSlice() []*AnalyticsNode {
	return []*AnalyticsNode{}
}
