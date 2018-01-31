package models

// InterfaceRouteTable

import "encoding/json"

// InterfaceRouteTable
type InterfaceRouteTable struct {
	FQName                    []string        `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType    `json:"id_perms,omitempty"`
	DisplayName               string          `json:"display_name,omitempty"`
	Perms2                    *PermType2      `json:"perms2,omitempty"`
	ParentType                string          `json:"parent_type,omitempty"`
	InterfaceRouteTableRoutes *RouteTableType `json:"interface_route_table_routes,omitempty"`
	UUID                      string          `json:"uuid,omitempty"`
	ParentUUID                string          `json:"parent_uuid,omitempty"`
	Annotations               *KeyValuePairs  `json:"annotations,omitempty"`

	ServiceInstanceRefs []*InterfaceRouteTableServiceInstanceRef `json:"service_instance_refs,omitempty"`
}

// InterfaceRouteTableServiceInstanceRef references each other
type InterfaceRouteTableServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceInterfaceTag
}

// String returns json representation of the object
func (model *InterfaceRouteTable) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeInterfaceRouteTable makes InterfaceRouteTable
func MakeInterfaceRouteTable() *InterfaceRouteTable {
	return &InterfaceRouteTable{
		//TODO(nati): Apply default
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Perms2:                    MakePermType2(),
		ParentType:                "",
		InterfaceRouteTableRoutes: MakeRouteTableType(),
		FQName:      []string{},
		ParentUUID:  "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
	}
}

// MakeInterfaceRouteTableSlice() makes a slice of InterfaceRouteTable
func MakeInterfaceRouteTableSlice() []*InterfaceRouteTable {
	return []*InterfaceRouteTable{}
}
