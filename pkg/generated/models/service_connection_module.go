
package models
// ServiceConnectionModule



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceConnectionModule_service_type int = iota
    propServiceConnectionModule_e2_service int = iota
    propServiceConnectionModule_uuid int = iota
    propServiceConnectionModule_parent_uuid int = iota
    propServiceConnectionModule_annotations int = iota
    propServiceConnectionModule_perms2 int = iota
    propServiceConnectionModule_parent_type int = iota
    propServiceConnectionModule_fq_name int = iota
    propServiceConnectionModule_id_perms int = iota
    propServiceConnectionModule_display_name int = iota
)

// ServiceConnectionModule 
type ServiceConnectionModule struct {

    ServiceType ServiceConnectionType `json:"service_type,omitempty"`
    E2Service E2servicetype `json:"e2_service,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`

    ServiceObjectRefs []*ServiceConnectionModuleServiceObjectRef `json:"service_object_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// ServiceConnectionModuleServiceObjectRef references each other
type ServiceConnectionModuleServiceObjectRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *ServiceConnectionModule) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceConnectionModule makes ServiceConnectionModule
func MakeServiceConnectionModule() *ServiceConnectionModule{
    return &ServiceConnectionModule{
    //TODO(nati): Apply default
    ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ServiceType: MakeServiceConnectionType(),
        E2Service: MakeE2servicetype(),
        UUID: "",
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeServiceConnectionModuleSlice makes a slice of ServiceConnectionModule
func MakeServiceConnectionModuleSlice() []*ServiceConnectionModule {
    return []*ServiceConnectionModule{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceConnectionModule) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *ServiceConnectionModule) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceConnectionModule) GetDefaultName() string {
    return strings.Replace("default-service_connection_module", "_", "-", -1)
}

func (model *ServiceConnectionModule) GetType() string {
    return strings.Replace("service_connection_module", "_", "-", -1)
}

func (model *ServiceConnectionModule) GetFQName() []string {
    return model.FQName
}

func (model *ServiceConnectionModule) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceConnectionModule) GetParentType() string {
    return model.ParentType
}

func (model *ServiceConnectionModule) GetUuid() string {
    return model.UUID
}

func (model *ServiceConnectionModule) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceConnectionModule) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceConnectionModule) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceConnectionModule) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceConnectionModule) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceConnectionModule_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propServiceConnectionModule_service_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceType as service_type")
        }
        msg["service_type"] = &val
    }
    
    if model.modified.Bit(propServiceConnectionModule_e2_service) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.E2Service); err != nil {
            return nil, errors.Wrap(err, "Marshal of: E2Service as e2_service")
        }
        msg["e2_service"] = &val
    }
    
    if model.modified.Bit(propServiceConnectionModule_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propServiceConnectionModule_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propServiceConnectionModule_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propServiceConnectionModule_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propServiceConnectionModule_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propServiceConnectionModule_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propServiceConnectionModule_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceConnectionModule) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceConnectionModule) UpdateReferences() error {
    return nil
}


