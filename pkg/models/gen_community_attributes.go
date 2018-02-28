package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeCommunityAttributes makes CommunityAttributes
// nolint
func MakeCommunityAttributes() *CommunityAttributes {
	return &CommunityAttributes{
	//TODO(nati): Apply default

	}
}

// MakeCommunityAttributes makes CommunityAttributes
// nolint
func InterfaceToCommunityAttributes(i interface{}) *CommunityAttributes {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &CommunityAttributes{
	//TODO(nati): Apply default

	}
}

// MakeCommunityAttributesSlice() makes a slice of CommunityAttributes
// nolint
func MakeCommunityAttributesSlice() []*CommunityAttributes {
	return []*CommunityAttributes{}
}

// InterfaceToCommunityAttributesSlice() makes a slice of CommunityAttributes
// nolint
func InterfaceToCommunityAttributesSlice(i interface{}) []*CommunityAttributes {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*CommunityAttributes{}
	for _, item := range list {
		result = append(result, InterfaceToCommunityAttributes(item))
	}
	return result
}
