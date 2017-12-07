package models

// VirtualNetworkPolicyType

import "encoding/json"

// VirtualNetworkPolicyType
type VirtualNetworkPolicyType struct {
	Timer    *TimerType    `json:"timer"`
	Sequence *SequenceType `json:"sequence"`
}

//  parents relation object

// String returns json representation of the object
func (model *VirtualNetworkPolicyType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualNetworkPolicyType makes VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyType() *VirtualNetworkPolicyType {
	return &VirtualNetworkPolicyType{
		//TODO(nati): Apply default
		Timer:    MakeTimerType(),
		Sequence: MakeSequenceType(),
	}
}

// InterfaceToVirtualNetworkPolicyType makes VirtualNetworkPolicyType from interface
func InterfaceToVirtualNetworkPolicyType(iData interface{}) *VirtualNetworkPolicyType {
	data := iData.(map[string]interface{})
	return &VirtualNetworkPolicyType{
		Timer: InterfaceToTimerType(data["timer"]),

		//{"Title":"","Description":"Timer to specify when the policy can be active","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"end_time":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"end_time","Item":null,"GoName":"EndTime","GoType":"string","GoPremitive":true},"off_interval":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"off_interval","Item":null,"GoName":"OffInterval","GoType":"string","GoPremitive":true},"on_interval":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"on_interval","Item":null,"GoName":"OnInterval","GoType":"string","GoPremitive":true},"start_time":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"start_time","Item":null,"GoName":"StartTime","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TimerType","CollectionType":"","Column":"","Item":null,"GoName":"Timer","GoType":"TimerType","GoPremitive":false}
		Sequence: InterfaceToSequenceType(data["sequence"]),

		//{"Title":"","Description":"Sequence number to specify order of policy attachment to network","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"major":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"major","Item":null,"GoName":"Major","GoType":"int","GoPremitive":true},"minor":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"minor","Item":null,"GoName":"Minor","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SequenceType","CollectionType":"","Column":"","Item":null,"GoName":"Sequence","GoType":"SequenceType","GoPremitive":false}

	}
}

// InterfaceToVirtualNetworkPolicyTypeSlice makes a slice of VirtualNetworkPolicyType from interface
func InterfaceToVirtualNetworkPolicyTypeSlice(data interface{}) []*VirtualNetworkPolicyType {
	list := data.([]interface{})
	result := MakeVirtualNetworkPolicyTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualNetworkPolicyType(item))
	}
	return result
}

// MakeVirtualNetworkPolicyTypeSlice() makes a slice of VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyTypeSlice() []*VirtualNetworkPolicyType {
	return []*VirtualNetworkPolicyType{}
}
