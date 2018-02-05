package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	Status                string                   `json:"status,omitempty"`
	StatusDescription     string                   `json:"status_description,omitempty"`
	Protocol              LoadbalancerProtocolType `json:"protocol,omitempty"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type,omitempty"`
	AdminState            bool                     `json:"admin_state"`
	ProtocolPort          int                      `json:"protocol_port,omitempty"`
	SubnetID              UuidStringType           `json:"subnet_id,omitempty"`
	PersistenceCookieName string                   `json:"persistence_cookie_name,omitempty"`
	ConnectionLimit       int                      `json:"connection_limit,omitempty"`
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
		Status:                "",
		StatusDescription:     "",
		Protocol:              MakeLoadbalancerProtocolType(),
		PersistenceType:       MakeSessionPersistenceType(),
		AdminState:            false,
		ProtocolPort:          0,
		SubnetID:              MakeUuidStringType(),
		PersistenceCookieName: "",
		ConnectionLimit:       0,
		Address:               MakeIpAddressType(),
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
