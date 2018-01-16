package models

// QosIdForwardingClassPair

import "encoding/json"

// QosIdForwardingClassPair
type QosIdForwardingClassPair struct {
	ForwardingClassID ForwardingClassId `json:"forwarding_class_id,omitempty"`
	Key               int               `json:"key,omitempty"`
}

// String returns json representation of the object
func (model *QosIdForwardingClassPair) String() string {
	b, _ := json.Marshal(model)
	return string(b)
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
