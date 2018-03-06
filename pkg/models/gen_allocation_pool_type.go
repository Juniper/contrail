package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAllocationPoolType makes AllocationPoolType
// nolint
func MakeAllocationPoolType() *AllocationPoolType {
	return &AllocationPoolType{
		//TODO(nati): Apply default
		VrouterSpecificPool: false,
		Start:               "",
		End:                 "",
	}
}

// MakeAllocationPoolType makes AllocationPoolType
// nolint
func InterfaceToAllocationPoolType(i interface{}) *AllocationPoolType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AllocationPoolType{
		//TODO(nati): Apply default
		VrouterSpecificPool: common.InterfaceToBool(m["vrouter_specific_pool"]),
		Start:               common.InterfaceToString(m["start"]),
		End:                 common.InterfaceToString(m["end"]),
	}
}

// MakeAllocationPoolTypeSlice() makes a slice of AllocationPoolType
// nolint
func MakeAllocationPoolTypeSlice() []*AllocationPoolType {
	return []*AllocationPoolType{}
}

// InterfaceToAllocationPoolTypeSlice() makes a slice of AllocationPoolType
// nolint
func InterfaceToAllocationPoolTypeSlice(i interface{}) []*AllocationPoolType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AllocationPoolType{}
	for _, item := range list {
		result = append(result, InterfaceToAllocationPoolType(item))
	}
	return result
}
