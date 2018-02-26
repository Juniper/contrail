package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeAllocationPoolType makes AllocationPoolType
func MakeAllocationPoolType() *AllocationPoolType {
	return &AllocationPoolType{
		//TODO(nati): Apply default
		VrouterSpecificPool: false,
		Start:               "",
		End:                 "",
	}
}

// MakeAllocationPoolType makes AllocationPoolType
func InterfaceToAllocationPoolType(i interface{}) *AllocationPoolType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AllocationPoolType{
		//TODO(nati): Apply default
		VrouterSpecificPool: schema.InterfaceToBool(m["vrouter_specific_pool"]),
		Start:               schema.InterfaceToString(m["start"]),
		End:                 schema.InterfaceToString(m["end"]),
	}
}

// MakeAllocationPoolTypeSlice() makes a slice of AllocationPoolType
func MakeAllocationPoolTypeSlice() []*AllocationPoolType {
	return []*AllocationPoolType{}
}

// InterfaceToAllocationPoolTypeSlice() makes a slice of AllocationPoolType
func InterfaceToAllocationPoolTypeSlice(i interface{}) []*AllocationPoolType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AllocationPoolType{}
	for _, item := range list {
		result = append(result, InterfaceToAllocationPoolType(item))
	}
	return result
}
