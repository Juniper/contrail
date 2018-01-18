package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	Address               IpAddressType            `json:"address,omitempty"`
	ProtocolPort          int                      `json:"protocol_port,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
	ConnectionLimit       int                      `json:"connection_limit,omitempty"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type,omitempty"`
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	AdminState            bool                     `json:"admin_state"`
	Status                string                   `json:"status,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
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
		ProtocolPort:          0,
		StatusDescription:     "",
		ConnectionLimit:       0,
		SubnetID:              MakeUuidStringType(),
		PersistenceCookieName: "",
		AdminState:            false,
		Status:                "",
		Protocol:              MakeLoadbalancerProtocolType(),
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
