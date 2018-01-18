package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	Address               IpAddressType            `json:"address,omitempty"`
	ProtocolPort          int                      `json:"protocol_port,omitempty"`
	Status                string                   `json:"status,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	ConnectionLimit       int                      `json:"connection_limit,omitempty"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type,omitempty"`
	AdminState            bool                     `json:"admin_state"`
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
		Address:               MakeIpAddressType(),
		ProtocolPort:          0,
		Status:                "",
		StatusDescription:     "",
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
		PersistenceCookieName: "",
		ConnectionLimit:       0,
		PersistenceType:       MakeSessionPersistenceType(),
		AdminState:            false,
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
