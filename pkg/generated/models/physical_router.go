package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakePhysicalRouter makes PhysicalRouter
func MakePhysicalRouter() *PhysicalRouter {
	return &PhysicalRouter{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
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
func InterfaceToPhysicalRouter(i interface{}) *PhysicalRouter {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PhysicalRouter{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		PhysicalRouterManagementIP:      schema.InterfaceToString(m["physical_router_management_ip"]),
		PhysicalRouterSNMPCredentials:   InterfaceToSNMPCredentials(m["physical_router_snmp_credentials"]),
		PhysicalRouterRole:              schema.InterfaceToString(m["physical_router_role"]),
		PhysicalRouterUserCredentials:   InterfaceToUserCredentials(m["physical_router_user_credentials"]),
		PhysicalRouterVendorName:        schema.InterfaceToString(m["physical_router_vendor_name"]),
		PhysicalRouterVNCManaged:        schema.InterfaceToBool(m["physical_router_vnc_managed"]),
		PhysicalRouterProductName:       schema.InterfaceToString(m["physical_router_product_name"]),
		PhysicalRouterLLDP:              schema.InterfaceToBool(m["physical_router_lldp"]),
		PhysicalRouterLoopbackIP:        schema.InterfaceToString(m["physical_router_loopback_ip"]),
		PhysicalRouterImageURI:          schema.InterfaceToString(m["physical_router_image_uri"]),
		TelemetryInfo:                   InterfaceToTelemetryStateInfo(m["telemetry_info"]),
		PhysicalRouterSNMP:              schema.InterfaceToBool(m["physical_router_snmp"]),
		PhysicalRouterDataplaneIP:       schema.InterfaceToString(m["physical_router_dataplane_ip"]),
		PhysicalRouterJunosServicePorts: InterfaceToJunosServicePorts(m["physical_router_junos_service_ports"]),
	}
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}

// InterfaceToPhysicalRouterSlice() makes a slice of PhysicalRouter
func InterfaceToPhysicalRouterSlice(i interface{}) []*PhysicalRouter {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PhysicalRouter{}
	for _, item := range list {
		result = append(result, InterfaceToPhysicalRouter(item))
	}
	return result
}
