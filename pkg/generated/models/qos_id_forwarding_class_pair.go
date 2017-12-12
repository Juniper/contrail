package models

// QosIdForwardingClassPair

import "encoding/json"

// QosIdForwardingClassPair
type QosIdForwardingClassPair struct {
	Key               int               `json:"key"`
	ForwardingClassID ForwardingClassId `json:"forwarding_class_id"`
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

// InterfaceToQosIdForwardingClassPair makes QosIdForwardingClassPair from interface
func InterfaceToQosIdForwardingClassPair(iData interface{}) *QosIdForwardingClassPair {
	data := iData.(map[string]interface{})
	return &QosIdForwardingClassPair{
		Key: data["key"].(int),

		//{"description":"QoS bit value (DSCP or Vlan priority or EXP bit value","type":"integer"}
		ForwardingClassID: InterfaceToForwardingClassId(data["forwarding_class_id"]),

		//{"default":"0","type":"integer","minimum":0,"maximum":255}

	}
}

// InterfaceToQosIdForwardingClassPairSlice makes a slice of QosIdForwardingClassPair from interface
func InterfaceToQosIdForwardingClassPairSlice(data interface{}) []*QosIdForwardingClassPair {
	list := data.([]interface{})
	result := MakeQosIdForwardingClassPairSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosIdForwardingClassPair(item))
	}
	return result
}

// MakeQosIdForwardingClassPairSlice() makes a slice of QosIdForwardingClassPair
func MakeQosIdForwardingClassPairSlice() []*QosIdForwardingClassPair {
	return []*QosIdForwardingClassPair{}
}
