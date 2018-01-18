package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	Status                string                   `json:"status,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	AdminState            bool                     `json:"admin_state,omitempty"`
	ProtocolPort          int                      `json:"protocol_port,omitempty"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type,omitempty"`
	Address               IpAddressType            `json:"address,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	ConnectionLimit       int                      `json:"connection_limit,omitempty"`
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
		PersistenceType:       MakeSessionPersistenceType(),
		Address:               MakeIpAddressType(),
		StatusDescription:     "",
		SubnetID:              MakeUuidStringType(),
		PersistenceCookieName: "",
		ConnectionLimit:       0,
		Status:                "",
		Protocol:              MakeLoadbalancerProtocolType(),
		AdminState:            false,
		ProtocolPort:          0,
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
