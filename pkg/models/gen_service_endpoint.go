package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceEndpoint makes ServiceEndpoint
// nolint
func MakeServiceEndpoint() *ServiceEndpoint {
	return &ServiceEndpoint{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeServiceEndpoint makes ServiceEndpoint
// nolint
func InterfaceToServiceEndpoint(i interface{}) *ServiceEndpoint {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceEndpoint{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
	}
}

// MakeServiceEndpointSlice() makes a slice of ServiceEndpoint
// nolint
func MakeServiceEndpointSlice() []*ServiceEndpoint {
	return []*ServiceEndpoint{}
}

// InterfaceToServiceEndpointSlice() makes a slice of ServiceEndpoint
// nolint
func InterfaceToServiceEndpointSlice(i interface{}) []*ServiceEndpoint {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceEndpoint{}
	for _, item := range list {
		result = append(result, InterfaceToServiceEndpoint(item))
	}
	return result
}
