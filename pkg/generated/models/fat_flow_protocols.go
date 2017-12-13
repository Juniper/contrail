package models

// FatFlowProtocols

import "encoding/json"

// FatFlowProtocols
type FatFlowProtocols struct {
	FatFlowProtocol []*ProtocolType `json:"fat_flow_protocol"`
}

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

		//{"type":"array","item":{"type":"object","properties":{"port":{"type":"integer"},"protocol":{"type":"string"}}}}

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
