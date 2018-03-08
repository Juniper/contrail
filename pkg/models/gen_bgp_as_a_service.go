package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeBGPAsAService makes BGPAsAService
// nolint
func MakeBGPAsAService() *BGPAsAService {
	return &BGPAsAService{
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
		BgpaasShared:                     false,
		BgpaasSessionAttributes:          "",
		BgpaasSuppressRouteAdvertisement: false,
		BgpaasIpv4MappedIpv6Nexthop:      false,
		BgpaasIPAddress:                  "",
		AutonomousSystem:                 0,
	}
}

// MakeBGPAsAService makes BGPAsAService
// nolint
func InterfaceToBGPAsAService(i interface{}) *BGPAsAService {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BGPAsAService{
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
		BgpaasShared:                     common.InterfaceToBool(m["bgpaas_shared"]),
		BgpaasSessionAttributes:          common.InterfaceToString(m["bgpaas_session_attributes"]),
		BgpaasSuppressRouteAdvertisement: common.InterfaceToBool(m["bgpaas_suppress_route_advertisement"]),
		BgpaasIpv4MappedIpv6Nexthop:      common.InterfaceToBool(m["bgpaas_ipv4_mapped_ipv6_nexthop"]),
		BgpaasIPAddress:                  common.InterfaceToString(m["bgpaas_ip_address"]),
		AutonomousSystem:                 common.InterfaceToInt64(m["autonomous_system"]),

		VirtualMachineInterfaceRefs: InterfaceToBGPAsAServiceVirtualMachineInterfaceRefs(m["virtual_machine_interface_refs"]),

		ServiceHealthCheckRefs: InterfaceToBGPAsAServiceServiceHealthCheckRefs(m["service_health_check_refs"]),
	}
}

func InterfaceToBGPAsAServiceVirtualMachineInterfaceRefs(i interface{}) []*BGPAsAServiceVirtualMachineInterfaceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*BGPAsAServiceVirtualMachineInterfaceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &BGPAsAServiceVirtualMachineInterfaceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToBGPAsAServiceServiceHealthCheckRefs(i interface{}) []*BGPAsAServiceServiceHealthCheckRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*BGPAsAServiceServiceHealthCheckRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &BGPAsAServiceServiceHealthCheckRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
// nolint
func MakeBGPAsAServiceSlice() []*BGPAsAService {
	return []*BGPAsAService{}
}

// InterfaceToBGPAsAServiceSlice() makes a slice of BGPAsAService
// nolint
func InterfaceToBGPAsAServiceSlice(i interface{}) []*BGPAsAService {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BGPAsAService{}
	for _, item := range list {
		result = append(result, InterfaceToBGPAsAService(item))
	}
	return result
}
