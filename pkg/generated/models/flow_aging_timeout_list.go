package models

// FlowAgingTimeoutList

import "encoding/json"

type FlowAgingTimeoutList struct {
	FlowAgingTimeout []*FlowAgingTimeout `json:"flow_aging_timeout"`
}

func (model *FlowAgingTimeoutList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFlowAgingTimeoutList() *FlowAgingTimeoutList {
	return &FlowAgingTimeoutList{
		//TODO(nati): Apply default

		FlowAgingTimeout: MakeFlowAgingTimeoutSlice(),
	}
}

func InterfaceToFlowAgingTimeoutList(iData interface{}) *FlowAgingTimeoutList {
	data := iData.(map[string]interface{})
	return &FlowAgingTimeoutList{

		FlowAgingTimeout: InterfaceToFlowAgingTimeoutSlice(data["flow_aging_timeout"]),

		//{"Title":"","Description":"List of (ip protocol, port number, timeout in seconds)","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Port","GoType":"int"},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"},"timeout_in_seconds":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TimeoutInSeconds","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FlowAgingTimeout","CollectionType":"","Column":"","Item":null,"GoName":"FlowAgingTimeout","GoType":"FlowAgingTimeout"},"GoName":"FlowAgingTimeout","GoType":"[]*FlowAgingTimeout"}

	}
}

func InterfaceToFlowAgingTimeoutListSlice(data interface{}) []*FlowAgingTimeoutList {
	list := data.([]interface{})
	result := MakeFlowAgingTimeoutListSlice()
	for _, item := range list {
		result = append(result, InterfaceToFlowAgingTimeoutList(item))
	}
	return result
}

func MakeFlowAgingTimeoutListSlice() []*FlowAgingTimeoutList {
	return []*FlowAgingTimeoutList{}
}
