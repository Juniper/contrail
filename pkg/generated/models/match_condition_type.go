package models

// MatchConditionType

import "encoding/json"

// MatchConditionType
type MatchConditionType struct {
	SRCPort    *PortType    `json:"src_port"`
	SRCAddress *AddressType `json:"src_address"`
	Ethertype  EtherType    `json:"ethertype"`
	DSTAddress *AddressType `json:"dst_address"`
	DSTPort    *PortType    `json:"dst_port"`
	Protocol   string       `json:"protocol"`
}

// String returns json representation of the object
func (model *MatchConditionType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMatchConditionType makes MatchConditionType
func MakeMatchConditionType() *MatchConditionType {
	return &MatchConditionType{
		//TODO(nati): Apply default
		Protocol:   "",
		SRCPort:    MakePortType(),
		SRCAddress: MakeAddressType(),
		Ethertype:  MakeEtherType(),
		DSTAddress: MakeAddressType(),
		DSTPort:    MakePortType(),
	}
}

// InterfaceToMatchConditionType makes MatchConditionType from interface
func InterfaceToMatchConditionType(iData interface{}) *MatchConditionType {
	data := iData.(map[string]interface{})
	return &MatchConditionType{
		SRCAddress: InterfaceToAddressType(data["src_address"]),

		//{"description":"Source ip matching criteria","type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}
		Ethertype: InterfaceToEtherType(data["ethertype"]),

		//{"type":"string","enum":["IPv4","IPv6"]}
		DSTAddress: InterfaceToAddressType(data["dst_address"]),

		//{"description":"Destination ip matching criteria","type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}
		DSTPort: InterfaceToPortType(data["dst_port"]),

		//{"description":"Range of destination  port for layer 4 protocol","type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}
		Protocol: data["protocol"].(string),

		//{"description":"Layer 4 protocol in ip packet","type":"string"}
		SRCPort: InterfaceToPortType(data["src_port"]),

		//{"description":"Range of source port for layer 4 protocol","type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}

	}
}

// InterfaceToMatchConditionTypeSlice makes a slice of MatchConditionType from interface
func InterfaceToMatchConditionTypeSlice(data interface{}) []*MatchConditionType {
	list := data.([]interface{})
	result := MakeMatchConditionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMatchConditionType(item))
	}
	return result
}

// MakeMatchConditionTypeSlice() makes a slice of MatchConditionType
func MakeMatchConditionTypeSlice() []*MatchConditionType {
	return []*MatchConditionType{}
}
