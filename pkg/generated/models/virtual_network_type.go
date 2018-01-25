package models

// VirtualNetworkType

// VirtualNetworkType
//proteus:generate
type VirtualNetworkType struct {
	ForwardingMode         ForwardingModeType         `json:"forwarding_mode,omitempty"`
	AllowTransit           bool                       `json:"allow_transit"`
	NetworkID              int                        `json:"network_id,omitempty"`
	MirrorDestination      bool                       `json:"mirror_destination"`
	VxlanNetworkIdentifier VxlanNetworkIdentifierType `json:"vxlan_network_identifier,omitempty"`
	RPF                    RpfModeType                `json:"rpf,omitempty"`
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
