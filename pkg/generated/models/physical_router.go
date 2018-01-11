package models

// PhysicalRouter

import "encoding/json"

// PhysicalRouter
type PhysicalRouter struct {
	PhysicalRouterSNMPCredentials   *SNMPCredentials    `json:"physical_router_snmp_credentials"`
	PhysicalRouterRole              PhysicalRouterRole  `json:"physical_router_role"`
	PhysicalRouterLLDP              bool                `json:"physical_router_lldp"`
	UUID                            string              `json:"uuid"`
	ParentType                      string              `json:"parent_type"`
	Annotations                     *KeyValuePairs      `json:"annotations"`
	PhysicalRouterManagementIP      string              `json:"physical_router_management_ip"`
	PhysicalRouterUserCredentials   *UserCredentials    `json:"physical_router_user_credentials"`
	PhysicalRouterVNCManaged        bool                `json:"physical_router_vnc_managed"`
	PhysicalRouterLoopbackIP        string              `json:"physical_router_loopback_ip"`
	PhysicalRouterImageURI          string              `json:"physical_router_image_uri"`
	PhysicalRouterJunosServicePorts *JunosServicePorts  `json:"physical_router_junos_service_ports"`
	IDPerms                         *IdPermsType        `json:"id_perms"`
	ParentUUID                      string              `json:"parent_uuid"`
	FQName                          []string            `json:"fq_name"`
	PhysicalRouterProductName       string              `json:"physical_router_product_name"`
	TelemetryInfo                   *TelemetryStateInfo `json:"telemetry_info"`
	PhysicalRouterDataplaneIP       string              `json:"physical_router_dataplane_ip"`
	Perms2                          *PermType2          `json:"perms2"`
	PhysicalRouterVendorName        string              `json:"physical_router_vendor_name"`
	PhysicalRouterSNMP              bool                `json:"physical_router_snmp"`
	DisplayName                     string              `json:"display_name"`

	VirtualRouterRefs  []*PhysicalRouterVirtualRouterRef  `json:"virtual_router_refs"`
	VirtualNetworkRefs []*PhysicalRouterVirtualNetworkRef `json:"virtual_network_refs"`
	BGPRouterRefs      []*PhysicalRouterBGPRouterRef      `json:"bgp_router_refs"`

	LogicalInterfaces  []*LogicalInterface  `json:"logical_interfaces"`
	PhysicalInterfaces []*PhysicalInterface `json:"physical_interfaces"`
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

// PhysicalRouterVirtualRouterRef references each other
type PhysicalRouterVirtualRouterRef struct {
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
		TelemetryInfo:             MakeTelemetryStateInfo(),
		PhysicalRouterDataplaneIP: "",
		Perms2: MakePermType2(),
		PhysicalRouterProductName:       "",
		PhysicalRouterSNMP:              false,
		DisplayName:                     "",
		PhysicalRouterVendorName:        "",
		PhysicalRouterRole:              MakePhysicalRouterRole(),
		PhysicalRouterLLDP:              false,
		UUID:                            "",
		ParentType:                      "",
		PhysicalRouterSNMPCredentials:   MakeSNMPCredentials(),
		PhysicalRouterUserCredentials:   MakeUserCredentials(),
		PhysicalRouterVNCManaged:        false,
		PhysicalRouterLoopbackIP:        "",
		PhysicalRouterImageURI:          "",
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
		IDPerms:                    MakeIdPermsType(),
		Annotations:                MakeKeyValuePairs(),
		PhysicalRouterManagementIP: "",
		FQName:     []string{},
		ParentUUID: "",
	}
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}
