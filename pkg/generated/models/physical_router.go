package models

// PhysicalRouter

import "encoding/json"

// PhysicalRouter
type PhysicalRouter struct {
	FQName                          []string            `json:"fq_name,omitempty"`
	DisplayName                     string              `json:"display_name,omitempty"`
	PhysicalRouterVNCManaged        bool                `json:"physical_router_vnc_managed"`
	PhysicalRouterLLDP              bool                `json:"physical_router_lldp"`
	PhysicalRouterLoopbackIP        string              `json:"physical_router_loopback_ip,omitempty"`
	TelemetryInfo                   *TelemetryStateInfo `json:"telemetry_info,omitempty"`
	Annotations                     *KeyValuePairs      `json:"annotations,omitempty"`
	UUID                            string              `json:"uuid,omitempty"`
	PhysicalRouterSNMPCredentials   *SNMPCredentials    `json:"physical_router_snmp_credentials,omitempty"`
	PhysicalRouterUserCredentials   *UserCredentials    `json:"physical_router_user_credentials,omitempty"`
	PhysicalRouterJunosServicePorts *JunosServicePorts  `json:"physical_router_junos_service_ports,omitempty"`
	IDPerms                         *IdPermsType        `json:"id_perms,omitempty"`
	ParentUUID                      string              `json:"parent_uuid,omitempty"`
	PhysicalRouterManagementIP      string              `json:"physical_router_management_ip,omitempty"`
	PhysicalRouterVendorName        string              `json:"physical_router_vendor_name,omitempty"`
	PhysicalRouterImageURI          string              `json:"physical_router_image_uri,omitempty"`
	PhysicalRouterSNMP              bool                `json:"physical_router_snmp"`
	Perms2                          *PermType2          `json:"perms2,omitempty"`
	PhysicalRouterRole              PhysicalRouterRole  `json:"physical_router_role,omitempty"`
	PhysicalRouterProductName       string              `json:"physical_router_product_name,omitempty"`
	PhysicalRouterDataplaneIP       string              `json:"physical_router_dataplane_ip,omitempty"`
	ParentType                      string              `json:"parent_type,omitempty"`

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
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		PhysicalRouterSNMPCredentials:   MakeSNMPCredentials(),
		PhysicalRouterUserCredentials:   MakeUserCredentials(),
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
		IDPerms:                    MakeIdPermsType(),
		ParentUUID:                 "",
		PhysicalRouterManagementIP: "",
		PhysicalRouterVendorName:   "",
		PhysicalRouterImageURI:     "",
		PhysicalRouterSNMP:         false,
		Perms2:                     MakePermType2(),
		PhysicalRouterRole:         MakePhysicalRouterRole(),
		PhysicalRouterProductName:  "",
		PhysicalRouterDataplaneIP:  "",
		ParentType:                 "",
		FQName:                     []string{},
		DisplayName:                "",
		PhysicalRouterVNCManaged:   false,
		PhysicalRouterLLDP:         false,
		PhysicalRouterLoopbackIP:   "",
		TelemetryInfo:              MakeTelemetryStateInfo(),
	}
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}
