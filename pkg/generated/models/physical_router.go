package models

// PhysicalRouter

import "encoding/json"

// PhysicalRouter
type PhysicalRouter struct {
	PhysicalRouterUserCredentials   *UserCredentials    `json:"physical_router_user_credentials,omitempty"`
	ParentUUID                      string              `json:"parent_uuid,omitempty"`
	DisplayName                     string              `json:"display_name,omitempty"`
	PhysicalRouterSNMPCredentials   *SNMPCredentials    `json:"physical_router_snmp_credentials,omitempty"`
	PhysicalRouterProductName       string              `json:"physical_router_product_name,omitempty"`
	PhysicalRouterImageURI          string              `json:"physical_router_image_uri,omitempty"`
	TelemetryInfo                   *TelemetryStateInfo `json:"telemetry_info,omitempty"`
	Annotations                     *KeyValuePairs      `json:"annotations,omitempty"`
	Perms2                          *PermType2          `json:"perms2,omitempty"`
	UUID                            string              `json:"uuid,omitempty"`
	FQName                          []string            `json:"fq_name,omitempty"`
	PhysicalRouterLLDP              bool                `json:"physical_router_lldp"`
	PhysicalRouterSNMP              bool                `json:"physical_router_snmp"`
	PhysicalRouterJunosServicePorts *JunosServicePorts  `json:"physical_router_junos_service_ports,omitempty"`
	ParentType                      string              `json:"parent_type,omitempty"`
	PhysicalRouterManagementIP      string              `json:"physical_router_management_ip,omitempty"`
	PhysicalRouterRole              PhysicalRouterRole  `json:"physical_router_role,omitempty"`
	PhysicalRouterVendorName        string              `json:"physical_router_vendor_name,omitempty"`
	PhysicalRouterVNCManaged        bool                `json:"physical_router_vnc_managed"`
	PhysicalRouterLoopbackIP        string              `json:"physical_router_loopback_ip,omitempty"`
	PhysicalRouterDataplaneIP       string              `json:"physical_router_dataplane_ip,omitempty"`
	IDPerms                         *IdPermsType        `json:"id_perms,omitempty"`

	VirtualNetworkRefs []*PhysicalRouterVirtualNetworkRef `json:"virtual_network_refs,omitempty"`
	BGPRouterRefs      []*PhysicalRouterBGPRouterRef      `json:"bgp_router_refs,omitempty"`
	VirtualRouterRefs  []*PhysicalRouterVirtualRouterRef  `json:"virtual_router_refs,omitempty"`

	LogicalInterfaces  []*LogicalInterface  `json:"logical_interfaces,omitempty"`
	PhysicalInterfaces []*PhysicalInterface `json:"physical_interfaces,omitempty"`
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

// String returns json representation of the object
func (model *PhysicalRouter) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePhysicalRouter makes PhysicalRouter
func MakePhysicalRouter() *PhysicalRouter {
	return &PhysicalRouter{
		//TODO(nati): Apply default
		PhysicalRouterManagementIP: "",
		PhysicalRouterRole:         MakePhysicalRouterRole(),
		PhysicalRouterVendorName:   "",
		PhysicalRouterVNCManaged:   false,
		PhysicalRouterLoopbackIP:   "",
		PhysicalRouterDataplaneIP:  "",
		IDPerms:                    MakeIdPermsType(),
		PhysicalRouterUserCredentials: MakeUserCredentials(),
		ParentUUID:                    "",
		DisplayName:                   "",
		FQName:                        []string{},
		PhysicalRouterSNMPCredentials:   MakeSNMPCredentials(),
		PhysicalRouterProductName:       "",
		PhysicalRouterImageURI:          "",
		TelemetryInfo:                   MakeTelemetryStateInfo(),
		Annotations:                     MakeKeyValuePairs(),
		Perms2:                          MakePermType2(),
		UUID:                            "",
		PhysicalRouterLLDP:              false,
		PhysicalRouterSNMP:              false,
		PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
		ParentType:                      "",
	}
}

// MakePhysicalRouterSlice() makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
	return []*PhysicalRouter{}
}
