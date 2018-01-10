package models

// ServiceEndpoint

import "encoding/json"

// ServiceEndpoint
type ServiceEndpoint struct {
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`

	ServiceConnectionModuleRefs []*ServiceEndpointServiceConnectionModuleRef `json:"service_connection_module_refs"`
	PhysicalRouterRefs          []*ServiceEndpointPhysicalRouterRef          `json:"physical_router_refs"`
	ServiceObjectRefs           []*ServiceEndpointServiceObjectRef           `json:"service_object_refs"`
}

// ServiceEndpointServiceConnectionModuleRef references each other
type ServiceEndpointServiceConnectionModuleRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// ServiceEndpointPhysicalRouterRef references each other
type ServiceEndpointPhysicalRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// ServiceEndpointServiceObjectRef references each other
type ServiceEndpointServiceObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *ServiceEndpoint) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceEndpoint makes ServiceEndpoint
func MakeServiceEndpoint() *ServiceEndpoint {
	return &ServiceEndpoint{
		//TODO(nati): Apply default
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
	}
}

// InterfaceToServiceEndpoint makes ServiceEndpoint from interface
func InterfaceToServiceEndpoint(iData interface{}) *ServiceEndpoint {
	data := iData.(map[string]interface{})
	return &ServiceEndpoint{
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
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

	}
}

// InterfaceToServiceEndpointSlice makes a slice of ServiceEndpoint from interface
func InterfaceToServiceEndpointSlice(data interface{}) []*ServiceEndpoint {
	list := data.([]interface{})
	result := MakeServiceEndpointSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceEndpoint(item))
	}
	return result
}

// MakeServiceEndpointSlice() makes a slice of ServiceEndpoint
func MakeServiceEndpointSlice() []*ServiceEndpoint {
	return []*ServiceEndpoint{}
}
