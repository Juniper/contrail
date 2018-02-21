package models


// MakeBGPaaServiceParametersType makes BGPaaServiceParametersType
func MakeBGPaaServiceParametersType() *BGPaaServiceParametersType{
    return &BGPaaServiceParametersType{
    //TODO(nati): Apply default
    PortStart: 0,
        PortEnd: 0,
        
    }
}

// MakeBGPaaServiceParametersTypeSlice() makes a slice of BGPaaServiceParametersType
func MakeBGPaaServiceParametersTypeSlice() []*BGPaaServiceParametersType {
    return []*BGPaaServiceParametersType{}
}


