package models

// InterfaceRouteTable

import "encoding/json"

// InterfaceRouteTable
type InterfaceRouteTable struct {
	IDPerms                   *IdPermsType    `json:"id_perms,omitempty"`
	Perms2                    *PermType2      `json:"perms2,omitempty"`
	ParentUUID                string          `json:"parent_uuid,omitempty"`
	ParentType                string          `json:"parent_type,omitempty"`
	FQName                    []string        `json:"fq_name,omitempty"`
	InterfaceRouteTableRoutes *RouteTableType `json:"interface_route_table_routes,omitempty"`
	DisplayName               string          `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs  `json:"annotations,omitempty"`
	UUID                      string          `json:"uuid,omitempty"`

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
		ParentType: "",
		FQName:     []string{},
		IDPerms:    MakeIdPermsType(),
		Perms2:     MakePermType2(),
		ParentUUID: "",
		UUID:       "",
		InterfaceRouteTableRoutes: MakeRouteTableType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
	}
}

// MakeInterfaceRouteTableSlice() makes a slice of InterfaceRouteTable
func MakeInterfaceRouteTableSlice() []*InterfaceRouteTable {
	return []*InterfaceRouteTable{}
}
