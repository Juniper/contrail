package models

// PolicyEntriesType

import "encoding/json"

// PolicyEntriesType
type PolicyEntriesType struct {
	PolicyRule []*PolicyRuleType `json:"policy_rule"`
}

// String returns json representation of the object
func (model *PolicyEntriesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePolicyEntriesType makes PolicyEntriesType
func MakePolicyEntriesType() *PolicyEntriesType {
	return &PolicyEntriesType{
		//TODO(nati): Apply default

		PolicyRule: MakePolicyRuleTypeSlice(),
	}
}

// InterfaceToPolicyEntriesType makes PolicyEntriesType from interface
func InterfaceToPolicyEntriesType(iData interface{}) *PolicyEntriesType {
	data := iData.(map[string]interface{})
	return &PolicyEntriesType{

		PolicyRule: InterfaceToPolicyRuleTypeSlice(data["policy_rule"]),

		//{"description":"List of policy rules","type":"array","item":{"type":"object","properties":{"action_list":{"type":"object","properties":{"alert":{"type":"boolean"},"apply_service":{"type":"array","item":{"type":"string"}},"assign_routing_instance":{"type":"string"},"gateway_name":{"type":"string"},"log":{"type":"boolean"},"mirror_to":{"type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}},"qos_action":{"type":"string"},"simple_action":{"type":"string","enum":["deny","pass"]}}},"application":{"type":"array","item":{"type":"string"}},"created":{"type":"string"},"direction":{"type":"string","enum":["\u003e","\u003c\u003e"]},"dst_addresses":{"type":"array","item":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}},"dst_ports":{"type":"array","item":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}},"ethertype":{"type":"string","enum":["IPv4","IPv6"]},"last_modified":{"type":"string"},"protocol":{"type":"string"},"rule_sequence":{"type":"object","properties":{"major":{"type":"integer"},"minor":{"type":"integer"}}},"rule_uuid":{"type":"string"},"src_addresses":{"type":"array","item":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}},"src_ports":{"type":"array","item":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}}}}

	}
}

// InterfaceToPolicyEntriesTypeSlice makes a slice of PolicyEntriesType from interface
func InterfaceToPolicyEntriesTypeSlice(data interface{}) []*PolicyEntriesType {
	list := data.([]interface{})
	result := MakePolicyEntriesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPolicyEntriesType(item))
	}
	return result
}

// MakePolicyEntriesTypeSlice() makes a slice of PolicyEntriesType
func MakePolicyEntriesTypeSlice() []*PolicyEntriesType {
	return []*PolicyEntriesType{}
}
