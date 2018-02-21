
package models
// ServiceInstance



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceInstance_fq_name int = iota
    propServiceInstance_display_name int = iota
    propServiceInstance_annotations int = iota
    propServiceInstance_perms2 int = iota
    propServiceInstance_uuid int = iota
    propServiceInstance_parent_type int = iota
    propServiceInstance_service_instance_bindings int = iota
    propServiceInstance_service_instance_properties int = iota
    propServiceInstance_id_perms int = iota
    propServiceInstance_parent_uuid int = iota
)

// ServiceInstance 
type ServiceInstance struct {

    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ServiceInstanceBindings *KeyValuePairs `json:"service_instance_bindings,omitempty"`
    ServiceInstanceProperties *ServiceInstanceType `json:"service_instance_properties,omitempty"`

    ServiceTemplateRefs []*ServiceInstanceServiceTemplateRef `json:"service_template_refs,omitempty"`
    InstanceIPRefs []*ServiceInstanceInstanceIPRef `json:"instance_ip_refs,omitempty"`

    PortTuples []*PortTuple `json:"port_tuples,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// ServiceInstanceInstanceIPRef references each other
type ServiceInstanceInstanceIPRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *ServiceInterfaceTag
    
}

// ServiceInstanceServiceTemplateRef references each other
type ServiceInstanceServiceTemplateRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *ServiceInstance) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceInstance makes ServiceInstance
func MakeServiceInstance() *ServiceInstance{
    return &ServiceInstance{
    //TODO(nati): Apply default
    FQName: []string{},
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentType: "",
        ServiceInstanceBindings: MakeKeyValuePairs(),
        ServiceInstanceProperties: MakeServiceInstanceType(),
        IDPerms: MakeIdPermsType(),
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeServiceInstanceSlice makes a slice of ServiceInstance
func MakeServiceInstanceSlice() []*ServiceInstance {
    return []*ServiceInstance{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceInstance) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4202e5a40])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *ServiceInstance) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *ServiceInstance) GetDefaultName() string {
    return strings.Replace("default-service_instance", "_", "-", -1)
}

func (model *ServiceInstance) GetType() string {
    return strings.Replace("service_instance", "_", "-", -1)
}

func (model *ServiceInstance) GetFQName() []string {
    return model.FQName
}

func (model *ServiceInstance) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceInstance) GetParentType() string {
    return model.ParentType
}

func (model *ServiceInstance) GetUuid() string {
    return model.UUID
}

func (model *ServiceInstance) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceInstance) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceInstance) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceInstance) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceInstance) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceInstance_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propServiceInstance_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propServiceInstance_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propServiceInstance_service_instance_bindings) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceInstanceBindings); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceInstanceBindings as service_instance_bindings")
        }
        msg["service_instance_bindings"] = &val
    }
    
    if model.modified.Bit(propServiceInstance_service_instance_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceInstanceProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceInstanceProperties as service_instance_properties")
        }
        msg["service_instance_properties"] = &val
    }
    
    if model.modified.Bit(propServiceInstance_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propServiceInstance_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propServiceInstance_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propServiceInstance_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propServiceInstance_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceInstance) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceInstance) UpdateReferences() error {
    return nil
}


