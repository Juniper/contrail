package models

// InterfaceRouteTable

import "encoding/json"

// InterfaceRouteTable
type InterfaceRouteTable struct {
	FQName                    []string        `json:"fq_name"`
	Annotations               *KeyValuePairs  `json:"annotations"`
	UUID                      string          `json:"uuid"`
	ParentUUID                string          `json:"parent_uuid"`
	ParentType                string          `json:"parent_type"`
	IDPerms                   *IdPermsType    `json:"id_perms"`
	DisplayName               string          `json:"display_name"`
	InterfaceRouteTableRoutes *RouteTableType `json:"interface_route_table_routes"`
	Perms2                    *PermType2      `json:"perms2"`

	ServiceInstanceRefs []*InterfaceRouteTableServiceInstanceRef `json:"service_instance_refs"`
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
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		InterfaceRouteTableRoutes: MakeRouteTableType(),
		Perms2:      MakePermType2(),
		FQName:      []string{},
		Annotations: MakeKeyValuePairs(),
	}
}

// InterfaceToInterfaceRouteTable makes InterfaceRouteTable from interface
func InterfaceToInterfaceRouteTable(iData interface{}) *InterfaceRouteTable {
	data := iData.(map[string]interface{})
	return &InterfaceRouteTable{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		InterfaceRouteTableRoutes: InterfaceToRouteTableType(data["interface_route_table_routes"]),

		//{"description":"Interface route table used same structure as route table, however the next hop field is irrelevant.","type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}

	}
}

// InterfaceToInterfaceRouteTableSlice makes a slice of InterfaceRouteTable from interface
func InterfaceToInterfaceRouteTableSlice(data interface{}) []*InterfaceRouteTable {
	list := data.([]interface{})
	result := MakeInterfaceRouteTableSlice()
	for _, item := range list {
		result = append(result, InterfaceToInterfaceRouteTable(item))
	}
	return result
}

// MakeInterfaceRouteTableSlice() makes a slice of InterfaceRouteTable
func MakeInterfaceRouteTableSlice() []*InterfaceRouteTable {
	return []*InterfaceRouteTable{}
}
