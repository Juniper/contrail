package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFlowAgingTimeoutList makes FlowAgingTimeoutList
// nolint
func MakeFlowAgingTimeoutList() *FlowAgingTimeoutList {
	return &FlowAgingTimeoutList{
		//TODO(nati): Apply default

		FlowAgingTimeout: MakeFlowAgingTimeoutSlice(),
	}
}

// MakeFlowAgingTimeoutList makes FlowAgingTimeoutList
// nolint
func InterfaceToFlowAgingTimeoutList(i interface{}) *FlowAgingTimeoutList {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FlowAgingTimeoutList{
		//TODO(nati): Apply default

		FlowAgingTimeout: InterfaceToFlowAgingTimeoutSlice(m["flow_aging_timeout"]),
	}
}

// MakeFlowAgingTimeoutListSlice() makes a slice of FlowAgingTimeoutList
// nolint
func MakeFlowAgingTimeoutListSlice() []*FlowAgingTimeoutList {
	return []*FlowAgingTimeoutList{}
}

// InterfaceToFlowAgingTimeoutListSlice() makes a slice of FlowAgingTimeoutList
// nolint
func InterfaceToFlowAgingTimeoutListSlice(i interface{}) []*FlowAgingTimeoutList {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FlowAgingTimeoutList{}
	for _, item := range list {
		result = append(result, InterfaceToFlowAgingTimeoutList(item))
	}
	return result
}
