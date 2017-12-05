package models

// LoadbalancerMemberType

import "encoding/json"

type LoadbalancerMemberType struct {
	Address           IpAddressType `json:"address"`
	ProtocolPort      int           `json:"protocol_port"`
	Status            string        `json:"status"`
	StatusDescription string        `json:"status_description"`
	Weight            int           `json:"weight"`
	AdminState        bool          `json:"admin_state"`
}

func (model *LoadbalancerMemberType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeLoadbalancerMemberType() *LoadbalancerMemberType {
	return &LoadbalancerMemberType{
		//TODO(nati): Apply default
		Status:            "",
		StatusDescription: "",
		Weight:            0,
		AdminState:        false,
		Address:           MakeIpAddressType(),
		ProtocolPort:      0,
	}
}

func InterfaceToLoadbalancerMemberType(iData interface{}) *LoadbalancerMemberType {
	data := iData.(map[string]interface{})
	return &LoadbalancerMemberType{
		ProtocolPort: data["protocol_port"].(int),

		//{"Title":"","Description":"Destination port for the application on the member.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProtocolPort","GoType":"int"}
		Status: data["status"].(string),

		//{"Title":"","Description":"Operational status of the member.","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Status","GoType":"string"}
		StatusDescription: data["status_description"].(string),

		//{"Title":"","Description":"Operational status description of the member.","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StatusDescription","GoType":"string"}
		Weight: data["weight"].(int),

		//{"Title":"","Description":"Weight for load balancing","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Weight","GoType":"int"}
		AdminState: data["admin_state"].(bool),

		//{"Title":"","Description":"Administrative up or down.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AdminState","GoType":"bool"}
		Address: InterfaceToIpAddressType(data["address"]),

		//{"Title":"","Description":"Ip address of the member","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"Address","GoType":"IpAddressType"}

	}
}

func InterfaceToLoadbalancerMemberTypeSlice(data interface{}) []*LoadbalancerMemberType {
	list := data.([]interface{})
	result := MakeLoadbalancerMemberTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerMemberType(item))
	}
	return result
}

func MakeLoadbalancerMemberTypeSlice() []*LoadbalancerMemberType {
	return []*LoadbalancerMemberType{}
}
