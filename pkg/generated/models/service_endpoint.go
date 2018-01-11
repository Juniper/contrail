package models

// ServiceEndpoint

import "encoding/json"

// ServiceEndpoint
type ServiceEndpoint struct {
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`

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
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
	}
}

// MakeServiceEndpointSlice() makes a slice of ServiceEndpoint
func MakeServiceEndpointSlice() []*ServiceEndpoint {
	return []*ServiceEndpoint{}
}
