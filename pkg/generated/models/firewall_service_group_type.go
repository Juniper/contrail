package models

// FirewallServiceGroupType

import "encoding/json"

// FirewallServiceGroupType
type FirewallServiceGroupType struct {
	FirewallService []*FirewallServiceType `json:"firewall_service"`
}

// String returns json representation of the object
func (model *FirewallServiceGroupType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallServiceGroupType makes FirewallServiceGroupType
func MakeFirewallServiceGroupType() *FirewallServiceGroupType {
	return &FirewallServiceGroupType{
		//TODO(nati): Apply default

		FirewallService: MakeFirewallServiceTypeSlice(),
	}
}

// InterfaceToFirewallServiceGroupType makes FirewallServiceGroupType from interface
func InterfaceToFirewallServiceGroupType(iData interface{}) *FirewallServiceGroupType {
	data := iData.(map[string]interface{})
	return &FirewallServiceGroupType{

		FirewallService: InterfaceToFirewallServiceTypeSlice(data["firewall_service"]),

		//{"type":"array","item":{"type":"object","properties":{"dst_ports":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}},"protocol":{"type":"string"},"protocol_id":{"type":"integer"},"src_ports":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}}}

	}
}

// InterfaceToFirewallServiceGroupTypeSlice makes a slice of FirewallServiceGroupType from interface
func InterfaceToFirewallServiceGroupTypeSlice(data interface{}) []*FirewallServiceGroupType {
	list := data.([]interface{})
	result := MakeFirewallServiceGroupTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallServiceGroupType(item))
	}
	return result
}

// MakeFirewallServiceGroupTypeSlice() makes a slice of FirewallServiceGroupType
func MakeFirewallServiceGroupTypeSlice() []*FirewallServiceGroupType {
	return []*FirewallServiceGroupType{}
}
