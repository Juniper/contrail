package models

// QosIdForwardingClassPairs

import "encoding/json"

// QosIdForwardingClassPairs
type QosIdForwardingClassPairs struct {
	QosIDForwardingClassPair []*QosIdForwardingClassPair `json:"qos_id_forwarding_class_pair,omitempty"`
}

// String returns json representation of the object
func (model *QosIdForwardingClassPairs) String() string {
	b, _ := json.Marshal(model)
	return string(b)
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
