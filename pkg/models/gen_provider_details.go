package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeProviderDetails makes ProviderDetails
// nolint
func MakeProviderDetails() *ProviderDetails {
	return &ProviderDetails{
		//TODO(nati): Apply default
		SegmentationID:  0,
		PhysicalNetwork: "",
	}
}

// MakeProviderDetails makes ProviderDetails
// nolint
func InterfaceToProviderDetails(i interface{}) *ProviderDetails {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ProviderDetails{
		//TODO(nati): Apply default
		SegmentationID:  common.InterfaceToInt64(m["segmentation_id"]),
		PhysicalNetwork: common.InterfaceToString(m["physical_network"]),
	}
}

// MakeProviderDetailsSlice() makes a slice of ProviderDetails
// nolint
func MakeProviderDetailsSlice() []*ProviderDetails {
	return []*ProviderDetails{}
}

// InterfaceToProviderDetailsSlice() makes a slice of ProviderDetails
// nolint
func InterfaceToProviderDetailsSlice(i interface{}) []*ProviderDetails {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ProviderDetails{}
	for _, item := range list {
		result = append(result, InterfaceToProviderDetails(item))
	}
	return result
}
