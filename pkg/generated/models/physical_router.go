package models

// PhysicalRouter

import "encoding/json"

// PhysicalRouter
type PhysicalRouter struct {
	PhysicalRouterImageURI          string              `json:"physical_router_image_uri,omitempty"`
	PhysicalRouterJunosServicePorts *JunosServicePorts  `json:"physical_router_junos_service_ports,omitempty"`
	IDPerms                         *IdPermsType        `json:"id_perms,omitempty"`
	Perms2                          *PermType2          `json:"perms2,omitempty"`
	PhysicalRouterUserCredentials   *UserCredentials    `json:"physical_router_user_credentials,omitempty"`
	PhysicalRouterLLDP              bool                `json:"physical_router_lldp,omitempty"`
	PhysicalRouterVendorName        string              `json:"physical_router_vendor_name,omitempty"`
	PhysicalRouterLoopbackIP        string              `json:"physical_router_loopback_ip,omitempty"`
	TelemetryInfo                   *TelemetryStateInfo `json:"telemetry_info,omitempty"`
	FQName                          []string            `json:"fq_name,omitempty"`
	Annotations                     *KeyValuePairs      `json:"annotations,omitempty"`
	ParentUUID                      string              `json:"parent_uuid,omitempty"`
	PhysicalRouterRole              PhysicalRouterRole  `json:"physical_router_role,omitempty"`
	PhysicalRouterProductName       string              `json:"physical_router_product_name,omitempty"`
	PhysicalRouterVNCManaged        bool                `json:"physical_router_vnc_managed,omitempty"`
	PhysicalRouterSNMP              bool                `json:"physical_router_snmp,omitempty"`
	PhysicalRouterDataplaneIP       string              `json:"physical_router_dataplane_ip,omitempty"`
	DisplayName                     string              `json:"display_name,omitempty"`
	UUID                            string              `json:"uuid,omitempty"`
	ParentType                      string              `json:"parent_type,omitempty"`
	PhysicalRouterManagementIP      string              `json:"physical_router_management_ip,omitempty"`
	PhysicalRouterSNMPCredentials   *SNMPCredentials    `json:"physical_router_snmp_credentials,omitempty"`

	VirtualNetworkRefs []*PhysicalRouterVirtualNetworkRef `json:"virtual_network_refs,omitempty"`
	BGPRouterRefs      []*PhysicalRouterBGPRouterRef      `json:"bgp_router_refs,omitempty"`
	VirtualRouterRefs  []*PhysicalRouterVirtualRouterRef  `json:"virtual_router_refs,omitempty"`

	LogicalInterfaces  []*LogicalInterface  `json:"logical_interfaces,omitempty"`
	PhysicalInterfaces []*PhysicalInterface `json:"physical_interfaces,omitempty"`
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
		PhysicalRouterRole:              MakePhysicalRouterRole(),
		PhysicalRouterProductName:       "",
		PhysicalRouterLoopbackIP:        "",
		TelemetryInfo:                   MakeTelemetryStateInfo(),
		FQName:                          []string{},
		Annotations:                     MakeKeyValuePairs(),
		ParentUUID:                      "",
		PhysicalRouterManagementIP:      "",
		PhysicalRouterSNMPCredentials:   MakeSNMPCredentials(),
		PhysicalRouterVNCManaged:        false,
		PhysicalRouterSNMP:              false,
		PhysicalRouterDataplaneIP:       "",
		DisplayName:                     "",
		UUID:                            "",
		ParentType:                      "",
		PhysicalRouterUserCredentials:   MakeUserCredentials(),
		PhysicalRouterLLDP:              false,
		PhysicalRouterImageURI:          "",
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
		IDPerms: MakeIdPermsType(),
		Perms2:  MakePermType2(),
		PhysicalRouterVendorName: "",
	}
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}
