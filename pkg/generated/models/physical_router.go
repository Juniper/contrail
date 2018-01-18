package models

// PhysicalRouter

import "encoding/json"

// PhysicalRouter
type PhysicalRouter struct {
	PhysicalRouterProductName       string              `json:"physical_router_product_name,omitempty"`
	PhysicalRouterSNMP              bool                `json:"physical_router_snmp"`
	Perms2                          *PermType2          `json:"perms2,omitempty"`
	PhysicalRouterRole              PhysicalRouterRole  `json:"physical_router_role,omitempty"`
	PhysicalRouterUserCredentials   *UserCredentials    `json:"physical_router_user_credentials,omitempty"`
	ParentUUID                      string              `json:"parent_uuid,omitempty"`
	UUID                            string              `json:"uuid,omitempty"`
	PhysicalRouterVNCManaged        bool                `json:"physical_router_vnc_managed"`
	PhysicalRouterLLDP              bool                `json:"physical_router_lldp"`
	TelemetryInfo                   *TelemetryStateInfo `json:"telemetry_info,omitempty"`
	ParentType                      string              `json:"parent_type,omitempty"`
	IDPerms                         *IdPermsType        `json:"id_perms,omitempty"`
	DisplayName                     string              `json:"display_name,omitempty"`
	PhysicalRouterManagementIP      string              `json:"physical_router_management_ip,omitempty"`
	PhysicalRouterVendorName        string              `json:"physical_router_vendor_name,omitempty"`
	PhysicalRouterImageURI          string              `json:"physical_router_image_uri,omitempty"`
	PhysicalRouterDataplaneIP       string              `json:"physical_router_dataplane_ip,omitempty"`
	PhysicalRouterJunosServicePorts *JunosServicePorts  `json:"physical_router_junos_service_ports,omitempty"`
	FQName                          []string            `json:"fq_name,omitempty"`
	Annotations                     *KeyValuePairs      `json:"annotations,omitempty"`
	PhysicalRouterSNMPCredentials   *SNMPCredentials    `json:"physical_router_snmp_credentials,omitempty"`
	PhysicalRouterLoopbackIP        string              `json:"physical_router_loopback_ip,omitempty"`

	BGPRouterRefs      []*PhysicalRouterBGPRouterRef      `json:"bgp_router_refs,omitempty"`
	VirtualRouterRefs  []*PhysicalRouterVirtualRouterRef  `json:"virtual_router_refs,omitempty"`
	VirtualNetworkRefs []*PhysicalRouterVirtualNetworkRef `json:"virtual_network_refs,omitempty"`

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
		ParentUUID:                      "",
		UUID:                            "",
		TelemetryInfo:                   MakeTelemetryStateInfo(),
		ParentType:                      "",
		IDPerms:                         MakeIdPermsType(),
		DisplayName:                     "",
		PhysicalRouterManagementIP:      "",
		PhysicalRouterVendorName:        "",
		PhysicalRouterVNCManaged:        false,
		PhysicalRouterLLDP:              false,
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
		FQName:                        []string{},
		Annotations:                   MakeKeyValuePairs(),
		PhysicalRouterSNMPCredentials: MakeSNMPCredentials(),
		PhysicalRouterLoopbackIP:      "",
		PhysicalRouterImageURI:        "",
		PhysicalRouterDataplaneIP:     "",
		Perms2:                        MakePermType2(),
		PhysicalRouterRole:            MakePhysicalRouterRole(),
		PhysicalRouterUserCredentials: MakeUserCredentials(),
		PhysicalRouterProductName:     "",
		PhysicalRouterSNMP:            false,
	}
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}
