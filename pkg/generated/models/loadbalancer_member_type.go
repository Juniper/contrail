package models

// LoadbalancerMemberType

import "encoding/json"

// LoadbalancerMemberType
type LoadbalancerMemberType struct {
	Address           IpAddressType `json:"address"`
	ProtocolPort      int           `json:"protocol_port"`
	Status            string        `json:"status"`
	StatusDescription string        `json:"status_description"`
	Weight            int           `json:"weight"`
	AdminState        bool          `json:"admin_state"`
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
		Weight:            0,
		AdminState:        false,
		Address:           MakeIpAddressType(),
		ProtocolPort:      0,
		Status:            "",
		StatusDescription: "",
	}
}

// MakeLoadbalancerMemberTypeSlice() makes a slice of LoadbalancerMemberType
func MakeLoadbalancerMemberTypeSlice() []*LoadbalancerMemberType {
	return []*LoadbalancerMemberType{}
}
