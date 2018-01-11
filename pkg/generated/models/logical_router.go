package models

// LogicalRouter

import "encoding/json"

// LogicalRouter
type LogicalRouter struct {
	Perms2                    *PermType2       `json:"perms2"`
	UUID                      string           `json:"uuid"`
	ParentUUID                string           `json:"parent_uuid"`
	ParentType                string           `json:"parent_type"`
	FQName                    []string         `json:"fq_name"`
	Annotations               *KeyValuePairs   `json:"annotations"`
	VxlanNetworkIdentifier    string           `json:"vxlan_network_identifier"`
	ConfiguredRouteTargetList *RouteTargetList `json:"configured_route_target_list"`
	IDPerms                   *IdPermsType     `json:"id_perms"`
	DisplayName               string           `json:"display_name"`

	ServiceInstanceRefs         []*LogicalRouterServiceInstanceRef         `json:"service_instance_refs"`
	RouteTableRefs              []*LogicalRouterRouteTableRef              `json:"route_table_refs"`
	VirtualNetworkRefs          []*LogicalRouterVirtualNetworkRef          `json:"virtual_network_refs"`
	PhysicalRouterRefs          []*LogicalRouterPhysicalRouterRef          `json:"physical_router_refs"`
	BGPVPNRefs                  []*LogicalRouterBGPVPNRef                  `json:"bgpvpn_refs"`
	RouteTargetRefs             []*LogicalRouterRouteTargetRef             `json:"route_target_refs"`
	VirtualMachineInterfaceRefs []*LogicalRouterVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
}

// LogicalRouterRouteTargetRef references each other
type LogicalRouterRouteTargetRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterVirtualMachineInterfaceRef references each other
type LogicalRouterVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterServiceInstanceRef references each other
type LogicalRouterServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterRouteTableRef references each other
type LogicalRouterRouteTableRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterVirtualNetworkRef references each other
type LogicalRouterVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterPhysicalRouterRef references each other
type LogicalRouterPhysicalRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterBGPVPNRef references each other
type LogicalRouterBGPVPNRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *LogicalRouter) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLogicalRouter makes LogicalRouter
func MakeLogicalRouter() *LogicalRouter {
	return &LogicalRouter{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ConfiguredRouteTargetList: MakeRouteTargetList(),
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		VxlanNetworkIdentifier:    "",
	}
}

// MakeLogicalRouterSlice() makes a slice of LogicalRouter
func MakeLogicalRouterSlice() []*LogicalRouter {
	return []*LogicalRouter{}
}
