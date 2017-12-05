package models

// VirtualNetworkType

import "encoding/json"

type VirtualNetworkType struct {
	VxlanNetworkIdentifier VxlanNetworkIdentifierType `json:"vxlan_network_identifier"`
	RPF                    RpfModeType                `json:"rpf"`
	ForwardingMode         ForwardingModeType         `json:"forwarding_mode"`
	AllowTransit           bool                       `json:"allow_transit"`
	NetworkID              int                        `json:"network_id"`
	MirrorDestination      bool                       `json:"mirror_destination"`
}

func (model *VirtualNetworkType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeVirtualNetworkType() *VirtualNetworkType {
	return &VirtualNetworkType{
		//TODO(nati): Apply default
		VxlanNetworkIdentifier: MakeVxlanNetworkIdentifierType(),
		RPF:               MakeRpfModeType(),
		ForwardingMode:    MakeForwardingModeType(),
		AllowTransit:      false,
		NetworkID:         0,
		MirrorDestination: false,
	}
}

func InterfaceToVirtualNetworkType(iData interface{}) *VirtualNetworkType {
	data := iData.(map[string]interface{})
	return &VirtualNetworkType{
		NetworkID: data["network_id"].(int),

		//{"Title":"","Description":"Not currently in used","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkID","GoType":"int"}
		MirrorDestination: data["mirror_destination"].(bool),

		//{"Title":"","Description":"Flag to mark the virtual network as mirror destination network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MirrorDestination","GoType":"bool"}
		VxlanNetworkIdentifier: InterfaceToVxlanNetworkIdentifierType(data["vxlan_network_identifier"]),

		//{"Title":"","Description":"VxLAN VNI value for this network","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"VxlanNetworkIdentifier","GoType":"VxlanNetworkIdentifierType"}
		RPF: InterfaceToRpfModeType(data["rpf"]),

		//{"Title":"","Description":"Flag used to disable Reverse Path Forwarding(RPF) check for this network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["enable","disable"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RpfModeType","CollectionType":"","Column":"","Item":null,"GoName":"RPF","GoType":"RpfModeType"}
		ForwardingMode: InterfaceToForwardingModeType(data["forwarding_mode"]),

		//{"Title":"","Description":"Packet forwarding mode for this virtual network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["l2_l3","l2","l3"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ForwardingModeType","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingMode","GoType":"ForwardingModeType"}
		AllowTransit: data["allow_transit"].(bool),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AllowTransit","GoType":"bool"}

	}
}

func InterfaceToVirtualNetworkTypeSlice(data interface{}) []*VirtualNetworkType {
	list := data.([]interface{})
	result := MakeVirtualNetworkTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetworkType(item))
	}
	return result
}

func MakeVirtualNetworkTypeSlice() []*VirtualNetworkType {
	return []*VirtualNetworkType{}
}
