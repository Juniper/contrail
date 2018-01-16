package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type,omitempty"`
	ProtocolPort          int                      `json:"protocol_port,omitempty"`
	Status                string                   `json:"status,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
	ConnectionLimit       int                      `json:"connection_limit,omitempty"`
	AdminState            bool                     `json:"admin_state,omitempty"`
	Address               IpAddressType            `json:"address,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
}

// String returns json representation of the object
func (model *VirtualIpType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualIpType makes VirtualIpType
func MakeVirtualIpType() *VirtualIpType {
	return &VirtualIpType{
		//TODO(nati): Apply default
		ConnectionLimit:       0,
		AdminState:            false,
		Address:               MakeIpAddressType(),
		StatusDescription:     "",
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
		ProtocolPort:          0,
		Status:                "",
		PersistenceCookieName: "",
		PersistenceType:       MakeSessionPersistenceType(),
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
