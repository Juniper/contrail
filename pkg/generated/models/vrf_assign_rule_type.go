package models

// VrfAssignRuleType

import "encoding/json"

// VrfAssignRuleType
type VrfAssignRuleType struct {
	VlanTag         int                 `json:"vlan_tag"`
	IgnoreACL       bool                `json:"ignore_acl"`
	RoutingInstance string              `json:"routing_instance"`
	MatchCondition  *MatchConditionType `json:"match_condition"`
}

// String returns json representation of the object
func (model *VrfAssignRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVrfAssignRuleType makes VrfAssignRuleType
func MakeVrfAssignRuleType() *VrfAssignRuleType {
	return &VrfAssignRuleType{
		//TODO(nati): Apply default
		RoutingInstance: "",
		MatchCondition:  MakeMatchConditionType(),
		VlanTag:         0,
		IgnoreACL:       false,
	}
}

// InterfaceToVrfAssignRuleType makes VrfAssignRuleType from interface
func InterfaceToVrfAssignRuleType(iData interface{}) *VrfAssignRuleType {
	data := iData.(map[string]interface{})
	return &VrfAssignRuleType{
		IgnoreACL: data["ignore_acl"].(bool),

		//{"type":"boolean"}
		RoutingInstance: data["routing_instance"].(string),

		//{"type":"string"}
		MatchCondition: InterfaceToMatchConditionType(data["match_condition"]),

		//{"type":"object","properties":{"dst_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"dst_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}},"ethertype":{"type":"string","enum":["IPv4","IPv6"]},"protocol":{"type":"string"},"src_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"src_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}}
		VlanTag: data["vlan_tag"].(int),

		//{"type":"integer"}

	}
}

// InterfaceToVrfAssignRuleTypeSlice makes a slice of VrfAssignRuleType from interface
func InterfaceToVrfAssignRuleTypeSlice(data interface{}) []*VrfAssignRuleType {
	list := data.([]interface{})
	result := MakeVrfAssignRuleTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVrfAssignRuleType(item))
	}
	return result
}

// MakeVrfAssignRuleTypeSlice() makes a slice of VrfAssignRuleType
func MakeVrfAssignRuleTypeSlice() []*VrfAssignRuleType {
	return []*VrfAssignRuleType{}
}
