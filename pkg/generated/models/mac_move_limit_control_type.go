package models


// MakeMACMoveLimitControlType makes MACMoveLimitControlType
func MakeMACMoveLimitControlType() *MACMoveLimitControlType{
    return &MACMoveLimitControlType{
    //TODO(nati): Apply default
    MacMoveTimeWindow: 0,
        MacMoveLimit: 0,
        MacMoveLimitAction: "",
        
    }
}

// MakeMACMoveLimitControlTypeSlice() makes a slice of MACMoveLimitControlType
func MakeMACMoveLimitControlTypeSlice() []*MACMoveLimitControlType {
    return []*MACMoveLimitControlType{}
}


