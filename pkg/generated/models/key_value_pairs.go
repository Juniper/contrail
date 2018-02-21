package models


// MakeKeyValuePairs makes KeyValuePairs
func MakeKeyValuePairs() *KeyValuePairs{
    return &KeyValuePairs{
    //TODO(nati): Apply default
    
            
                KeyValuePair:  MakeKeyValuePairSlice(),
            
        
    }
}

// MakeKeyValuePairsSlice() makes a slice of KeyValuePairs
func MakeKeyValuePairsSlice() []*KeyValuePairs {
    return []*KeyValuePairs{}
}


