package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeFloatingIP makes FloatingIP
func MakeFloatingIP() *FloatingIP {
	return &FloatingIP{
		//TODO(nati): Apply default
		UUID:                         "",
		ParentUUID:                   "",
		ParentType:                   "",
		FQName:                       []string{},
		IDPerms:                      MakeIdPermsType(),
		DisplayName:                  "",
		Annotations:                  MakeKeyValuePairs(),
		Perms2:                       MakePermType2(),
		FloatingIPAddressFamily:      "",
		FloatingIPPortMappings:       MakePortMappings(),
		FloatingIPIsVirtualIP:        false,
		FloatingIPAddress:            "",
		FloatingIPPortMappingsEnable: false,
		FloatingIPFixedIPAddress:     "",
		FloatingIPTrafficDirection:   "",
	}
}

// MakeFloatingIP makes FloatingIP
func InterfaceToFloatingIP(i interface{}) *FloatingIP {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FloatingIP{
		//TODO(nati): Apply default
		UUID:                         schema.InterfaceToString(m["uuid"]),
		ParentUUID:                   schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                   schema.InterfaceToString(m["parent_type"]),
		FQName:                       schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                      InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                  schema.InterfaceToString(m["display_name"]),
		Annotations:                  InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                       InterfaceToPermType2(m["perms2"]),
		FloatingIPAddressFamily:      schema.InterfaceToString(m["floating_ip_address_family"]),
		FloatingIPPortMappings:       InterfaceToPortMappings(m["floating_ip_port_mappings"]),
		FloatingIPIsVirtualIP:        schema.InterfaceToBool(m["floating_ip_is_virtual_ip"]),
		FloatingIPAddress:            schema.InterfaceToString(m["floating_ip_address"]),
		FloatingIPPortMappingsEnable: schema.InterfaceToBool(m["floating_ip_port_mappings_enable"]),
		FloatingIPFixedIPAddress:     schema.InterfaceToString(m["floating_ip_fixed_ip_address"]),
		FloatingIPTrafficDirection:   schema.InterfaceToString(m["floating_ip_traffic_direction"]),
	}
}

// MakeFloatingIPSlice() makes a slice of FloatingIP
func MakeFloatingIPSlice() []*FloatingIP {
	return []*FloatingIP{}
}

// InterfaceToFloatingIPSlice() makes a slice of FloatingIP
func InterfaceToFloatingIPSlice(i interface{}) []*FloatingIP {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FloatingIP{}
	for _, item := range list {
		result = append(result, InterfaceToFloatingIP(item))
	}
	return result
}
