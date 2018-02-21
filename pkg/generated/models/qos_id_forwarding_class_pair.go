package models


// MakeQosIdForwardingClassPair makes QosIdForwardingClassPair
func MakeQosIdForwardingClassPair() *QosIdForwardingClassPair{
    return &QosIdForwardingClassPair{
    //TODO(nati): Apply default
    Key: 0,
        ForwardingClassID: 0,
        
    }
}

// MakeQosIdForwardingClassPairSlice() makes a slice of QosIdForwardingClassPair
func MakeQosIdForwardingClassPairSlice() []*QosIdForwardingClassPair {
    return []*QosIdForwardingClassPair{}
}


