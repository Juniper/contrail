package models

// PhysicalRouter

import "encoding/json"

// PhysicalRouter
type PhysicalRouter struct {
	PhysicalRouterDataplaneIP       string              `json:"physical_router_dataplane_ip"`
	DisplayName                     string              `json:"display_name"`
	Annotations                     *KeyValuePairs      `json:"annotations"`
	FQName                          []string            `json:"fq_name"`
	PhysicalRouterManagementIP      string              `json:"physical_router_management_ip"`
	PhysicalRouterVendorName        string              `json:"physical_router_vendor_name"`
	PhysicalRouterLLDP              bool                `json:"physical_router_lldp"`
	PhysicalRouterSNMP              bool                `json:"physical_router_snmp"`
	PhysicalRouterJunosServicePorts *JunosServicePorts  `json:"physical_router_junos_service_ports"`
	IDPerms                         *IdPermsType        `json:"id_perms"`
	UUID                            string              `json:"uuid"`
	PhysicalRouterRole              PhysicalRouterRole  `json:"physical_router_role"`
	PhysicalRouterUserCredentials   *UserCredentials    `json:"physical_router_user_credentials"`
	PhysicalRouterProductName       string              `json:"physical_router_product_name"`
	PhysicalRouterLoopbackIP        string              `json:"physical_router_loopback_ip"`
	Perms2                          *PermType2          `json:"perms2"`
	PhysicalRouterSNMPCredentials   *SNMPCredentials    `json:"physical_router_snmp_credentials"`
	PhysicalRouterVNCManaged        bool                `json:"physical_router_vnc_managed"`
	PhysicalRouterImageURI          string              `json:"physical_router_image_uri"`
	TelemetryInfo                   *TelemetryStateInfo `json:"telemetry_info"`

	// virtual_network <common.Reference Value>
	VirtualNetworkRefs []*PhysicalRouterVirtualNetworkRef `json:"virtual_network_refs"`
	// bgp_router <common.Reference Value>
	BGPRouterRefs []*PhysicalRouterBGPRouterRef `json:"bgp_router_refs"`
	// virtual_router <common.Reference Value>
	VirtualRouterRefs []*PhysicalRouterVirtualRouterRef `json:"virtual_router_refs"`

	GlobalSystemConfigs []*PhysicalRouterGlobalSystemConfig
	Locations           []*PhysicalRouterLocation
}

// PhysicalRouterVirtualRouterRef references each other
type PhysicalRouterVirtualRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// PhysicalRouterVirtualNetworkRef references each other
type PhysicalRouterVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// PhysicalRouterBGPRouterRef references each other
type PhysicalRouterBGPRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// PhysicalRouter parents relation object

type PhysicalRouterLocation struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

type PhysicalRouterGlobalSystemConfig struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *PhysicalRouter) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePhysicalRouter makes PhysicalRouter
func MakePhysicalRouter() *PhysicalRouter {
	return &PhysicalRouter{
		//TODO(nati): Apply default
		PhysicalRouterLLDP:        false,
		PhysicalRouterSNMP:        false,
		PhysicalRouterDataplaneIP: "",
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		FQName:                    []string{},
		PhysicalRouterManagementIP:      "",
		PhysicalRouterVendorName:        "",
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
		IDPerms:                   MakeIdPermsType(),
		PhysicalRouterProductName: "",
		PhysicalRouterLoopbackIP:  "",
		UUID:                          "",
		PhysicalRouterRole:            MakePhysicalRouterRole(),
		PhysicalRouterUserCredentials: MakeUserCredentials(),
		PhysicalRouterImageURI:        "",
		TelemetryInfo:                 MakeTelemetryStateInfo(),
		Perms2:                        MakePermType2(),
		PhysicalRouterSNMPCredentials: MakeSNMPCredentials(),
		PhysicalRouterVNCManaged:      false,
	}
}

// InterfaceToPhysicalRouter makes PhysicalRouter from interface
func InterfaceToPhysicalRouter(iData interface{}) *PhysicalRouter {
	data := iData.(map[string]interface{})
	return &PhysicalRouter{
		PhysicalRouterJunosServicePorts: InterfaceToJunosServicePorts(data["physical_router_junos_service_ports"]),

		//{"Title":"","Description":"Juniper JUNOS specific service interfaces name  to perform services like NAT.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"service_port":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"service_port","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServicePort","GoType":"string","GoPremitive":true},"GoName":"ServicePort","GoType":"[]string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/JunosServicePorts","CollectionType":"","Column":"","Item":null,"GoName":"PhysicalRouterJunosServicePorts","GoType":"JunosServicePorts","GoPremitive":false}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string","GoPremitive":true},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string","GoPremitive":true},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string","GoPremitive":true},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string","GoPremitive":true},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"permissions_owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"permissions_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType","GoPremitive":false},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType","GoPremitive":false}
		PhysicalRouterRole: InterfaceToPhysicalRouterRole(data["physical_router_role"]),

		//{"Title":"","Description":"Physical router role (e.g spine or leaf), used by the device manager to provision physical router, for e.g device manager may choose to configure physical router based on its role.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["spine","leaf","e2-access","e2-provider","e2-internet","e2-vrr"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PhysicalRouterRole","CollectionType":"","Column":"physical_router_role","Item":null,"GoName":"PhysicalRouterRole","GoType":"PhysicalRouterRole","GoPremitive":false}
		PhysicalRouterUserCredentials: InterfaceToUserCredentials(data["physical_router_user_credentials"]),

		//{"Title":"","Description":"Username and password for netconf to the physical router by device manager.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"password":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"password","Item":null,"GoName":"Password","GoType":"string","GoPremitive":true},"username":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"username","Item":null,"GoName":"Username","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/UserCredentials","CollectionType":"","Column":"","Item":null,"GoName":"PhysicalRouterUserCredentials","GoType":"UserCredentials","GoPremitive":false}
		PhysicalRouterProductName: data["physical_router_product_name"].(string),

		//{"Title":"","Description":"Model name of the physical router (e.g juniper). Used by the device manager to select driver.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_router_product_name","Item":null,"GoName":"PhysicalRouterProductName","GoType":"string","GoPremitive":true}
		PhysicalRouterLoopbackIP: data["physical_router_loopback_ip"].(string),

		//{"Title":"","Description":"This is ip address of loopback interface of physical router. Used by the device manager to configure physical router loopback interface.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_router_loopback_ip","Item":null,"GoName":"PhysicalRouterLoopbackIP","GoType":"string","GoPremitive":true}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string","GoPremitive":true}
		PhysicalRouterSNMPCredentials: InterfaceToSNMPCredentials(data["physical_router_snmp_credentials"]),

		//{"Title":"","Description":"SNMP credentials for the physical router used by SNMP collector.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"local_port":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"local_port","Item":null,"GoName":"LocalPort","GoType":"int","GoPremitive":true},"retries":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"retries","Item":null,"GoName":"Retries","GoType":"int","GoPremitive":true},"timeout":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"timeout","Item":null,"GoName":"Timeout","GoType":"int","GoPremitive":true},"v2_community":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v2_community","Item":null,"GoName":"V2Community","GoType":"string","GoPremitive":true},"v3_authentication_password":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_authentication_password","Item":null,"GoName":"V3AuthenticationPassword","GoType":"string","GoPremitive":true},"v3_authentication_protocol":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_authentication_protocol","Item":null,"GoName":"V3AuthenticationProtocol","GoType":"string","GoPremitive":true},"v3_context":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_context","Item":null,"GoName":"V3Context","GoType":"string","GoPremitive":true},"v3_context_engine_id":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_context_engine_id","Item":null,"GoName":"V3ContextEngineID","GoType":"string","GoPremitive":true},"v3_engine_boots":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_engine_boots","Item":null,"GoName":"V3EngineBoots","GoType":"int","GoPremitive":true},"v3_engine_id":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_engine_id","Item":null,"GoName":"V3EngineID","GoType":"string","GoPremitive":true},"v3_engine_time":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_engine_time","Item":null,"GoName":"V3EngineTime","GoType":"int","GoPremitive":true},"v3_privacy_password":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_privacy_password","Item":null,"GoName":"V3PrivacyPassword","GoType":"string","GoPremitive":true},"v3_privacy_protocol":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_privacy_protocol","Item":null,"GoName":"V3PrivacyProtocol","GoType":"string","GoPremitive":true},"v3_security_engine_id":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_security_engine_id","Item":null,"GoName":"V3SecurityEngineID","GoType":"string","GoPremitive":true},"v3_security_level":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_security_level","Item":null,"GoName":"V3SecurityLevel","GoType":"string","GoPremitive":true},"v3_security_name":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"v3_security_name","Item":null,"GoName":"V3SecurityName","GoType":"string","GoPremitive":true},"version":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"version","Item":null,"GoName":"Version","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SNMPCredentials","CollectionType":"","Column":"","Item":null,"GoName":"PhysicalRouterSNMPCredentials","GoType":"SNMPCredentials","GoPremitive":false}
		PhysicalRouterVNCManaged: data["physical_router_vnc_managed"].(bool),

		//{"Title":"","Description":"This physical router is enabled to be configured by device manager.","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_router_vnc_managed","Item":null,"GoName":"PhysicalRouterVNCManaged","GoType":"bool","GoPremitive":true}
		PhysicalRouterImageURI: data["physical_router_image_uri"].(string),

		//{"Title":"","Description":"Physical router OS image uri","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_router_image_uri","Item":null,"GoName":"PhysicalRouterImageURI","GoType":"string","GoPremitive":true}
		TelemetryInfo: InterfaceToTelemetryStateInfo(data["telemetry_info"]),

		//{"Title":"","Description":"Telemetry info of router. Check TelemetryStateInfo","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"resource":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"resource","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Name","GoType":"string","GoPremitive":true},"path":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Path","GoType":"string","GoPremitive":true},"rate":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Rate","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TelemetryResourceInfo","CollectionType":"","Column":"","Item":null,"GoName":"Resource","GoType":"TelemetryResourceInfo","GoPremitive":false},"GoName":"Resource","GoType":"[]*TelemetryResourceInfo","GoPremitive":true},"server_ip":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"server_ip","Item":null,"GoName":"ServerIP","GoType":"string","GoPremitive":true},"server_port":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"server_port","Item":null,"GoName":"ServerPort","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TelemetryStateInfo","CollectionType":"","Column":"","Item":null,"GoName":"TelemetryInfo","GoType":"TelemetryStateInfo","GoPremitive":false}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType","GoPremitive":false},"GoName":"Share","GoType":"[]*ShareType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2","GoPremitive":false}
		PhysicalRouterManagementIP: data["physical_router_management_ip"].(string),

		//{"Title":"","Description":"Management ip for this physical router. It is used by the device manager to perform netconf and by SNMP collector if enabled.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_router_management_ip","Item":null,"GoName":"PhysicalRouterManagementIP","GoType":"string","GoPremitive":true}
		PhysicalRouterVendorName: data["physical_router_vendor_name"].(string),

		//{"Title":"","Description":"Vendor name of the physical router (e.g juniper). Used by the device manager to select driver.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"required","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_router_vendor_name","Item":null,"GoName":"PhysicalRouterVendorName","GoType":"string","GoPremitive":true}
		PhysicalRouterLLDP: data["physical_router_lldp"].(bool),

		//{"Title":"","Description":"LLDP support on this router","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_router_lldp","Item":null,"GoName":"PhysicalRouterLLDP","GoType":"bool","GoPremitive":true}
		PhysicalRouterSNMP: data["physical_router_snmp"].(bool),

		//{"Title":"","Description":"SNMP support on this router","SQL":"bool","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_router_snmp","Item":null,"GoName":"PhysicalRouterSNMP","GoType":"bool","GoPremitive":true}
		PhysicalRouterDataplaneIP: data["physical_router_dataplane_ip"].(string),

		//{"Title":"","Description":"This is ip address in the ip-fabric(underlay) network that can be used in data plane by physical router. Usually it is the VTEP address in VxLAN for the TOR switch.","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"physical_router_dataplane_ip","Item":null,"GoName":"PhysicalRouterDataplaneIP","GoType":"string","GoPremitive":true}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string","GoPremitive":true}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair","GoPremitive":false},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs","GoPremitive":false}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string","GoPremitive":true},"GoName":"FQName","GoType":"[]string","GoPremitive":true}

	}
}

// InterfaceToPhysicalRouterSlice makes a slice of PhysicalRouter from interface
func InterfaceToPhysicalRouterSlice(data interface{}) []*PhysicalRouter {
	list := data.([]interface{})
	result := MakePhysicalRouterSlice()
	for _, item := range list {
		result = append(result, InterfaceToPhysicalRouter(item))
	}
	return result
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}
