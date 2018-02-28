package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeOpenStackImageProperty makes OpenStackImageProperty
// nolint
func MakeOpenStackImageProperty() *OpenStackImageProperty {
	return &OpenStackImageProperty{
		//TODO(nati): Apply default
		ID:    "",
		Links: MakeOpenStackLink(),
	}
}

// MakeOpenStackImageProperty makes OpenStackImageProperty
// nolint
func InterfaceToOpenStackImageProperty(i interface{}) *OpenStackImageProperty {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OpenStackImageProperty{
		//TODO(nati): Apply default
		ID:    common.InterfaceToString(m["id"]),
		Links: InterfaceToOpenStackLink(m["links"]),
	}
}

// MakeOpenStackImagePropertySlice() makes a slice of OpenStackImageProperty
// nolint
func MakeOpenStackImagePropertySlice() []*OpenStackImageProperty {
	return []*OpenStackImageProperty{}
}

// InterfaceToOpenStackImagePropertySlice() makes a slice of OpenStackImageProperty
// nolint
func InterfaceToOpenStackImagePropertySlice(i interface{}) []*OpenStackImageProperty {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OpenStackImageProperty{}
	for _, item := range list {
		result = append(result, InterfaceToOpenStackImageProperty(item))
	}
	return result
}
