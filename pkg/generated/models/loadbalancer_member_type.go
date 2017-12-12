package models

// LoadbalancerMemberType

import "encoding/json"

// LoadbalancerMemberType
type LoadbalancerMemberType struct {
	Weight            int           `json:"weight"`
	AdminState        bool          `json:"admin_state"`
	Address           IpAddressType `json:"address"`
	ProtocolPort      int           `json:"protocol_port"`
	Status            string        `json:"status"`
	StatusDescription string        `json:"status_description"`
}

// String returns json representation of the object
func (model *LoadbalancerMemberType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerMemberType makes LoadbalancerMemberType
func MakeLoadbalancerMemberType() *LoadbalancerMemberType {
	return &LoadbalancerMemberType{
		//TODO(nati): Apply default
		ProtocolPort:      0,
		Status:            "",
		StatusDescription: "",
		Weight:            0,
		AdminState:        false,
		Address:           MakeIpAddressType(),
	}
}

// InterfaceToLoadbalancerMemberType makes LoadbalancerMemberType from interface
func InterfaceToLoadbalancerMemberType(iData interface{}) *LoadbalancerMemberType {
	data := iData.(map[string]interface{})
	return &LoadbalancerMemberType{
		Weight: data["weight"].(int),

		//{"description":"Weight for load balancing","type":"integer"}
		AdminState: data["admin_state"].(bool),

		//{"description":"Administrative up or down.","type":"boolean"}
		Address: InterfaceToIpAddressType(data["address"]),

		//{"description":"Ip address of the member","type":"string"}
		ProtocolPort: data["protocol_port"].(int),

		//{"description":"Destination port for the application on the member.","type":"integer"}
		Status: data["status"].(string),

		//{"description":"Operational status of the member.","type":"string"}
		StatusDescription: data["status_description"].(string),

		//{"description":"Operational status description of the member.","type":"string"}

	}
}

// InterfaceToLoadbalancerMemberTypeSlice makes a slice of LoadbalancerMemberType from interface
func InterfaceToLoadbalancerMemberTypeSlice(data interface{}) []*LoadbalancerMemberType {
	list := data.([]interface{})
	result := MakeLoadbalancerMemberTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerMemberType(item))
	}
	return result
}

// MakeLoadbalancerMemberTypeSlice() makes a slice of LoadbalancerMemberType
func MakeLoadbalancerMemberTypeSlice() []*LoadbalancerMemberType {
	return []*LoadbalancerMemberType{}
}
