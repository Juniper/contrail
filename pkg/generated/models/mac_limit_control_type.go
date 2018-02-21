package models


// MakeMACLimitControlType makes MACLimitControlType
func MakeMACLimitControlType() *MACLimitControlType{
    return &MACLimitControlType{
    //TODO(nati): Apply default
    MacLimit: 0,
        MacLimitAction: "",
        
    }
}

// MakeMACLimitControlTypeSlice() makes a slice of MACLimitControlType
func MakeMACLimitControlTypeSlice() []*MACLimitControlType {
    return []*MACLimitControlType{}
}


