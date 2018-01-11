package models

// FlowAgingTimeout

import "encoding/json"

// FlowAgingTimeout
type FlowAgingTimeout struct {
	Port             int    `json:"port"`
	TimeoutInSeconds int    `json:"timeout_in_seconds"`
	Protocol         string `json:"protocol"`
}

// String returns json representation of the object
func (model *FlowAgingTimeout) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFlowAgingTimeout makes FlowAgingTimeout
func MakeFlowAgingTimeout() *FlowAgingTimeout {
	return &FlowAgingTimeout{
		//TODO(nati): Apply default
		TimeoutInSeconds: 0,
		Protocol:         "",
		Port:             0,
	}
}

// MakeFlowAgingTimeoutSlice() makes a slice of FlowAgingTimeout
func MakeFlowAgingTimeoutSlice() []*FlowAgingTimeout {
	return []*FlowAgingTimeout{}
}
