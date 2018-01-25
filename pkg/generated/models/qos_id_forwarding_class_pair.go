package models

// QosIdForwardingClassPair

// QosIdForwardingClassPair
//proteus:generate
type QosIdForwardingClassPair struct {
	Key               int               `json:"key,omitempty"`
	ForwardingClassID ForwardingClassId `json:"forwarding_class_id,omitempty"`
}

// MakeQosIdForwardingClassPair makes QosIdForwardingClassPair
func MakeQosIdForwardingClassPair() *QosIdForwardingClassPair {
	return &QosIdForwardingClassPair{
		//TODO(nati): Apply default
		Key:               0,
		ForwardingClassID: MakeForwardingClassId(),
	}
}

// MakeQosIdForwardingClassPairSlice() makes a slice of QosIdForwardingClassPair
func MakeQosIdForwardingClassPairSlice() []*QosIdForwardingClassPair {
	return []*QosIdForwardingClassPair{}
}
