package models

// QosIdForwardingClassPair

import "encoding/json"

type QosIdForwardingClassPair struct {
	Key               int               `json:"key"`
	ForwardingClassID ForwardingClassId `json:"forwarding_class_id"`
}

func (model *QosIdForwardingClassPair) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeQosIdForwardingClassPair() *QosIdForwardingClassPair {
	return &QosIdForwardingClassPair{
		//TODO(nati): Apply default
		Key:               0,
		ForwardingClassID: MakeForwardingClassId(),
	}
}

func InterfaceToQosIdForwardingClassPair(iData interface{}) *QosIdForwardingClassPair {
	data := iData.(map[string]interface{})
	return &QosIdForwardingClassPair{
		Key: data["key"].(int),

		//{"Title":"","Description":"QoS bit value (DSCP or Vlan priority or EXP bit value","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"int"}
		ForwardingClassID: InterfaceToForwardingClassId(data["forwarding_class_id"]),

		//{"Title":"","Description":"","SQL":"","Default":"0","Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":255,"Ref":"types.json#/definitions/ForwardingClassId","CollectionType":"","Column":"","Item":null,"GoName":"ForwardingClassID","GoType":"ForwardingClassId"}

	}
}

func InterfaceToQosIdForwardingClassPairSlice(data interface{}) []*QosIdForwardingClassPair {
	list := data.([]interface{})
	result := MakeQosIdForwardingClassPairSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosIdForwardingClassPair(item))
	}
	return result
}

func MakeQosIdForwardingClassPairSlice() []*QosIdForwardingClassPair {
	return []*QosIdForwardingClassPair{}
}
