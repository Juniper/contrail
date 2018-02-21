package models


// MakeAllowedAddressPairs makes AllowedAddressPairs
func MakeAllowedAddressPairs() *AllowedAddressPairs{
    return &AllowedAddressPairs{
    //TODO(nati): Apply default
    
            
                AllowedAddressPair:  MakeAllowedAddressPairSlice(),
            
        
    }
}

// MakeAllowedAddressPairsSlice() makes a slice of AllowedAddressPairs
func MakeAllowedAddressPairsSlice() []*AllowedAddressPairs {
    return []*AllowedAddressPairs{}
}


