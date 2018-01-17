package models

// VirtualNetworkType

import "encoding/json"

// VirtualNetworkType
type VirtualNetworkType struct {
	VxlanNetworkIdentifier VxlanNetworkIdentifierType `json:"vxlan_network_identifier,omitempty"`
	RPF                    RpfModeType                `json:"rpf,omitempty"`
	ForwardingMode         ForwardingModeType         `json:"forwarding_mode,omitempty"`
	AllowTransit           bool                       `json:"allow_transit,omitempty"`
	NetworkID              int                        `json:"network_id,omitempty"`
	MirrorDestination      bool                       `json:"mirror_destination,omitempty"`
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
		ForwardingMode:         MakeForwardingModeType(),
		AllowTransit:           false,
		NetworkID:              0,
		MirrorDestination:      false,
		VxlanNetworkIdentifier: MakeVxlanNetworkIdentifierType(),
		RPF: MakeRpfModeType(),
	}
}

// MakeVirtualNetworkTypeSlice() makes a slice of VirtualNetworkType
func MakeVirtualNetworkTypeSlice() []*VirtualNetworkType {
	return []*VirtualNetworkType{}
}
