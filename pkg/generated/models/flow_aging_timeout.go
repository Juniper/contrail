package models

// FlowAgingTimeout

import "encoding/json"

// FlowAgingTimeout
type FlowAgingTimeout struct {
	Protocol         string `json:"protocol,omitempty"`
	Port             int    `json:"port,omitempty"`
	TimeoutInSeconds int    `json:"timeout_in_seconds,omitempty"`
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
		Port:             0,
		TimeoutInSeconds: 0,
		Protocol:         "",
	}
}

// MakeFlowAgingTimeoutSlice() makes a slice of FlowAgingTimeout
func MakeFlowAgingTimeoutSlice() []*FlowAgingTimeout {
	return []*FlowAgingTimeout{}
}
