package models

// PhysicalRouter

import "encoding/json"

// PhysicalRouter
type PhysicalRouter struct {
	TelemetryInfo                   *TelemetryStateInfo `json:"telemetry_info,omitempty"`
	ParentType                      string              `json:"parent_type,omitempty"`
	PhysicalRouterUserCredentials   *UserCredentials    `json:"physical_router_user_credentials,omitempty"`
	PhysicalRouterVendorName        string              `json:"physical_router_vendor_name,omitempty"`
	PhysicalRouterProductName       string              `json:"physical_router_product_name,omitempty"`
	PhysicalRouterLLDP              bool                `json:"physical_router_lldp"`
	PhysicalRouterLoopbackIP        string              `json:"physical_router_loopback_ip,omitempty"`
	FQName                          []string            `json:"fq_name,omitempty"`
	IDPerms                         *IdPermsType        `json:"id_perms,omitempty"`
	PhysicalRouterManagementIP      string              `json:"physical_router_management_ip,omitempty"`
	PhysicalRouterSNMP              bool                `json:"physical_router_snmp"`
	DisplayName                     string              `json:"display_name,omitempty"`
	Perms2                          *PermType2          `json:"perms2,omitempty"`
	UUID                            string              `json:"uuid,omitempty"`
	PhysicalRouterSNMPCredentials   *SNMPCredentials    `json:"physical_router_snmp_credentials,omitempty"`
	PhysicalRouterJunosServicePorts *JunosServicePorts  `json:"physical_router_junos_service_ports,omitempty"`
	Annotations                     *KeyValuePairs      `json:"annotations,omitempty"`
	PhysicalRouterRole              PhysicalRouterRole  `json:"physical_router_role,omitempty"`
	PhysicalRouterVNCManaged        bool                `json:"physical_router_vnc_managed"`
	PhysicalRouterImageURI          string              `json:"physical_router_image_uri,omitempty"`
	PhysicalRouterDataplaneIP       string              `json:"physical_router_dataplane_ip,omitempty"`
	ParentUUID                      string              `json:"parent_uuid,omitempty"`

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
		PhysicalRouterLLDP:            false,
		PhysicalRouterLoopbackIP:      "",
		TelemetryInfo:                 MakeTelemetryStateInfo(),
		ParentType:                    "",
		PhysicalRouterUserCredentials: MakeUserCredentials(),
		PhysicalRouterVendorName:      "",
		PhysicalRouterProductName:     "",
		Perms2:                          MakePermType2(),
		UUID:                            "",
		FQName:                          []string{},
		IDPerms:                         MakeIdPermsType(),
		PhysicalRouterManagementIP:      "",
		PhysicalRouterSNMP:              false,
		DisplayName:                     "",
		PhysicalRouterSNMPCredentials:   MakeSNMPCredentials(),
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
		Annotations:                     MakeKeyValuePairs(),
		PhysicalRouterDataplaneIP:       "",
		ParentUUID:                      "",
		PhysicalRouterRole:              MakePhysicalRouterRole(),
		PhysicalRouterVNCManaged:        false,
		PhysicalRouterImageURI:          "",
	}
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}
