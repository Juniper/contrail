
package models
// ServiceTemplate



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceTemplate_perms2 int = iota
    propServiceTemplate_parent_type int = iota
    propServiceTemplate_display_name int = iota
    propServiceTemplate_annotations int = iota
    propServiceTemplate_service_template_properties int = iota
    propServiceTemplate_uuid int = iota
    propServiceTemplate_parent_uuid int = iota
    propServiceTemplate_fq_name int = iota
    propServiceTemplate_id_perms int = iota
)

// ServiceTemplate 
type ServiceTemplate struct {

    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ServiceTemplateProperties *ServiceTemplateType `json:"service_template_properties,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`

    ServiceApplianceSetRefs []*ServiceTemplateServiceApplianceSetRef `json:"service_appliance_set_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// ServiceTemplateServiceApplianceSetRef references each other
type ServiceTemplateServiceApplianceSetRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *ServiceTemplate) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceTemplate makes ServiceTemplate
func MakeServiceTemplate() *ServiceTemplate{
    return &ServiceTemplate{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        ParentType: "",
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        ServiceTemplateProperties: MakeServiceTemplateType(),
        UUID: "",
        ParentUUID: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        
        modified: big.NewInt(0),
    }
}



// MakeServiceTemplateSlice makes a slice of ServiceTemplate
func MakeServiceTemplateSlice() []*ServiceTemplate {
    return []*ServiceTemplate{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceTemplate) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[domain:0xc4202e5b80])
    fqn := []string{}
    
    fqn = Domain{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *ServiceTemplate) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-domain", "_", "-", -1)
}

func (model *ServiceTemplate) GetDefaultName() string {
    return strings.Replace("default-service_template", "_", "-", -1)
}

func (model *ServiceTemplate) GetType() string {
    return strings.Replace("service_template", "_", "-", -1)
}

func (model *ServiceTemplate) GetFQName() []string {
    return model.FQName
}

func (model *ServiceTemplate) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceTemplate) GetParentType() string {
    return model.ParentType
}

func (model *ServiceTemplate) GetUuid() string {
    return model.UUID
}

func (model *ServiceTemplate) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceTemplate) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceTemplate) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceTemplate) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceTemplate) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceTemplate_service_template_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceTemplateProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceTemplateProperties as service_template_properties")
        }
        msg["service_template_properties"] = &val
    }
    
    if model.modified.Bit(propServiceTemplate_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propServiceTemplate_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propServiceTemplate_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propServiceTemplate_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propServiceTemplate_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propServiceTemplate_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propServiceTemplate_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propServiceTemplate_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceTemplate) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceTemplate) UpdateReferences() error {
    return nil
}


