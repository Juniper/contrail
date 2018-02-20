package models

// LoadbalancerMemberType

// LoadbalancerMemberType
//proteus:generate
type LoadbalancerMemberType struct {
	Status            string        `json:"status,omitempty"`
	StatusDescription string        `json:"status_description,omitempty"`
	Weight            int           `json:"weight,omitempty"`
	AdminState        bool          `json:"admin_state"`
	Address           IpAddressType `json:"address,omitempty"`
	ProtocolPort      int           `json:"protocol_port,omitempty"`
}

// MakeLoadbalancerMemberType makes LoadbalancerMemberType
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

// MakeLoadbalancerMemberTypeSlice() makes a slice of LoadbalancerMemberType
func MakeLoadbalancerMemberTypeSlice() []*LoadbalancerMemberType {
	return []*LoadbalancerMemberType{}
}
