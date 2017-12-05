package models

// QosIdForwardingClassPairs

import "encoding/json"

type QosIdForwardingClassPairs struct {
	QosIDForwardingClassPair []*QosIdForwardingClassPair `json:"qos_id_forwarding_class_pair"`
}

func (model *QosIdForwardingClassPairs) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeQosIdForwardingClassPairs() *QosIdForwardingClassPairs {
	return &QosIdForwardingClassPairs{
		//TODO(nati): Apply default

		QosIDForwardingClassPair: MakeQosIdForwardingClassPairSlice(),
	}
}

func InterfaceToQosIdForwardingClassPairs(iData interface{}) *QosIdForwardingClassPairs {
	data := iData.(map[string]interface{})
	return &QosIdForwardingClassPairs{

		QosIDForwardingClassPair: InterfaceToQosIdForwardingClassPairSlice(data["qos_id_forwarding_class_pair"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"forwarding_class_id":{"Title":"","Description":"","SQL":"","Default":"0","Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingClassID","GoType":"ForwardingClassId"},"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/QosIdForwardingClassPair","CollectionType":"","Column":"","Item":null,"GoName":"QosIDForwardingClassPair","GoType":"QosIdForwardingClassPair"},"GoName":"QosIDForwardingClassPair","GoType":"[]*QosIdForwardingClassPair"}

	}
}

func InterfaceToQosIdForwardingClassPairsSlice(data interface{}) []*QosIdForwardingClassPairs {
	list := data.([]interface{})
	result := MakeQosIdForwardingClassPairsSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosIdForwardingClassPairs(item))
	}
	return result
}

func MakeQosIdForwardingClassPairsSlice() []*QosIdForwardingClassPairs {
	return []*QosIdForwardingClassPairs{}
}
