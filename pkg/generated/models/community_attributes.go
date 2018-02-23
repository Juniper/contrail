package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeCommunityAttributes makes CommunityAttributes
func MakeCommunityAttributes() *CommunityAttributes {
	return &CommunityAttributes{
		//TODO(nati): Apply default

	}
}

// MakeCommunityAttributes makes CommunityAttributes
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
func MakeCommunityAttributesSlice() []*CommunityAttributes {
	return []*CommunityAttributes{}
}

// InterfaceToCommunityAttributesSlice() makes a slice of CommunityAttributes
func InterfaceToCommunityAttributesSlice(i interface{}) []*CommunityAttributes {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*CommunityAttributes{}
	for _, item := range list {
		result = append(result, InterfaceToCommunityAttributes(item))
	}
	return result
}
