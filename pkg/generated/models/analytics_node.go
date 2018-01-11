package models

// AnalyticsNode

import "encoding/json"

// AnalyticsNode
type AnalyticsNode struct {
	AnalyticsNodeIPAddress IpAddressType  `json:"analytics_node_ip_address"`
	Perms2                 *PermType2     `json:"perms2"`
	ParentUUID             string         `json:"parent_uuid"`
	IDPerms                *IdPermsType   `json:"id_perms"`
	DisplayName            string         `json:"display_name"`
	Annotations            *KeyValuePairs `json:"annotations"`
	UUID                   string         `json:"uuid"`
	ParentType             string         `json:"parent_type"`
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
		Perms2:                 MakePermType2(),
		ParentUUID:             "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Annotations:            MakeKeyValuePairs(),
		UUID:                   "",
		ParentType:             "",
	}
}

// MakeAnalyticsNodeSlice() makes a slice of AnalyticsNode
func MakeAnalyticsNodeSlice() []*AnalyticsNode {
	return []*AnalyticsNode{}
}
