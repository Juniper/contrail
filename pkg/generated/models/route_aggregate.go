package models

// RouteAggregate

import "encoding/json"

// RouteAggregate
type RouteAggregate struct {
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`

	ServiceInstanceRefs []*RouteAggregateServiceInstanceRef `json:"service_instance_refs"`
}

// RouteAggregateServiceInstanceRef references each other
type RouteAggregateServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceInterfaceTag
}

// String returns json representation of the object
func (model *RouteAggregate) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteAggregate makes RouteAggregate
func MakeRouteAggregate() *RouteAggregate {
	return &RouteAggregate{
		//TODO(nati): Apply default
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
	}
}

// InterfaceToRouteAggregate makes RouteAggregate from interface
func InterfaceToRouteAggregate(iData interface{}) *RouteAggregate {
	data := iData.(map[string]interface{})
	return &RouteAggregate{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToRouteAggregateSlice makes a slice of RouteAggregate from interface
func InterfaceToRouteAggregateSlice(data interface{}) []*RouteAggregate {
	list := data.([]interface{})
	result := MakeRouteAggregateSlice()
	for _, item := range list {
		result = append(result, InterfaceToRouteAggregate(item))
	}
	return result
}

// MakeRouteAggregateSlice() makes a slice of RouteAggregate
func MakeRouteAggregateSlice() []*RouteAggregate {
	return []*RouteAggregate{}
}
