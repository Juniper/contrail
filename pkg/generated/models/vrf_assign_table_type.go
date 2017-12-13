package models

// VrfAssignTableType

import "encoding/json"

// VrfAssignTableType
type VrfAssignTableType struct {
	VRFAssignRule []*VrfAssignRuleType `json:"vrf_assign_rule"`
}

// String returns json representation of the object
func (model *VrfAssignTableType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVrfAssignTableType makes VrfAssignTableType
func MakeVrfAssignTableType() *VrfAssignTableType {
	return &VrfAssignTableType{
		//TODO(nati): Apply default

		VRFAssignRule: MakeVrfAssignRuleTypeSlice(),
	}
}

// InterfaceToVrfAssignTableType makes VrfAssignTableType from interface
func InterfaceToVrfAssignTableType(iData interface{}) *VrfAssignTableType {
	data := iData.(map[string]interface{})
	return &VrfAssignTableType{

		VRFAssignRule: InterfaceToVrfAssignRuleTypeSlice(data["vrf_assign_rule"]),

		//{"type":"array","item":{"type":"object","properties":{"ignore_acl":{"type":"boolean"},"match_condition":{"type":"object","properties":{"dst_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"dst_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}},"ethertype":{"type":"string","enum":["IPv4","IPv6"]},"protocol":{"type":"string"},"src_address":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}},"src_port":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}},"routing_instance":{"type":"string"},"vlan_tag":{"type":"integer"}}}}

	}
}

// InterfaceToVrfAssignTableTypeSlice makes a slice of VrfAssignTableType from interface
func InterfaceToVrfAssignTableTypeSlice(data interface{}) []*VrfAssignTableType {
	list := data.([]interface{})
	result := MakeVrfAssignTableTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVrfAssignTableType(item))
	}
	return result
}

// MakeVrfAssignTableTypeSlice() makes a slice of VrfAssignTableType
func MakeVrfAssignTableTypeSlice() []*VrfAssignTableType {
	return []*VrfAssignTableType{}
}
