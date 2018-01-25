package models
// ServiceEndpoint



import "encoding/json"

// ServiceEndpoint 
//proteus:generate
type ServiceEndpoint struct {

    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`

    ServiceConnectionModuleRefs []*ServiceEndpointServiceConnectionModuleRef `json:"service_connection_module_refs,omitempty"`
    PhysicalRouterRefs []*ServiceEndpointPhysicalRouterRef `json:"physical_router_refs,omitempty"`
    ServiceObjectRefs []*ServiceEndpointServiceObjectRef `json:"service_object_refs,omitempty"`

}


// ServiceEndpointServiceConnectionModuleRef references each other
type ServiceEndpointServiceConnectionModuleRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// ServiceEndpointPhysicalRouterRef references each other
type ServiceEndpointPhysicalRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// ServiceEndpointServiceObjectRef references each other
type ServiceEndpointServiceObjectRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *ServiceEndpoint) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceEndpoint makes ServiceEndpoint
func MakeServiceEndpoint() *ServiceEndpoint{
    return &ServiceEndpoint{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        
    }
}



// MakeServiceEndpointSlice() makes a slice of ServiceEndpoint
func MakeServiceEndpointSlice() []*ServiceEndpoint {
    return []*ServiceEndpoint{}
}
