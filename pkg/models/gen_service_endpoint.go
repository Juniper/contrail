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
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ConfigurationVersion: 0,
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
		UUID:                 common.InterfaceToString(m["uuid"]),
		ParentUUID:           common.InterfaceToString(m["parent_uuid"]),
		ParentType:           common.InterfaceToString(m["parent_type"]),
		FQName:               common.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          common.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion: common.InterfaceToInt64(m["configuration_version"]),

		ServiceConnectionModuleRefs: InterfaceToServiceEndpointServiceConnectionModuleRefs(m["service_connection_module_refs"]),

		PhysicalRouterRefs: InterfaceToServiceEndpointPhysicalRouterRefs(m["physical_router_refs"]),

		ServiceObjectRefs: InterfaceToServiceEndpointServiceObjectRefs(m["service_object_refs"]),
	}
}

func InterfaceToServiceEndpointPhysicalRouterRefs(i interface{}) []*ServiceEndpointPhysicalRouterRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ServiceEndpointPhysicalRouterRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ServiceEndpointPhysicalRouterRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToServiceEndpointServiceObjectRefs(i interface{}) []*ServiceEndpointServiceObjectRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ServiceEndpointServiceObjectRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ServiceEndpointServiceObjectRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToServiceEndpointServiceConnectionModuleRefs(i interface{}) []*ServiceEndpointServiceConnectionModuleRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ServiceEndpointServiceConnectionModuleRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ServiceEndpointServiceConnectionModuleRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
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
