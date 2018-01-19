package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	Status                string                   `json:"status,omitempty"`
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type,omitempty"`
	AdminState            bool                     `json:"admin_state"`
	Address               IpAddressType            `json:"address,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	ConnectionLimit       int                      `json:"connection_limit,omitempty"`
	ProtocolPort          int                      `json:"protocol_port,omitempty"`
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
		StatusDescription:     "",
		Protocol:              MakeLoadbalancerProtocolType(),
		PersistenceCookieName: "",
		ConnectionLimit:       0,
		ProtocolPort:          0,
		Status:                "",
		SubnetID:              MakeUuidStringType(),
		PersistenceType:       MakeSessionPersistenceType(),
		AdminState:            false,
		Address:               MakeIpAddressType(),
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
