package models


// MakeFatFlowProtocols makes FatFlowProtocols
func MakeFatFlowProtocols() *FatFlowProtocols{
    return &FatFlowProtocols{
    //TODO(nati): Apply default
    
            
                FatFlowProtocol:  MakeProtocolTypeSlice(),
            
        
    }
}

// MakeFatFlowProtocolsSlice() makes a slice of FatFlowProtocols
func MakeFatFlowProtocolsSlice() []*FatFlowProtocols {
    return []*FatFlowProtocols{}
}


