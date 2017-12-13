package models

// FlowAgingTimeoutList

import "encoding/json"

// FlowAgingTimeoutList
type FlowAgingTimeoutList struct {
	FlowAgingTimeout []*FlowAgingTimeout `json:"flow_aging_timeout"`
}

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

		//{"description":"List of (ip protocol, port number, timeout in seconds)","type":"array","item":{"type":"object","properties":{"port":{"type":"integer"},"protocol":{"type":"string"},"timeout_in_seconds":{"type":"integer"}}}}

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
