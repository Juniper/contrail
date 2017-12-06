package models

// FlowAgingTimeoutList

import "encoding/json"

// FlowAgingTimeoutList
type FlowAgingTimeoutList struct {
	FlowAgingTimeout []*FlowAgingTimeout `json:"flow_aging_timeout"`
}

//  parents relation object

// String returns json representation of the object
func (model *FlowAgingTimeoutList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFlowAgingTimeoutList makes FlowAgingTimeoutList
func MakeFlowAgingTimeoutList() *FlowAgingTimeoutList {
	return &FlowAgingTimeoutList{
		//TODO(nati): Apply default

		FlowAgingTimeout: MakeFlowAgingTimeoutSlice(),
	}
}

// InterfaceToFlowAgingTimeoutList makes FlowAgingTimeoutList from interface
func InterfaceToFlowAgingTimeoutList(iData interface{}) *FlowAgingTimeoutList {
	data := iData.(map[string]interface{})
	return &FlowAgingTimeoutList{

		FlowAgingTimeout: InterfaceToFlowAgingTimeoutSlice(data["flow_aging_timeout"]),

		//{"Title":"","Description":"List of (ip protocol, port number, timeout in seconds)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Port","GoType":"int","GoPremitive":true},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string","GoPremitive":true},"timeout_in_seconds":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TimeoutInSeconds","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FlowAgingTimeout","CollectionType":"","Column":"","Item":null,"GoName":"FlowAgingTimeout","GoType":"FlowAgingTimeout","GoPremitive":false},"GoName":"FlowAgingTimeout","GoType":"[]*FlowAgingTimeout","GoPremitive":true}

	}
}

// InterfaceToFlowAgingTimeoutListSlice makes a slice of FlowAgingTimeoutList from interface
func InterfaceToFlowAgingTimeoutListSlice(data interface{}) []*FlowAgingTimeoutList {
	list := data.([]interface{})
	result := MakeFlowAgingTimeoutListSlice()
	for _, item := range list {
		result = append(result, InterfaceToFlowAgingTimeoutList(item))
	}
	return result
}

// MakeFlowAgingTimeoutListSlice() makes a slice of FlowAgingTimeoutList
func MakeFlowAgingTimeoutListSlice() []*FlowAgingTimeoutList {
	return []*FlowAgingTimeoutList{}
}
