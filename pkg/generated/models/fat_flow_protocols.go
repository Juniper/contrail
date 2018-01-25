package models
// FatFlowProtocols



import "encoding/json"

// FatFlowProtocols 
//proteus:generate
type FatFlowProtocols struct {

    FatFlowProtocol []*ProtocolType `json:"fat_flow_protocol,omitempty"`


}



// String returns json representation of the object
func (model *FatFlowProtocols) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

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
