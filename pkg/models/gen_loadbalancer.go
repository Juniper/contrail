package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancer makes Loadbalancer
// nolint
func MakeLoadbalancer() *Loadbalancer {
	return &Loadbalancer{
		//TODO(nati): Apply default
		UUID:                   "",
		ParentUUID:             "",
		ParentType:             "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Annotations:            MakeKeyValuePairs(),
		Perms2:                 MakePermType2(),
		ConfigurationVersion:   0,
		LoadbalancerProperties: MakeLoadbalancerType(),
		LoadbalancerProvider:   "",
	}
}

// MakeLoadbalancer makes Loadbalancer
// nolint
func InterfaceToLoadbalancer(i interface{}) *Loadbalancer {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Loadbalancer{
		//TODO(nati): Apply default
		UUID:                   common.InterfaceToString(m["uuid"]),
		ParentUUID:             common.InterfaceToString(m["parent_uuid"]),
		ParentType:             common.InterfaceToString(m["parent_type"]),
		FQName:                 common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:            common.InterfaceToString(m["display_name"]),
		Annotations:            InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                 InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:   common.InterfaceToInt64(m["configuration_version"]),
		LoadbalancerProperties: InterfaceToLoadbalancerType(m["loadbalancer_properties"]),
		LoadbalancerProvider:   common.InterfaceToString(m["loadbalancer_provider"]),

		VirtualMachineInterfaceRefs: InterfaceToLoadbalancerVirtualMachineInterfaceRefs(m["virtual_machine_interface_refs"]),

		ServiceInstanceRefs: InterfaceToLoadbalancerServiceInstanceRefs(m["service_instance_refs"]),

		ServiceApplianceSetRefs: InterfaceToLoadbalancerServiceApplianceSetRefs(m["service_appliance_set_refs"]),
	}
}

func InterfaceToLoadbalancerServiceApplianceSetRefs(i interface{}) []*LoadbalancerServiceApplianceSetRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LoadbalancerServiceApplianceSetRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LoadbalancerServiceApplianceSetRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLoadbalancerVirtualMachineInterfaceRefs(i interface{}) []*LoadbalancerVirtualMachineInterfaceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LoadbalancerVirtualMachineInterfaceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LoadbalancerVirtualMachineInterfaceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToLoadbalancerServiceInstanceRefs(i interface{}) []*LoadbalancerServiceInstanceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*LoadbalancerServiceInstanceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &LoadbalancerServiceInstanceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeLoadbalancerSlice() makes a slice of Loadbalancer
// nolint
func MakeLoadbalancerSlice() []*Loadbalancer {
	return []*Loadbalancer{}
}

// InterfaceToLoadbalancerSlice() makes a slice of Loadbalancer
// nolint
func InterfaceToLoadbalancerSlice(i interface{}) []*Loadbalancer {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Loadbalancer{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancer(item))
	}
	return result
}
