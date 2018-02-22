package models


// MakeKeyValuePair makes KeyValuePair
func MakeKeyValuePair() *KeyValuePair{
    return &KeyValuePair{
    //TODO(nati): Apply default
    Value: "",
        Key: "",
        
    }
}

// MakeKeyValuePairSlice() makes a slice of KeyValuePair
func MakeKeyValuePairSlice() []*KeyValuePair {
    return []*KeyValuePair{}
}


