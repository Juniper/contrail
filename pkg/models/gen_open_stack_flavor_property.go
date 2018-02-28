package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeOpenStackFlavorProperty makes OpenStackFlavorProperty
// nolint
func MakeOpenStackFlavorProperty() *OpenStackFlavorProperty {
	return &OpenStackFlavorProperty{
		//TODO(nati): Apply default
		ID:    "",
		Links: MakeOpenStackLink(),
	}
}

// MakeOpenStackFlavorProperty makes OpenStackFlavorProperty
// nolint
func InterfaceToOpenStackFlavorProperty(i interface{}) *OpenStackFlavorProperty {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OpenStackFlavorProperty{
		//TODO(nati): Apply default
		ID:    common.InterfaceToString(m["id"]),
		Links: InterfaceToOpenStackLink(m["links"]),
	}
}

// MakeOpenStackFlavorPropertySlice() makes a slice of OpenStackFlavorProperty
// nolint
func MakeOpenStackFlavorPropertySlice() []*OpenStackFlavorProperty {
	return []*OpenStackFlavorProperty{}
}

// InterfaceToOpenStackFlavorPropertySlice() makes a slice of OpenStackFlavorProperty
// nolint
func InterfaceToOpenStackFlavorPropertySlice(i interface{}) []*OpenStackFlavorProperty {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OpenStackFlavorProperty{}
	for _, item := range list {
		result = append(result, InterfaceToOpenStackFlavorProperty(item))
	}
	return result
}
