package models

// PolicyRuleType

import "encoding/json"

// PolicyRuleType
type PolicyRuleType struct {
	RuleUUID     string          `json:"rule_uuid"`
	LastModified string          `json:"last_modified"`
	SRCPorts     []*PortType     `json:"src_ports"`
	Application  []string        `json:"application"`
	Ethertype    EtherType       `json:"ethertype"`
	Direction    DirectionType   `json:"direction"`
	Protocol     string          `json:"protocol"`
	DSTAddresses []*AddressType  `json:"dst_addresses"`
	ActionList   *ActionListType `json:"action_list"`
	Created      string          `json:"created"`
	DSTPorts     []*PortType     `json:"dst_ports"`
	SRCAddresses []*AddressType  `json:"src_addresses"`
	RuleSequence *SequenceType   `json:"rule_sequence"`
}

// String returns json representation of the object
func (model *PolicyRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePolicyRuleType makes PolicyRuleType
func MakePolicyRuleType() *PolicyRuleType {
	return &PolicyRuleType{
		//TODO(nati): Apply default
		RuleUUID:     "",
		LastModified: "",

		SRCPorts: MakePortTypeSlice(),

		Direction: MakeDirectionType(),
		Protocol:  "",

		DSTAddresses: MakeAddressTypeSlice(),

		ActionList: MakeActionListType(),
		Created:    "",

		DSTPorts: MakePortTypeSlice(),

		Application: []string{},
		Ethertype:   MakeEtherType(),

		SRCAddresses: MakeAddressTypeSlice(),

		RuleSequence: MakeSequenceType(),
	}
}

// InterfaceToPolicyRuleType makes PolicyRuleType from interface
func InterfaceToPolicyRuleType(iData interface{}) *PolicyRuleType {
	data := iData.(map[string]interface{})
	return &PolicyRuleType{
		Application: data["application"].([]string),

		//{"description":"Optionally application can be specified instead of protocol and port. not currently implemented","type":"array","item":{"type":"string"}}
		Ethertype: InterfaceToEtherType(data["ethertype"]),

		//{"type":"string","enum":["IPv4","IPv6"]}
		Direction: InterfaceToDirectionType(data["direction"]),

		//{"type":"string","enum":["\u003e","\u003c\u003e"]}
		Protocol: data["protocol"].(string),

		//{"description":"Layer 4 protocol in ip packet","type":"string"}

		DSTAddresses: InterfaceToAddressTypeSlice(data["dst_addresses"]),

		//{"description":"Destination ip matching criteria","type":"array","item":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}}
		ActionList: InterfaceToActionListType(data["action_list"]),

		//{"description":"Actions to be performed if packets match condition","type":"object","properties":{"alert":{"type":"boolean"},"apply_service":{"type":"array","item":{"type":"string"}},"assign_routing_instance":{"type":"string"},"gateway_name":{"type":"string"},"log":{"type":"boolean"},"mirror_to":{"type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}},"qos_action":{"type":"string"},"simple_action":{"type":"string","enum":["deny","pass"]}}}
		Created: data["created"].(string),

		//{"description":"timestamp when security group rule object gets created","type":"string"}

		DSTPorts: InterfaceToPortTypeSlice(data["dst_ports"]),

		//{"description":"Range of destination  port for layer 4 protocol","type":"array","item":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}

		SRCAddresses: InterfaceToAddressTypeSlice(data["src_addresses"]),

		//{"description":"Source ip matching criteria","type":"array","item":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}}
		RuleSequence: InterfaceToSequenceType(data["rule_sequence"]),

		//{"description":"Deprecated, Will be removed because rules themselves are already an ordered list","type":"object","properties":{"major":{"type":"integer"},"minor":{"type":"integer"}}}
		RuleUUID: data["rule_uuid"].(string),

		//{"description":"Rule UUID is identifier used in flow records to identify rule","type":"string"}
		LastModified: data["last_modified"].(string),

		//{"description":"timestamp when security group rule object gets updated","type":"string"}

		SRCPorts: InterfaceToPortTypeSlice(data["src_ports"]),

		//{"description":"Range of source port for layer 4 protocol","type":"array","item":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}

	}
}

// InterfaceToPolicyRuleTypeSlice makes a slice of PolicyRuleType from interface
func InterfaceToPolicyRuleTypeSlice(data interface{}) []*PolicyRuleType {
	list := data.([]interface{})
	result := MakePolicyRuleTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPolicyRuleType(item))
	}
	return result
}

// MakePolicyRuleTypeSlice() makes a slice of PolicyRuleType
func MakePolicyRuleTypeSlice() []*PolicyRuleType {
	return []*PolicyRuleType{}
}
