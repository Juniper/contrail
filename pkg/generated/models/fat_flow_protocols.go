package models

// FatFlowProtocols

import "encoding/json"

type FatFlowProtocols struct {
	FatFlowProtocol []*ProtocolType `json:"fat_flow_protocol"`
}

func (model *FatFlowProtocols) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFatFlowProtocols() *FatFlowProtocols {
	return &FatFlowProtocols{
		//TODO(nati): Apply default

		FatFlowProtocol: MakeProtocolTypeSlice(),
	}
}

func InterfaceToFatFlowProtocols(iData interface{}) *FatFlowProtocols {
	data := iData.(map[string]interface{})
	return &FatFlowProtocols{

		FatFlowProtocol: InterfaceToProtocolTypeSlice(data["fat_flow_protocol"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Port","GoType":"int"},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ProtocolType","CollectionType":"","Column":"","Item":null,"GoName":"FatFlowProtocol","GoType":"ProtocolType"},"GoName":"FatFlowProtocol","GoType":"[]*ProtocolType"}

	}
}

func InterfaceToFatFlowProtocolsSlice(data interface{}) []*FatFlowProtocols {
	list := data.([]interface{})
	result := MakeFatFlowProtocolsSlice()
	for _, item := range list {
		result = append(result, InterfaceToFatFlowProtocols(item))
	}
	return result
}

func MakeFatFlowProtocolsSlice() []*FatFlowProtocols {
	return []*FatFlowProtocols{}
}
