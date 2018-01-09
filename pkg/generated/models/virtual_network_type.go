package models

// VirtualNetworkType

import "encoding/json"

// VirtualNetworkType
type VirtualNetworkType struct {
	MirrorDestination      bool                       `json:"mirror_destination"`
	VxlanNetworkIdentifier VxlanNetworkIdentifierType `json:"vxlan_network_identifier"`
	RPF                    RpfModeType                `json:"rpf"`
	ForwardingMode         ForwardingModeType         `json:"forwarding_mode"`
	AllowTransit           bool                       `json:"allow_transit"`
	NetworkID              int                        `json:"network_id"`
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

// InterfaceToVirtualNetworkType makes VirtualNetworkType from interface
func InterfaceToVirtualNetworkType(iData interface{}) *VirtualNetworkType {
	data := iData.(map[string]interface{})
	return &VirtualNetworkType{
		ForwardingMode: InterfaceToForwardingModeType(data["forwarding_mode"]),

		//{"description":"Packet forwarding mode for this virtual network","type":"string","enum":["l2_l3","l2","l3"]}
		AllowTransit: data["allow_transit"].(bool),

		//{"type":"boolean"}
		NetworkID: data["network_id"].(int),

		//{"description":"Not currently in used","type":"integer"}
		MirrorDestination: data["mirror_destination"].(bool),

		//{"description":"Flag to mark the virtual network as mirror destination network","type":"boolean"}
		VxlanNetworkIdentifier: InterfaceToVxlanNetworkIdentifierType(data["vxlan_network_identifier"]),

		//{"description":"VxLAN VNI value for this network","type":"integer","minimum":1,"maximum":16777215}
		RPF: InterfaceToRpfModeType(data["rpf"]),

		//{"description":"Flag used to disable Reverse Path Forwarding(RPF) check for this network","type":"string","enum":["enable","disable"]}

	}
}

// InterfaceToVirtualNetworkTypeSlice makes a slice of VirtualNetworkType from interface
func InterfaceToVirtualNetworkTypeSlice(data interface{}) []*VirtualNetworkType {
	list := data.([]interface{})
	result := MakeVirtualNetworkTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetworkType(item))
	}
	return result
}

// MakeVirtualNetworkTypeSlice() makes a slice of VirtualNetworkType
func MakeVirtualNetworkTypeSlice() []*VirtualNetworkType {
	return []*VirtualNetworkType{}
}
