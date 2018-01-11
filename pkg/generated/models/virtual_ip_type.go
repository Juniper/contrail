package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	Protocol              LoadbalancerProtocolType `json:"protocol"`
	SubnetID              UuidStringType           `json:"subnet_id"`
	PersistenceCookieName string                   `json:"persistence_cookie_name"`
	ConnectionLimit       int                      `json:"connection_limit"`
	ProtocolPort          int                      `json:"protocol_port"`
	Status                string                   `json:"status"`
	StatusDescription     string                   `json:"status_description"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type"`
	AdminState            bool                     `json:"admin_state"`
	Address               IpAddressType            `json:"address"`
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
		PersistenceType:       MakeSessionPersistenceType(),
		AdminState:            false,
		Address:               MakeIpAddressType(),
		Protocol:              MakeLoadbalancerProtocolType(),
		SubnetID:              MakeUuidStringType(),
		PersistenceCookieName: "",
		ConnectionLimit:       0,
		ProtocolPort:          0,
	}
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
