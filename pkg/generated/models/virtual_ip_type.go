package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	Status                string                   `json:"status,omitempty"`
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type,omitempty"`
	ProtocolPort          int                      `json:"protocol_port,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
	ConnectionLimit       int                      `json:"connection_limit,omitempty"`
	AdminState            bool                     `json:"admin_state"`
	Address               IpAddressType            `json:"address,omitempty"`
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
		ProtocolPort:          0,
		Status:                "",
		PersistenceCookieName: "",
		PersistenceType:       MakeSessionPersistenceType(),
		ConnectionLimit:       0,
		AdminState:            false,
		Address:               MakeIpAddressType(),
		StatusDescription:     "",
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
