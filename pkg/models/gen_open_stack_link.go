package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeOpenStackLink makes OpenStackLink
// nolint
func MakeOpenStackLink() *OpenStackLink {
	return &OpenStackLink{
		//TODO(nati): Apply default
		Href: "",
		Rel:  "",
		Type: "",
	}
}

// MakeOpenStackLink makes OpenStackLink
// nolint
func InterfaceToOpenStackLink(i interface{}) *OpenStackLink {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OpenStackLink{
		//TODO(nati): Apply default
		Href: common.InterfaceToString(m["href"]),
		Rel:  common.InterfaceToString(m["rel"]),
		Type: common.InterfaceToString(m["type"]),
	}
}

// MakeOpenStackLinkSlice() makes a slice of OpenStackLink
// nolint
func MakeOpenStackLinkSlice() []*OpenStackLink {
	return []*OpenStackLink{}
}

// InterfaceToOpenStackLinkSlice() makes a slice of OpenStackLink
// nolint
func InterfaceToOpenStackLinkSlice(i interface{}) []*OpenStackLink {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OpenStackLink{}
	for _, item := range list {
		result = append(result, InterfaceToOpenStackLink(item))
	}
	return result
}
