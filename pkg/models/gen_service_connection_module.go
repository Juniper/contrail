package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceConnectionModule makes ServiceConnectionModule
// nolint
func MakeServiceConnectionModule() *ServiceConnectionModule {
	return &ServiceConnectionModule{
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
		ServiceType:          "",
		E2Service:            "",
	}
}

// MakeServiceConnectionModule makes ServiceConnectionModule
// nolint
func InterfaceToServiceConnectionModule(i interface{}) *ServiceConnectionModule {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceConnectionModule{
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
		ServiceType:          common.InterfaceToString(m["service_type"]),
		E2Service:            common.InterfaceToString(m["e2_service"]),

		ServiceObjectRefs: InterfaceToServiceConnectionModuleServiceObjectRefs(m["service_object_refs"]),
	}
}

func InterfaceToServiceConnectionModuleServiceObjectRefs(i interface{}) []*ServiceConnectionModuleServiceObjectRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ServiceConnectionModuleServiceObjectRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ServiceConnectionModuleServiceObjectRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeServiceConnectionModuleSlice() makes a slice of ServiceConnectionModule
// nolint
func MakeServiceConnectionModuleSlice() []*ServiceConnectionModule {
	return []*ServiceConnectionModule{}
}

// InterfaceToServiceConnectionModuleSlice() makes a slice of ServiceConnectionModule
// nolint
func InterfaceToServiceConnectionModuleSlice(i interface{}) []*ServiceConnectionModule {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceConnectionModule{}
	for _, item := range list {
		result = append(result, InterfaceToServiceConnectionModule(item))
	}
	return result
}
