package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceHealthCheck makes ServiceHealthCheck
// nolint
func MakeServiceHealthCheck() *ServiceHealthCheck {
	return &ServiceHealthCheck{
		//TODO(nati): Apply default
		UUID:                         "",
		ParentUUID:                   "",
		ParentType:                   "",
		FQName:                       []string{},
		IDPerms:                      MakeIdPermsType(),
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		Perms2:                       MakePermType2(),
		ConfigurationVersion:         0,
		ServiceHealthCheckProperties: MakeServiceHealthCheckType(),
	}
}

// MakeServiceHealthCheck makes ServiceHealthCheck
// nolint
func InterfaceToServiceHealthCheck(i interface{}) *ServiceHealthCheck {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceHealthCheck{
		//TODO(nati): Apply default
		UUID:                         common.InterfaceToString(m["uuid"]),
		ParentUUID:                   common.InterfaceToString(m["parent_uuid"]),
		ParentType:                   common.InterfaceToString(m["parent_type"]),
		FQName:                       common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                      InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                  common.InterfaceToString(m["display_name"]),
		Annotations:                  InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                       InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:         common.InterfaceToInt64(m["configuration_version"]),
		ServiceHealthCheckProperties: InterfaceToServiceHealthCheckType(m["service_health_check_properties"]),
	}
}

// MakeServiceHealthCheckSlice() makes a slice of ServiceHealthCheck
// nolint
func MakeServiceHealthCheckSlice() []*ServiceHealthCheck {
	return []*ServiceHealthCheck{}
}

// InterfaceToServiceHealthCheckSlice() makes a slice of ServiceHealthCheck
// nolint
func InterfaceToServiceHealthCheckSlice(i interface{}) []*ServiceHealthCheck {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceHealthCheck{}
	for _, item := range list {
		result = append(result, InterfaceToServiceHealthCheck(item))
	}
	return result
}
