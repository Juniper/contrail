package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFlowAgingTimeout makes FlowAgingTimeout
// nolint
func MakeFlowAgingTimeout() *FlowAgingTimeout {
	return &FlowAgingTimeout{
		//TODO(nati): Apply default
		TimeoutInSeconds: 0,
		Protocol:         "",
		Port:             0,
	}
}

// MakeFlowAgingTimeout makes FlowAgingTimeout
// nolint
func InterfaceToFlowAgingTimeout(i interface{}) *FlowAgingTimeout {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FlowAgingTimeout{
		//TODO(nati): Apply default
		TimeoutInSeconds: common.InterfaceToInt64(m["timeout_in_seconds"]),
		Protocol:         common.InterfaceToString(m["protocol"]),
		Port:             common.InterfaceToInt64(m["port"]),
	}
}

// MakeFlowAgingTimeoutSlice() makes a slice of FlowAgingTimeout
// nolint
func MakeFlowAgingTimeoutSlice() []*FlowAgingTimeout {
	return []*FlowAgingTimeout{}
}

// InterfaceToFlowAgingTimeoutSlice() makes a slice of FlowAgingTimeout
// nolint
func InterfaceToFlowAgingTimeoutSlice(i interface{}) []*FlowAgingTimeout {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FlowAgingTimeout{}
	for _, item := range list {
		result = append(result, InterfaceToFlowAgingTimeout(item))
	}
	return result
}
