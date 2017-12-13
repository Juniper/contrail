package models

// QosIdForwardingClassPairs

import "encoding/json"

// QosIdForwardingClassPairs
type QosIdForwardingClassPairs struct {
	QosIDForwardingClassPair []*QosIdForwardingClassPair `json:"qos_id_forwarding_class_pair"`
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

// InterfaceToQosIdForwardingClassPairs makes QosIdForwardingClassPairs from interface
func InterfaceToQosIdForwardingClassPairs(iData interface{}) *QosIdForwardingClassPairs {
	data := iData.(map[string]interface{})
	return &QosIdForwardingClassPairs{

		QosIDForwardingClassPair: InterfaceToQosIdForwardingClassPairSlice(data["qos_id_forwarding_class_pair"]),

		//{"type":"array","item":{"type":"object","properties":{"forwarding_class_id":{"default":"0","type":"integer","minimum":0,"maximum":255},"key":{"type":"integer"}}}}

	}
}

// InterfaceToQosIdForwardingClassPairsSlice makes a slice of QosIdForwardingClassPairs from interface
func InterfaceToQosIdForwardingClassPairsSlice(data interface{}) []*QosIdForwardingClassPairs {
	list := data.([]interface{})
	result := MakeQosIdForwardingClassPairsSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosIdForwardingClassPairs(item))
	}
	return result
}

// MakeQosIdForwardingClassPairsSlice() makes a slice of QosIdForwardingClassPairs
func MakeQosIdForwardingClassPairsSlice() []*QosIdForwardingClassPairs {
	return []*QosIdForwardingClassPairs{}
}
