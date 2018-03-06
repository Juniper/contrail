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
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
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
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
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
	}
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
