package models

// VirtualNetworkType

import "encoding/json"

// VirtualNetworkType
type VirtualNetworkType struct {
	RPF                    RpfModeType                `json:"rpf"`
	ForwardingMode         ForwardingModeType         `json:"forwarding_mode"`
	AllowTransit           bool                       `json:"allow_transit"`
	NetworkID              int                        `json:"network_id"`
	MirrorDestination      bool                       `json:"mirror_destination"`
	VxlanNetworkIdentifier VxlanNetworkIdentifierType `json:"vxlan_network_identifier"`
}

//  parents relation object

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

		//{"Title":"","Description":"Packet forwarding mode for this virtual network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["l2_l3","l2","l3"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ForwardingModeType","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingMode","GoType":"ForwardingModeType","GoPremitive":false}
		AllowTransit: data["allow_transit"].(bool),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AllowTransit","GoType":"bool","GoPremitive":true}
		NetworkID: data["network_id"].(int),

		//{"Title":"","Description":"Not currently in used","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkID","GoType":"int","GoPremitive":true}
		MirrorDestination: data["mirror_destination"].(bool),

		//{"Title":"","Description":"Flag to mark the virtual network as mirror destination network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MirrorDestination","GoType":"bool","GoPremitive":true}
		VxlanNetworkIdentifier: InterfaceToVxlanNetworkIdentifierType(data["vxlan_network_identifier"]),

		//{"Title":"","Description":"VxLAN VNI value for this network","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"VxlanNetworkIdentifier","GoType":"VxlanNetworkIdentifierType","GoPremitive":false}
		RPF: InterfaceToRpfModeType(data["rpf"]),

		//{"Title":"","Description":"Flag used to disable Reverse Path Forwarding(RPF) check for this network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["enable","disable"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RpfModeType","CollectionType":"","Column":"","Item":null,"GoName":"RPF","GoType":"RpfModeType","GoPremitive":false}

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
