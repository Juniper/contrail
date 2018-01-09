package models

// FlowAgingTimeout

import "encoding/json"

// FlowAgingTimeout
type FlowAgingTimeout struct {
	TimeoutInSeconds int    `json:"timeout_in_seconds"`
	Protocol         string `json:"protocol"`
	Port             int    `json:"port"`
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
		Protocol:         "",
		Port:             0,
		TimeoutInSeconds: 0,
	}
}

// InterfaceToFlowAgingTimeout makes FlowAgingTimeout from interface
func InterfaceToFlowAgingTimeout(iData interface{}) *FlowAgingTimeout {
	data := iData.(map[string]interface{})
	return &FlowAgingTimeout{
		TimeoutInSeconds: data["timeout_in_seconds"].(int),

		//{"type":"integer"}
		Protocol: data["protocol"].(string),

		//{"type":"string"}
		Port: data["port"].(int),

		//{"type":"integer"}

	}
}

// InterfaceToFlowAgingTimeoutSlice makes a slice of FlowAgingTimeout from interface
func InterfaceToFlowAgingTimeoutSlice(data interface{}) []*FlowAgingTimeout {
	list := data.([]interface{})
	result := MakeFlowAgingTimeoutSlice()
	for _, item := range list {
		result = append(result, InterfaceToFlowAgingTimeout(item))
	}
	return result
}

// MakeFlowAgingTimeoutSlice() makes a slice of FlowAgingTimeout
func MakeFlowAgingTimeoutSlice() []*FlowAgingTimeout {
	return []*FlowAgingTimeout{}
}
