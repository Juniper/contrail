package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualMachineInterface makes VirtualMachineInterface
// nolint
func MakeVirtualMachineInterface() *VirtualMachineInterface {
	return &VirtualMachineInterface{
		//TODO(nati): Apply default
		UUID:                                       "",
		ParentUUID:                                 "",
		ParentType:                                 "",
		FQName:                                     []string{},
		IDPerms:                                    MakeIdPermsType(),
		DisplayName:                                "",
		Annotations:                                MakeKeyValuePairs(),
		Perms2:                                     MakePermType2(),
		ConfigurationVersion:                       0,
		EcmpHashingIncludeFields:                   MakeEcmpHashingIncludeFields(),
		VirtualMachineInterfaceHostRoutes:          MakeRouteTableType(),
		VirtualMachineInterfaceMacAddresses:        MakeMacAddressesType(),
		VirtualMachineInterfaceDHCPOptionList:      MakeDhcpOptionsListType(),
		VirtualMachineInterfaceBindings:            MakeKeyValuePairs(),
		VirtualMachineInterfaceDisablePolicy:       false,
		VirtualMachineInterfaceAllowedAddressPairs: MakeAllowedAddressPairs(),
		VirtualMachineInterfaceFatFlowProtocols:    MakeFatFlowProtocols(),
		VlanTagBasedBridgeDomain:                   false,
		VirtualMachineInterfaceDeviceOwner:         "",
		VRFAssignTable:                             MakeVrfAssignTableType(),
		PortSecurityEnabled:                        false,
		VirtualMachineInterfaceProperties:          MakeVirtualMachineInterfacePropertiesType(),
	}
}

// MakeVirtualMachineInterface makes VirtualMachineInterface
// nolint
func InterfaceToVirtualMachineInterface(i interface{}) *VirtualMachineInterface {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualMachineInterface{
		//TODO(nati): Apply default
		UUID:                                       common.InterfaceToString(m["uuid"]),
		ParentUUID:                                 common.InterfaceToString(m["parent_uuid"]),
		ParentType:                                 common.InterfaceToString(m["parent_type"]),
		FQName:                                     common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                                    InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                                common.InterfaceToString(m["display_name"]),
		Annotations:                                InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                                     InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:                       common.InterfaceToInt64(m["configuration_version"]),
		EcmpHashingIncludeFields:                   InterfaceToEcmpHashingIncludeFields(m["ecmp_hashing_include_fields"]),
		VirtualMachineInterfaceHostRoutes:          InterfaceToRouteTableType(m["virtual_machine_interface_host_routes"]),
		VirtualMachineInterfaceMacAddresses:        InterfaceToMacAddressesType(m["virtual_machine_interface_mac_addresses"]),
		VirtualMachineInterfaceDHCPOptionList:      InterfaceToDhcpOptionsListType(m["virtual_machine_interface_dhcp_option_list"]),
		VirtualMachineInterfaceBindings:            InterfaceToKeyValuePairs(m["virtual_machine_interface_bindings"]),
		VirtualMachineInterfaceDisablePolicy:       common.InterfaceToBool(m["virtual_machine_interface_disable_policy"]),
		VirtualMachineInterfaceAllowedAddressPairs: InterfaceToAllowedAddressPairs(m["virtual_machine_interface_allowed_address_pairs"]),
		VirtualMachineInterfaceFatFlowProtocols:    InterfaceToFatFlowProtocols(m["virtual_machine_interface_fat_flow_protocols"]),
		VlanTagBasedBridgeDomain:                   common.InterfaceToBool(m["vlan_tag_based_bridge_domain"]),
		VirtualMachineInterfaceDeviceOwner:         common.InterfaceToString(m["virtual_machine_interface_device_owner"]),
		VRFAssignTable:                             InterfaceToVrfAssignTableType(m["vrf_assign_table"]),
		PortSecurityEnabled:                        common.InterfaceToBool(m["port_security_enabled"]),
		VirtualMachineInterfaceProperties:          InterfaceToVirtualMachineInterfacePropertiesType(m["virtual_machine_interface_properties"]),

		PortTupleRefs: InterfaceToVirtualMachineInterfacePortTupleRefs(m["port_tuple_refs"]),

		VirtualNetworkRefs: InterfaceToVirtualMachineInterfaceVirtualNetworkRefs(m["virtual_network_refs"]),

		BGPRouterRefs: InterfaceToVirtualMachineInterfaceBGPRouterRefs(m["bgp_router_refs"]),

		SecurityLoggingObjectRefs: InterfaceToVirtualMachineInterfaceSecurityLoggingObjectRefs(m["security_logging_object_refs"]),

		RoutingInstanceRefs: InterfaceToVirtualMachineInterfaceRoutingInstanceRefs(m["routing_instance_refs"]),

		QosConfigRefs: InterfaceToVirtualMachineInterfaceQosConfigRefs(m["qos_config_refs"]),

		VirtualMachineInterfaceRefs: InterfaceToVirtualMachineInterfaceVirtualMachineInterfaceRefs(m["virtual_machine_interface_refs"]),

		VirtualMachineRefs: InterfaceToVirtualMachineInterfaceVirtualMachineRefs(m["virtual_machine_refs"]),

		ServiceHealthCheckRefs: InterfaceToVirtualMachineInterfaceServiceHealthCheckRefs(m["service_health_check_refs"]),

		InterfaceRouteTableRefs: InterfaceToVirtualMachineInterfaceInterfaceRouteTableRefs(m["interface_route_table_refs"]),

		PhysicalInterfaceRefs: InterfaceToVirtualMachineInterfacePhysicalInterfaceRefs(m["physical_interface_refs"]),

		BridgeDomainRefs: InterfaceToVirtualMachineInterfaceBridgeDomainRefs(m["bridge_domain_refs"]),

		SecurityGroupRefs: InterfaceToVirtualMachineInterfaceSecurityGroupRefs(m["security_group_refs"]),

		ServiceEndpointRefs: InterfaceToVirtualMachineInterfaceServiceEndpointRefs(m["service_endpoint_refs"]),
	}
}

func InterfaceToVirtualMachineInterfaceBGPRouterRefs(i interface{}) []*VirtualMachineInterfaceBGPRouterRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceBGPRouterRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceBGPRouterRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceSecurityLoggingObjectRefs(i interface{}) []*VirtualMachineInterfaceSecurityLoggingObjectRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceSecurityLoggingObjectRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceSecurityLoggingObjectRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceRoutingInstanceRefs(i interface{}) []*VirtualMachineInterfaceRoutingInstanceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceRoutingInstanceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceRoutingInstanceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),

			Attr: InterfaceToPolicyBasedForwardingRuleType(m["attr"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceQosConfigRefs(i interface{}) []*VirtualMachineInterfaceQosConfigRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceQosConfigRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceQosConfigRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfacePortTupleRefs(i interface{}) []*VirtualMachineInterfacePortTupleRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfacePortTupleRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfacePortTupleRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceVirtualNetworkRefs(i interface{}) []*VirtualMachineInterfaceVirtualNetworkRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceVirtualNetworkRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceVirtualNetworkRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceVirtualMachineInterfaceRefs(i interface{}) []*VirtualMachineInterfaceVirtualMachineInterfaceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceVirtualMachineInterfaceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceVirtualMachineInterfaceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceVirtualMachineRefs(i interface{}) []*VirtualMachineInterfaceVirtualMachineRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceVirtualMachineRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceVirtualMachineRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceServiceHealthCheckRefs(i interface{}) []*VirtualMachineInterfaceServiceHealthCheckRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceServiceHealthCheckRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceServiceHealthCheckRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceInterfaceRouteTableRefs(i interface{}) []*VirtualMachineInterfaceInterfaceRouteTableRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceInterfaceRouteTableRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceInterfaceRouteTableRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfacePhysicalInterfaceRefs(i interface{}) []*VirtualMachineInterfacePhysicalInterfaceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfacePhysicalInterfaceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfacePhysicalInterfaceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceBridgeDomainRefs(i interface{}) []*VirtualMachineInterfaceBridgeDomainRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceBridgeDomainRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceBridgeDomainRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),

			Attr: InterfaceToBridgeDomainMembershipType(m["attr"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceSecurityGroupRefs(i interface{}) []*VirtualMachineInterfaceSecurityGroupRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceSecurityGroupRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceSecurityGroupRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToVirtualMachineInterfaceServiceEndpointRefs(i interface{}) []*VirtualMachineInterfaceServiceEndpointRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*VirtualMachineInterfaceServiceEndpointRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &VirtualMachineInterfaceServiceEndpointRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeVirtualMachineInterfaceSlice() makes a slice of VirtualMachineInterface
// nolint
func MakeVirtualMachineInterfaceSlice() []*VirtualMachineInterface {
	return []*VirtualMachineInterface{}
}

// InterfaceToVirtualMachineInterfaceSlice() makes a slice of VirtualMachineInterface
// nolint
func InterfaceToVirtualMachineInterfaceSlice(i interface{}) []*VirtualMachineInterface {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualMachineInterface{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualMachineInterface(item))
	}
	return result
}
