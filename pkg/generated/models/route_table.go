package models

// RouteTable

import "encoding/json"

// RouteTable
type RouteTable struct {
	ParentType  string          `json:"parent_type"`
	Routes      *RouteTableType `json:"routes"`
	FQName      []string        `json:"fq_name"`
	Annotations *KeyValuePairs  `json:"annotations"`
	Perms2      *PermType2      `json:"perms2"`
	UUID        string          `json:"uuid"`
	IDPerms     *IdPermsType    `json:"id_perms"`
	DisplayName string          `json:"display_name"`
	ParentUUID  string          `json:"parent_uuid"`
}

// String returns json representation of the object
func (model *RouteTable) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteTable makes RouteTable
func MakeRouteTable() *RouteTable {
	return &RouteTable{
		//TODO(nati): Apply default
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		ParentUUID:  "",
		Routes:      MakeRouteTableType(),
		FQName:      []string{},
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentType:  "",
	}
}

// InterfaceToRouteTable makes RouteTable from interface
func InterfaceToRouteTable(iData interface{}) *RouteTable {
	data := iData.(map[string]interface{})
	return &RouteTable{
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		Routes: InterfaceToRouteTableType(data["routes"]),

		//{"description":"Routes in the route table are configured in following way.","type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}

	}
}

// InterfaceToRouteTableSlice makes a slice of RouteTable from interface
func InterfaceToRouteTableSlice(data interface{}) []*RouteTable {
	list := data.([]interface{})
	result := MakeRouteTableSlice()
	for _, item := range list {
		result = append(result, InterfaceToRouteTable(item))
	}
	return result
}

// MakeRouteTableSlice() makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
	return []*RouteTable{}
}
