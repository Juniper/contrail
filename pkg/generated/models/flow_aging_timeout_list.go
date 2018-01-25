package models

// FlowAgingTimeoutList

// FlowAgingTimeoutList
//proteus:generate
type FlowAgingTimeoutList struct {
	FlowAgingTimeout []*FlowAgingTimeout `json:"flow_aging_timeout,omitempty"`
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
