package models

// AclRuleType

import "encoding/json"

// AclRuleType
type AclRuleType struct {
	RuleUUID       string              `json:"rule_uuid"`
	MatchCondition *MatchConditionType `json:"match_condition"`
	Direction      DirectionType       `json:"direction"`
	ActionList     *ActionListType     `json:"action_list"`
}

// String returns json representation of the object
func (model *AclRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAclRuleType makes AclRuleType
func MakeAclRuleType() *AclRuleType {
	return &AclRuleType{
		//TODO(nati): Apply default
		RuleUUID:       "",
		MatchCondition: MakeMatchConditionType(),
		Direction:      MakeDirectionType(),
		ActionList:     MakeActionListType(),
	}
}

// InterfaceToAclRuleType makes AclRuleType from interface
func InterfaceToAclRuleType(iData interface{}) *AclRuleType {
	data := iData.(map[string]interface{})
	return &AclRuleType{
		RuleUUID: data["rule_uuid"].(string),

		//{"description":"Rule UUID is identifier used in flow records to identify rule","type":"string"}
		MatchCondition: InterfaceToMatchConditionType(data["match_condition"]),

		//{"description":"Match condition for packets","type":"object","properties":{"dst_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"dst_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}},"ethertype":{"type":"string","enum":["IPv4","IPv6"]},"protocol":{"type":"string"},"src_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"src_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}}
		Direction: InterfaceToDirectionType(data["direction"]),

		//{"description":"Direction in the rule","type":"string","enum":["\u003e","\u003c\u003e"]}
		ActionList: InterfaceToActionListType(data["action_list"]),

		//{"description":"Actions to be performed if packets match condition","type":"object","properties":{"alert":{"type":"boolean"},"apply_service":{"type":"array","item":{"type":"string"}},"assign_routing_instance":{"type":"string"},"gateway_name":{"type":"string"},"log":{"type":"boolean"},"mirror_to":{"type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}},"qos_action":{"type":"string"},"simple_action":{"type":"string","enum":["deny","pass"]}}}

	}
}

// InterfaceToAclRuleTypeSlice makes a slice of AclRuleType from interface
func InterfaceToAclRuleTypeSlice(data interface{}) []*AclRuleType {
	list := data.([]interface{})
	result := MakeAclRuleTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAclRuleType(item))
	}
	return result
}

// MakeAclRuleTypeSlice() makes a slice of AclRuleType
func MakeAclRuleTypeSlice() []*AclRuleType {
	return []*AclRuleType{}
}
