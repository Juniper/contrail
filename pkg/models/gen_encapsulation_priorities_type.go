package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeEncapsulationPrioritiesType makes EncapsulationPrioritiesType
// nolint
func MakeEncapsulationPrioritiesType() *EncapsulationPrioritiesType {
	return &EncapsulationPrioritiesType{
	//TODO(nati): Apply default

	}
}

// MakeEncapsulationPrioritiesType makes EncapsulationPrioritiesType
// nolint
func InterfaceToEncapsulationPrioritiesType(i interface{}) *EncapsulationPrioritiesType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &EncapsulationPrioritiesType{
	//TODO(nati): Apply default

	}
}

// MakeEncapsulationPrioritiesTypeSlice() makes a slice of EncapsulationPrioritiesType
// nolint
func MakeEncapsulationPrioritiesTypeSlice() []*EncapsulationPrioritiesType {
	return []*EncapsulationPrioritiesType{}
}

// InterfaceToEncapsulationPrioritiesTypeSlice() makes a slice of EncapsulationPrioritiesType
// nolint
func InterfaceToEncapsulationPrioritiesTypeSlice(i interface{}) []*EncapsulationPrioritiesType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*EncapsulationPrioritiesType{}
	for _, item := range list {
		result = append(result, InterfaceToEncapsulationPrioritiesType(item))
	}
	return result
}
