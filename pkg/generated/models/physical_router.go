package models

// PhysicalRouter

import "encoding/json"

// PhysicalRouter
type PhysicalRouter struct {
	DisplayName                     string              `json:"display_name,omitempty"`
	PhysicalRouterSNMPCredentials   *SNMPCredentials    `json:"physical_router_snmp_credentials,omitempty"`
	PhysicalRouterVendorName        string              `json:"physical_router_vendor_name,omitempty"`
	PhysicalRouterJunosServicePorts *JunosServicePorts  `json:"physical_router_junos_service_ports,omitempty"`
	Perms2                          *PermType2          `json:"perms2,omitempty"`
	PhysicalRouterManagementIP      string              `json:"physical_router_management_ip,omitempty"`
	PhysicalRouterUserCredentials   *UserCredentials    `json:"physical_router_user_credentials,omitempty"`
	PhysicalRouterDataplaneIP       string              `json:"physical_router_dataplane_ip,omitempty"`
	IDPerms                         *IdPermsType        `json:"id_perms,omitempty"`
	ParentUUID                      string              `json:"parent_uuid,omitempty"`
	FQName                          []string            `json:"fq_name,omitempty"`
	PhysicalRouterProductName       string              `json:"physical_router_product_name,omitempty"`
	PhysicalRouterLLDP              bool                `json:"physical_router_lldp,omitempty"`
	TelemetryInfo                   *TelemetryStateInfo `json:"telemetry_info,omitempty"`
	PhysicalRouterSNMP              bool                `json:"physical_router_snmp,omitempty"`
	UUID                            string              `json:"uuid,omitempty"`
	ParentType                      string              `json:"parent_type,omitempty"`
	Annotations                     *KeyValuePairs      `json:"annotations,omitempty"`
	PhysicalRouterRole              PhysicalRouterRole  `json:"physical_router_role,omitempty"`
	PhysicalRouterVNCManaged        bool                `json:"physical_router_vnc_managed,omitempty"`
	PhysicalRouterLoopbackIP        string              `json:"physical_router_loopback_ip,omitempty"`
	PhysicalRouterImageURI          string              `json:"physical_router_image_uri,omitempty"`

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
		PhysicalRouterImageURI: "",
		UUID:                     "",
		ParentType:               "",
		Annotations:              MakeKeyValuePairs(),
		PhysicalRouterRole:       MakePhysicalRouterRole(),
		PhysicalRouterVNCManaged: false,
		PhysicalRouterLoopbackIP: "",
		Perms2:                          MakePermType2(),
		DisplayName:                     "",
		PhysicalRouterSNMPCredentials:   MakeSNMPCredentials(),
		PhysicalRouterVendorName:        "",
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
		IDPerms:                       MakeIdPermsType(),
		PhysicalRouterManagementIP:    "",
		PhysicalRouterUserCredentials: MakeUserCredentials(),
		PhysicalRouterDataplaneIP:     "",
		PhysicalRouterSNMP:            false,
		ParentUUID:                    "",
		FQName:                        []string{},
		PhysicalRouterProductName: "",
		PhysicalRouterLLDP:        false,
		TelemetryInfo:             MakeTelemetryStateInfo(),
	}
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}
