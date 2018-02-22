package models


// MakeQosIdForwardingClassPairs makes QosIdForwardingClassPairs
func MakeQosIdForwardingClassPairs() *QosIdForwardingClassPairs{
    return &QosIdForwardingClassPairs{
    //TODO(nati): Apply default
    
            
                QosIDForwardingClassPair:  MakeQosIdForwardingClassPairSlice(),
            
        
    }
}

// MakeQosIdForwardingClassPairsSlice() makes a slice of QosIdForwardingClassPairs
func MakeQosIdForwardingClassPairsSlice() []*QosIdForwardingClassPairs {
    return []*QosIdForwardingClassPairs{}
}


