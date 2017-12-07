package models

// LoadbalancerType

import "encoding/json"

// LoadbalancerType
type LoadbalancerType struct {
	ProvisioningStatus string         `json:"provisioning_status"`
	AdminState         bool           `json:"admin_state"`
	VipAddress         IpAddressType  `json:"vip_address"`
	VipSubnetID        UuidStringType `json:"vip_subnet_id"`
	OperatingStatus    string         `json:"operating_status"`
	Status             string         `json:"status"`
}

//  parents relation object

// String returns json representation of the object
func (model *LoadbalancerType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerType makes LoadbalancerType
func MakeLoadbalancerType() *LoadbalancerType {
	return &LoadbalancerType{
		//TODO(nati): Apply default
		Status:             "",
		ProvisioningStatus: "",
		AdminState:         false,
		VipAddress:         MakeIpAddressType(),
		VipSubnetID:        MakeUuidStringType(),
		OperatingStatus:    "",
	}
}

// InterfaceToLoadbalancerType makes LoadbalancerType from interface
func InterfaceToLoadbalancerType(iData interface{}) *LoadbalancerType {
	data := iData.(map[string]interface{})
	return &LoadbalancerType{
		VipAddress: InterfaceToIpAddressType(data["vip_address"]),

		//{"Title":"","Description":"Virtual ip for this LBaaS","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"VipAddress","GoType":"IpAddressType","GoPremitive":false}
		VipSubnetID: InterfaceToUuidStringType(data["vip_subnet_id"]),

		//{"Title":"","Description":"Subnet UUID of the subnet of VIP, representing virtual network.","SQL":"","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UuidStringType","CollectionType":"","Column":"","Item":null,"GoName":"VipSubnetID","GoType":"UuidStringType","GoPremitive":false}
		OperatingStatus: data["operating_status"].(string),

		//{"Title":"","Description":"Operational status of the load balancer updated by system.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"OperatingStatus","GoType":"string","GoPremitive":true}
		Status: data["status"].(string),

		//{"Title":"","Description":"Operational status of the load balancer updated by system.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Status","GoType":"string","GoPremitive":true}
		ProvisioningStatus: data["provisioning_status"].(string),

		//{"Title":"","Description":"Provisioning  status of the load balancer updated by system.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProvisioningStatus","GoType":"string","GoPremitive":true}
		AdminState: data["admin_state"].(bool),

		//{"Title":"","Description":"Administrative up or down","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AdminState","GoType":"bool","GoPremitive":true}

	}
}

// InterfaceToLoadbalancerTypeSlice makes a slice of LoadbalancerType from interface
func InterfaceToLoadbalancerTypeSlice(data interface{}) []*LoadbalancerType {
	list := data.([]interface{})
	result := MakeLoadbalancerTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerType(item))
	}
	return result
}

// MakeLoadbalancerTypeSlice() makes a slice of LoadbalancerType
func MakeLoadbalancerTypeSlice() []*LoadbalancerType {
	return []*LoadbalancerType{}
}
