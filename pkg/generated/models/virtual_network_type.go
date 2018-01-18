package models

// VirtualNetworkType

import "encoding/json"

// VirtualNetworkType
type VirtualNetworkType struct {
	AllowTransit           bool                       `json:"allow_transit"`
	NetworkID              int                        `json:"network_id,omitempty"`
	MirrorDestination      bool                       `json:"mirror_destination"`
	VxlanNetworkIdentifier VxlanNetworkIdentifierType `json:"vxlan_network_identifier,omitempty"`
	RPF                    RpfModeType                `json:"rpf,omitempty"`
	ForwardingMode         ForwardingModeType         `json:"forwarding_mode,omitempty"`
}

// String returns json representation of the object
func (model *VirtualNetworkType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualNetworkType makes VirtualNetworkType
func MakeVirtualNetworkType() *VirtualNetworkType {
	return &VirtualNetworkType{
		//TODO(nati): Apply default
		NetworkID:              0,
		MirrorDestination:      false,
		VxlanNetworkIdentifier: MakeVxlanNetworkIdentifierType(),
		RPF:            MakeRpfModeType(),
		ForwardingMode: MakeForwardingModeType(),
		AllowTransit:   false,
	}
}

// MakeVirtualNetworkTypeSlice() makes a slice of VirtualNetworkType
func MakeVirtualNetworkTypeSlice() []*VirtualNetworkType {
	return []*VirtualNetworkType{}
}
