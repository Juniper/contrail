package models

// FlowAgingTimeout

// FlowAgingTimeout
//proteus:generate
type FlowAgingTimeout struct {
	TimeoutInSeconds int    `json:"timeout_in_seconds,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	Port             int    `json:"port,omitempty"`
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
