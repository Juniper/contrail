package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePhysicalRouter makes PhysicalRouter
// nolint
func MakePhysicalRouter() *PhysicalRouter {
	return &PhysicalRouter{
		//TODO(nati): Apply default
		UUID:                            "",
		ParentUUID:                      "",
		ParentType:                      "",
		FQName:                          []string{},
		IDPerms:                         MakeIdPermsType(),
		DisplayName:                     "",
		Annotations:                     MakeKeyValuePairs(),
		Perms2:                          MakePermType2(),
		ConfigurationVersion:            0,
		PhysicalRouterManagementIP:      "",
		PhysicalRouterSNMPCredentials:   MakeSNMPCredentials(),
		PhysicalRouterRole:              "",
		PhysicalRouterUserCredentials:   MakeUserCredentials(),
		PhysicalRouterVendorName:        "",
		PhysicalRouterVNCManaged:        false,
		PhysicalRouterProductName:       "",
		PhysicalRouterLLDP:              false,
		PhysicalRouterLoopbackIP:        "",
		PhysicalRouterImageURI:          "",
		TelemetryInfo:                   MakeTelemetryStateInfo(),
		PhysicalRouterSNMP:              false,
		PhysicalRouterDataplaneIP:       "",
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
	}
}

// MakePhysicalRouter makes PhysicalRouter
// nolint
func InterfaceToPhysicalRouter(i interface{}) *PhysicalRouter {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PhysicalRouter{
		//TODO(nati): Apply default
		UUID:                            common.InterfaceToString(m["uuid"]),
		ParentUUID:                      common.InterfaceToString(m["parent_uuid"]),
		ParentType:                      common.InterfaceToString(m["parent_type"]),
		FQName:                          common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                         InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                     common.InterfaceToString(m["display_name"]),
		Annotations:                     InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                          InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:            common.InterfaceToInt64(m["configuration_version"]),
		PhysicalRouterManagementIP:      common.InterfaceToString(m["physical_router_management_ip"]),
		PhysicalRouterSNMPCredentials:   InterfaceToSNMPCredentials(m["physical_router_snmp_credentials"]),
		PhysicalRouterRole:              common.InterfaceToString(m["physical_router_role"]),
		PhysicalRouterUserCredentials:   InterfaceToUserCredentials(m["physical_router_user_credentials"]),
		PhysicalRouterVendorName:        common.InterfaceToString(m["physical_router_vendor_name"]),
		PhysicalRouterVNCManaged:        common.InterfaceToBool(m["physical_router_vnc_managed"]),
		PhysicalRouterProductName:       common.InterfaceToString(m["physical_router_product_name"]),
		PhysicalRouterLLDP:              common.InterfaceToBool(m["physical_router_lldp"]),
		PhysicalRouterLoopbackIP:        common.InterfaceToString(m["physical_router_loopback_ip"]),
		PhysicalRouterImageURI:          common.InterfaceToString(m["physical_router_image_uri"]),
		TelemetryInfo:                   InterfaceToTelemetryStateInfo(m["telemetry_info"]),
		PhysicalRouterSNMP:              common.InterfaceToBool(m["physical_router_snmp"]),
		PhysicalRouterDataplaneIP:       common.InterfaceToString(m["physical_router_dataplane_ip"]),
		PhysicalRouterJunosServicePorts: InterfaceToJunosServicePorts(m["physical_router_junos_service_ports"]),
	}
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
// nolint
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}

// InterfaceToPhysicalRouterSlice() makes a slice of PhysicalRouter
// nolint
func InterfaceToPhysicalRouterSlice(i interface{}) []*PhysicalRouter {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PhysicalRouter{}
	for _, item := range list {
		result = append(result, InterfaceToPhysicalRouter(item))
	}
	return result
}
