package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerPool makes LoadbalancerPool
// nolint
func MakeLoadbalancerPool() *LoadbalancerPool {
	return &LoadbalancerPool{
		//TODO(nati): Apply default
		UUID:                             "",
		ParentUUID:                       "",
		ParentType:                       "",
		FQName:                           []string{},
		IDPerms:                          MakeIdPermsType(),
		DisplayName:                      "",
		Annotations:                      MakeKeyValuePairs(),
		Perms2:                           MakePermType2(),
		ConfigurationVersion:             0,
		LoadbalancerPoolProperties:       MakeLoadbalancerPoolType(),
		LoadbalancerPoolCustomAttributes: MakeKeyValuePairs(),
		LoadbalancerPoolProvider:         "",
	}
}

// MakeLoadbalancerPool makes LoadbalancerPool
// nolint
func InterfaceToLoadbalancerPool(i interface{}) *LoadbalancerPool {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerPool{
		//TODO(nati): Apply default
		UUID:                             common.InterfaceToString(m["uuid"]),
		ParentUUID:                       common.InterfaceToString(m["parent_uuid"]),
		ParentType:                       common.InterfaceToString(m["parent_type"]),
		FQName:                           common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                          InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                      common.InterfaceToString(m["display_name"]),
		Annotations:                      InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                           InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:             common.InterfaceToInt64(m["configuration_version"]),
		LoadbalancerPoolProperties:       InterfaceToLoadbalancerPoolType(m["loadbalancer_pool_properties"]),
		LoadbalancerPoolCustomAttributes: InterfaceToKeyValuePairs(m["loadbalancer_pool_custom_attributes"]),
		LoadbalancerPoolProvider:         common.InterfaceToString(m["loadbalancer_pool_provider"]),

		ServiceApplianceSetRefs: InterfaceToLoadbalancerPoolServiceApplianceSetRefs(m["service_appliance_set_refs"]),

		VirtualMachineInterfaceRefs: InterfaceToLoadbalancerPoolVirtualMachineInterfaceRefs(m["virtual_machine_interface_refs"]),

		LoadbalancerListenerRefs: InterfaceToLoadbalancerPoolLoadbalancerListenerRefs(m["loadbalancer_listener_refs"]),

		ServiceInstanceRefs: InterfaceToLoadbalancerPoolServiceInstanceRefs(m["service_instance_refs"]),

		LoadbalancerHealthmonitorRefs: InterfaceToLoadbalancerPoolLoadbalancerHealthmonitorRefs(m["loadbalancer_healthmonitor_refs"]),
	}
}

func InterfaceToLoadbalancerPoolVirtualMachineInterfaceRefs(i interface{}) []*LoadbalancerPoolVirtualMachineInterfaceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LoadbalancerPoolVirtualMachineInterfaceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LoadbalancerPoolVirtualMachineInterfaceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLoadbalancerPoolLoadbalancerListenerRefs(i interface{}) []*LoadbalancerPoolLoadbalancerListenerRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LoadbalancerPoolLoadbalancerListenerRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LoadbalancerPoolLoadbalancerListenerRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLoadbalancerPoolServiceInstanceRefs(i interface{}) []*LoadbalancerPoolServiceInstanceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LoadbalancerPoolServiceInstanceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LoadbalancerPoolServiceInstanceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLoadbalancerPoolLoadbalancerHealthmonitorRefs(i interface{}) []*LoadbalancerPoolLoadbalancerHealthmonitorRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LoadbalancerPoolLoadbalancerHealthmonitorRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LoadbalancerPoolLoadbalancerHealthmonitorRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLoadbalancerPoolServiceApplianceSetRefs(i interface{}) []*LoadbalancerPoolServiceApplianceSetRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LoadbalancerPoolServiceApplianceSetRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LoadbalancerPoolServiceApplianceSetRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
// nolint
func MakeLoadbalancerPoolSlice() []*LoadbalancerPool {
	return []*LoadbalancerPool{}
}

// InterfaceToLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
// nolint
func InterfaceToLoadbalancerPoolSlice(i interface{}) []*LoadbalancerPool {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerPool{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerPool(item))
	}
	return result
}
