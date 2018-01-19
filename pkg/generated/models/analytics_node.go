package models

// AnalyticsNode

import "encoding/json"

// AnalyticsNode
type AnalyticsNode struct {
	DisplayName            string         `json:"display_name,omitempty"`
	Annotations            *KeyValuePairs `json:"annotations,omitempty"`
	IDPerms                *IdPermsType   `json:"id_perms,omitempty"`
	Perms2                 *PermType2     `json:"perms2,omitempty"`
	UUID                   string         `json:"uuid,omitempty"`
	AnalyticsNodeIPAddress IpAddressType  `json:"analytics_node_ip_address,omitempty"`
	ParentUUID             string         `json:"parent_uuid,omitempty"`
	ParentType             string         `json:"parent_type,omitempty"`
	FQName                 []string       `json:"fq_name,omitempty"`
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
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		Perms2:      MakePermType2(),
		UUID:        "",
		AnalyticsNodeIPAddress: MakeIpAddressType(),
		ParentUUID:             "",
	}
}

// MakeAnalyticsNodeSlice() makes a slice of AnalyticsNode
func MakeAnalyticsNodeSlice() []*AnalyticsNode {
	return []*AnalyticsNode{}
}
