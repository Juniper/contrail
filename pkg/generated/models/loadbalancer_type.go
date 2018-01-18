package models

// LoadbalancerType

import "encoding/json"

// LoadbalancerType
type LoadbalancerType struct {
	VipAddress         IpAddressType  `json:"vip_address,omitempty"`
	VipSubnetID        UuidStringType `json:"vip_subnet_id,omitempty"`
	OperatingStatus    string         `json:"operating_status,omitempty"`
	Status             string         `json:"status,omitempty"`
	ProvisioningStatus string         `json:"provisioning_status,omitempty"`
	AdminState         bool           `json:"admin_state"`
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
		VipSubnetID:        MakeUuidStringType(),
		OperatingStatus:    "",
		Status:             "",
		ProvisioningStatus: "",
		AdminState:         false,
		VipAddress:         MakeIpAddressType(),
	}
}

// MakeLoadbalancerTypeSlice() makes a slice of LoadbalancerType
func MakeLoadbalancerTypeSlice() []*LoadbalancerType {
	return []*LoadbalancerType{}
}
