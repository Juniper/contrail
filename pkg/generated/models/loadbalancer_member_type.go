package models

// LoadbalancerMemberType

import "encoding/json"

// LoadbalancerMemberType
type LoadbalancerMemberType struct {
	AdminState        bool          `json:"admin_state"`
	Address           IpAddressType `json:"address"`
	ProtocolPort      int           `json:"protocol_port"`
	Status            string        `json:"status"`
	StatusDescription string        `json:"status_description"`
	Weight            int           `json:"weight"`
}

//  parents relation object

// String returns json representation of the object
func (model *LoadbalancerMemberType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerMemberType makes LoadbalancerMemberType
func MakeLoadbalancerMemberType() *LoadbalancerMemberType {
	return &LoadbalancerMemberType{
		//TODO(nati): Apply default
		StatusDescription: "",
		Weight:            0,
		AdminState:        false,
		Address:           MakeIpAddressType(),
		ProtocolPort:      0,
		Status:            "",
	}
}

// InterfaceToLoadbalancerMemberType makes LoadbalancerMemberType from interface
func InterfaceToLoadbalancerMemberType(iData interface{}) *LoadbalancerMemberType {
	data := iData.(map[string]interface{})
	return &LoadbalancerMemberType{
		Status: data["status"].(string),

		//{"Title":"","Description":"Operational status of the member.","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Status","GoType":"string","GoPremitive":true}
		StatusDescription: data["status_description"].(string),

		//{"Title":"","Description":"Operational status description of the member.","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StatusDescription","GoType":"string","GoPremitive":true}
		Weight: data["weight"].(int),

		//{"Title":"","Description":"Weight for load balancing","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Weight","GoType":"int","GoPremitive":true}
		AdminState: data["admin_state"].(bool),

		//{"Title":"","Description":"Administrative up or down.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AdminState","GoType":"bool","GoPremitive":true}
		Address: InterfaceToIpAddressType(data["address"]),

		//{"Title":"","Description":"Ip address of the member","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"Address","GoType":"IpAddressType","GoPremitive":false}
		ProtocolPort: data["protocol_port"].(int),

		//{"Title":"","Description":"Destination port for the application on the member.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProtocolPort","GoType":"int","GoPremitive":true}

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
