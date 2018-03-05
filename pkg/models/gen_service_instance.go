package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceInstance makes ServiceInstance
// nolint
func MakeServiceInstance() *ServiceInstance {
	return &ServiceInstance{
		//TODO(nati): Apply default
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		ConfigurationVersion:      0,
		ServiceInstanceBindings:   MakeKeyValuePairs(),
		ServiceInstanceProperties: MakeServiceInstanceType(),
	}
}

// MakeServiceInstance makes ServiceInstance
// nolint
func InterfaceToServiceInstance(i interface{}) *ServiceInstance {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceInstance{
		//TODO(nati): Apply default
		UUID:                      common.InterfaceToString(m["uuid"]),
		ParentUUID:                common.InterfaceToString(m["parent_uuid"]),
		ParentType:                common.InterfaceToString(m["parent_type"]),
		FQName:                    common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               common.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:      common.InterfaceToInt64(m["configuration_version"]),
		ServiceInstanceBindings:   InterfaceToKeyValuePairs(m["service_instance_bindings"]),
		ServiceInstanceProperties: InterfaceToServiceInstanceType(m["service_instance_properties"]),
	}
}

// MakeServiceInstanceSlice() makes a slice of ServiceInstance
// nolint
func MakeServiceInstanceSlice() []*ServiceInstance {
	return []*ServiceInstance{}
}

// InterfaceToServiceInstanceSlice() makes a slice of ServiceInstance
// nolint
func InterfaceToServiceInstanceSlice(i interface{}) []*ServiceInstance {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceInstance{}
	for _, item := range list {
		result = append(result, InterfaceToServiceInstance(item))
	}
	return result
}
