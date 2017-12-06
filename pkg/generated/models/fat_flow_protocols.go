package models

// FatFlowProtocols

import "encoding/json"

// FatFlowProtocols
type FatFlowProtocols struct {
	FatFlowProtocol []*ProtocolType `json:"fat_flow_protocol"`
}

//  parents relation object

// String returns json representation of the object
func (model *FatFlowProtocols) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFatFlowProtocols makes FatFlowProtocols
func MakeFatFlowProtocols() *FatFlowProtocols {
	return &FatFlowProtocols{
		//TODO(nati): Apply default

		FatFlowProtocol: MakeProtocolTypeSlice(),
	}
}

// InterfaceToFatFlowProtocols makes FatFlowProtocols from interface
func InterfaceToFatFlowProtocols(iData interface{}) *FatFlowProtocols {
	data := iData.(map[string]interface{})
	return &FatFlowProtocols{

		FatFlowProtocol: InterfaceToProtocolTypeSlice(data["fat_flow_protocol"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Port","GoType":"int","GoPremitive":true},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ProtocolType","CollectionType":"","Column":"","Item":null,"GoName":"FatFlowProtocol","GoType":"ProtocolType","GoPremitive":false},"GoName":"FatFlowProtocol","GoType":"[]*ProtocolType","GoPremitive":true}

	}
}

// InterfaceToFatFlowProtocolsSlice makes a slice of FatFlowProtocols from interface
func InterfaceToFatFlowProtocolsSlice(data interface{}) []*FatFlowProtocols {
	list := data.([]interface{})
	result := MakeFatFlowProtocolsSlice()
	for _, item := range list {
		result = append(result, InterfaceToFatFlowProtocols(item))
	}
	return result
}

// MakeFatFlowProtocolsSlice() makes a slice of FatFlowProtocols
func MakeFatFlowProtocolsSlice() []*FatFlowProtocols {
	return []*FatFlowProtocols{}
}
