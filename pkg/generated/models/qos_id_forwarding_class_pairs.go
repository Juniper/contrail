package models

// QosIdForwardingClassPairs

import "encoding/json"

// QosIdForwardingClassPairs
type QosIdForwardingClassPairs struct {
	QosIDForwardingClassPair []*QosIdForwardingClassPair `json:"qos_id_forwarding_class_pair"`
}

//  parents relation object

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

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"forwarding_class_id":{"Title":"","Description":"","SQL":"","Default":"0","Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingClassID","GoType":"ForwardingClassId","GoPremitive":false},"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPair","CollectionType":"","Column":"","Item":null,"GoName":"QosIDForwardingClassPair","GoType":"QosIdForwardingClassPair","GoPremitive":false},"GoName":"QosIDForwardingClassPair","GoType":"[]*QosIdForwardingClassPair","GoPremitive":true}

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
