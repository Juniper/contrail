package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	ConnectionLimit       int                      `json:"connection_limit,omitempty"`
	AdminState            bool                     `json:"admin_state"`
	ProtocolPort          int                      `json:"protocol_port,omitempty"`
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type,omitempty"`
	Address               IpAddressType            `json:"address,omitempty"`
	Status                string                   `json:"status,omitempty"`
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
		PersistenceCookieName: "",
		ConnectionLimit:       0,
		AdminState:            false,
		ProtocolPort:          0,
		PersistenceType:       MakeSessionPersistenceType(),
		Address:               MakeIpAddressType(),
		Status:                "",
		StatusDescription:     "",
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
