package models

// AclEntriesType

import "encoding/json"

// AclEntriesType
type AclEntriesType struct {
	Dynamic bool           `json:"dynamic"`
	ACLRule []*AclRuleType `json:"acl_rule"`
}

// String returns json representation of the object
func (model *AclEntriesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAclEntriesType makes AclEntriesType
func MakeAclEntriesType() *AclEntriesType {
	return &AclEntriesType{
		//TODO(nati): Apply default
		Dynamic: false,

		ACLRule: MakeAclRuleTypeSlice(),
	}
}

// InterfaceToAclEntriesType makes AclEntriesType from interface
func InterfaceToAclEntriesType(iData interface{}) *AclEntriesType {
	data := iData.(map[string]interface{})
	return &AclEntriesType{
		Dynamic: data["dynamic"].(bool),

		//{"description":"For Internal use only","type":"boolean"}

		ACLRule: InterfaceToAclRuleTypeSlice(data["acl_rule"]),

		//{"description":"For Internal use only","type":"array","item":{"type":"object","properties":{"action_list":{"type":"object","properties":{"alert":{"type":"boolean"},"apply_service":{"type":"array","item":{"type":"string"}},"assign_routing_instance":{"type":"string"},"gateway_name":{"type":"string"},"log":{"type":"boolean"},"mirror_to":{"type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}},"qos_action":{"type":"string"},"simple_action":{"type":"string","enum":["deny","pass"]}}},"direction":{"type":"string","enum":["\u003e","\u003c\u003e"]},"match_condition":{"type":"object","properties":{"dst_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"dst_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}},"ethertype":{"type":"string","enum":["IPv4","IPv6"]},"protocol":{"type":"string"},"src_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"src_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}},"rule_uuid":{"type":"string"}}}}

	}
}

// InterfaceToAclEntriesTypeSlice makes a slice of AclEntriesType from interface
func InterfaceToAclEntriesTypeSlice(data interface{}) []*AclEntriesType {
	list := data.([]interface{})
	result := MakeAclEntriesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAclEntriesType(item))
	}
	return result
}

// MakeAclEntriesTypeSlice() makes a slice of AclEntriesType
func MakeAclEntriesTypeSlice() []*AclEntriesType {
	return []*AclEntriesType{}
}
