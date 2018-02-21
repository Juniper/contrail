
package models
// ServiceHealthCheck



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceHealthCheck_parent_uuid int = iota
    propServiceHealthCheck_id_perms int = iota
    propServiceHealthCheck_display_name int = iota
    propServiceHealthCheck_annotations int = iota
    propServiceHealthCheck_perms2 int = iota
    propServiceHealthCheck_uuid int = iota
    propServiceHealthCheck_parent_type int = iota
    propServiceHealthCheck_fq_name int = iota
    propServiceHealthCheck_service_health_check_properties int = iota
)

// ServiceHealthCheck 
type ServiceHealthCheck struct {

    ServiceHealthCheckProperties *ServiceHealthCheckType `json:"service_health_check_properties,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`

    ServiceInstanceRefs []*ServiceHealthCheckServiceInstanceRef `json:"service_instance_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// ServiceHealthCheckServiceInstanceRef references each other
type ServiceHealthCheckServiceInstanceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *ServiceInterfaceTag
    
}


// String returns json representation of the object
func (model *ServiceHealthCheck) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceHealthCheck makes ServiceHealthCheck
func MakeServiceHealthCheck() *ServiceHealthCheck{
    return &ServiceHealthCheck{
    //TODO(nati): Apply default
    Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ParentUUID: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        ServiceHealthCheckProperties: MakeServiceHealthCheckType(),
        UUID: "",
        ParentType: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeServiceHealthCheckSlice makes a slice of ServiceHealthCheck
func MakeServiceHealthCheckSlice() []*ServiceHealthCheck {
    return []*ServiceHealthCheck{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceHealthCheck) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4202e5860])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *ServiceHealthCheck) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *ServiceHealthCheck) GetDefaultName() string {
    return strings.Replace("default-service_health_check", "_", "-", -1)
}

func (model *ServiceHealthCheck) GetType() string {
    return strings.Replace("service_health_check", "_", "-", -1)
}

func (model *ServiceHealthCheck) GetFQName() []string {
    return model.FQName
}

func (model *ServiceHealthCheck) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceHealthCheck) GetParentType() string {
    return model.ParentType
}

func (model *ServiceHealthCheck) GetUuid() string {
    return model.UUID
}

func (model *ServiceHealthCheck) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceHealthCheck) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceHealthCheck) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceHealthCheck) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceHealthCheck) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceHealthCheck_service_health_check_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceHealthCheckProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceHealthCheckProperties as service_health_check_properties")
        }
        msg["service_health_check_properties"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheck_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheck_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheck_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheck_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheck_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheck_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheck_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propServiceHealthCheck_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceHealthCheck) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceHealthCheck) UpdateReferences() error {
    return nil
}


