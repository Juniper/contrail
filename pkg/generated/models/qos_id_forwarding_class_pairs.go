package models

// QosIdForwardingClassPairs

// QosIdForwardingClassPairs
//proteus:generate
type QosIdForwardingClassPairs struct {
	QosIDForwardingClassPair []*QosIdForwardingClassPair `json:"qos_id_forwarding_class_pair,omitempty"`
}

// MakeQosIdForwardingClassPairs makes QosIdForwardingClassPairs
func MakeQosIdForwardingClassPairs() *QosIdForwardingClassPairs {
	return &QosIdForwardingClassPairs{
		//TODO(nati): Apply default

		QosIDForwardingClassPair: MakeQosIdForwardingClassPairSlice(),
	}
}

// MakeQosIdForwardingClassPairsSlice() makes a slice of QosIdForwardingClassPairs
func MakeQosIdForwardingClassPairsSlice() []*QosIdForwardingClassPairs {
	return []*QosIdForwardingClassPairs{}
}
