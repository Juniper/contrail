package models

// FatFlowProtocols

// FatFlowProtocols
//proteus:generate
type FatFlowProtocols struct {
	FatFlowProtocol []*ProtocolType `json:"fat_flow_protocol,omitempty"`
}

// MakeFatFlowProtocols makes FatFlowProtocols
func MakeFatFlowProtocols() *FatFlowProtocols {
	return &FatFlowProtocols{
		//TODO(nati): Apply default

		FatFlowProtocol: MakeProtocolTypeSlice(),
	}
}

// MakeFatFlowProtocolsSlice() makes a slice of FatFlowProtocols
func MakeFatFlowProtocolsSlice() []*FatFlowProtocols {
	return []*FatFlowProtocols{}
}
