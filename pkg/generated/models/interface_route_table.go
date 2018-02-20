package models

// InterfaceRouteTable

// InterfaceRouteTable
//proteus:generate
type InterfaceRouteTable struct {
	UUID                      string          `json:"uuid,omitempty"`
	ParentUUID                string          `json:"parent_uuid,omitempty"`
	ParentType                string          `json:"parent_type,omitempty"`
	FQName                    []string        `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType    `json:"id_perms,omitempty"`
	DisplayName               string          `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs  `json:"annotations,omitempty"`
	Perms2                    *PermType2      `json:"perms2,omitempty"`
	InterfaceRouteTableRoutes *RouteTableType `json:"interface_route_table_routes,omitempty"`

	ServiceInstanceRefs []*InterfaceRouteTableServiceInstanceRef `json:"service_instance_refs,omitempty"`
}

// InterfaceRouteTableServiceInstanceRef references each other
type InterfaceRouteTableServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceInterfaceTag
}

// MakeInterfaceRouteTable makes InterfaceRouteTable
func MakeInterfaceRouteTable() *InterfaceRouteTable {
	return &InterfaceRouteTable{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		InterfaceRouteTableRoutes: MakeRouteTableType(),
	}
}

// MakeInterfaceRouteTableSlice() makes a slice of InterfaceRouteTable
func MakeInterfaceRouteTableSlice() []*InterfaceRouteTable {
	return []*InterfaceRouteTable{}
}
