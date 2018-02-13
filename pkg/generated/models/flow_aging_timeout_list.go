package models

// FlowAgingTimeoutList

import "encoding/json"

// FlowAgingTimeoutList
//proteus:generate
type FlowAgingTimeoutList struct {
	FlowAgingTimeout []*FlowAgingTimeout `json:"flow_aging_timeout,omitempty"`
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

// MakeFlowAgingTimeoutListSlice() makes a slice of FlowAgingTimeoutList
func MakeFlowAgingTimeoutListSlice() []*FlowAgingTimeoutList {
	return []*FlowAgingTimeoutList{}
}
