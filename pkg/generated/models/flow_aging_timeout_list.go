package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeFlowAgingTimeoutList makes FlowAgingTimeoutList
func MakeFlowAgingTimeoutList() *FlowAgingTimeoutList {
	return &FlowAgingTimeoutList{
		//TODO(nati): Apply default

		FlowAgingTimeout: MakeFlowAgingTimeoutSlice(),
	}
}

// MakeFlowAgingTimeoutList makes FlowAgingTimeoutList
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
func MakeFlowAgingTimeoutListSlice() []*FlowAgingTimeoutList {
	return []*FlowAgingTimeoutList{}
}

// InterfaceToFlowAgingTimeoutListSlice() makes a slice of FlowAgingTimeoutList
func InterfaceToFlowAgingTimeoutListSlice(i interface{}) []*FlowAgingTimeoutList {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FlowAgingTimeoutList{}
	for _, item := range list {
		result = append(result, InterfaceToFlowAgingTimeoutList(item))
	}
	return result
}
