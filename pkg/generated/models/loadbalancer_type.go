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
		VipSubnetID: InterfaceToUuidStringType(data["vip_subnet_id"]),

		//{"description":"Subnet UUID of the subnet of VIP, representing virtual network.","type":"string"}
		OperatingStatus: data["operating_status"].(string),

		//{"description":"Operational status of the load balancer updated by system.","type":"string"}
		Status: data["status"].(string),

		//{"description":"Operational status of the load balancer updated by system.","type":"string"}
		ProvisioningStatus: data["provisioning_status"].(string),

		//{"description":"Provisioning  status of the load balancer updated by system.","type":"string"}
		AdminState: data["admin_state"].(bool),

		//{"description":"Administrative up or down","type":"boolean"}
		VipAddress: InterfaceToIpAddressType(data["vip_address"]),

		//{"description":"Virtual ip for this LBaaS","type":"string"}

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
