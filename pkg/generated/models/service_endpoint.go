
package models
// ServiceEndpoint



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceEndpoint_parent_type int = iota
    propServiceEndpoint_fq_name int = iota
    propServiceEndpoint_id_perms int = iota
    propServiceEndpoint_display_name int = iota
    propServiceEndpoint_annotations int = iota
    propServiceEndpoint_perms2 int = iota
    propServiceEndpoint_uuid int = iota
    propServiceEndpoint_parent_uuid int = iota
)

// ServiceEndpoint 
type ServiceEndpoint struct {

    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`

    ServiceObjectRefs []*ServiceEndpointServiceObjectRef `json:"service_object_refs,omitempty"`
    ServiceConnectionModuleRefs []*ServiceEndpointServiceConnectionModuleRef `json:"service_connection_module_refs,omitempty"`
    PhysicalRouterRefs []*ServiceEndpointPhysicalRouterRef `json:"physical_router_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
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
    Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        
        modified: big.NewInt(0),
    }
}



// MakeServiceEndpointSlice makes a slice of ServiceEndpoint
func MakeServiceEndpointSlice() []*ServiceEndpoint {
    return []*ServiceEndpoint{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceEndpoint) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *ServiceEndpoint) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceEndpoint) GetDefaultName() string {
    return strings.Replace("default-service_endpoint", "_", "-", -1)
}

func (model *ServiceEndpoint) GetType() string {
    return strings.Replace("service_endpoint", "_", "-", -1)
}

func (model *ServiceEndpoint) GetFQName() []string {
    return model.FQName
}

func (model *ServiceEndpoint) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceEndpoint) GetParentType() string {
    return model.ParentType
}

func (model *ServiceEndpoint) GetUuid() string {
    return model.UUID
}

func (model *ServiceEndpoint) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceEndpoint) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceEndpoint) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceEndpoint) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceEndpoint) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceEndpoint_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propServiceEndpoint_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propServiceEndpoint_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propServiceEndpoint_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propServiceEndpoint_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propServiceEndpoint_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propServiceEndpoint_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propServiceEndpoint_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceEndpoint) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceEndpoint) UpdateReferences() error {
    return nil
}


